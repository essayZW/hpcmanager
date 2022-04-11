<script setup lang="ts">
import { reactive } from 'vue';
import { NodeDistributeBill } from '../api/fee';
import { paginationGetNodeDistributeBill } from '../service/fee';

const tableData = reactive<{
  data: NodeDistributeBill[];
  count: number;
}>({
  data: [],
  count: 0,
});

// 加载表格数据
const loadTableData = async (pageIndex: number, pageSize: number) => {
  try {
    const data = await paginationGetNodeDistributeBill(pageIndex, pageSize);
    tableData.data = data.Data;
    tableData.count = data.Count;
  } catch (error) {
    ElMessage({
      message: `${error}`,
      type: 'error',
    });
  }
};

// 分页信息
const paginationInfo = reactive<{
  pageIndex: number;
  pageSize: number;
}>({
  pageIndex: 1,
  pageSize: 5,
});

const refreshTableData = () => {
  loadTableData(paginationInfo.pageIndex, paginationInfo.pageSize);
};

refreshTableData();

const handleSizeChange = (size: number) => {
  paginationInfo.pageSize = size;
  refreshTableData();
};

const handleCurrentChange = (index: number) => {
  paginationInfo.pageIndex = index;
  refreshTableData();
};
</script>
<template>
  <el-row justify="space-between" class="button-row">
    <div>
      <slot name="tool"></slot>
    </div>
    <el-button type="primary" @click="refreshTableData">
      <el-icon class="el-icon--left">
        <i-ic-round-refresh />
      </el-icon>
      刷新
    </el-button>
  </el-row>
  <el-row justify="center">
    <el-col :span="24">
      <el-table border table-layout="auto" :data="tableData.data">
        <el-table-column label="ID" align="center" prop="id"></el-table-column>
        <el-table-column
          label="工单ID"
          align="center"
          prop="applyID"
        ></el-table-column>
        <el-table-column label="应缴费用" align="center">
          <template #default="props"> {{ props.row.fee }}元 </template>
        </el-table-column>
        <el-table-column
          label="用户ID"
          align="center"
          prop="userID"
        ></el-table-column>
        <el-table-column
          label="用户帐号"
          align="center"
          prop="userUsername"
        ></el-table-column>
        <el-table-column
          label="用户姓名"
          align="center"
          prop="userName"
        ></el-table-column>
        <el-table-column label="操作" align="center">
          <template #default="props">
            <el-button v-if="!props.row.payFlag" type="primary">缴费</el-button>
            <span v-else class="green">已缴费</span>
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
.button-row {
  margin-top: 16px;
  margin-bottom: 16px;
}
.green {
  color: green;
}
</style>
