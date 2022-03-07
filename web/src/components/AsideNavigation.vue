<script setup lang="ts">
import { ref } from 'vue';
import { registryRouter, NavigationItem } from '../service/navigation';
import { UserLevels } from '../service/user';
import { useRouter } from 'vue-router';
import { watchEffect } from 'vue';

const router = useRouter();
console.log(router);

const props = defineProps<{
  levels: number[];
}>();

let routerNum = ref(new Map<UserLevels, NavigationItem[]>());

watchEffect(() => {
  routerNum.value = registryRouter(
    'Main',
    router,
    props.levels as UserLevels[]
  );
});
</script>
<template>
  <el-menu class="aside-menu" router>
    <el-sub-menu v-if="routerNum.has(UserLevels.SuperAdmin)" index="1">
      <template #title>
        <span>超级管理员操作</span>
      </template>
      <el-menu-item
        v-for="item in routerNum.get(UserLevels.SuperAdmin)"
        :key="item.name"
        router
        index="{{ item.to }}"
        >{{ item.name }}</el-menu-item
      >
    </el-sub-menu>
    <el-sub-menu v-if="routerNum.has(UserLevels.CommonAdmin)" index="2">
      <template #title>
        <span>普通管理员操作</span>
      </template>
      <el-menu-item
        v-for="item in routerNum.get(UserLevels.CommonAdmin)"
        :key="item.name"
        router
        index="{{ item.to }}"
        >{{ item.name }}</el-menu-item
      >
    </el-sub-menu>
    <el-sub-menu v-if="routerNum.has(UserLevels.Tutor)" index="3">
      <template #title>
        <span>导师操作</span>
      </template>
      <el-menu-item
        v-for="item in routerNum.get(UserLevels.Tutor)"
        :key="item.name"
        router
        index="{{ item.to }}"
        >{{ item.name }}</el-menu-item
      >
    </el-sub-menu>
    <el-sub-menu v-if="routerNum.has(UserLevels.Common)" index="4">
      <template #title>
        <span>学生操作</span>
      </template>
      <el-menu-item
        v-for="item in routerNum.get(UserLevels.Common)"
        :key="item.name"
        router
        index="{{ item.to }}"
        >{{ item.name }}</el-menu-item
      >
    </el-sub-menu>
    <el-sub-menu v-if="routerNum.has(UserLevels.Guest)" index="5">
      <template #title>
        <span>游客操作</span>
      </template>
      <el-menu-item
        v-for="item in routerNum.get(UserLevels.Guest)"
        :key="item.name"
        :route="item.to"
        index="{{ item.to }}"
        >{{ item.name }}</el-menu-item
      >
    </el-sub-menu>
  </el-menu>
</template>
<style lang="less">
.aside-menu {
  border: none;
}
</style>
