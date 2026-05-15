<template>
  <div
    class="min-h-screen bg-neutral-50 text-neutral-950 transition-colors dark:bg-[#07100f] dark:text-white"
  >
    <div class="pointer-events-none fixed inset-0">
      <div
        class="absolute inset-0 bg-[radial-gradient(circle_at_22%_22%,rgba(23,23,23,0.08),transparent_28%),radial-gradient(circle_at_78%_18%,rgba(115,115,115,0.12),transparent_30%),linear-gradient(135deg,#ffffff_0%,#f5f5f5_52%,#e5e5e5_100%)] dark:bg-[radial-gradient(circle_at_20%_20%,rgba(45,212,191,0.22),transparent_28%),radial-gradient(circle_at_80%_10%,rgba(59,130,246,0.18),transparent_30%),linear-gradient(135deg,#07100f_0%,#0c1721_48%,#05070b_100%)]"
      ></div>
      <div
        class="absolute inset-0 bg-[linear-gradient(rgba(23,23,23,0.045)_1px,transparent_1px),linear-gradient(90deg,rgba(23,23,23,0.045)_1px,transparent_1px)] bg-[size:72px_72px] opacity-35 dark:bg-[linear-gradient(rgba(255,255,255,0.045)_1px,transparent_1px),linear-gradient(90deg,rgba(255,255,255,0.045)_1px,transparent_1px)] dark:opacity-40"
      ></div>
    </div>

    <header class="relative z-10 border-b border-neutral-950/10 px-5 py-5 backdrop-blur dark:border-white/10 sm:px-8">
      <nav class="mx-auto flex max-w-5xl items-center justify-between">
        <router-link to="/home" class="flex items-center gap-3">
          <span
            class="flex h-10 w-10 items-center justify-center rounded-xl border border-neutral-950/10 bg-white/85 shadow-sm dark:border-cyan-300/25 dark:bg-white/8"
          >
            <img :src="siteLogo || defaultLogo" alt="AitdAPI" class="h-full w-full object-cover" />
          </span>
          <span class="text-base font-semibold text-neutral-950 dark:text-white">{{ siteName }}</span>
        </router-link>
        <div class="flex items-center gap-2">
          <button
            type="button"
            class="rounded-lg border border-neutral-950/10 bg-white/75 p-2 text-neutral-700 shadow-sm transition hover:border-neutral-950/20 hover:bg-white dark:border-white/10 dark:bg-white/5 dark:text-cyan-100 dark:hover:border-cyan-300/40 dark:hover:bg-white/10"
            :title="isDark ? t('home.switchToLight') : t('home.switchToDark')"
            @click="toggleTheme"
          >
            <Icon v-if="isDark" name="sun" size="md" />
            <Icon v-else name="moon" size="md" />
          </button>
          <router-link
            to="/home"
            class="rounded-lg px-3 py-2 text-sm font-medium text-neutral-600 transition hover:bg-neutral-900/5 hover:text-neutral-950 dark:text-cyan-100/80 dark:hover:bg-white/10 dark:hover:text-white"
          >
            {{ localText("返回首页", "Home") }}
          </router-link>
        </div>
      </nav>
    </header>

    <main class="relative z-10 mx-auto max-w-5xl px-5 py-10 sm:px-8">
      <div class="mb-8">
        <p class="text-sm font-semibold uppercase tracking-[0.22em] text-neutral-500 dark:text-cyan-200/80">
          API Documentation
        </p>
        <h1 class="mt-3 text-4xl font-black tracking-normal text-neutral-950 dark:text-white">
          AitdAPI
        </h1>
      </div>

      <div
        v-if="loading"
        class="flex min-h-[320px] items-center justify-center rounded-2xl border border-neutral-950/10 bg-white/70 dark:border-white/10 dark:bg-white/[0.06]"
      >
        <div class="h-8 w-8 animate-spin rounded-full border-2 border-cyan-400 border-t-transparent"></div>
      </div>

      <div
        v-else-if="error"
        class="rounded-2xl border border-red-200 bg-red-50 p-6 text-sm text-red-700 dark:border-red-900/40 dark:bg-red-950/30 dark:text-red-200"
      >
        {{ error }}
      </div>

      <article
        v-else
        class="markdown-public-doc rounded-2xl border border-neutral-950/10 bg-white/80 p-6 shadow-sm backdrop-blur dark:border-white/10 dark:bg-white/[0.065] sm:p-9"
        v-html="renderedHtml"
      ></article>
    </main>
  </div>
</template>

<script setup lang="ts">
import { computed, onMounted, ref } from 'vue'
import { useI18n } from 'vue-i18n'
import { marked } from 'marked'
import DOMPurify from 'dompurify'
import { useAppStore } from '@/stores'
import Icon from '@/components/icons/Icon.vue'

