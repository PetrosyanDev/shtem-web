import { fileURLToPath, URL } from 'node:url'

import { defineConfig } from 'vite'
import vue from '@vitejs/plugin-vue'

export default defineConfig({
  build: {
    rollupOptions: {
      input: 'sources/includes/ts/main.ts',
      output: {
        entryFileNames: 'main.min.js',
        compact: true
      },
      external: [/^node:.*/]
    }
  },
  plugins: [
    vue({
      reactivityTransform: true
    })
  ],
  base: './',
  resolve: {
    alias: {
      '@': fileURLToPath(new URL('./sources/includes', import.meta.url))
    }
  }
})