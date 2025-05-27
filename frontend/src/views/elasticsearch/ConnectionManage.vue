<script lang="ts" setup>
import { useRouter } from 'vue-router'
import { ref } from 'vue'
import { ArrowLeftBold as Back } from '@element-plus/icons-vue'
import {GeTest} from "../../../wailsjs/go/elasticsearch/ESModule";
import {TestClient, SetConnection, GetNodes} from "../../../wailsjs/go/elasticsearch/ESModule";
import {
  Menu as IconMenu,
} from '@element-plus/icons-vue'


const router = useRouter()

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

type Node = {
  ip: string;
  name: string;
};

const nodes = ref<Array<Node>>([])

async function fetchConnections() {
  try {
    const data = await GeTest()
    connections.value = data.data
  } catch (error) {
    console.error("加载失败:", error)
  }
}

fetchConnections()

async function getConnectionAndNodes(id : number, host: string, username: string, password: string) {
  try {
    console.log("click....")
    SetConnection(id, host, username, password)
    TestClient(id, host, username, password)
    const data = await GetNodes(id)
    nodes.value = data.data
    console.log(nodes.value)
  } catch (error) {
    console.error("加载失败:", error)
  }
}

const goMain = () => {
  router.push('/')
}

const currentConnection = ref({
  id: 0,
  name: '',
  host: '',
  port: '',
  username: '',
  password: '',
})

const isCollapse = ref(true)
const handleOpen = (key: string, keyPath: string[]) => {
  console.log(key, keyPath)
}
const handleClose = (key: string, keyPath: string[]) => {
  console.log(key, keyPath)
}

// 添加选择处理方
const handleSelect = (index: string) => {  // 参数类型应为 string
  const selectedId = parseInt(index)
  const selectedConn = connections.value.find((conn: { id: number }) => conn.id === selectedId)
  if (selectedConn) {
    currentConnection.value = selectedConn
  }
}

// 新增折叠状态
const isCollapsed = ref(true)
const toggleCollapse = () => {
  isCollapsed.value = !isCollapsed.value
}
// 在脚本部分添加导入
import ConnectionDetail from './ConnectionDetail.vue'
import { validate } from '@babel/types';
</script>

<template>
  <div class="common-layout">
    <el-container>
      <el-aside width="200px">
        <el-icon 
      @click="goMain" 
      class="global-back-icon"
    >
      <Back />
    </el-icon>

    <el-radio-group v-model="isCollapse" style="margin-bottom: 20px">
      <el-radio-button :value="false">expand</el-radio-button>
      <el-radio-button :value="true">collapse</el-radio-button>
    </el-radio-group>
    <el-menu
      default-active="2"
      class="el-menu-vertical-demo"
      :collapse="isCollapse"
      @select="handleSelect"
      @open="handleOpen"
      @close="handleClose"
    >
      <el-menu-item 
        v-for="conn in connections" 
        :key="conn.id"
        :index="conn.id.toString()"
        @click="handleSelect(conn.id.toString())"
      >
        <el-icon><icon-menu /></el-icon>
        <template #title>
          <div class="menu-item-content" @click="getConnectionAndNodes(conn.id, conn.host, conn.username, conn.password)">
            <span class="connection-name">{{ conn.name }}</span>
            <span class="connection-host">{{ conn.host }}</span>
          </div>
        </template>
      </el-menu-item>
    </el-menu>
      </el-aside>
      <el-container>
        <el-header>ElasticSearch</el-header>
        <el-main class="connection-detail">
           <ConnectionDetail :connection="currentConnection" :nodes="{ Nodes: nodes }" />
        </el-main>
      </el-container>
    </el-container>
  </div>
</template>

<style scoped>
/* 修改折叠按钮定位方式 */
.collapse-control {
  position: relative;
  overflow: visible; /* 允许按钮溢出 */
}

.collapse-icon {
  position: fixed;
  left: 280px; /* 初始位置 */
  top: 50%;
  transform: translateY(-50%);
  transition: left 0.3s ease;
  z-index: 2000;
  background: #fff;
  box-shadow: 0 2px 12px rgba(0,0,0,0.1);
}
.connection-detail {
  transition: margin-left 0.3s ease;
}
.connection-list[width="0"] + .el-main {
  margin-left: 0;
}

.connection-layout {
  height: 90vh;
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
}.el-menu-vertical-demo:not(.el-menu--collapse) {
  width: 200px;
  min-height: 500px;
}
</style>
