<script setup lang="ts">
import { HpcUsageTime } from '../api/node';
import { reactive } from 'vue';
import { paginationGetNodeUsageTime } from '../service/node';
import dayjs from 'dayjs';
import { timeSecondFormat } from '../utils/obj';

// 表格数据
const tableData = reactive<{
  data: HpcUsageTime[];
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
    const data = await paginationGetNodeUsageTime(
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
      <el-table border table-layout="auto" :data="tableData.data">
        <el-table-column
          label="用户学号"
          prop="username"
          align="center"
        ></el-table-column>
        <el-table-column
          label="用户姓名"
          prop="name"
          align="center"
        ></el-table-column>
        <el-table-column
          label="导师工号"
          prop="tutorUsername"
          align="center"
        ></el-table-column>
        <el-table-column
          label="导师姓名"
          prop="tutorName"
          align="center"
        ></el-table-column>
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
        <el-table-column label="开始时间" align="center">
          <template #default="props">
            {{ dayjs(props.row.startTime * 1000).format('YYYY-MM-DD') }}
          </template>
        </el-table-column>
        <el-table-column label="结束时间" align="center">
          <template #default="props">
            {{ dayjs(props.row.endTime * 1000).format('YYYY-MM-DD') }}
          </template>
        </el-table-column>
        <!-- TODO: 待后端数据库变更中引入作业ID以及作业名数据之后,将数据添加到详情信息中 -->
        <!--
        <el-table-column label="详情" type="expand" align="center">
          <template #default="props">
            {{ dayjs(props.row.endTime * 1000).format('YYYY-MM-DD') }}
          </template>
        </el-table-column>
        -->
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
