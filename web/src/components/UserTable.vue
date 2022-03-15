<script setup lang="ts">
import { reactive } from 'vue';
import { UserInfo } from '../api/user';
import { GroupInfo } from '../api/group';
import { getGroupInfoByID } from '../service/group';
import { paginationGetUserInfo } from '../service/user';
import dayjs from 'dayjs';
import { HpcUser } from '../api/hpc';
import { getHpcUserInfoByID } from '../service/hpc';
import { zeroWithDefault } from '../utils/obj';

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

// 表的扩展字段属性
const tableRowExtraInfo = reactive<{
  [id: number]: {
    group?: GroupInfo;
    hpcUser?: HpcUser;
    loading: boolean;
  };
}>({});
// table columnt expand事件处理
const expandChangeHandler = async (row: UserInfo) => {
  if (!tableRowExtraInfo[row.id]) {
    tableRowExtraInfo[row.id] = {
      loading: false,
    };
  }
  tableRowExtraInfo[row.id].loading = true;
  if (row.groupId && !tableRowExtraInfo[row.id].group) {
    // 加载组信息
    try {
      const group = await getGroupInfoByID(row.groupId);
      tableRowExtraInfo[row.id].group = group;
    } catch (error) {
      ElMessage({
        type: 'error',
        message: `${error}`,
      });
    }
  }
  if (row.hpcUserID && !tableRowExtraInfo[row.id].hpcUser) {
    // 加载 hpc_user 信息
    try {
      const hpcUser = await getHpcUserInfoByID(row.hpcUserID);
      tableRowExtraInfo[row.id].hpcUser = hpcUser;
    } catch (error) {
      ElMessage({
        type: 'error',
        message: `${error}`,
      });
    }
  }
  tableRowExtraInfo[row.id].loading = false;
};
</script>
<template>
  <el-row justify="end" class="refresh-button-row">
    <el-button type="primary" @click="refreshTable">
      <el-icon class="el-icon--left">
        <i-ic-round-refresh />
      </el-icon>
      刷新</el-button
    >
  </el-row>
  <el-row :span="24" justify="center">
    <el-col>
      <el-table
        v-loading="tableData.loading"
        border
        :data="tableData.data"
        @expand-change="expandChangeHandler"
      >
        <el-table-column label="ID" prop="id"></el-table-column>
        <el-table-column label="姓名" prop="name"></el-table-column>
        <el-table-column label="学号" prop="username"></el-table-column>
        <el-table-column label="学院">
          <template #default="props">
            {{ zeroWithDefault(props.row.college, '无') }}
          </template>
        </el-table-column>
        <el-table-column label="电话">
          <template #default="props">
            {{ zeroWithDefault(props.row.tel, '无') }}
          </template>
        </el-table-column>
        <el-table-column label="邮箱" prop="email">
          <template #default="props">
            {{ zeroWithDefault(props.row.email, '无') }}
          </template>
        </el-table-column>
        <el-table-column label="创建时间">
          <template #default="props">
            {{ dayjs(props.row.createTime * 1000).format('YYYY-MM-DD HH:mm') }}
          </template>
        </el-table-column>
        <el-table-column label="更多" type="expand">
          <template #default="props">
            <div v-loading="tableRowExtraInfo[props.row.id].loading">
              <p><strong>用户组信息:</strong></p>
              <p v-if="props.row.groupId" class="info">
                <span
                  ><strong>导师姓名: </strong
                  >{{ tableRowExtraInfo[props.row.id].group?.tutorName }}</span
                >
                <span
                  ><strong>导师用户名: </strong
                  >{{
                    tableRowExtraInfo[props.row.id].group?.tutorUsername
                  }}</span
                >
                <span
                  ><strong>导师ID: </strong
                  >{{ tableRowExtraInfo[props.row.id].group?.tutorID }}</span
                >
              </p>
              <p v-else class="info">未加入用户组</p>
            </div>
            <div>
              <p><strong>计算节点用户信息</strong></p>
              <p v-if="props.row.hpcUserID" class="info">
                <span>
                  <strong>UID: </strong>
                  {{ tableRowExtraInfo[props.row.id].hpcUser?.nodeUID }}
                </span>
                <span
                  ><strong>账户名: </strong
                  >{{
                    tableRowExtraInfo[props.row.id].hpcUser?.nodeUsername
                  }}</span
                >
              </p>
              <p v-else class="info">未创建计算节点账户</p>
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
<style lang="less" scoped>
.pagination-row {
  margin: 16px 0px;
}
p.info {
  padding-left: 16px;
}
.refresh-button-row {
  margin: 16px 0px;
}
</style>
