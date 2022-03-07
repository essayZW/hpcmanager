<script setup lang="ts">
import { registryRouter } from '../service/navigation';
import { UserLevels } from '../service/user';
import { useRouter } from 'vue-router';

const router = useRouter();

const props = defineProps<{
  levels: number[];
}>();

const routerNum = registryRouter('Main', router, props.levels as UserLevels[]);
</script>
<template>
  <el-menu class="aside-menu">
    <el-sub-menu v-if="routerNum.has(UserLevels.SuperAdmin)" index="1">
      <template #title>
        <span>超级管理员操作台</span>
      </template>
      <el-menu-item>xss</el-menu-item>
    </el-sub-menu>
    <el-sub-menu v-if="routerNum.has(UserLevels.CommonAdmin)" index="2">
      <template #title>
        <span>普通管理员操作台</span>
      </template>
      <el-menu-item>xss</el-menu-item>
    </el-sub-menu>
    <el-sub-menu v-if="routerNum.has(UserLevels.Tutor)" index="3">
      <template #title>
        <span>导师操作台</span>
      </template>
      <el-menu-item>xss</el-menu-item>
    </el-sub-menu>
    <el-sub-menu v-if="routerNum.has(UserLevels.Common)" index="4">
      <template #title>
        <span>学生操作台</span>
      </template>
      <el-menu-item>xss</el-menu-item>
    </el-sub-menu>
    <el-sub-menu v-if="routerNum.has(UserLevels.Guest)" index="5">
      <template #title>
        <span>游客操作台</span>
      </template>
      <el-menu-item>xss</el-menu-item>
    </el-sub-menu>
  </el-menu>
</template>
<style lang="less">
.aside-menu {
  border: none;
}
</style>
