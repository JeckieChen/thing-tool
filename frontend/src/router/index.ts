import { createRouter, createWebHistory } from 'vue-router'

const routes = [
  {
    path: '/',
    component: () => import('@/components/Main.vue')  // 确保根路径指向Main组件
  },
  {
    path: '/redis/connection',
    component: () => import('@/views/redis/ConnectionManage.vue')
  },
  {
    path: '/elasticsearch/connection',
    component: () => import('@/views/elasticsearch/ConnectionManage.vue')
  }
]

const router = createRouter({
  history: createWebHistory(),
  routes
})

export default router