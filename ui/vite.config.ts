import { fileURLToPath, URL } from 'node:url'
import { defineConfig, splitVendorChunkPlugin } from 'vite'
import vue from '@vitejs/plugin-vue'
import vueJsx from '@vitejs/plugin-vue-jsx'

// const appRootFolder = fileURLToPath(new URL('./src/apps/', import.meta.url))

// https://vitejs.dev/config/
export default defineConfig({
  plugins: [vue(), vueJsx(), splitVendorChunkPlugin()],
  resolve: {
    alias: {
      '@': fileURLToPath(new URL('./src', import.meta.url))
    }
  },
  build: {
    minify: 'esbuild',
    copyPublicDir: true,
    cssCodeSplit: true,
    cssMinify: true,
    modulePreload: true,
    rollupOptions: {
      input: {
        oauthLogin: fileURLToPath(new URL('./src/login/index.html', import.meta.url)),
        userInfo: fileURLToPath(new URL('./src/profile/index.html', import.meta.url))
      },
      output: {
        dir: 'dist',
        format: 'es',
        strict: true
      }
    }
  }
})
