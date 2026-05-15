<template>
  <div
    class="relative flex min-h-screen items-center justify-center overflow-hidden bg-neutral-50 p-4 text-neutral-950 transition-colors dark:bg-[#07100f] dark:text-white"
  >
    <!-- Background -->
    <div class="pointer-events-none absolute inset-0 overflow-hidden">
      <div
        class="absolute inset-0 bg-[radial-gradient(circle_at_22%_22%,rgba(23,23,23,0.08),transparent_28%),radial-gradient(circle_at_78%_18%,rgba(115,115,115,0.12),transparent_30%),linear-gradient(135deg,#ffffff_0%,#f5f5f5_52%,#e5e5e5_100%)] dark:bg-[radial-gradient(circle_at_20%_20%,rgba(45,212,191,0.22),transparent_28%),radial-gradient(circle_at_80%_10%,rgba(59,130,246,0.18),transparent_30%),linear-gradient(135deg,#07100f_0%,#0c1721_48%,#05070b_100%)]"
      ></div>
      <div
        class="absolute inset-0 bg-[linear-gradient(rgba(23,23,23,0.045)_1px,transparent_1px),linear-gradient(90deg,rgba(23,23,23,0.045)_1px,transparent_1px)] bg-[size:72px_72px] opacity-35 dark:bg-[linear-gradient(rgba(255,255,255,0.045)_1px,transparent_1px),linear-gradient(90deg,rgba(255,255,255,0.045)_1px,transparent_1px)] dark:opacity-40"
      ></div>
      <div class="absolute inset-0 bg-[radial-gradient(ellipse_at_24%_34%,rgba(23,23,23,0.045),transparent_34%),radial-gradient(ellipse_at_36%_48%,rgba(255,255,255,0.72),transparent_38%)] dark:hidden"></div>
      <div class="absolute left-0 right-0 top-0 h-px bg-gradient-to-r from-transparent via-neutral-500/70 to-transparent dark:via-cyan-300/80"></div>
    </div>

    <!-- Content Container -->
    <div class="relative z-10 w-full max-w-md py-12">
      <!-- Logo/Brand -->
      <div class="mb-8 text-center">
        <!-- Custom Logo or Default Logo -->
        <template v-if="settingsLoaded">
          <div
            class="mb-4 inline-flex h-16 w-16 items-center justify-center overflow-hidden rounded-2xl border border-neutral-950/10 bg-white/85 shadow-sm backdrop-blur dark:border-cyan-300/25 dark:bg-white/8 dark:shadow-[0_0_34px_rgba(45,212,191,0.28)]"
          >
            <img :src="siteLogo || '/aitd-logo.svg'" alt="AitdAPI" class="h-full w-full object-cover" />
          </div>
          <h1 class="mb-2 text-3xl font-black tracking-normal text-neutral-900 dark:text-cyan-100">
            {{ siteName }}
          </h1>
          <p class="text-sm text-neutral-600 dark:text-slate-300">
            {{ siteSubtitle }}
          </p>
        </template>
      </div>

      <!-- Card Container -->
      <div
        class="rounded-2xl border border-neutral-950/10 bg-white/70 p-8 shadow-sm backdrop-blur dark:border-white/10 dark:bg-white/[0.06]"
      >
        <slot />
      </div>

      <!-- Footer Links -->
      <div class="mt-6 text-center text-sm">
        <slot name="footer" />
      </div>

      <!-- Copyright -->
      <div class="mt-8 text-center text-xs text-neutral-500 dark:text-slate-400">
        &copy; {{ currentYear }} {{ siteName }}. All rights reserved.
      </div>
    </div>
  </div>
</template>

<script setup lang="ts">
import { computed, onMounted } from 'vue'
import { useAppStore } from '@/stores'
import { sanitizeUrl } from '@/utils/url'

const appStore = useAppStore()

const siteName = computed(() => appStore.siteName || 'AitdAPI')
const siteLogo = computed(() => sanitizeUrl(appStore.siteLogo || '', { allowRelative: true, allowDataUrl: true }))
const siteSubtitle = computed(() => appStore.cachedPublicSettings?.site_subtitle || 'Subscription to API Conversion Platform')
const settingsLoaded = computed(() => appStore.publicSettingsLoaded)

const currentYear = computed(() => new Date().getFullYear())

onMounted(() => {
  appStore.fetchPublicSettings()
})
</script>
