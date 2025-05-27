<script lang="ts" setup>
import { ref } from 'vue'
interface Connection {
  id: number
  name: string
  host: string
  port: string
  username: string
  password: string
}

interface Node {
    ip: string;
    name: string;
}

interface Nodes {
  Nodes: Array<Node>
}

type ConnectionOptions = {
  Name: string
}

const props = defineProps<{
  connection: Connection
  nodes: Nodes
}>()


const activeIndex = ref('1')
const activeIndex2 = ref('1')
const handleSelect = (key: string, keyPath: string[]) => {
  activeIndex.value = key
}

</script>

<template>
    <el-menu
      :default-active="activeIndex"
      class="el-menu-demo"
      mode="horizontal"
      @select="handleSelect"
    >
      <el-menu-item index="1">节点</el-menu-item>
      <el-menu-item index="3">索引</el-menu-item>
      <el-menu-item index="4">Orders</el-menu-item>
      <el-sub-menu index="2">
        <template #title>Workspace</template>
        <el-menu-item index="2-1">item one</el-menu-item>
        <el-menu-item index="2-2">item two</el-menu-item>
        <el-menu-item index="2-3">item three</el-menu-item>
        <el-sub-menu index="2-4">
          <template #title>item four</template>
          <el-menu-item index="2-4-1">item one</el-menu-item>
          <el-menu-item index="2-4-2">item two</el-menu-item>
          <el-menu-item index="2-4-3">item three</el-menu-item>
        </el-sub-menu>
      </el-sub-menu>
    </el-menu>

    <div v-if="activeIndex === '1'">
      <el-row :gutter="20">
              <el-col 
            v-for="(node, index) in props.nodes?.Nodes" 
            :key="index" 
            :span="8"
          >
            <div class="grid-content ep-bg-purple">
              <el-card style="max-width: 480px">
                <template #header>
                  <div class="card-header">
                    <span>{{ node.name }}</span>
                  </div>
                </template>
                <div class="node-info">
                  <p>IP 地址：{{ node.ip }}</p>
                  <p>节点名称：{{ node.name }}</p>
                </div>
              </el-card>
            </div>
          </el-col>
        </el-row>
    </div>
    <div v-if="activeIndex === '3'">
      <!-- 任务管理内容 -->
      <el-empty description="任务功能开发中" />
    </div>
</template>

<style scoped>
.el-row {
  margin-bottom: 20px;
}
.el-row:last-child {
  margin-bottom: 0;
}
.el-col {
  border-radius: 4px;
}

.grid-content {
  border-radius: 4px;
  min-height: 36px;
}
</style>
