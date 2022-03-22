<script setup lang="ts">
import { reactive } from 'vue';
import { NodeApplyInfo } from '../api/node';
import { paginationGetNodeApplyInfo, nodeTypeToName } from '../service/node';
import dayjs from 'dayjs';
import { ProjectInfo } from '../api/project';
import { UserInfo } from '../api/user';
import { getUserInfoById } from '../service/user';
import { getProjectInfoByID } from '../service/project';
import { zeroWithDefault } from '../utils/obj';

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

defineExpose({
  refreshTableData,
});

const handleCurrentChange = (pageIndex: number) => {
  paginationInfo.pageIndex = pageIndex;
};

const handleSizeChange = (pageSize: number) => {
  paginationInfo.pageSize = pageSize;
};

// 表格展开行中的数据
const tableExpandRowInfo = reactive<{
  [id: number]: {
    projectInfo?: ProjectInfo;
    applierInfo?: UserInfo;
    tutorInfo?: UserInfo;
    loading: boolean;
  };
}>({});

// 表格行展开时候的回调事件
const handlerTableExpand = async (row: NodeApplyInfo) => {
  if (!tableExpandRowInfo[row.id]) {
    tableExpandRowInfo[row.id] = {
      loading: true,
    };
  }
  tableExpandRowInfo[row.id].loading = true;
  if (!tableExpandRowInfo[row.id].applierInfo) {
    const userInfo = await getUserInfoById(row.createrID);
    if (typeof userInfo == 'string') {
      ElMessage({
        type: 'error',
        message: userInfo as string,
      });
    } else {
      tableExpandRowInfo[row.id].applierInfo = userInfo as UserInfo;
    }
  }

  if (!tableExpandRowInfo[row.id].tutorInfo) {
    const userInfo = await getUserInfoById(row.tutorID);
    if (typeof userInfo == 'string') {
      ElMessage({
        type: 'error',
        message: userInfo as string,
      });
    } else {
      tableExpandRowInfo[row.id].tutorInfo = userInfo as UserInfo;
    }
  }

  if (!tableExpandRowInfo[row.id].projectInfo) {
    try {
      const projectInfo = await getProjectInfoByID(row.projectID);
      tableExpandRowInfo[row.id].projectInfo = projectInfo;
    } catch (error) {
      ElMessage({
        type: 'error',
        message: `${error}`,
      });
    }
  }
  tableExpandRowInfo[row.id].loading = false;
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
        @expand-change="handlerTableExpand"
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
        <el-table-column label="节点类型">
          <template #default="props">
            {{ nodeTypeToName(props.row.nodeType) }}
          </template>
        </el-table-column>
        <el-table-column label="节点数目" prop="nodeNum"></el-table-column>
        <el-table-column label="详情" type="expand">
          <template #default="props">
            <div
              v-loading="tableExpandRowInfo[props.row.id].loading"
              class="table-expand-area"
            >
              <div><strong>申请人信息:</strong></div>
              <div class="info">
                <p>
                  <span
                    ><strong>姓名: </strong
                    >{{
                      tableExpandRowInfo[props.row.id].applierInfo?.name
                    }}</span
                  >
                  <span
                    ><strong>学号: </strong
                    >{{
                      tableExpandRowInfo[props.row.id].applierInfo?.username
                    }}</span
                  >
                  <span
                    ><strong>学院: </strong
                    >{{
                      zeroWithDefault(
                        tableExpandRowInfo[props.row.id].applierInfo?.college,
                        '无'
                      )
                    }}</span
                  >
                </p>
                <p>
                  <span
                    ><strong>电话: </strong
                    >{{
                      zeroWithDefault(
                        tableExpandRowInfo[props.row.id].applierInfo?.tel,
                        '无'
                      )
                    }}</span
                  >
                  <span
                    ><strong>邮箱地址: </strong
                    >{{
                      zeroWithDefault(
                        tableExpandRowInfo[props.row.id].applierInfo?.email,
                        '无'
                      )
                    }}</span
                  >
                </p>
              </div>
              <div><strong>导师信息:</strong></div>
              <div class="info">
                <p>
                  <span
                    ><strong>姓名: </strong
                    >{{
                      tableExpandRowInfo[props.row.id].tutorInfo?.name
                    }}</span
                  >
                  <span
                    ><strong>学号: </strong
                    >{{
                      tableExpandRowInfo[props.row.id].tutorInfo?.username
                    }}</span
                  >
                  <span
                    ><strong>学院: </strong
                    >{{
                      zeroWithDefault(
                        tableExpandRowInfo[props.row.id].tutorInfo?.college,
                        '无'
                      )
                    }}</span
                  >
                </p>
                <p>
                  <span
                    ><strong>电话: </strong
                    >{{
                      zeroWithDefault(
                        tableExpandRowInfo[props.row.id].tutorInfo?.tel,
                        '无'
                      )
                    }}</span
                  >
                  <span
                    ><strong>邮箱地址: </strong
                    >{{
                      zeroWithDefault(
                        tableExpandRowInfo[props.row.id].tutorInfo?.email,
                        '无'
                      )
                    }}</span
                  >
                </p>
              </div>
              <div><strong>项目信息: </strong></div>
              <div class="info">
                <p>
                  <span
                    ><strong>名称: </strong
                    >{{
                      zeroWithDefault(
                        tableExpandRowInfo[props.row.id].projectInfo?.name,
                        '无'
                      )
                    }}</span
                  >
                  <span
                    ><strong>来源: </strong
                    >{{
                      zeroWithDefault(
                        tableExpandRowInfo[props.row.id].projectInfo?.from,
                        '无'
                      )
                    }}</span
                  >
                  <span
                    ><strong>编号: </strong
                    >{{
                      zeroWithDefault(
                        tableExpandRowInfo[props.row.id].projectInfo?.numbering,
                        '无'
                      )
                    }}</span
                  >
                  <span
                    ><strong>经费: </strong
                    >{{
                      zeroWithDefault(
                        tableExpandRowInfo[props.row.id].projectInfo?.expenses,
                        '无'
                      )
                    }}</span
                  >
                </p>
                <p>
                  <span
                    ><strong>描述: </strong
                    >{{
                      zeroWithDefault(
                        tableExpandRowInfo[props.row.id].projectInfo
                          ?.description,
                        '无'
                      )
                    }}</span
                  >
                </p>
              </div>
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
</style>
