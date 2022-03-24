<script setup lang="ts">
import { reactive } from 'vue';
import { NodeApplyInfo, NodeDistribute } from '../../api/node';
import {
  paginationGetNodeDistributeInfo,
  getNodeApplyByID,
  nodeTypeToName,
  handlerNodeDistributeByID,
} from '../../service/node';
import dayjs from 'dayjs';
import { zeroWithDefault } from '../../utils/obj';

import PageTitle from '../PageTitle.vue';

const tableData = reactive<{
  data: NodeDistribute[];
  count: number;
  loading: boolean;
}>({
  data: [],
  count: 0,
  loading: false,
});

// 加载表格指定页的数据
const loadTableData = async (pageIndex: number, pageSize: number) => {
  tableData.loading = true;
  try {
    const resp = await paginationGetNodeDistributeInfo(pageIndex, pageSize);
    tableData.count = resp.Count;
    tableData.data = resp.Data;
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

const refreshTable = () => {
  loadTableData(paginationInfo.pageIndex, paginationInfo.pageSize);
};

refreshTable();

const handleCurrentChange = (pageIndex: number) => {
  paginationInfo.pageIndex = pageIndex;
  refreshTable();
};

const handleSizeChange = (pageSize: number) => {
  paginationInfo.pageSize = pageSize;
  refreshTable();
};

// 表格扩展字段的数据
const tableRowExpandData = reactive<{
  [id: number]: {
    applyInfo?: NodeApplyInfo;
  };
}>({});
const tableRowExpandHandler = async (row: NodeDistribute) => {
  if (!tableRowExpandData[row.id]) {
    tableRowExpandData[row.id] = {};
  }
  if (!tableRowExpandData[row.id].applyInfo) {
    try {
      const info = await getNodeApplyByID(row.applyID);
      tableRowExpandData[row.id].applyInfo = info;
    } catch (error) {
      ElMessage({
        type: 'error',
        message: `${error}`,
      });
    }
  }
};

const handlerFinishWorkOrder = async (id: number) => {
  if (!confirm(`确认标记ID为${id}的工单为已经处理?`))
    try {
      await handlerNodeDistributeByID(id);
      refreshTable();
    } catch (error) {
      ElMessage({
        type: 'error',
        message: `${error}`,
      });
    }
};
</script>
<template>
  <page-title
    title="节点分配工单管理"
    des="查看并处理机器节点分配的工单"
  ></page-title>
  <el-row justify="end" class="tool-button-row">
    <el-button type="primary" @click="refreshTable">
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
        @expand-change="tableRowExpandHandler"
      >
        <el-table-column label="ID" prop="id"></el-table-column>
        <el-table-column label="申请记录ID" prop="applyID"></el-table-column>
        <el-table-column label="是否处理">
          <template #default="props">
            <span v-if="!props.row.handlerFlag" class="red">未处理</span>
            <span v-else class="green">已处理</span>
          </template>
        </el-table-column>
        <el-table-column label="处理人ID">
          <template #default="props">
            {{ zeroWithDefault(props.row.handlerUserID, '无') }}
          </template>
        </el-table-column>
        <el-table-column label="处理人工号">
          <template #default="props">
            {{ zeroWithDefault(props.row.handlerUsername, '无') }}
          </template>
        </el-table-column>
        <el-table-column label="处理人姓名" prop="handlerUserName">
          <template #default="props">
            {{ zeroWithDefault(props.row.handlerName, '无') }}
          </template>
        </el-table-column>
        <el-table-column label="创建时间">
          <template #default="props">
            {{
              dayjs(props.row.createTime * 1000).format('YYYY-MM-DD HH:mm:ss')
            }}
          </template>
        </el-table-column>
        <el-table-column label="详情" type="expand">
          <template #default="props">
            <div class="table-expand-area">
              <p><strong>机器节点申请信息: </strong></p>
              <p class="info">
                <span
                  ><strong>申请人ID: </strong
                  >{{
                    tableRowExpandData[props.row.id].applyInfo?.createrID
                  }}</span
                >
                <span
                  ><strong>申请人姓名: </strong
                  >{{
                    tableRowExpandData[props.row.id].applyInfo?.createrName
                  }}</span
                >
                <span
                  ><strong>申请人学号: </strong
                  >{{
                    tableRowExpandData[props.row.id].applyInfo?.createrUsername
                  }}</span
                >
              </p>
              <p class="info">
                <span
                  ><strong>导师ID: </strong
                  >{{
                    tableRowExpandData[props.row.id].applyInfo?.tutorID
                  }}</span
                >
                <span
                  ><strong>导师姓名: </strong
                  >{{
                    tableRowExpandData[props.row.id].applyInfo?.tutorName
                  }}</span
                >
                <span
                  ><strong>导师工号: </strong
                  >{{
                    tableRowExpandData[props.row.id].applyInfo?.tutorUsername
                  }}</span
                >
              </p>
              <p class="info">
                <span
                  ><strong>机器节点类型: </strong
                  >{{
                  nodeTypeToName(
                    tableRowExpandData[props.row.id].applyInfo?.nodeType as string
                  )
                  }}</span
                >
                <span
                  ><strong>申请数量: </strong
                  >{{
                    tableRowExpandData[props.row.id].applyInfo?.nodeNum
                  }}</span
                >
              </p>
              <p class="info">
                <span
                  ><strong>申请时间: </strong>
                  {{
                  dayjs(
                    tableRowExpandData[props.row.id].applyInfo?.createTime as number *
                      1000
                  ).format('YYYY-MM-DD HH:mm:ss')
                  }}
                </span>
                <span
                  ><strong>审核通过时间: </strong>
                  {{
                  dayjs(
                    tableRowExpandData[props.row.id].applyInfo?.managerCheckTime as number *
                      1000
                  ).format('YYYY-MM-DD HH:mm:ss')
                  }}
                </span>
              </p>
              <p v-if="!props.row.handlerFlag"><strong>操作面板: </strong></p>
              <p v-if="!props.row.handlerFlag">
                <span>
                  <el-button
                    type="primary"
                    size="small"
                    @click="handlerFinishWorkOrder(props.row.id)"
                    >处理工单</el-button
                  >
                </span>
              </p>
            </div>
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
.red {
  color: red;
}
.green {
  color: green;
}
p.info {
  padding-left: 16px;
  span {
    margin: 8px 8px;
  }
}
.tool-button-row {
  margin: 16px 0px;
}
.table-expand-area {
  padding-left: 16px;
}
</style>
