<script lang="ts" setup>
interface Connection {
  id: number
  name: string
  host: string
}

const props = defineProps<{
  connection: Connection
}>()
</script>

<template>
    <h1>{{ props }}</h1>
</template>

<style scoped>
</style>
