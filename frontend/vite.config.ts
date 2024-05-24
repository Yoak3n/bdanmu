import { defineConfig } from 'vite'
import vue from '@vitejs/plugin-vue'
import path from 'path'
// https://vitejs.dev/config/
export default defineConfig({
  plugins: [vue()],
  define: {
    __VUE_PROD_HYDRATION_MISMATCH_DETAILS__:"true",
  },
  server: {
    port: 5173
  },
  resolve:{
    alias:{
      "@": path.resolve(__dirname,"./src")
    }
  },
})
