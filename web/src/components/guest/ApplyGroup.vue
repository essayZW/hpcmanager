<script setup lang="ts">
import PageTitle from '../PageTitle.vue';
import { onMounted, reactive, ref } from 'vue';
import { useRouter } from 'vue-router';
import { FormInstance, requiredWithLength } from '../../utils/validateRule';
import { searchTutorInfo, applyJoinGroup } from '../../service/group';
import { zeroWithDefault } from '../../utils/obj';

import ApplyGroupTable from '../ApplyGroupTable.vue';

// 表格实例
const tableElem = ref<typeof ApplyGroupTable | null>(null);

const router = useRouter();

// 当前激活的tab的名称
const activeName = ref('createApply');

onMounted(() => {
  if (router.currentRoute.value.query?.view == 'showApplies') {
    activeName.value = 'showApplies';
    if (tableElem.value) {
      tableElem.value.refreshTableData();
    }
  }
});
// tab点击事件处理
// eslint-disable-next-line @typescript-eslint/no-unused-vars, @typescript-eslint/no-explicit-any
const tabClickHandler = (tab: any, event: Event) => {
  if (!tableElem.value) {
    return;
  }
  if (tab.paneName == 'showApplies') {
    tableElem.value.refreshTableData();
  }
};

// 搜索导师信息表单数据
const searchTutorFormData = reactive<{
  tutorUsername: string;
}>({
  tutorUsername: '',
});

const searchTutorFormElem = ref<FormInstance>();

// 搜索导师信息表单验证规则
const searchTutorFormRules = {
  tutorUsername: requiredWithLength('导师用户名', 6, 32),
};

// 搜索到的导师基本信息
const searchedUserInfo = reactive<{
  tutorName: string;
  tutorID: number;
  groupID: number;
  groupName: string;
  flag: boolean;
  loading: boolean;
}>({
  tutorName: '',
  tutorID: 0,
  groupID: 0,
  groupName: '',
  flag: false,
  loading: false,
});

// 搜索导师信息表单处理函数
const searchTutorFormHandler = (elem: FormInstance | undefined) => {
  if (!elem) {
    return;
  }
  elem.validate(async (valid) => {
    if (!valid) {
      searchedUserInfo.flag = false;
      return valid;
    }
    try {
      searchedUserInfo.loading = true;
      const userInfo = await searchTutorInfo(searchTutorFormData.tutorUsername);
      searchedUserInfo.tutorName = userInfo.tutorName;
      searchedUserInfo.groupID = userInfo.groupID;
      searchedUserInfo.tutorID = userInfo.tutorID;
      searchedUserInfo.groupName = userInfo.groupName;
      searchedUserInfo.flag = true;
    } catch (error) {
      ElMessage({
        type: 'error',
        message: `${error}`,
      });
      searchedUserInfo.flag = false;
    } finally {
      searchedUserInfo.loading = false;
    }
    return valid;
  });
};

// 提交加入组申请
const submitJoinGroupApply = async () => {
  if (!searchedUserInfo.flag) {
    ElMessage({
      type: 'error',
      message: '请先搜索导师信息以确认导师信息无误',
    });
    return;
  }
  if (!searchedUserInfo.groupID) {
    ElMessage({
      type: 'error',
      message: '错误的用户组信息',
    });
    return;
  }

  searchedUserInfo.loading = true;
  try {
    const applyID = await applyJoinGroup(searchedUserInfo.groupID);
    ElMessage({
      type: 'success',
      message: `申请成功,申请记录ID:${applyID}`,
    });
    searchedUserInfo.flag = false;
  } catch (error) {
    ElMessage({
      type: 'error',
      message: `${error}`,
    });
  } finally {
    searchedUserInfo.loading = false;
  }
};
</script>
<template>
  <page-title
    title="用户组申请管理"
    des="查看并管理用户加入用户组申请"
  ></page-title>
  <el-row
    v-loading="searchedUserInfo.loading"
    justify="start"
    class="tab-panes-row"
  >
    <el-tabs v-model="activeName" type="card" @tab-click="tabClickHandler">
      <el-tab-pane label="创建申请" name="createApply">
        <el-col :sm="22" :lg="16">
          <el-form
            ref="searchTutorFormElem"
            :inline="true"
            :model="searchTutorFormData"
            :rules="searchTutorFormRules"
          >
            <el-form-item label="导师工号" prop="tutorUsername">
              <el-input
                v-model="searchTutorFormData.tutorUsername"
                placeholder="导师工号"
                type="text"
              ></el-input>
            </el-form-item>
            <el-form-item>
              <el-button
                type="primary"
                @click="searchTutorFormHandler(searchTutorFormElem)"
                >查找</el-button
              >
            </el-form-item>
          </el-form>
          <div v-if="searchedUserInfo.flag">
            <p><strong>导师信息:</strong></p>
            <p class="info">
              <span
                ><strong>姓名: </strong
                >{{ zeroWithDefault(searchedUserInfo.tutorName, '无') }}</span
              >
              <span
                ><strong>ID: </strong
                >{{ zeroWithDefault(searchedUserInfo.tutorID, 0) }}</span
              >
            </p>
            <p><strong>用户组信息: </strong></p>
            <p class="info">
              <span
                ><strong>组名: </strong
                >{{ zeroWithDefault(searchedUserInfo.groupName, '无') }}</span
              >
              <span
                ><strong>ID: </strong
                >{{ zeroWithDefault(searchedUserInfo.groupID, 0) }}</span
              >
            </p>
            <p>
              <el-button type="primary" @click="submitJoinGroupApply"
                >提交申请</el-button
              >
            </p>
          </div>
          <div v-else>
            <p>请先搜索导师信息以确认导师信息无误</p>
          </div>
        </el-col>
      </el-tab-pane>
      <el-tab-pane label="查看申请" name="showApplies">
        <apply-group-table ref="tableElem"></apply-group-table>
      </el-tab-pane>
    </el-tabs>
  </el-row>
</template>
<style lang="less" scoped>
.tab-panes-row {
  margin-top: 16px;
  .el-tabs {
    flex-grow: 2;
  }
}
.info {
  padding-left: 16px;
  span {
    margin: 0px 12px;
  }
}
</style>
