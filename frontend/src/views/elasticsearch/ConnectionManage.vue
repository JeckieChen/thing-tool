<script lang="ts" setup>
import { useRouter } from 'vue-router'
import { ref, watch } from 'vue'
import { ArrowLeftBold as Back } from '@element-plus/icons-vue'
// 新增图标导入
import { Fold, Expand } from '@element-plus/icons-vue'
import {GeTest} from "../../../wailsjs/go/elasticsearch/ESModule";
// import type { Connection } from "../../../wailsjs/go/main/App"
const router = useRouter()

// let connections = ref<Array<{ Id: number, Name: string, Host: string }>>([])
// let test = ref<string>("")

type Result = {
  code: string
  data: Connection[]
}

type Connection = {
  id: number;
  name: string;
  host: string;
  port: string;
  username: string;
  password: string;
};
const connections = ref<Connection[]>([])

async function fetchConnections() {
  try {
    const data = await GeTest()
    connections.value = data.data
  } catch (error) {
    console.error("加载失败:", error)
  }
}

fetchConnections()

const goMain = () => {
  router.push('/')
}

const currentConnection = ref({
  id: 0,
  name: '',
  host: ''
})

// 添加选择处理方
const handleSelect = (index: string) => {
  const selectedId = parseInt(index)
  const selectedConn = connections.value.find((conn: { id: number }) => conn.id === selectedId)
  if (selectedConn) {
    currentConnection.value = selectedConn
  }
}

// 新增折叠状态
const isCollapsed = ref(false)
const toggleCollapse = () => {
  isCollapsed.value = !isCollapsed.value
}
// 在脚本部分添加导入
import ConnectionDetail from './ConnectionDetail.vue'
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
        <ConnectionDetail :connection="currentConnection" />
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
