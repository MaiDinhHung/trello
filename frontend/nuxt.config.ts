// https://nuxt.com/docs/api/configuration/nuxt-config
export default defineNuxtConfig({
  devtools: { enabled: true },

  modules: [
    '@nuxt/ui',
    '@nuxt/eslint',
    '@nuxt/icon'
  ],

  css: ['~/assets/css/main.css'],
  build: {
    transpile: ['@vuepic/vue-datepicker']
  },

  future: {
    compatibilityVersion: 4
  },

  compatibilityDate: '2024-11-27'
})