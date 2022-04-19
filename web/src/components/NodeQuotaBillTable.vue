<script lang="ts" setup>
import { reactive } from 'vue';
import { NodeQuotaBill } from '../api/fee';
import { paginationGetNodeQuotaBill, payTypeToString } from '../service/fee';
import { zeroWithDefault } from '../utils/obj';
import dayjs from 'dayjs';

const tableData = reactive<{
  data: NodeQuotaBill[];
  count: number;
  loading: boolean;
}>({
  data: [],
  count: 0,
  loading: false,
});

// 加载表格某一页的数据
const loadTableData = async (pageIndex: number, pageSize: number) => {
  tableData.loading = true;
  try {
    const data = await paginationGetNodeQuotaBill(pageIndex, pageSize);
    tableData.data = data.Data;
    tableData.count = data.Count;
  } catch (error) {
    ElMessage({
      type: 'error',
      message: `${error}`,
    });
  }
  tableData.loading = false;
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
  <el-row justify="end" class="operator-tool-row">
    <el-button type="primary" @click="refreshTableData">
      <el-icon class="el-icon--left">
        <i-ic-round-refresh />
      </el-icon>
      刷新
    </el-button>
  </el-row>
  <el-row justify="center">
    <el-col :span="24">
      <el-table
        v-loading="tableData.loading"
        table-layout="auto"
        border
        :data="tableData.data"
      >
        <el-table-column
          label="用户学(工)号"
          align="center"
          prop="username"
        ></el-table-column>
        <el-table-column
          label="用户姓名"
          align="center"
          prop="name"
        ></el-table-column>
        <el-table-column
          label="用户组ID"
          align="center"
          prop="userGroupID"
        ></el-table-column>
        <el-table-column label="容量变化" align="center">
          <template #default="props">
            {{ props.row.oldSize }}TB
            <el-icon><i-ic-baseline-arrow-right-alt /></el-icon>
            {{ props.row.newSize }}TB
          </template>
        </el-table-column>
        <el-table-column label="结束日期变化" align="center">
          <template #default="props">
            {{ dayjs(props.row.oldEndTimeUnix * 1000).format('YYYY-MM-DD') }}
            <el-icon><i-ic-baseline-arrow-right-alt /></el-icon>
            {{ dayjs(props.row.newEndTimeUnix * 1000).format('YYYY-MM-DD') }}
          </template>
        </el-table-column>
        <el-table-column label="应缴费用" align="center">
          <template #default="props"> {{ props.row.fee }}元 </template>
        </el-table-column>
        <el-table-column label="缴费状态" align="center">
          <template #default="props">
            <span v-if="props.row.payFlag" class="green"
              >已缴费 {{ zeroWithDefault(props.row.payFee, 0) }}元</span
            >
            <span v-else class="red">未缴费</span>
          </template>
        </el-table-column>
        <el-table-column label="缴费方式" align="center">
          <template #default="props">
            <span v-if="props.row.payFlag">{{
              payTypeToString(props.row.payType)
            }}</span>
            <span v-else>未缴费</span>
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
<style lang="less" scoped>
.operator-tool-row {
  margin: 16px 0px;
}
.pagination-row {
  margin: 16px 0px;
  .pagination-control-panel {
    margin: 0px auto;
    justify-content: center;
  }
}
.green {
  color: green;
}
.red {
  color: red;
}
</style>
