import { defineConfig } from 'vite'
import vue from '@vitejs/plugin-vue'
import path from 'path' // 确保已安装@types/node

export default defineConfig({
  plugins: [vue()],
  resolve: {
    alias: {
      '@': path.resolve(__dirname, './src'),
      // 添加其他需要的别名
    }
  }
})
