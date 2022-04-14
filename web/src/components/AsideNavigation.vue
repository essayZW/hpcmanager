<script setup lang="ts">
import { ref } from 'vue';
import { getAvailableNavigation, NavigationItem } from '../service/navigation';
import { UserLevels } from '../service/user';
import { useRouter } from 'vue-router';
import { getUserInfoFromStorage } from '../service/user';

const router = useRouter();

let routerNum = ref(new Map<UserLevels, NavigationItem[]>());

const info = getUserInfoFromStorage();
if (info == null) {
  router.push({
    path: '/login',
  });
} else {
  routerNum.value = getAvailableNavigation(info.Levels);
}
</script>
<template>
  <el-menu class="aside-menu" router active-text-color="#000">
    <el-sub-menu v-if="routerNum.has(UserLevels.SuperAdmin)" index="1">
      <template #title>
        <span>超级管理员操作</span>
      </template>
      <el-menu-item
        v-for="item in routerNum.get(UserLevels.SuperAdmin)"
        :key="item.name"
        :route="item.to"
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
        :route="item.to"
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
        :route="item.to"
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
        :route="item.to"
        index="{{ item.to }}"
        >{{ item.name }}</el-menu-item
      >
    </el-sub-menu>
    <el-menu-item-group v-if="routerNum.has(UserLevels.Guest)">
      <el-menu-item
        v-for="item in routerNum.get(UserLevels.Guest)"
        :key="item.name"
        :route="item.to"
        index="{{ item.to }}"
        >{{ item.name }}</el-menu-item
      >
    </el-menu-item-group>
  </el-menu>
</template>
<style lang="less" scoped>
.aside-menu {
  border: none;
  .el-menu-item {
    min-width: 100px;
  }
}
</style>
