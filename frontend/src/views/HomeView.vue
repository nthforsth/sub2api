<template>
  <div v-if="homeContent" class="min-h-screen">
    <iframe
      v-if="isHomeContentUrl"
      :src="homeContent.trim()"
      class="h-screen w-full border-0"
      allowfullscreen
    ></iframe>
    <div v-else v-html="homeContent"></div>
  </div>

  <div
    v-else
    class="relative min-h-screen overflow-hidden bg-neutral-50 text-neutral-950 transition-colors dark:bg-[#07100f] dark:text-white"
  >
    <div class="pointer-events-none absolute inset-0">
      <div class="absolute inset-0 bg-[radial-gradient(circle_at_22%_22%,rgba(23,23,23,0.08),transparent_28%),radial-gradient(circle_at_78%_18%,rgba(115,115,115,0.12),transparent_30%),linear-gradient(135deg,#ffffff_0%,#f5f5f5_52%,#e5e5e5_100%)] dark:bg-[radial-gradient(circle_at_20%_20%,rgba(45,212,191,0.22),transparent_28%),radial-gradient(circle_at_80%_10%,rgba(59,130,246,0.18),transparent_30%),linear-gradient(135deg,#07100f_0%,#0c1721_48%,#05070b_100%)]"></div>
      <div class="absolute inset-0 bg-[linear-gradient(rgba(23,23,23,0.045)_1px,transparent_1px),linear-gradient(90deg,rgba(23,23,23,0.045)_1px,transparent_1px)] bg-[size:72px_72px] opacity-35 dark:bg-[linear-gradient(rgba(255,255,255,0.045)_1px,transparent_1px),linear-gradient(90deg,rgba(255,255,255,0.045)_1px,transparent_1px)] dark:opacity-40"></div>
      <div class="absolute inset-0 bg-[radial-gradient(ellipse_at_72%_34%,rgba(23,23,23,0.045),transparent_34%),radial-gradient(ellipse_at_64%_48%,rgba(255,255,255,0.72),transparent_38%)] dark:hidden"></div>
      <div class="absolute left-0 right-0 top-0 h-px bg-gradient-to-r from-transparent via-neutral-500/70 to-transparent dark:via-cyan-300/80"></div>
    </div>

    <header class="relative z-30 px-5 py-5 sm:px-8">
      <nav class="mx-auto flex max-w-7xl items-center justify-between">
        <router-link to="/home" class="flex items-center gap-3">
          <span class="flex h-10 w-10 items-center justify-center rounded-xl border border-neutral-950/10 bg-white/85 shadow-sm backdrop-blur dark:border-cyan-300/25 dark:bg-white/8 dark:shadow-[0_0_28px_rgba(45,212,191,0.24)]">
            <img :src="siteLogo || defaultLogo" alt="AitdAPI" class="h-full w-full object-cover" />
          </span>
          <span class="text-base font-semibold tracking-wide text-neutral-950 dark:text-white">{{ siteName }}</span>
        </router-link>

        <div class="flex items-center gap-2 sm:gap-3">
          <LocaleSwitcher />
          <a
            v-if="docUrl"
            :href="docUrl"
            target="_blank"
            rel="noopener noreferrer"
            class="hidden rounded-lg px-3 py-2 text-sm font-medium text-neutral-600 transition hover:bg-neutral-900/5 hover:text-neutral-950 dark:text-cyan-100/80 dark:hover:bg-white/10 dark:hover:text-white sm:inline-flex"
          >
            {{ t('home.docs') }}
          </a>
          <router-link
            to="/docs"
            class="hidden rounded-lg px-3 py-2 text-sm font-medium text-neutral-600 transition hover:bg-neutral-900/5 hover:text-neutral-950 dark:text-cyan-100/80 dark:hover:bg-white/10 dark:hover:text-white sm:inline-flex"
          >
            API 文档
          </router-link>
          <button
            @click="toggleTheme"
            class="rounded-lg border border-neutral-950/10 bg-white/75 p-2 text-neutral-700 shadow-sm transition hover:border-neutral-950/20 hover:bg-white dark:border-white/10 dark:bg-white/5 dark:text-cyan-100 dark:hover:border-cyan-300/40 dark:hover:bg-white/10"
            :title="isDark ? t('home.switchToLight') : t('home.switchToDark')"
          >
            <Icon v-if="isDark" name="sun" size="md" />
            <Icon v-else name="moon" size="md" />
          </button>
          <router-link
            :to="isAuthenticated ? dashboardPath : '/login'"
            class="inline-flex items-center gap-2 rounded-lg bg-neutral-950 px-4 py-2 text-sm font-semibold text-white shadow-sm transition hover:bg-neutral-800 dark:bg-cyan-300 dark:text-slate-950 dark:shadow-[0_0_28px_rgba(103,232,249,0.28)] dark:hover:bg-cyan-200"
          >
            {{ isAuthenticated ? t('home.goToDashboard') : t('home.login') }}
            <Icon name="arrowRight" size="sm" />
          </router-link>
        </div>
      </nav>
    </header>

    <main class="relative z-10 px-5 pb-16 pt-12 sm:px-8 sm:pt-20">
      <section class="mx-auto max-w-7xl">
        <div class="max-w-4xl">
          <div class="mb-6 inline-flex items-center gap-2 rounded-full border border-neutral-950/15 bg-white/50 px-3 py-1.5 text-xs font-semibold uppercase tracking-[0.22em] text-neutral-700 backdrop-blur dark:border-cyan-300/20 dark:bg-cyan-300/10 dark:text-cyan-100">
            AI Gateway / Routing / Billing
          </div>
          <h1 class="text-5xl font-black leading-[0.95] tracking-normal text-neutral-950 dark:text-white sm:text-6xl lg:text-7xl">
            {{ siteName }}
          </h1>
          <p class="mt-6 max-w-2xl text-lg leading-8 text-neutral-600 dark:text-slate-300 sm:text-xl">
            {{ t('home.aitdHero.description') }}
          </p>

          <div class="mt-9 flex flex-wrap gap-3">
            <router-link
              :to="isAuthenticated ? dashboardPath : '/login'"
              class="inline-flex items-center justify-center gap-2 rounded-xl bg-neutral-950 px-5 py-3 text-sm font-bold text-white shadow-lg shadow-neutral-950/10 transition hover:bg-neutral-800 dark:bg-white dark:text-slate-950 dark:hover:bg-cyan-100"
            >
              {{ isAuthenticated ? t('home.goToDashboard') : t('home.getStarted') }}
              <Icon name="arrowRight" size="sm" />
            </router-link>
            <router-link
              to="/docs"
              class="inline-flex items-center justify-center gap-2 rounded-xl border border-neutral-950/10 bg-white/70 px-5 py-3 text-sm font-bold text-neutral-800 shadow-sm backdrop-blur transition hover:border-neutral-950/20 hover:bg-white dark:border-white/10 dark:bg-white/[0.08] dark:text-cyan-100 dark:hover:border-cyan-300/40 dark:hover:bg-white/[0.12]"
            >
              API 文档
              <Icon name="book" size="sm" />
            </router-link>
          </div>

          <div class="mt-10 grid max-w-2xl grid-cols-3 gap-3">
            <div
              v-for="metric in metrics"
              :key="metric.labelKey"
              class="rounded-xl border border-neutral-950/10 bg-white/70 p-4 shadow-sm backdrop-blur dark:border-white/10 dark:bg-white/[0.06]"
            >
              <div class="text-2xl font-black text-neutral-800 dark:text-cyan-200">{{ metric.value }}</div>
              <div class="mt-1 text-xs font-medium text-neutral-500 dark:text-slate-400">{{ t(metric.labelKey) }}</div>
            </div>
          </div>
        </div>
      </section>

      <section class="mx-auto mt-16 grid max-w-7xl gap-4 md:grid-cols-3">
        <article
          v-for="feature in features"
          :key="feature.titleKey"
          class="rounded-2xl border border-neutral-950/10 bg-white/70 p-6 shadow-sm backdrop-blur transition hover:border-neutral-950/20 hover:bg-white dark:border-white/10 dark:bg-white/[0.055] dark:hover:border-cyan-300/35 dark:hover:bg-white/[0.08]"
        >
          <div class="mb-5 flex h-11 w-11 items-center justify-center rounded-xl bg-neutral-950/5 text-neutral-700 dark:bg-cyan-300/12 dark:text-cyan-200">
            <Icon :name="feature.icon" size="md" />
          </div>
          <h2 class="text-lg font-bold text-neutral-950 dark:text-white">{{ t(feature.titleKey) }}</h2>
          <p class="mt-3 text-sm leading-6 text-neutral-600 dark:text-slate-400">{{ t(feature.descriptionKey) }}</p>
        </article>
      </section>
    </main>

    <footer class="relative z-10 border-t border-neutral-950/10 px-5 py-6 text-center text-sm text-neutral-500 dark:border-white/10">
      &copy; {{ currentYear }} {{ siteName }}. {{ t('home.footer.allRightsReserved') }}
    </footer>
  </div>
