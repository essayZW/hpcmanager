<script lang="ts" setup>
import { UserInfo } from '../api/user';
import {
  getUserInfoById,
  getUserInfoFromStorage,
  isSuperAdmin,
  servicePing as userPing,
} from '../service/user';
import { servicePing as permissionPing } from '../service/permission';
import { servicePing as nodePing } from '../service/node';
import { servicePing as projectPing } from '../service/project';
import { servicePing as feePing } from '../service/fee';
import { servicePing as hpcPing } from '../service/hpc';

import { zeroWithDefault } from '../utils/obj';
import { onMounted, onUnmounted, reactive, ref } from 'vue';

document.title = '计算平台管理系统';

const userInfo = ref<UserInfo | null>(null);

const serviceState = reactive<{
  user: boolean;
  permission: boolean;
  node: boolean;
  project: boolean;
  fee: boolean;
  hpc: boolean;
}>({
  user: false,
  permission: false,
  node: false,
  project: false,
  fee: false,
  hpc: false,
});
onMounted(async () => {
  const loginUserInfo = getUserInfoFromStorage();
  if (!loginUserInfo) {
    return;
  }
  const res = await getUserInfoById(loginUserInfo.UserId);
  if (typeof res == 'string') {
    ElMessage({
      type: 'error',
      message: res as string,
    });
    return;
  }
  userInfo.value = res as UserInfo;
});
const refreshServiceState = async () => {
  if (!isSuperAdmin()) {
    return;
  }
  console.log('refresh');
  serviceState.user = await userPing();
  serviceState.permission = await permissionPing();
  serviceState.node = await nodePing();
  serviceState.project = await projectPing();
  serviceState.fee = await feePing();
  serviceState.hpc = await hpcPing();
};
const interval = setInterval(refreshServiceState, 5000);
refreshServiceState();
onUnmounted(() => {
  console.log('leave');
  clearInterval(interval);
});
</script>
<template>
  <el-row justify="center">
    <el-col :lg="10" :sm="22">
      <el-card v-if="userInfo">
        <template #header>
          <span>用户信息</span>
        </template>
        <p><strong>姓名:&nbsp;</strong>{{ userInfo.name }}</p>
        <p><strong>学(工)号:&nbsp;</strong>{{ userInfo.username }}</p>
        <p>
          <strong>联系方式:&nbsp;</strong
          >{{ zeroWithDefault(userInfo.tel, '无') }}
        </p>
        <p>
          <strong>邮箱地址:&nbsp;</strong
          >{{ zeroWithDefault(userInfo.email, '无') }}
        </p>
        <p>
          <strong>所属学院:&nbsp;</strong
          >{{ zeroWithDefault(userInfo.college, '无') }}
        </p>
      </el-card>
    </el-col>
    <el-col :sm="0" :lg="1"> </el-col>
    <el-col v-if="isSuperAdmin()" :lg="10" :sm="22">
      <el-card>
        <template #header>
          <span>系统状态: </span>
        </template>
        <div>
          <div v-if="serviceState.user" class="success-state"></div>
          <div v-else class="error-state"></div>
          <strong>用户服务</strong>
        </div>
        <div>
          <div v-if="serviceState.permission" class="success-state"></div>
          <div v-else class="error-state"></div>
          <strong>权限服务</strong>
        </div>
        <div>
          <div v-if="serviceState.node" class="success-state"></div>
          <div v-else class="error-state"></div>
          <strong>机器节点服务</strong>
        </div>
        <div>
          <div v-if="serviceState.project" class="success-state"></div>
          <div v-else class="error-state"></div>
          <strong>项目服务</strong>
        </div>
        <div>
          <div v-if="serviceState.fee" class="success-state"></div>
          <div v-else class="error-state"></div>
          <strong>费用服务</strong>
        </div>
        <div>
          <div v-if="serviceState.hpc" class="success-state"></div>
          <div v-else class="error-state"></div>
          <strong>作业调度服务</strong>
        </div>
      </el-card>
    </el-col>
  </el-row>
</template>
<style lang="less">
.success-state {
  width: 10px;
  height: 10px;
  background-color: green;
  border-radius: 50%;
  display: inline-block;
  margin-left: 8px;
  margin-right: 8px;
}

.error-state {
  width: 10px;
  height: 10px;
  background-color: red;
  border-radius: 50%;
  display: inline-block;
  margin-left: 8px;
  margin-right: 8px;
}
</style>
