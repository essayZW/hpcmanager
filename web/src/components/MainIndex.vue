<script lang="ts" setup>
import { UserInfo } from '../api/user';
import {
  getUserInfoById,
  getUserInfoFromStorage,
  isSuperAdmin,
  servicePing as userPing,
} from '../service/user';
import { UserQuotaInfo } from '../api/hpc';

import { servicePing as permissionPing } from '../service/permission';
import { servicePing as nodePing } from '../service/node';
import { servicePing as projectPing } from '../service/project';
import { servicePing as feePing } from '../service/fee';
import { servicePing as hpcPing, getHpcUserQuotaInfo } from '../service/hpc';
import { servicePing as fssPing } from '../service/fss';
import { servicePing as awardPing } from '../service/award';

import { zeroWithDefault } from '../utils/obj';
import { onMounted, onUnmounted, reactive, ref } from 'vue';
import dayjs from 'dayjs';

document.title = '计算平台管理系统';

const userInfo = ref<UserInfo | null>(null);

const serviceState = reactive<{
  user: boolean;
  permission: boolean;
  node: boolean;
  project: boolean;
  fee: boolean;
  hpc: boolean;
  fss: boolean;
  award: boolean;
}>({
  user: false,
  permission: false,
  node: false,
  project: false,
  fee: false,
  hpc: false,
  fss: false,
  award: false,
});

const quotaInfo = ref<UserQuotaInfo | undefined>(undefined);
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

  if (!userInfo.value.tel || !userInfo.value.email || !userInfo.value.college) {
    ElNotification({
      title: '消息完善',
      message: '请修改用户信息完善用户联系方式、邮箱以及学院信息',
      type: 'warning',
      offset: 80,
    });
  }
  if (userInfo.value) {
    try {
      const info = await getHpcUserQuotaInfo(userInfo.value.hpcUserID);
      quotaInfo.value = info;
    } catch (error) {
      console.log(`no quota info for current user`);
    }
  }
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
  serviceState.fss = await fssPing();
  serviceState.award = await awardPing();
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
    <el-col v-if="isSuperAdmin()" :sm="0" :lg="1"> </el-col>
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
        <div>
          <div v-if="serviceState.fss" class="success-state"></div>
          <div v-else class="error-state"></div>
          <strong>文件存储服务</strong>
        </div>
        <div>
          <div v-if="serviceState.award" class="success-state"></div>
          <div v-else class="error-state"></div>
          <strong>奖励服务</strong>
        </div>
      </el-card>
    </el-col>
  </el-row>
  <el-row v-if="quotaInfo" justify="center" class="second-row">
    <el-col :lg="10" :sm="22">
      <el-card>
        <template #header>
          <span>用户存储情况: </span>
        </template>
        <p>
          <strong>已用空间: {{ quotaInfo.used }}</strong>
        </p>
        <p>
          <strong>最大容量: {{ quotaInfo.max }}</strong>
        </p>
        <p>
          <strong>使用期限: </strong>
          {{ dayjs(quotaInfo.startTimeUnix * 1000).format('YYYY-MM-DD') }}至{{
            dayjs(quotaInfo?.endTimeUnix * 1000).format('YYYY-MM-DD')
          }}
        </p>
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
.second-row {
  margin-top: 16px;
}
</style>
