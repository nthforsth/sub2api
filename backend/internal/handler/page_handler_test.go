package handler

import (
	"net/http"
	"net/http/httptest"
	"os"
	"path/filepath"
	"strings"
	"testing"

	"github.com/gin-gonic/gin"
)

func TestCleanPageImageRelativePath(t *testing.T) {
	tests := []struct {
		name string
		in   string
		want string
		ok   bool
	}{
		{name: "single filename", in: "logo.png", want: "logo.png", ok: true},
		{name: "nested path", in: "images/logo.png", want: filepath.Join("images", "logo.png"), ok: true},
		{name: "dot prefix", in: "./logo.png", want: "logo.png", ok: true},
		{name: "url escaped slash", in: "images%2Flogo.png", want: filepath.Join("images", "logo.png"), ok: true},
		{name: "parent traversal", in: "../secret.png", ok: false},
		{name: "encoded parent traversal", in: "%2e%2e/secret.png", ok: false},
		{name: "backslash traversal", in: `images\secret.png`, ok: false},
		{name: "absolute path", in: "/etc/passwd", ok: false},
		{name: "encoded absolute path", in: "%2fetc/passwd", ok: false},
		{name: "encoded nul byte", in: "logo.png%00", ok: false},
		{name: "invalid escape", in: "logo.png%zz", ok: false},
		{name: "empty path", in: "", ok: false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, ok := cleanPageImageRelativePath(tt.in)
			if ok != tt.ok {
				t.Fatalf("ok = %v, want %v", ok, tt.ok)
			}
			if got != tt.want {
				t.Fatalf("path = %q, want %q", got, tt.want)
			}
		})
	}
}

func TestResolvePageImagePath(t *testing.T) {
	root := t.TempDir()
	pagesDir := filepath.Join(root, "pages")
	base := filepath.Join(pagesDir, "guide")
	if err := os.MkdirAll(filepath.Join(base, "images"), 0755); err != nil {
		t.Fatalf("create images dir: %v", err)
	}
	if err := os.WriteFile(filepath.Join(base, "logo.png"), []byte("fake"), 0644); err != nil {
		t.Fatalf("create direct image: %v", err)
	}
	if err := os.WriteFile(filepath.Join(base, "images", "logo.png"), []byte("fake"), 0644); err != nil {
		t.Fatalf("create image: %v", err)
	}

	got, ok := resolvePageImagePath(pagesDir, base, "logo.png")
	if !ok {
		t.Fatal("expected direct image path to be accepted")
	}
	want := filepath.Join(base, "logo.png")
	if got != want {
		t.Fatalf("path = %q, want %q", got, want)
	}

	got, ok = resolvePageImagePath(pagesDir, base, "images/logo.png")
	if !ok {
		t.Fatal("expected nested image path to be accepted")
	}
	want = filepath.Join(base, "images", "logo.png")
	if got != want {
		t.Fatalf("path = %q, want %q", got, want)
	}

	if got, ok := resolvePageImagePath(pagesDir, base, "../guide.md"); ok {
		t.Fatalf("expected traversal to be rejected, got %q", got)
	}
}

func TestResolvePageImagePathRejectsSymlinkEscape(t *testing.T) {
	root := t.TempDir()
	pagesDir := filepath.Join(root, "pages")
	base := filepath.Join(pagesDir, "guide")
	outside := filepath.Join(root, "outside")

	if err := os.MkdirAll(base, 0755); err != nil {
		t.Fatalf("create page dir: %v", err)
	}
	if err := os.MkdirAll(outside, 0755); err != nil {
		t.Fatalf("create outside dir: %v", err)
	}
	if err := os.WriteFile(filepath.Join(outside, "secret.png"), []byte("secret"), 0644); err != nil {
		t.Fatalf("create outside file: %v", err)
	}
	if err := os.Symlink(outside, filepath.Join(base, "images")); err != nil {
		t.Skipf("symlink not supported: %v", err)
	}

	if got, ok := resolvePageImagePath(pagesDir, base, "images/secret.png"); ok {
		t.Fatalf("expected symlink escape to be rejected, got %q", got)
	}
}

func TestResolvePageMarkdownPath(t *testing.T) {
	h := NewPageHandler(t.TempDir(), nil)

	got, ok := h.resolvePageMarkdownPath("api-docs")
	if !ok {
		t.Fatal("expected valid slug to be accepted")
	}
	if filepath.Base(got) != "api-docs.md" {
		t.Fatalf("path = %q, want api-docs.md", got)
	}

	for _, slug := range []string{"", "../secret", "bad/slug", strings.Repeat("a", 65)} {
		if got, ok := h.resolvePageMarkdownPath(slug); ok {
			t.Fatalf("expected slug %q to be rejected, got %q", slug, got)
		}
	}
}

func TestAdminPageContentRoundTrip(t *testing.T) {
	gin.SetMode(gin.TestMode)
	h := NewPageHandler(t.TempDir(), nil)
	r := gin.New()
	r.GET("/admin/pages/:slug", h.GetAdminPageContent)
	r.PUT("/admin/pages/:slug", h.UpdateAdminPageContent)

	body := `{"content":"# API Docs\n\nHello"}`
	req := httptest.NewRequest(http.MethodPut, "/admin/pages/api-docs", strings.NewReader(body))
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)
	if w.Code != http.StatusOK {
		t.Fatalf("PUT status = %d, body = %s", w.Code, w.Body.String())
	}

	req = httptest.NewRequest(http.MethodGet, "/admin/pages/api-docs", nil)
	w = httptest.NewRecorder()
	r.ServeHTTP(w, req)
	if w.Code != http.StatusOK {
		t.Fatalf("GET status = %d, body = %s", w.Code, w.Body.String())
	}
	if !strings.Contains(w.Body.String(), "# API Docs") {
		t.Fatalf("GET body missing content: %s", w.Body.String())
	}
}

func TestPublicAPIDocs(t *testing.T) {
	gin.SetMode(gin.TestMode)
	dataDir := t.TempDir()
	pagesDir := filepath.Join(dataDir, "pages")
	if err := os.MkdirAll(pagesDir, 0755); err != nil {
		t.Fatalf("create pages dir: %v", err)
	}
	if err := os.WriteFile(filepath.Join(pagesDir, "api-docs.md"), []byte("# Public API Docs"), 0644); err != nil {
		t.Fatalf("create api docs: %v", err)
	}

	h := NewPageHandler(dataDir, nil)
	r := gin.New()
	r.GET("/public/pages/api-docs", h.GetPublicAPIDocs)

	req := httptest.NewRequest(http.MethodGet, "/public/pages/api-docs", nil)
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)
	if w.Code != http.StatusOK {
		t.Fatalf("GET status = %d, body = %s", w.Code, w.Body.String())
	}
	if !strings.Contains(w.Body.String(), "# Public API Docs") {
		t.Fatalf("GET body missing content: %s", w.Body.String())
	}
}
