<script setup lang="ts">
import { createUserByAdmin, isAdmin } from '../../service/user';
import { FormInstance, requiredWithLength } from '../../utils/validateRule';
import { reactive, ref } from 'vue';

import PageTitle from '../PageTitle.vue';
import UserTable from '../UserTable.vue';

const tableElem = ref<typeof UserTable | null>(null);

const refreshTable = () => {
  if (tableElem.value) {
    tableElem.value.refreshTable();
  }
};

const createUserDialogShowFlag = ref<boolean>(false);

const hideCreateUserDialog = () => {
  createUserDialogShowFlag.value = false;
};

const showCreateUserDialog = () => {
  createUserDialogShowFlag.value = true;
};

// 创建用户的表单数据
const createUserFormData = reactive<{
  username: string;
  name: string;
  password: string;
  tel: string;
  email: string;
  college: string;
  groupId: number | string;
}>({
  username: '',
  name: '',
  password: '',
  tel: '',
  email: '',
  college: '',
  groupId: 0,
});
const clearForm = () => {
  createUserFormData.username = '';
  createUserFormData.password = '';
  createUserFormData.name = '';
  createUserFormData.tel = '';
  createUserFormData.email = '';
  createUserFormData.college = '';
  createUserFormData.groupId = 0;
};
const createUserFormElem = ref<FormInstance>();

const createUserFormRules = {
  username: requiredWithLength('学号', 6, 32),
  name: requiredWithLength('姓名', 2, 32),
  password: requiredWithLength('密码', 6, 16),
  tel: [
    {
      pattern: new RegExp(/^1[3-9]\d{9}$/),
      message: '电话号码格式错误',
      trigger: 'change',
    },
  ],
  email: [
    {
      pattern: new RegExp(/^\w+([-+.]\w+)*@\w+([-.]\w+)*\.\w+([-.]\w+)*$/),
      message: '邮箱格式错误',
      trigger: 'change',
    },
  ],
};

const createUserFormSubmitHandler = (elem: FormInstance | undefined) => {
  if (!elem) {
    return;
  }
  elem.validate(async (valid) => {
    if (valid) {
      try {
        const id = await createUserByAdmin({
          username: createUserFormData.username,
          password: createUserFormData.password,
          name: createUserFormData.name,
          tel: createUserFormData.tel,
          email: createUserFormData.email,
          collegeName: createUserFormData.college,
          groupID: parseInt(createUserFormData.groupId as string),
        });
        ElMessage({
          type: 'success',
          message: `创建成功,新用户ID : ${id}`,
        });
        refreshTable();
        hideCreateUserDialog();
        clearForm();
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
  <page-title title="用户管理" des="查看所有的用户信息"></page-title>
  <user-table ref="tableElem">
    <template #tool>
      <el-button v-if="isAdmin()" type="primary" @click="showCreateUserDialog"
        >新建用户</el-button
      >
    </template>
  </user-table>
  <el-dialog
    v-if="isAdmin()"
    v-model="createUserDialogShowFlag"
    title="新建用户"
  >
    <el-form
      ref="createUserFormElem"
      :model="createUserFormData"
      :rules="createUserFormRules"
    >
      <el-form-item label="学号" prop="username">
        <el-input
          v-model="createUserFormData.username"
          type="text"
          placeholder="用户的学号"
        ></el-input>
      </el-form-item>
      <el-form-item label="姓名" prop="name">
        <el-input
          v-model="createUserFormData.name"
          type="text"
          placeholder="用户的姓名"
        ></el-input>
      </el-form-item>
      <el-form-item label="密码" prop="password">
        <el-input
          v-model="createUserFormData.password"
          type="password"
          placeholder="用户密码"
        ></el-input>
      </el-form-item>
      <el-form-item label="电话" prop="tel">
        <el-input
          v-model="createUserFormData.tel"
          type="text"
          placeholder="电话号码"
        ></el-input>
      </el-form-item>
      <el-form-item label="邮箱地址" prop="email">
        <el-input
          v-model="createUserFormData.email"
          type="email"
          placeholder="邮箱地址"
        ></el-input>
      </el-form-item>
      <el-form-item label="所属学院" prop="college">
        <el-input
          v-model="createUserFormData.college"
          type="text"
          placeholder="所属的学院名称"
        ></el-input>
      </el-form-item>
      <el-alert
        title="如果用户组ID为0,则创建一个用户,而不加入任何用户组,需要另行申请加入用户组"
        type="info"
        effect="dark"
        :closable="false"
        class="alert-info"
      />
      <el-form-item label="用户组ID" prop="groupID">
        <el-input
          v-model="createUserFormData.groupId"
          type="number"
          :min="1"
          placeholder="需要加入的用户组的ID"
        ></el-input>
      </el-form-item>
    </el-form>
    <template #footer>
      <span class="dialog-footer">
        <el-button @click="hideCreateUserDialog">取消</el-button>
        <el-button
          type="primary"
          @click="createUserFormSubmitHandler(createUserFormElem)"
          >确认</el-button
        >
      </span>
    </template>
  </el-dialog>
</template>
<style lang="less" scoped>
.alert-info {
  margin-bottom: 8px;
}
</style>