const { locale, t } = useI18n()
const appStore = useAppStore()

const defaultLogo = '/aitd-logo.svg'
const siteName = computed(() => appStore.cachedPublicSettings?.site_name || appStore.siteName || 'AitdAPI')
const siteLogo = computed(() => appStore.cachedPublicSettings?.site_logo || appStore.siteLogo || '')
const loading = ref(true)
const error = ref('')
const renderedHtml = ref('')
const isDark = ref(document.documentElement.classList.contains('dark'))

function localText(zh: string, en: string): string {
  return locale.value.startsWith('zh') ? zh : en
}

function toggleTheme() {
  isDark.value = !isDark.value
  document.documentElement.classList.toggle('dark', isDark.value)
  localStorage.setItem('theme', isDark.value ? 'dark' : 'light')
}

function initTheme() {
  const savedTheme = localStorage.getItem('theme')
  if (savedTheme !== 'light') {
    isDark.value = true
    document.documentElement.classList.add('dark')
  }
}

async function loadDocs() {
  loading.value = true
  error.value = ''
  try {
    const resp = await fetch('/api/v1/public/pages/api-docs')
    if (!resp.ok) {
      error.value = localText('API 文档暂未配置。', 'API docs are not configured yet.')
      return
    }
    const raw = await resp.text()
    renderedHtml.value = DOMPurify.sanitize(marked.parse(raw) as string)
  } catch (_err) {
    error.value = localText('加载 API 文档失败，请稍后重试。', 'Failed to load API docs. Please try again later.')
  } finally {
    loading.value = false
  }
}

onMounted(() => {
  initTheme()
  if (!appStore.publicSettingsLoaded) {
    appStore.fetchPublicSettings()
  }
  loadDocs()
})
</script>

<style scoped>
.markdown-public-doc :deep(h1),
.markdown-public-doc :deep(h2),
.markdown-public-doc :deep(h3) {
  color: inherit;
  font-weight: 800;
  line-height: 1.2;
}

.markdown-public-doc :deep(h1) {
  font-size: 2rem;
  margin-bottom: 1.25rem;
}

.markdown-public-doc :deep(h2) {
  font-size: 1.35rem;
  margin-top: 2rem;
  margin-bottom: 0.8rem;
}

.markdown-public-doc :deep(h3) {
  font-size: 1.05rem;
  margin-top: 1.5rem;
  margin-bottom: 0.6rem;
}

.markdown-public-doc :deep(p),
.markdown-public-doc :deep(li) {
  color: rgb(82 82 82);
  line-height: 1.75;
}

:global(.dark) .markdown-public-doc :deep(p),
:global(.dark) .markdown-public-doc :deep(li) {
  color: #ffffff !important;
}

.markdown-public-doc :deep(pre) {
  margin: 1rem 0;
  overflow-x: auto;
  border-radius: 0.85rem;
  background: rgb(15 23 42);
  padding: 1rem;
  color: rgb(226 232 240);
}

.markdown-public-doc :deep(code) {
  border-radius: 0.35rem;
  background: rgba(15, 23, 42, 0.08);
  padding: 0.15rem 0.35rem;
  font-size: 0.9em;
}

:global(.dark) .markdown-public-doc :deep(code) {
  background: rgba(255, 255, 255, 0.12);
  color: rgb(248 250 252);
}

.markdown-public-doc :deep(pre code) {
  background: transparent;
  padding: 0;
}

.markdown-public-doc :deep(table) {
  width: 100%;
  border-collapse: collapse;
  margin: 1rem 0;
  overflow: hidden;
}

.markdown-public-doc :deep(th),
.markdown-public-doc :deep(td) {
  border: 1px solid rgba(115, 115, 115, 0.25);
  padding: 0.65rem 0.8rem;
  text-align: left;
}

:global(.dark) .markdown-public-doc :deep(th),
:global(.dark) .markdown-public-doc :deep(td) {
  border-color: rgba(226, 232, 240, 0.22);
  color: #ffffff !important;
}

.markdown-public-doc :deep(a) {
  color: rgb(8 145 178);
  font-weight: 600;
}

:global(.dark) .markdown-public-doc :deep(a) {
  color: rgb(103 232 249);
}
</style>

<style>
.dark .markdown-public-doc,
.dark .markdown-public-doc p,
.dark .markdown-public-doc li,
.dark .markdown-public-doc ol,
.dark .markdown-public-doc ul,
.dark .markdown-public-doc table,
.dark .markdown-public-doc th,
.dark .markdown-public-doc td,
.dark .markdown-public-doc blockquote,
.dark .markdown-public-doc strong,
.dark .markdown-public-doc em {
  color: #ffffff !important;
}
</style>
