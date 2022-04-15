<script setup lang="ts">
import { reactive } from 'vue';
import { NodeWeekUsageBill } from '../api/fee';
import dayjs from 'dayjs';
import { paginationGetNodeWeekUsageBill } from '../service/fee';
import { timeSecondFormat } from '../utils/obj';

const tableData = reactive<{
  data: NodeWeekUsageBill[];
  count: number;
  timeRange: Date[];
}>({
  data: [],
  count: 0,
  timeRange: [dayjs(new Date().getTime()).add(-1, 'year').toDate(), new Date()],
});

// 加载表格某一页的数据
const loadTableData = async (
  pageIndex: number,
  pageSize: number,
  startTime: number,
  endTime: number
) => {
  try {
    const data = await paginationGetNodeWeekUsageBill(
      pageIndex,
      pageSize,
      startTime,
      endTime
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

const refreshTableDate = () => {
  loadTableData(
    paginationInfo.pageIndex,
    paginationInfo.pageSize,
    tableData.timeRange[0].getTime(),
    tableData.timeRange[1].getTime()
  );
};

refreshTableDate();

const handleCurrentChange = (pageIndex: number) => {
  paginationInfo.pageIndex = pageIndex;
  refreshTableDate();
};

const handleSizeChange = (pageSize: number) => {
  paginationInfo.pageSize = pageSize;
  refreshTableDate();
};
</script>
<template>
  <el-row justify="space-between" class="operator-tool-row">
    <div>
      <span>时间范围选择:</span>
      <el-date-picker
        v-model="tableData.timeRange"
        type="daterange"
        unlink-panels
        range-separator="To"
        start-placeholder="Start date"
        end-placeholder="End date"
      />
    </div>
    <el-button type="primary" @click="refreshTableDate">
      <el-icon class="el-icon--left">
        <i-ic-round-refresh />
      </el-icon>
      刷新
    </el-button>
  </el-row>
  <el-row justify="center">
    <el-col :span="24">
      <el-table table-layout="auto" border :data="tableData.data">
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
        <el-table-column label="时间范围" align="center">
          <template #default="props">
            {{ dayjs(props.row.startTime * 1000).format('YYYY-MM-DD') }}至{{
              dayjs(props.row.endTime * 1000).format('YYYY-MM-DD')
            }}
          </template>
        </el-table-column>
        <el-table-column label="CPU机时" align="center">
          <template #default="props">
            {{ timeSecondFormat(props.row.gwallTime) }}
          </template>
        </el-table-column>
        <el-table-column label="GPU机时" align="center">
          <template #default="props">
            {{ timeSecondFormat(props.row.gwallTime) }}
          </template>
        </el-table-column>
        <el-table-column label="应缴费用" align="center">
          <template #default="props"> {{ props.row.fee }}元 </template>
        </el-table-column>
        <el-table-column label="缴费状态" align="center">
          <template #default="props">
            <span v-if="props.row.payFlag">已缴费</span>
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
</style>
