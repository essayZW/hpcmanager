<script setup lang="ts">
import { ref } from 'vue';

// 获取当前窗口的高度
let WindowHeight = ref(window.innerHeight + 'px');
window.addEventListener('resize', () => {
  let currentHeight = window.innerHeight;
  if (currentHeight < 600) {
    currentHeight = 600;
  }
  WindowHeight.value = currentHeight + 'px';
});
</script>
<template>
  <router-view id="main-view"></router-view>
</template>

<style scoped lang="less">
#main-view {
  min-height: v-bind(WindowHeight);
  width: 100%;
}
</style>
