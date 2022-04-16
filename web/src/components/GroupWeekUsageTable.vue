<script lang="ts" setup>
import { reactive } from 'vue';
import { paginationGetGroupNodeUsageBill } from '../service/fee';
import { NodeWeekUsageBillForGroup } from '../api/fee';

const tableData = reactive<{
  data: NodeWeekUsageBillForGroup[];
  count: number;
}>({
  data: [],
  count: 0,
});

// 加载表格某一页的数据
const loadTableData = async (pageIndex: number, pageSize: number) => {
  try {
    const data = await paginationGetGroupNodeUsageBill(
      pageIndex,
      pageSize,
      false
    );
    tableData.data = data.Data;
    tableData.count = data.Count;
  } catch (error) {
    ElMessage({
      type: 'error',
      message: `${error}`,
    });
  }
};

const paginationInfo = reactive<{
  pageIndex: number;
  pageSize: number;
}>({
  pageIndex: 1,
  pageSize: 10,
});

const refreshTableData = () => {
  loadTableData(paginationInfo.pageIndex, paginationInfo.pageSize);
};

refreshTableData();

const handleCurrentChange = (pageIndex: number) => {
  paginationInfo.pageIndex = pageIndex;
  refreshTableData();
};

const handleSizeChange = (pageSize: number) => {
  paginationInfo.pageSize = pageSize;
  refreshTableData();
};
</script>
<template>
  <el-row justify="space-between" class="refresh-button-row">
    <el-button type="primary" @click="refreshTableData">
      <el-icon class="el-icon--left">
        <i-ic-round-refresh />
      </el-icon>
      刷新</el-button
    >
  </el-row>
  <el-row justify="center">
    <el-col :span="24">
      <el-table border table-layout="auto">
        <el-table-column
          label="组ID"
          prop="userGroupID"
          align="center"
        ></el-table-column>
        <el-table-column label="CPU机时" align="center"></el-table-column>
        <el-table-column label="GPU机时" align="center"></el-table-column>
        <el-table-column label="应缴费用" align="center"></el-table-column>
        <el-table-column label="未缴费用" align="center"></el-table-column>
        <el-table-column
          label="详情"
          type="expand"
          align="center"
        ></el-table-column>
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
<style lang="less" scoped>
.refresh-button-row {
  margin: 16px 0px;
}
.pagination-row {
  margin: 16px 0px;
  .pagination-control-panel {
    margin: 0px auto;
    justify-content: center;
  }
}
</style>
