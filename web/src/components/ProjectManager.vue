<script setup lang="ts">
import { reactive } from 'vue';
import { ProjectInfo } from '../api/project';
import { paginationGetProjectInfo } from '../service/project';

import PageTitle from './PageTitle.vue';

// 项目表格的数据
const tableData = reactive<{
  data: ProjectInfo[];
  count: number;
}>({
  data: [],
  count: 0,
});

// 表格当前的分页信息
const paginationInfo = reactive<{
  pageIndex: number;
  pageSize: number;
}>({
  pageIndex: 1,
  pageSize: 5,
});

// 分页加载表格的数据
const loadTableData = async (pageIndex: number, pageSize: number) => {
  try {
    const data = await paginationGetProjectInfo(pageIndex, pageSize);
    tableData.data = data.Data;
    tableData.count = data.Count;
  } catch (error) {
    ElMessage({
      type: 'error',
      message: `${error}`,
    });
  }
};

// 刷新表格的当前页的信息
const refreshTableData = () => {
  loadTableData(paginationInfo.pageIndex, paginationInfo.pageSize);
};

refreshTableData();
</script>
<template>
  <page-title title="项目管理" des="项目管理、查看页面"></page-title>
  <el-row justify="space-between" class="button-area-row">
    <el-button type="primary">新建项目</el-button>
    <el-button type="primary" @click="refreshTableData">
      <el-icon class="el-icon--left">
        <i-ic-round-refresh />
      </el-icon>
      刷新
    </el-button>
  </el-row>
  <el-row justify="center">
    <el-col :span="24">
      <el-table :data="tableData.data">
        <el-table-column label="ID"></el-table-column>
        <el-table-column label="名称"></el-table-column>
        <el-table-column label="来源"></el-table-column>
        <el-table-column label="编号"></el-table-column>
        <el-table-column label="经费"></el-table-column>
        <el-table-column label="创建者姓名"></el-table-column>
        <el-table-column label="创建者学(工)号"></el-table-column>
      </el-table>
    </el-col>
  </el-row>
</template>
<style lang="less" scoped>
.button-area-row {
  margin-top: 16px;
  margin-bottom: 8px;
}
</style>
