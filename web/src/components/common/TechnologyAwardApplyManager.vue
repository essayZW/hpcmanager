<script setup lang="ts">
import { reactive, ref } from 'vue';
import { getUserInfoFromStorage } from '../../service/user';
import { createTechnologyAwardApply, prizeLevels } from '../../service/award';
import { FormInstance } from '../../utils/validateRule';
import { ProjectInfo } from '../../api/project';
import { getProjectInfoByID } from '../../service/project';
import { zeroWithDefault } from '../../utils/obj';

import PageTitle from '../PageTitle.vue';
import TechnologyAwardApplyTable from '../TechnologyAwardApplyTable.vue';
import UploadImageFile from '../UploadImageFile.vue';

const tableElem = ref<typeof TechnologyAwardApplyTable | null>(null);

const refreshTableData = () => {
  if (!tableElem.value) {
    return;
  }
  tableElem.value.refreshTableData();
};

const technologyApplyDialogShowFlag = ref<boolean>(false);

const hideTechnologyApplyDrawer = () => {
  technologyApplyDialogShowFlag.value = false;
};

const showTechnologyApplyDrawer = () => {
  projectInfo.value = undefined;
  createTechnologyAwardApplyFormData.projectID = 0;
  createTechnologyAwardApplyFormData.remarkMessage = '';
  createTechnologyAwardApplyFormData.prizeImageName = '';
  technologyApplyDialogShowFlag.value = true;
};

const supportedTechnologyLevels = ref<string[]>(prizeLevels);
// 创建科技成果奖励申请的表单数据
const createTechnologyAwardApplyFormData = reactive<{
  projectID: number;
  prizeLevel: string;
  prizeImageName: string;
  remarkMessage: string;
}>({
  projectID: 0,
  prizeLevel: supportedTechnologyLevels.value[0],
  prizeImageName: '',
  remarkMessage: '',
});

const projectInfo = ref<ProjectInfo>();

const searchProjectInfo = async () => {
  try {
    const data = await getProjectInfoByID(
      createTechnologyAwardApplyFormData.projectID
    );
    projectInfo.value = data;
  } catch (error) {
    ElMessage({
      type: 'error',
      message: `${error}`,
    });
  }
};
const formRules = {
  title: [{ required: true, message: '科技成果标题不能为空', trigger: 'blur' }],
};

const formElem = ref<FormInstance | undefined>(undefined);

const createTechnologyAwardApplySubmit = (elem: FormInstance | undefined) => {
  if (!elem) {
    return;
  }
  elem.validate(async (valid) => {
    if (!valid) {
      return;
    }
    if (!projectInfo.value) {
      ElMessage({
        type: 'error',
        message: '请先搜索项目信息以进行关联',
      });
      return;
    }
    if (!createTechnologyAwardApplyFormData.prizeImageName) {
      ElMessage({
        type: 'error',
        message: '请上传奖项图片',
      });
      return;
    }
    try {
      await createTechnologyAwardApply(
        createTechnologyAwardApplyFormData.projectID,
        createTechnologyAwardApplyFormData.prizeLevel,
        createTechnologyAwardApplyFormData.prizeImageName,
        createTechnologyAwardApplyFormData.remarkMessage
      );
      ElMessage({
        type: 'success',
        message: '创建成功',
      });
      refreshTableData();
      // 由于文件上传的组件不太好用,所以这里直接跳转页面对文件上传组件进行重置
      // eslint-disable-next-line no-self-assign
      window.location.href = window.location.href;
      hideTechnologyApplyDrawer();
    } catch (error) {
      ElMessage({
        type: 'error',
        message: `${error}`,
      });
    }
  });
};

// 如果当前用户没有用户组则不显示申请科技成果奖励按钮
const hasGroup = ref<boolean>(true);
const userInfo = getUserInfoFromStorage();

if (userInfo) {
  if (!userInfo.GroupId) {
    hasGroup.value = false;
  }
}
</script>
<template>
  <page-title title="科技成果奖励申请" des="论文奖励申请管理"></page-title>
  <technology-award-apply-table ref="tableElem">
    <template #tool>
      <el-button
        v-if="hasGroup"
        type="primary"
        @click="showTechnologyApplyDrawer"
        >申请科技成果奖励</el-button
      >
    </template>
  </technology-award-apply-table>
  <el-drawer
    v-model="technologyApplyDialogShowFlag"
    size="40%"
    title="申请科技成果奖励"
  >
    <el-form
      ref="formElem"
      :model="createTechnologyAwardApplyFormData"
      :rules="formRules"
    >
      <el-divider content-position="left">项目信息</el-divider>
      <el-form-item label="项目ID">
        <el-input
          v-model.number="createTechnologyAwardApplyFormData.projectID"
          type="number"
          placeholder="关联的项目ID"
        ></el-input>
      </el-form-item>
      <el-form-item>
        <el-button type="primary" @click="searchProjectInfo">搜索</el-button>
      </el-form-item>
      <el-form-item label="项目名称">
        <span>{{ zeroWithDefault(projectInfo?.name, '无') }}</span>
      </el-form-item>
      <el-form-item label="项目描述">
        <span>{{ zeroWithDefault(projectInfo?.description, '无') }}</span>
      </el-form-item>
      <el-form-item label="项目创建人">
        <span
          >{{ zeroWithDefault(projectInfo?.createrName, '无') }}&nbsp;{{
            zeroWithDefault(projectInfo?.createrUsername, '无')
          }}</span
        >
      </el-form-item>
      <el-divider v-if="projectInfo" content-position="left"
        >科技成果信息</el-divider
      >
      <div v-if="projectInfo">
        <el-form-item label="奖项等级">
          <el-select v-model="createTechnologyAwardApplyFormData.prizeLevel">
            <el-option
              v-for="item in supportedTechnologyLevels"
              :key="item"
              :value="item"
              :label="item"
            />
          </el-select>
        </el-form-item>
        <el-form-item label="申请备注">
          <el-input
            v-model="createTechnologyAwardApplyFormData.remarkMessage"
            type="textarea"
            autosize
            placeholder="申请备注信息"
          ></el-input>
        </el-form-item>
        <el-divider content-position="left">科技成果图片信息上传</el-divider>
        <el-form-item label="科技成果奖项图片:">
          <upload-image-file
            v-model="createTechnologyAwardApplyFormData.prizeImageName"
          ></upload-image-file>
        </el-form-item>
      </div>
    </el-form>
    <template #footer>
      <span class="dialog-footer">
        <el-button @click="hideTechnologyApplyDrawer">取消</el-button>
        <el-button
          type="primary"
          @click="createTechnologyAwardApplySubmit(formElem)"
          >确认</el-button
        >
      </span>
    </template>
  </el-drawer>
</template>
<style lang="less" scoped>
.project-info {
  span {
    margin: 8px 8px;
  }
}
.date-picker {
  margin: 8px 0px;
}
</style>
