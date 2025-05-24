<script lang="ts" setup>
import { useRouter } from 'vue-router'

const router = useRouter()

const goRedis = () => {
  router.push('/redis/connection')
}
const goElasticSearch = () => {
  router.push('/elasticsearch/connection')
}
</script>
<template>
  <div class="mb-4">
    <el-button @click="goRedis">Redis</el-button>
    <el-button @click="goElasticSearch" type="primary" plain >ElasticSearch</el-button>
    <el-button type="success" plain>Docker</el-button>
    <el-button type="info" plain>MySQL</el-button>
    <el-button type="warning" plain>MongoDB</el-button>
  </div>
</template>

<style scoped>
.result {
  height: 20px;
  line-height: 20px;
  margin: 1.5rem auto;
}

.input-box .btn {
  width: 60px;
  height: 30px;
  line-height: 30px;
  border-radius: 3px;
  border: none;
  margin: 0 0 0 20px;
  padding: 0 8px;
  cursor: pointer;
}

.input-box .btn:hover {
  background-image: linear-gradient(to top, #cfd9df 0%, #e2ebf0 100%);
  color: #333333;
}

.input-box .input {
  border: none;
  border-radius: 3px;
  outline: none;
  height: 30px;
  line-height: 30px;
  padding: 0 10px;
  background-color: rgba(240, 240, 240, 1);
  -webkit-font-smoothing: antialiased;
}

.input-box .input:hover {
  border: none;
  background-color: rgba(255, 255, 255, 1);
}

.input-box .input:focus {
  border: none;
  background-color: rgba(255, 255, 255, 1);
}
</style>

insert into ucs_user values (56, '2yAmw9Ia', '15396284442', NULL, '9cbf8a4dcb8e30682b927f352d6559a0', '15396284442', NULL,NULL,NULL,NULL,1,1726301267335,1726301267335,1726301267335,23,114,NULL);

insert into ucs_user values (57, '2yAmw9Ib', '13794698214', NULL, '9cbf8a4dcb8e30682b927f352d6559a0', '13794698214', NULL,NULL,NULL,NULL,1,1726301267335,1726301267335,1726301267335,23,114,NULL);