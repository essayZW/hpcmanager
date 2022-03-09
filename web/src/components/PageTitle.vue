<script setup lang="ts">
import { defineProps } from 'vue';

const props = defineProps<{
  title: string;
  des: string;
}>();
</script>
<template>
  <el-row class="title-area">
    <el-col :span="24">
      <h1>{{ props.title }}</h1>
      <span>{{ props.des }}</span>
    </el-col>
  </el-row>
</template>
<style lang="less">
.title-area {
  h1 {
    display: inline-block;
    margin: 0px 8px;
    font-size: 28px;
    font-weight: 300;
  }
  span {
    vertical-align: text-bottom;
    color: gray;
    font-size: 16px;
    height: 16px;
  }
}
</style>
