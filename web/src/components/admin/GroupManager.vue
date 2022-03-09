<script setup lang="ts">
import PageTitle from '../PageTitle.vue';
import { reactive, ref } from 'vue';
import { GroupInfo } from '../../api/group';
import { paginationGetGroupInfo, createGroup } from '../../service/group';
import { PaginationQueryResponse } from '../../api/api';
import { getUserInfoById, getUserIdByUsername } from '../../service/user';
import { UserInfo } from '../../api/user';
import { requiredWithLength, FormInstance } from '../../utils/validateRule';

const tableData = reactive<{
  count: number;
  data: GroupInfo[];
}>({
  count: 0,
  data: [],
});

const paginationInfo = reactive<{
  pageIndex: number;
  pageSize: number;
}>({
  pageIndex: 1,
  pageSize: 10,
});

const tableLoadingFlag = ref(false);

const refreshTableData = async (pageIndex: number, pageSize: number) => {
  if (tableLoadingFlag.value) {
    console.log('table loading');
    return;
  }
  console.log(`loadtable page: ${pageIndex}, size: ${pageSize}`);
  tableLoadingFlag.value = true;
  const data = await paginationGetGroupInfo(pageIndex, pageSize);
  tableLoadingFlag.value = false;
  if (typeof data == 'string') {
    ElMessage({
      type: 'error',
      message: data,
    });
  }
  tableData.data = (data as PaginationQueryResponse<GroupInfo>).Data;
  tableData.count = (data as PaginationQueryResponse<GroupInfo>).Count;
};

refreshTableData(paginationInfo.pageIndex, paginationInfo.pageSize);

const handleSizeChange = (size: number) => {
  paginationInfo.pageSize = size;
  refreshTableData(paginationInfo.pageIndex, paginationInfo.pageSize);
};

const handleCurrentChange = (index: number) => {
  paginationInfo.pageIndex = index;
  refreshTableData(paginationInfo.pageIndex, paginationInfo.pageSize);
};

const tableRowExtraInfo = reactive<{
  [id: number]: {
    user: UserInfo;
  };
}>({});
const rowExpanded = async (row: GroupInfo) => {
  // TODO 查询信息是否已经存储在缓存中
  const tutorInfo = await getUserInfoById(row.tutorID);
  if (typeof tutorInfo == 'string') {
    ElMessage({
      type: 'error',
      message: tutorInfo,
    });
  } else {
    tableRowExtraInfo[row.id].user = tutorInfo as UserInfo;
  }
};

const refreshTable = () => {
  refreshTableData(paginationInfo.pageIndex, paginationInfo.pageSize);
};

const showFormDialog = () => {
  formDialogVisible.value = true;
};

const hideForamDialog = () => {
  formDialogVisible.value = false;
};

const formDialogVisible = ref(false);

const createGroupForm = reactive<{
  name: string;
  tutorUsername: string;
  queueName: string;
}>({
  name: '',
  tutorUsername: '',
  queueName: '',
});

const createGroupFormElem = ref<FormInstance>();

const createGroupFormRules = {
  name: requiredWithLength('用户组名', 2, 64),
  tutorUsername: requiredWithLength('导师工号', 6, 16),
  queueName: requiredWithLength('私有队列名', 2, 64),
};

const submitCreateGroupForm = (elem: FormInstance | undefined) => {
  if (!elem) {
    return;
  }
  elem.validate(async (valid) => {
    if (!valid) {
      return valid;
    }
    try {
      let id = await createGroup(
        createGroupForm.queueName,
        createGroupForm.name,
        createGroupForm.tutorUsername
      );
      refreshTable();
      ElMessage({
        type: 'success',
        message: `创建成功,新组ID: ${id}`,
      });
      hideForamDialog();
    } catch (error) {
      ElMessage({
        type: 'error',
        message: `${error}`,
      });
    }
    return valid;
  });
};
</script>
<template>
  <page-title title="用户组管理" des="管理员查看、新建用户分组"></page-title>
  <el-row class="operation-row">
    <el-button type="primary" @click="showFormDialog">新建用户组</el-button>
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
        v-loading="tableLoadingFlag"
        :data="tableData.data"
        border
        class="group-table"
        @expand-change="rowExpanded"
      >
        <el-table-column label="ID" prop="ID" align="center"></el-table-column>
        <el-table-column
          label="组名"
          prop="name"
          align="center"
        ></el-table-column>
        <el-table-column
          label="导师用户名"
          prop="tutorUserName"
          align="center"
        ></el-table-column>
        <el-table-column
          label="导师姓名"
          prop="tutorName"
          align="center"
        ></el-table-column>
        <el-table-column
          label="创建时间"
          prop="createTime"
          align="center"
        ></el-table-column>
        <el-table-column label="信息" type="expand" align="center">
          <template #default="props">
            <p>创建者帐号: {{ props.row.createrUsername }}</p>
            <p>创建者姓名: {{ props.row.createrName }}</p>
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
  <el-dialog v-model="formDialogVisible" title="创建用户组">
    <el-form
      ref="createGroupFormElem"
      :rules="createGroupFormRules"
      :model="createGroupForm"
    >
      <el-form-item label="用户组名" prop="name">
        <el-input
          v-model="createGroupForm.name"
          placeholder="用户组名"
          type="text"
        ></el-input>
      </el-form-item>
      <el-form-item label="导师工号" prop="tutorUsername">
        <el-input
          v-model="createGroupForm.tutorUsername"
          placeholder="导师工号"
          type="text"
        ></el-input>
      </el-form-item>
      <el-form-item label="私有队列名称" prop="queueName">
        <el-input
          v-model="createGroupForm.queueName"
          placeholder="私有队列名称"
          type="text"
        ></el-input>
      </el-form-item>
    </el-form>
    <template #footer>
      <span class="dialog-footer">
        <el-button @click="hideForamDialog">取消</el-button>
        <el-button
          type="primary"
          @click="submitCreateGroupForm(createGroupFormElem)"
          >确认</el-button
        >
      </span>
    </template>
  </el-dialog>
</template>
<style lang="less">
.group-table {
  width: 100%;
  margin-top: 24px;
}
.pagination-row {
  margin: 16px 0px;
  .pagination-control-panel {
    margin: 0px auto;
    justify-content: center;
  }
}
.operation-row {
  margin-top: 16px;
  justify-content: space-between;
}
</style>
