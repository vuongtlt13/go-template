// https://nuxt.com/docs/api/configuration/nuxt-config
import { NOTIFICATION_DURATION } from "./utils/constants";

const sitePage = process.env.SITE_PAGE || "user";

export default defineNuxtConfig({
  app: {
    head: {
      title: process.env.APP_NAME,
      htmlAttrs: {
        lang: "en",
      },
      link: [{ rel: "icon", type: "image/x-icon", href: "/favicon.ico" }],
    },
  },
  ssr: false,
  compatibilityDate: "2024-11-01",
  devtools: { enabled: true },
  build: {
    transpile: ["vuetify"],
  },
  dir: {
    pages: `pages/${sitePage}`,
  },
  components: [{ path: "~/components/base", pathPrefix: false }, "~/components"],
  css: [
    "~/assets/css/main.css",
    "~/assets/scss/main.scss",
    "notivue/notification.css", // Only needed if using built-in <Notification />
    "notivue/animations.css", // Only needed if using default animations
    "notivue/notification-progress.css",
  ],
  modules: [
    "@nuxtjs/i18n",
    "vuetify-nuxt-module",
    "@nuxtjs/tailwindcss",
    "@nuxt/eslint",
    "@pinia/nuxt",
    "@nuxt/devtools",
    "notivue/nuxt",
    "@vee-validate/nuxt",
  ],
  tailwindcss: {},
  eslint: {
    checker: true,
  },
  vuetify: {
    moduleOptions: {
      /* module specific options */
    },
    vuetifyOptions: {
      /* vuetify options */
    },
  },
  i18n: {
    vueI18n: "./utils/i18n.ts",
    defaultLocale: "en",
    strategy: "no_prefix", // or 'prefix_except_default'
  },
  pinia: {
    storesDirs: ["./stores/**"],
  },
  runtimeConfig: {
    // Public keys that are exposed to the client
    public: {
      baseURL: `${process.env.NUXT_PUBLIC_API_BASE_URL || "http://localhost:8000"}`,
    },
  },
  notivue: {
    position: "top-right",
    limit: 1000,
    enqueue: true,
    avoidDuplicates: true,
    notifications: {
      global: {
        duration: NOTIFICATION_DURATION,
      },
    },
  },
});
