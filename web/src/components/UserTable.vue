<script setup lang="ts">
import { reactive } from 'vue';
import { UserInfo } from '../api/user';
import { paginationGetUserInfo } from '../service/user';
import dayjs from 'dayjs';

// 表格数据
const tableData = reactive<{
  data: UserInfo[];
  count: number;
  loading: boolean;
}>({
  data: [],
  count: 0,
  loading: false,
});

// 分页信息
const paginationInfo = reactive<{
  pageIndex: number;
  pageSize: number;
}>({
  pageIndex: 1,
  pageSize: 5,
});

// 加载表格数据
const loadTableData = async (pageIndex: number, pageSize: number) => {
  tableData.loading = true;
  try {
    const resp = await paginationGetUserInfo(pageIndex, pageSize);
    tableData.count = resp.Count;
    tableData.data = resp.Data;
  } catch (error) {
    ElMessage({
      type: 'error',
      message: `${error}`,
    });
  } finally {
    tableData.loading = false;
  }
};

// 刷新表格数据
const refreshTable = () => {
  loadTableData(paginationInfo.pageIndex, paginationInfo.pageSize);
};

defineExpose({
  refreshTable,
});

refreshTable();

// 处理分页页码变化
const handleCurrentChange = (pageIndex: number) => {
  paginationInfo.pageIndex = pageIndex;
  refreshTable();
};

// 处理分页页大小变化
const handleSizeChange = (pageSize: number) => {
  paginationInfo.pageSize = pageSize;
  refreshTable();
};
</script>
<template>
  <el-row justify="end">
    <el-button type="primary" @click="refreshTable">
      <el-icon class="el-icon--left">
        <i-ic-round-refresh />
      </el-icon>
      刷新</el-button
    >
  </el-row>
  <el-row :span="24" justify="center">
    <el-col>
      <el-table v-loading="tableData.loading" :data="tableData.data">
        <el-table-column label="ID" prop="id"></el-table-column>
        <el-table-column label="姓名" prop="name"></el-table-column>
        <el-table-column label="学号" prop="username"></el-table-column>
        <el-table-column label="学院" prop="college"></el-table-column>
        <el-table-column label="电话" prop="tel"></el-table-column>
        <el-table-column label="邮箱" prop="email"></el-table-column>
        <el-table-column label="更多" type="expand">
          <template #default="props">
            {{
              dayjs(props.row.createTime * 1000).format('YYYY-MM-DD HH:mm:ss')
            }}
          </template>
        </el-table-column>
      </el-table>
    </el-col>
  </el-row>
  <el-row justify="center" class="pagination-row">
    <el-col :span="18">
      <el-pagination
        v-model:currentPage="paginationInfo.pageIndex"
        v-model:page-size="paginationInfo.pageSize"
        class="pagination-control-panel"
        :page-sizes="[5, 10, 25, 50]"
        background
        layout="total, sizes, prev, pager, next, jumper"
        :total="tableData.count"
        :hide-on-single-page="true"
        @size-change="handleSizeChange"
        @current-change="handleCurrentChange"
      >
      </el-pagination>
    </el-col>
  </el-row>
</template>
<style lang="less">
.pagination-row {
  margin: 16px 0px;
}
</style>
