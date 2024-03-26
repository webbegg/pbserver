// https://nuxt.com/docs/api/configuration/nuxt-config
export default defineNuxtConfig({
  devtools: { enabled: true },
  ssr: false,
  app: {
    baseURL: "/central/",
    buildAssetsDir: "/assets/"
  },
  nitro: {
    /* dev server proxy */
    devProxy: {
      "/api": {
        target: "http://localhost:8090/api",
        changeOrigin: true,
        prependPath: true,
      },
    },
    output: {
      publicDir: "dist",
    },
  },
})
