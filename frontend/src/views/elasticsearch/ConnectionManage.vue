<script lang="ts" setup>
import { useRouter } from 'vue-router'
import { ref } from 'vue'
import { ArrowLeftBold as Back } from '@element-plus/icons-vue'
// 新增图标导入
import { Fold, Expand } from '@element-plus/icons-vue'
const router = useRouter()

// 模拟连接数据
const connections = [
  { id: 1, name: '生产集群', host: 'es-prod.example.com' },
  { id: 2, name: '测试集群', host: 'es-test.example.com' },
  { id: 4, name: '本地开发', host: 'localhost:9200' },
  { id: 5, name: '本地开发', host: 'localhost:9200' },
  { id: 6, name: '本地开发', host: 'localhost:9200' },
  { id: 7, name: '本地开发', host: 'localhost:9200' },
  { id: 8, name: '本地开发', host: 'localhost:9200' },
  { id: 9, name: '本地开发', host: 'localhost:9200' },
  { id: 10, name: '本地开发', host: 'localhost:9200' },
  { id: 11, name: '本地开发', host: 'localhost:9200' },
  { id: 12, name: '本地开发', host: 'localhost:9200' },
  { id: 13, name: '本地开发', host: 'localhost:9200' },
  { id: 14, name: '本地开发', host: 'localhost:9200' },
  { id: 15, name: '本地开发', host: 'localhost:9200' },
  { id: 16, name: '本地开发', host: 'localhost:9200' },
  { id: 17, name: '本地开发', host: 'localhost:9200' },
  { id: 18, name: '本地开发', host: 'localhost:9200' },
]

const goMain = () => {
  router.push('/')
}

const currentConnection = ref(connections[0])

// 添加选择处理方法
const handleSelect = (index: string) => {
  const selectedId = parseInt(index)
  const selectedConn = connections.find(conn => conn.id === selectedId)
  if (selectedConn) {
    currentConnection.value = selectedConn
  }
}

// 新增折叠状态
const isCollapsed = ref(false)
const toggleCollapse = () => {
  isCollapsed.value = !isCollapsed.value
}
</script>

<template>
  <div class="connection-layout">
    <!-- 将返回按钮移到容器内部 -->
    <el-icon 
      @click="goMain" 
      class="global-back-icon"
    >
      <Back />
    </el-icon>
    
    <el-container>
      <el-aside :width="isCollapsed ? '0' : '280px'" class="connection-list">
        <div class="header-section">
          <div class="collapse-control">
            <!-- 修改折叠按钮位置 -->
            <el-icon @click="toggleCollapse" class="collapse-icon">
              <component :is="isCollapsed ? Expand : Fold" />
            </el-icon>
            <h2 v-show="!isCollapsed">ElasticSearch 连接</h2>
          </div>
        </div>
        
        <!-- 添加 v-show 控制菜单显示 -->
        <el-menu 
          v-show="!isCollapsed"
          :default-active="currentConnection.id.toString()"
          @select="handleSelect"
        >
          <el-menu-item 
            v-for="conn in connections" 
            :key="conn.id"
            :index="conn.id.toString()"
          >
            <span class="connection-name">{{ conn.name }}</span>
            <span class="connection-host">{{ conn.host }}</span>
          </el-menu-item>
        </el-menu>
      </el-aside>

      <!-- 右侧内容区域 -->
      <el-main class="connection-detail">
        <div class="detail-header">
          <h3>{{ currentConnection.name }}</h3>
          <el-tag type="info">{{ currentConnection.host }}</el-tag>
        </div>
        
        <!-- 连接详情内容 -->
        <div class="detail-content">

          <h1>aaaaaaaaaaaaa</h1>
        </div>
      </el-main>
    </el-container>
  </div>
</template>

<style scoped>
.connection-layout {
  height: 100vh;
  position: relative;
}

.connection-list {
  height: calc(100vh - 20px); /* 留出顶部返回按钮空间 */
}

.el-menu {
  height: calc(100% - 60px); /* 减去头部高度 */
  overflow-y: auto;
  /* 保留原有flex设置 */
}

.global-back-icon {
  top: 12px; /* 微调返回按钮位置 */
}
</style>
