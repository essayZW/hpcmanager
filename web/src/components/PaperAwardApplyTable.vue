<script setup lang="ts">
import { reactive, ref } from 'vue';
import { NodeApplyInfo } from '../api/node';
import { paginationGetNodeApplyInfo, nodeTypeToName } from '../service/node';
import dayjs from 'dayjs';
import { ProjectInfo } from '../api/project';
import { UserInfo } from '../api/user';
import { getUserInfoFromStorage, isAdmin, isTutor } from '../service/user';
import { getProjectInfoByID } from '../service/project';
import { zeroWithDefault, timeOrBlank } from '../utils/obj';

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

const userInfo = getUserInfoFromStorage();

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

defineExpose({
  refreshTableData,
});

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
      <el-table
        v-loading="tableData.loading"
        border
        table-layout="auto"
        :data="tableData.data"
      >
        <el-table-column
          label="ID"
          prop="id"
          sortable
          align="center"
        ></el-table-column>
        <el-table-column
          label="申请人学号"
          prop="createrUsername"
          align="center"
        ></el-table-column>
        <el-table-column
          label="申请人姓名"
          prop="createrName"
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
        <el-table-column label="申请时间" align="center">
          <template #default="props">
            {{
              dayjs(props.row.createTime * 1000).format('YYYY-MM-DD HH:mm:ss')
            }}
          </template>
        </el-table-column>
        <el-table-column label="审核状态" align="center">
          <template #default="props">
            <span v-if="props.row.checkStatus == -1">未审核</span>
            <span v-else-if="props.row.checkStatus == 1" class="green"
              >审核通过</span
            >
            <span v-else class="red">审核未通过</span>
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

.table-expand-area {
  padding-left: 12px;
  .info {
    padding-left: 16px;
    span {
      margin: 8px 8px;
    }
  }
}

.red {
  color: red;
}
.green {
  color: green;
}

.check-card {
  p {
    padding-left: 12px;
  }
  .box-title {
    padding-left: 0px;
    font-size: 16px;
  }
}
.operation-button-area {
  display: flex;
  justify-content: center;
}
.operation-button {
  display: inline-block;
}
.operation-button + span {
  display: none;
}
</style>
