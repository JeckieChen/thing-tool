<script lang="ts" setup>
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

</script>

<template>
    <h1>{{ props.nodes }}</h1>
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
