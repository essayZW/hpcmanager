<script setup lang="ts">
import { reactive, ref } from 'vue';
import { getUserInfoFromStorage } from '../../service/user';
import {
  getPaperCategory,
  getPaperPartition,
  createPaperApply,
} from '../../service/award';

import PageTitle from '../PageTitle.vue';
import PaperAwardApplyTable from '../PaperAwardApplyTable.vue';
import UploadImageFile from '../UploadImageFile.vue';
import { FormInstance } from '../../utils/validateRule';

const tableElem = ref<typeof PaperAwardApplyTable | null>(null);

const refreshTableData = () => {
  if (!tableElem.value) {
    return;
  }
  tableElem.value.refreshTableData();
};

const paperApplyDialogShowFlag = ref<boolean>(false);

const hidePaperApplyDrawer = () => {
  paperApplyDialogShowFlag.value = false;
};

const showPaperApplyDrawer = () => {
  paperApplyDialogShowFlag.value = true;
};

const supportedPaperCategory = ref<string[]>(getPaperCategory());
const supportedPaperPartition = ref<string[]>(getPaperPartition());
// 创建论文奖励申请的表单数据
const createPaperAwardApplyFormData = reactive<{
  title: string;
  category: string;
  partition: string;
  firstPageImageName: string;
  thanksPageImageName: string;
  remarkMessage: string;
}>({
  title: '',
  category: supportedPaperCategory.value[0],
  partition: supportedPaperPartition.value[0],
  firstPageImageName: '',
  thanksPageImageName: '',
  remarkMessage: '',
});

const formRules = {
  title: [{ required: true, message: '论文标题不能为空', trigger: 'blur' }],
};

const formElem = ref<FormInstance | undefined>(undefined);

const createPaperAwardApplySubmit = (elem: FormInstance | undefined) => {
  if (!elem) {
    return;
  }
  elem.validate(async (valid) => {
    if (!valid) {
      return;
    }
    if (!createPaperAwardApplyFormData.firstPageImageName) {
      ElMessage({
        type: 'error',
        message: '请上传首页图片',
      });
      return;
    }
    if (!createPaperAwardApplyFormData.thanksPageImageName) {
      ElMessage({
        type: 'error',
        message: '请上传致谢页图片',
      });
      return;
    }
    try {
      await createPaperApply(
        createPaperAwardApplyFormData.title,
        createPaperAwardApplyFormData.category,
        createPaperAwardApplyFormData.partition,
        createPaperAwardApplyFormData.firstPageImageName,
        createPaperAwardApplyFormData.thanksPageImageName,
        createPaperAwardApplyFormData.remarkMessage
      );
      ElMessage({
        type: 'success',
        message: '创建成功',
      });
      refreshTableData();
      // 由于文件上传的组件不太好用,所以这里直接跳转页面对文件上传组件进行重置
      // eslint-disable-next-line no-self-assign
      window.location.href = window.location.href;
      hidePaperApplyDrawer();
    } catch (error) {
      ElMessage({
        type: 'error',
        message: `${error}`,
      });
    }
  });
};

// 如果当前用户没有用户组则不显示申请论文奖励按钮
const hasGroup = ref<boolean>(true);
const userInfo = getUserInfoFromStorage();

if (userInfo) {
  if (!userInfo.GroupId) {
    hasGroup.value = false;
  }
}
</script>
<template>
  <page-title title="论文奖励申请" des="论文奖励申请管理"></page-title>
  <paper-award-apply-table ref="tableElem">
    <template #tool>
      <el-button v-if="hasGroup" type="primary" @click="showPaperApplyDrawer"
        >申请论文奖励</el-button
      >
    </template>
  </paper-award-apply-table>
  <el-drawer v-model="paperApplyDialogShowFlag" size="40%" title="申请论文奖励">
    <el-form
      ref="formElem"
      :model="createPaperAwardApplyFormData"
      :rules="formRules"
    >
      <el-divider content-position="left">论文信息</el-divider>
      <el-form-item label="论文题目" prop="title">
        <el-input
          v-model="createPaperAwardApplyFormData.title"
          type="text"
          placeholder="论文的题目"
        ></el-input>
      </el-form-item>
      <el-form-item label="论文分类">
        <el-select v-model="createPaperAwardApplyFormData.category">
          <el-option
            v-for="item in supportedPaperCategory"
            :key="item"
            :value="item"
            :label="item"
          />
        </el-select>
      </el-form-item>
      <el-form-item label="论文分区">
        <el-select v-model="createPaperAwardApplyFormData.partition">
          <el-option
            v-for="item in supportedPaperPartition"
            :key="item"
            :value="item"
            :label="item"
          />
        </el-select>
      </el-form-item>
      <el-form-item label="申请备注">
        <el-input
          v-model="createPaperAwardApplyFormData.remarkMessage"
          type="textarea"
          autosize
          placeholder="申请备注信息"
        ></el-input>
      </el-form-item>
      <el-divider content-position="left">论文图片信息上传</el-divider>
      <el-form-item label="论文首页图片 ">
        <upload-image-file
          v-model="createPaperAwardApplyFormData.firstPageImageName"
        ></upload-image-file>
      </el-form-item>
      <el-form-item label="论文致谢页图片">
        <upload-image-file
          v-model="createPaperAwardApplyFormData.thanksPageImageName"
        ></upload-image-file>
      </el-form-item>
    </el-form>
    <template #footer>
      <span class="dialog-footer">
        <el-button @click="hidePaperApplyDrawer">取消</el-button>
        <el-button type="primary" @click="createPaperAwardApplySubmit(formElem)"
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
