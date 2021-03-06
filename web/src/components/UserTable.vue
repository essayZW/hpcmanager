<script setup lang="ts">
import { reactive, ref } from 'vue';
import { UserInfo } from '../api/user';
import { GroupInfo } from '../api/group';
import { getGroupInfoByID } from '../service/group';
import {
  paginationGetUserInfo,
  isSuperAdmin,
  isAdmin,
  UserLevels,
} from '../service/user';
import dayjs from 'dayjs';
import { HpcUser, UserQuotaInfo } from '../api/hpc';
import {
  getHpcUserInfoByID,
  getHpcUserQuotaInfo,
  modifyUserQuotaInfo,
} from '../service/hpc';
import { zeroWithDefault } from '../utils/obj';
import { PermissionInfo } from '../api/permission';
import {
  getUserPermissionInfoByID,
  nameTransform,
  setAdminByUserID,
  delAdminByUserID,
} from '../service/permission';

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
// 表的扩展字段属性
const tableRowExtraInfo = reactive<{
  [id: number]: {
    group?: GroupInfo;
    hpcUser?: HpcUser;
    permission?: PermissionInfo[];
    quotaInfo?: UserQuotaInfo;
    loading?: boolean;
  };
}>({});

// 刷新表格数据
const refreshTable = () => {
  loadTableData(paginationInfo.pageIndex, paginationInfo.pageSize);
  // 清除缓存的表的扩展字段的属性
  for (const key in tableRowExtraInfo) {
    tableRowExtraInfo[key] = {};
  }
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

const canSetCommonAdmin = (id: number): boolean => {
  if (!id || !tableRowExtraInfo[id] || !tableRowExtraInfo[id].permission) {
    return false;
  }
  for (const item of tableRowExtraInfo[id].permission as PermissionInfo[]) {
    if (
      item.level == UserLevels.CommonAdmin ||
      item.level == UserLevels.SuperAdmin
    ) {
      return false;
    }
  }
  return true;
};
const canCancelCommonAdmin = (id: number): boolean => {
  if (!id || !tableRowExtraInfo[id] || !tableRowExtraInfo[id].permission) {
    return false;
  }
  for (const item of tableRowExtraInfo[id].permission as PermissionInfo[]) {
    if (item.level == UserLevels.SuperAdmin) {
      return false;
    }
  }
  return true;
};

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
  if (!tableRowExtraInfo[row.id].permission) {
    try {
      const permissionInfo = await getUserPermissionInfoByID(row.id);
      tableRowExtraInfo[row.id].permission = permissionInfo;
    } catch (error) {
      ElMessage({
        type: 'error',
        message: `${error}`,
      });
    }
  }
  if (
    !tableRowExtraInfo[row.id].quotaInfo &&
    tableRowExtraInfo[row.id].hpcUser?.quotaStartTime &&
    (tableRowExtraInfo[row.id].hpcUser as HpcUser).quotaStartTime > 0
  ) {
    try {
      const quotaInfo = await getHpcUserQuotaInfo(row.hpcUserID);
      tableRowExtraInfo[row.id].quotaInfo = quotaInfo;
    } catch (error) {
      tableRowExtraInfo[row.id].quotaInfo = undefined;
    }
  }
  tableRowExtraInfo[row.id].loading = false;
};

const setAdminHandler = async (id: number) => {
  if (!confirm('确认设置该用户为管理员吗?')) {
    return;
  }
  try {
    await setAdminByUserID(id);
    ElMessage({
      type: 'success',
      message: '设置成功',
    });
    refreshTable();
    // 让缓存失效
    tableRowExtraInfo[id].permission = undefined;
  } catch (error) {
    ElMessage({
      type: 'error',
      message: `${error}`,
    });
  }
};
const delAdminHandler = async (id: number) => {
  if (!confirm('确定删除该用户的管理员权限吗?')) {
    return;
  }
  try {
    await delAdminByUserID(id);
    ElMessage({
      type: 'success',
      message: '删除成功',
    });
    refreshTable();
    // 让缓存失效
    tableRowExtraInfo[id].permission = undefined;
  } catch (error) {
    ElMessage({
      type: 'error',
      message: `${error}`,
    });
  }
};

const modifyUserQuotaDialogVisible = ref<boolean>(false);

const dialogUserInfo = ref<UserInfo | undefined>(undefined);
const dialogUserQuotaInfo = ref<UserQuotaInfo | undefined>(undefined);

const showModifyUserQuotaDialog = async (row: UserInfo) => {
  dialogUserInfo.value = row;

  try {
    const data = await getHpcUserQuotaInfo(row.hpcUserID);
    dialogUserQuotaInfo.value = data;
  } catch (error) {
    ElMessage({
      type: 'error',
      message: `${error}`,
    });
    return;
  }
  dialogFormData.endTime = new Date(
    dialogUserQuotaInfo.value.endTimeUnix * 1000
  );
  dialogFormData.maxSize = parseInt(dialogUserQuotaInfo.value.max);
  modifyUserQuotaDialogVisible.value = true;
};
const hideModifyUserQuotaDialog = () => {
  modifyUserQuotaDialogVisible.value = false;
};

const dialogFormData = reactive<{
  operType: number;
  maxSize: number;
  endTime: Date;
}>({
  operType: 1,
  maxSize: 0,
  endTime: new Date(),
});

