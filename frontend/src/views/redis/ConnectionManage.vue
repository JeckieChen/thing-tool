<script lang="ts" setup>
import { useRouter } from 'vue-router'
import { ArrowLeftBold as Back } from '@element-plus/icons-vue'
const router = useRouter()

const goMain = () => {
  router.push('/')
}

</script>

<template>
  <el-icon 
    @click="goMain" 
    style="cursor: pointer; margin-right: 10px;"
    class="back-icon"
  >
    <Back />
  </el-icon>
  <h1>Redis</h1>
</template>

<style scoped>
.back-icon:hover {
  color: #409eff;  /* Element Plus 主题色 */
  transform: scale(1.1);
  transition: all 0.3s ease;
}
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
