import { defineConfig } from 'vite'
import react from '@vitejs/plugin-react'

// https://vite.dev/config/
export default defineConfig({
  plugins: [react()],
  base: '/web/',
  server: {
    proxy: {
      // Proxy requests starting with '/api' to http://localhost:9001
      '/api': {
        target: 'http://localhost:9001',
        changeOrigin: true, // Recommended for most proxy setups
        // Optional: Rewrite the path if needed (e.g., remove '/api' from the forwarded request)
        rewrite: (path) => path.replace(/^\/api/, '')
      },
    },
  },
})
