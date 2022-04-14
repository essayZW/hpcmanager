<script setup lang="ts">
import PageTitle from '../PageTitle.vue';
import { reactive, ref } from 'vue';
import { GroupInfo } from '../../api/group';
import {
  paginationGetGroupInfo,
  createGroup,
  addGroupBalance,
} from '../../service/group';
import { PaginationQueryResponse } from '../../api/api';
import { getUserInfoById, isAdmin } from '../../service/user';
import { UserInfo } from '../../api/user';
import { requiredWithLength, FormInstance } from '../../utils/validateRule';
import { zeroWithDefault } from '../../utils/obj';
import { HpcGroup } from '../../api/hpc';
import { getHpcGroupInfoByID } from '../../service/hpc';

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
    user?: UserInfo;
    group?: HpcGroup;
  };
}>({});
const rowExpanded = async (row: GroupInfo) => {
  if (tableRowExtraInfo[row.id] && tableRowExtraInfo[row.id].user) {
    return;
  }
  tableRowExtraInfo[row.id] = {};
  const tutorInfo = await getUserInfoById(row.tutorID);
  if (typeof tutorInfo == 'string') {
    ElMessage({
      type: 'error',
      message: tutorInfo,
    });
  } else {
    tableRowExtraInfo[row.id].user = tutorInfo;
  }

  const hpcGroupInfo = await getHpcGroupInfoByID(row.hpcGroupID);
  if (typeof hpcGroupInfo == 'string') {
    ElMessage({
      type: 'error',
      message: hpcGroupInfo,
    });
  } else {
    tableRowExtraInfo[row.id].group = hpcGroupInfo;
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

const groupBalanceUpdateDialog = ref<boolean>(false);

const showGroupBalanceUpdateDialog = (groupInfo: GroupInfo) => {
  groupBalanceUpdateFormData.addFee = 0;
  groupBalanceUpdateFormData.id = groupInfo.id;
  groupBalanceUpdateFormData.tutorName = groupInfo.tutorName;
  groupBalanceUpdateFormData.tutorUsername = groupInfo.tutorUsername;
  groupBalanceUpdateDialog.value = true;
};

const hideGroupBalanceUpdateDialog = () => {
  groupBalanceUpdateDialog.value = false;
};

const groupBalanceUpdateFormData = reactive<{
  id: number;
  tutorName: string;
  tutorUsername: string;
  oldFee: number;
  addFee: number;
}>({
  id: 0,
  tutorUsername: '无',
  tutorName: '无',
  oldFee: 0,
  addFee: 0,
});

const groupBalanceUpdateFormRule = {
  addFee: [
    { required: true, message: '新增余额不能为空', trigger: 'blur' },
    {
      validator: (
        _: unknown,
        value: number,
        callback: (arg0: Error | undefined) => void
      ) => {
        if (value <= 0) {
          callback(new Error('新增余额必须大于0'));
        }
        callback(undefined);
      },
      trigger: 'blur',
    },
  ],
};

const groupBalanceUpdateFormElem = ref<FormInstance>();

const groupBalanceUpdateFormSubmit = (elem: FormInstance | undefined) => {
  if (!elem) {
    return;
  }
  elem.validate(async (valid) => {
    if (valid) {
      try {
        await addGroupBalance(
          groupBalanceUpdateFormData.id,
          groupBalanceUpdateFormData.addFee
        );
        hideGroupBalanceUpdateDialog();
        ElMessage({
          type: 'success',
          message: '修改成功',
        });
        refreshTable();
      } catch (error) {
        ElMessage({
          type: 'error',
          message: `${error}`,
        });
      }
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
        <el-table-column label="id" prop="id" align="center"></el-table-column>
        <el-table-column
          label="组名"
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
        <el-table-column
          label="创建时间"
          prop="createTime"
          align="center"
        ></el-table-column>
        <el-table-column label="信息" type="expand" align="center">
          <template #default="props">
            <div class="row-expand-info">
              <p><strong>导师信息: </strong></p>
              <p v-if="tableRowExtraInfo[props.row.id]" class="info">
                <span
                  ><strong>工号: </strong>
                  {{ tableRowExtraInfo[props.row.id].user?.username }}</span
                >
                <span
                  ><strong>姓名: </strong>
                  {{ tableRowExtraInfo[props.row.id].user?.name }}</span
                >
                <span
                  ><strong>用户ID: </strong
                  >{{ tableRowExtraInfo[props.row.id].user?.id }}</span
                >
              </p>
              <p v-if="tableRowExtraInfo[props.row.id]" class="info">
                <span
                  ><strong>电话: </strong
                  >{{
                    zeroWithDefault(
                      tableRowExtraInfo[props.row.id].user?.tel,
                      '无'
                    )
                  }}</span
                >
                <span
                  ><strong>邮箱: </strong
                  >{{
                    zeroWithDefault(
                      tableRowExtraInfo[props.row.id].user?.email,
                      '无'
                    )
                  }}</span
                >
                <span
                  ><strong>学院: </strong
                  >{{
                    zeroWithDefault(
                      tableRowExtraInfo[props.row.id].user?.college,
                      '无'
                    )
                  }}</span
                >
              </p>
              <p><strong>用户组创建者信息: </strong></p>
              <p class="info">
                <span>
                  <strong>用户ID: </strong>{{ props.row.createrID }}
                </span>
                <span>
                  <strong>工号: </strong>{{ props.row.createrUsername }}
                </span>
                <span>
                  <strong>姓名: </strong>{{ props.row.createrName }}
                </span>
              </p>
              <p><strong>用户组信息: </strong></p>
              <p class="info">
                <span
                  ><strong>用户组余额: </strong>{{ props.row.balance }}元</span
                >
                <span
                  ><strong>私有队列名称: </strong
                  >{{
                    zeroWithDefault(
                      tableRowExtraInfo[props.row.id].group?.queueName,
                      '无'
                    )
                  }}</span
                >
                <span
                  ><strong>计算节点用户组名: </strong
                  >{{
                    zeroWithDefault(
                      tableRowExtraInfo[props.row.id].group?.name,
                      '无'
                    )
                  }}</span
                >
                <span
                  ><strong>计算节点GID: </strong
                  >{{
                    zeroWithDefault(
                      tableRowExtraInfo[props.row.id].group?.gID,
                      '无'
                    )
                  }}</span
                >
              </p>
              <p><strong>操作:</strong></p>
              <p class="info">
                <el-button
                  v-if="isAdmin()"
                  type="primary"
                  size="small"
                  @click="showGroupBalanceUpdateDialog(props.row)"
                  >余额充值</el-button
                >
                <span v-else>无</span>
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
  <el-dialog v-model="groupBalanceUpdateDialog" title="用户组余额充值">
    <el-form
      ref="groupBalanceUpdateFormElem"
      :model="groupBalanceUpdateFormData"
      :rules="groupBalanceUpdateFormRule"
      @submit.prevent
    >
      <el-form-item label="用户组ID:">
        <span>{{ groupBalanceUpdateFormData.id }}</span>
      </el-form-item>
      <el-form-item label="导师姓名:">
        <span>{{ groupBalanceUpdateFormData.tutorName }}</span>
      </el-form-item>
      <el-form-item label="导师工号:">
        <span>{{ groupBalanceUpdateFormData.tutorUsername }}</span>
      </el-form-item>
      <el-form-item label="现有余额">
        <span>{{ groupBalanceUpdateFormData.oldFee }}元</span>
      </el-form-item>
      <el-form-item label="添加的余额" prop="addFee">
        <el-input
          v-model.number="groupBalanceUpdateFormData.addFee"
          type="text"
        >
          <template #append>元</template>
        </el-input>
      </el-form-item>
    </el-form>
    <template #footer>
      <span class="dialog-footer">
        <el-button @click="hideGroupBalanceUpdateDialog">取消</el-button>
        <el-button
          type="primary"
          @click="groupBalanceUpdateFormSubmit(groupBalanceUpdateFormElem)"
          >确认</el-button
        >
      </span>
    </template>
  </el-dialog>
</template>
<style lang="less" scoped>
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
.row-expand-info {
  padding: 0px 12px;
  span {
    padding: 12px 12px 0px 0px;
  }
  .info {
    padding-left: 16px;
  }
}
</style>
