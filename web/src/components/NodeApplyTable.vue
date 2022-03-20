<script setup lang="ts">
import { reactive } from 'vue';
import { NodeApplyInfo } from '../api/node';
import { paginationGetNodeApplyInfo } from '../service/node';
import dayjs from 'dayjs';

// 表格数据
const tableData = reactive<{
  count: number;
  data: NodeApplyInfo[];
  loading: boolean;
}>({
  count: 0,
  data: [],
  loading: false,
});

// 加载表格数据
const loadTableData = async (pageIndex: number, pageSize: number) => {
  tableData.loading = true;
  try {
    const paginationData = await paginationGetNodeApplyInfo(
      pageIndex,
      pageSize
    );
    tableData.data = paginationData.Data;
    tableData.count = paginationData.Count;
  } catch (error) {
    ElMessage({
      type: 'error',
      message: `${error}`,
    });
  }
  tableData.loading = false;
};

// 分页信息
const paginationInfo = reactive<{
  pageIndex: number;
  pageSize: number;
}>({
  pageIndex: 1,
  pageSize: 5,
});

// 刷新表格当前页面的信息
const refreshTableData = () => {
  loadTableData(paginationInfo.pageIndex, paginationInfo.pageSize);
};

refreshTableData();

const handleCurrentChange = (pageIndex: number) => {
  paginationInfo.pageIndex = pageIndex;
};

const handleSizeChange = (pageSize: number) => {
  paginationInfo.pageSize = pageSize;
};
</script>
<template>
  <el-row justify="end" class="button-row">
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
        border
        table-layout="auto"
        :data="tableData.data"
      >
        <el-table-column label="ID" prop="id"></el-table-column>
        <el-table-column
          label="申请人学号"
          prop="createrUsername"
        ></el-table-column>
        <el-table-column
          label="申请人姓名"
          prop="createrName"
        ></el-table-column>
        <el-table-column
          label="导师工号"
          prop="tutorUsername"
        ></el-table-column>
        <el-table-column label="导师姓名" prop="tutorName"></el-table-column>
        <el-table-column label="申请时间">
          <template #default="props">
            {{
              dayjs(props.row.createTime * 1000).format('YYYY-MM-DD HH:mm:ss')
            }}
          </template>
        </el-table-column>
        <el-table-column label="节点类型" prop="nodeType"></el-table-column>
        <el-table-column label="节点数目" prop="nodeNum"></el-table-column>
        <el-table-column label="详情" type="expand">
          <template #default="props"></template>
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
.pagination-row {
  margin: 16px 0px;
  .pagination-control-panel {
    margin: 0px auto;
    justify-content: center;
  }
}
</style>
