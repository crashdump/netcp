<template>
  <h1>netcp.io</h1>

  <p>
    <a href="https://netcp.io/login">Login</a>
    |
    <a href="https://netcp.io/files">Files</a>
  </p>
</template>

<script setup>
import {
  defineProps,
  //reactive
} from 'vue'

defineProps({
  msg: String
})

//const state = reactive({ count: 0 })
</script>

<style scoped>
a {
  color: #42b983;
}
</style>