const modifyUserQuotaSubmit = async () => {
  if (!dialogUserInfo.value) {
    ElMessage({
      type: 'error',
      message: `用户信息加载失败,请刷新页面`,
    });
    return;
  }
  if (!dialogUserQuotaInfo.value) {
    ElMessage({
      type: 'error',
      message: `用户信息存储加载失败,请刷新页面`,
    });
    return;
  }
  try {
    await modifyUserQuotaInfo(
      dialogUserInfo.value.hpcUserID,
      parseInt(dialogUserQuotaInfo.value.max),
      dialogFormData.maxSize,
      dialogUserQuotaInfo.value?.endTimeUnix * 1000,
      dialogFormData.endTime.getTime(),
      dialogFormData.operType == 2
    );
    ElMessage({
      type: 'success',
      message: '修改成功',
    });
    hideModifyUserQuotaDialog();
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
  <el-row justify="space-between" class="refresh-button-row">
    <div>
      <slot name="tool"></slot>
    </div>
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
        table-layout="auto"
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
        <el-table-column label="用户组ID">
          <template #default="props">
            {{ zeroWithDefault(props.row.groupId, '无') }}
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
              <p><strong>用户存储信息: </strong></p>
              <div
                v-if="
                  props.row.hpcUserID &&
                  tableRowExtraInfo[props.row.id].quotaInfo
                "
                class="info"
              >
                <p>
                  <span>
                    <strong>使用容量: </strong>
                    {{ tableRowExtraInfo[props.row.id].quotaInfo?.used }}
                  </span>
                  <span
                    ><strong>最大容量: </strong
                    >{{ tableRowExtraInfo[props.row.id].quotaInfo?.max }}</span
                  >
                  <span>
                    <strong>使用期限: </strong>
                    {{
                      dayjs(zeroWithDefault(tableRowExtraInfo[props.row.id].quotaInfo?.startTimeUnix, 0) as number * 1000).format('YYYY-MM-DD')
                    }}至{{
                      dayjs(
                        zeroWithDefault(tableRowExtraInfo[props.row.id].quotaInfo?.endTimeUnix, 0) as number *
                          1000
                      ).format('YYYY-MM-DD')
                    }}
                  </span>
                </p>
                <p v-if="isAdmin()">
                  <span
                    ><strong>操作: </strong
                    ><el-button
                      size="small"
                      type="primary"
                      @click="showModifyUserQuotaDialog(props.row)"
                      >修改存储信息</el-button
                    ></span
                  >
                </p>
              </div>
              <p v-else class="info">无存储空间使用信息</p>
              <p><strong>用户权限信息: </strong></p>
              <p class="info">
                <span
                  ><strong>拥有权限: </strong
                  ><span
                    v-for="item in tableRowExtraInfo[props.row.id].permission"
                    :key="item.id"
                    :title="item.description"
                    >{{ nameTransform(item.name) }}</span
                  >
                </span>
              </p>
              <p v-if="isSuperAdmin()" class="info">
                <span
                  ><strong>操作: </strong
                  ><el-button
                    v-if="canSetCommonAdmin(props.row.id)"
                    type="primary"
                    size="small"
                    class="permission-admin-button"
                    @click="setAdminHandler(props.row.id)"
                    >设置为管理员</el-button
                  >
                  <el-button
                    v-else-if="canCancelCommonAdmin(props.row.id)"
                    type="warning"
                    size="small"
                    @click="delAdminHandler(props.row.id)"
                    >取消管理员</el-button
                  >
                  <span v-else>无</span>
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
  <el-dialog
    v-if="isAdmin()"
    v-model="modifyUserQuotaDialogVisible"
    title="修改用户存储信息"
  >
    <el-form inline>
      <el-form-item label="用户姓名: ">
        <span>{{ dialogUserInfo?.name }}</span>
      </el-form-item>
      <el-form-item label="用户学(工)号: "
        ><span>{{ dialogUserInfo?.username }}</span></el-form-item
      >
      <el-form-item label="当前最大容量: "
        ><span>{{ dialogUserQuotaInfo?.max }}</span></el-form-item
      >
      <el-form-item label="当前使用期限: "
        ><span
          >{{
            dayjs(zeroWithDefault(dialogUserQuotaInfo?.startTimeUnix, 0) as number * 1000).format(
              'YYYY-MM-DD'
            )
          }}至{{
            dayjs(zeroWithDefault(dialogUserQuotaInfo?.endTimeUnix, 0) as number * 1000).format('YYYY-MM-DD')
          }}</span
        ></el-form-item
      >
    </el-form>
    <el-form>
      <el-form-item label="操作: ">
        <el-radio-group v-model="dialogFormData.operType">
          <el-radio-button :label="1">扩容</el-radio-button>
          <el-radio-button :label="2">延期</el-radio-button>
        </el-radio-group>
      </el-form-item>
      <el-form-item v-if="dialogFormData.operType == 1" label="新的最大容量: ">
        <el-input
          v-model.number="dialogFormData.maxSize"
          type="number"
          placeholder="新的最大容量限制,不能小于现在的最大容量"
        >
          <template #append>TB</template>
        </el-input>
      </el-form-item>
      <el-form-item
        v-if="dialogFormData.operType == 2"
        label="新的期限结束日期: "
      >
        <el-date-picker
          v-model="dialogFormData.endTime"
          type="date"
          placeholder="选择新的期限结束日期"
        />
      </el-form-item>
    </el-form>
    <template #footer>
      <span class="dialog-footer">
        <el-button @click="hideModifyUserQuotaDialog">取消</el-button>
        <el-button type="primary" @click="modifyUserQuotaSubmit"
          >修改</el-button
        >
      </span>
    </template>
  </el-dialog>
</template>
<style lang="less" scoped>
.pagination-row {
  margin: 16px 0px;
  .pagination-control-panel {
    margin: 0px auto;
    justify-content: center;
  }
}
p.info,
div.info {
  padding-left: 16px;
  span {
    margin-right: 12px;
  }
}
.refresh-button-row {
  margin: 16px 0px;
}

.permission-admin-button {
  margin-left: 8px;
}
</style>