</template>

<script setup lang="ts">
import { computed, onMounted, ref } from 'vue'
import { useI18n } from 'vue-i18n'
import { useAuthStore, useAppStore } from '@/stores'
import LocaleSwitcher from '@/components/common/LocaleSwitcher.vue'
import Icon from '@/components/icons/Icon.vue'

const { t } = useI18n()
const authStore = useAuthStore()
const appStore = useAppStore()

const defaultLogo = '/aitd-logo.svg'
const siteName = computed(() => appStore.cachedPublicSettings?.site_name || appStore.siteName || 'AitdAPI')
const siteLogo = computed(() => appStore.cachedPublicSettings?.site_logo || appStore.siteLogo || '')
const docUrl = computed(() => appStore.cachedPublicSettings?.doc_url || appStore.docUrl || '')
const homeContent = computed(() => appStore.cachedPublicSettings?.home_content || '')

const isHomeContentUrl = computed(() => {
  const content = homeContent.value.trim()
  return content.startsWith('http://') || content.startsWith('https://')
})

const isDark = ref(document.documentElement.classList.contains('dark'))
const isAuthenticated = computed(() => authStore.isAuthenticated)
const isAdmin = computed(() => authStore.isAdmin)
const dashboardPath = computed(() => (isAdmin.value ? '/admin/dashboard' : '/dashboard'))
const currentYear = computed(() => new Date().getFullYear())

const metrics = [
  { value: '1 Key', labelKey: 'home.aitdHero.metrics.unifiedAccess' },
  { value: 'Multi', labelKey: 'home.aitdHero.metrics.modelRouting' },
  { value: 'Live', labelKey: 'home.aitdHero.metrics.usageControl' }
]

const features = [
  {
    icon: 'server',
    titleKey: 'home.aitdHero.features.unifiedGateway.title',
    descriptionKey: 'home.aitdHero.features.unifiedGateway.description'
  },
  {
    icon: 'swap',
    titleKey: 'home.aitdHero.features.smartDispatch.title',
    descriptionKey: 'home.aitdHero.features.smartDispatch.description'
  },
  {
    icon: 'chart',
    titleKey: 'home.aitdHero.features.clearOperations.title',
    descriptionKey: 'home.aitdHero.features.clearOperations.description'
  }
] as const

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

onMounted(() => {
  initTheme()
  authStore.checkAuth()
  if (!appStore.publicSettingsLoaded) {
    appStore.fetchPublicSettings()
  }
})
</script>
