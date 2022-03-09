<script setup lang="ts">
import { reactive, ref } from 'vue';
import { install } from '../service/sys';
import { useRouter } from 'vue-router';
import { requiredWithLength, FormInstance } from '../utils/validateRule';
import { InstallRequest } from '../api/sys';

const router = useRouter();

const formElem = ref<FormInstance>();
// 提交登录表单
let adminFormData = reactive<InstallRequest>({
  username: '',
  password: '',
  name: '',
  tel: '',
  email: '',
  collegeName: '',
});

const adminFormRules = reactive({
  username: requiredWithLength('用户名'),
  password: requiredWithLength('密码'),
  name: requiredWithLength('姓名', 2, 32),
});

function submit(elem: FormInstance | undefined) {
  if (!elem) {
    return;
  }
  elem.validate(async (valid) => {
    if (valid) {
      let { status, message } = await install(adminFormData as InstallRequest);
      if (!status) {
        ElMessage({
          type: 'error',
          message: message,
        });
      } else {
        ElMessage({
          type: 'success',
          message: '初始化成功,请登录',
        });
        // 跳转到登录
        router.push({
          path: '/login',
        });
      }
    }
    return valid;
  });
}
</script>
<template>
  <el-row justify="center">
    <el-col :span="12" class="title">
      <h1>计算平台管理系统初始化</h1>
    </el-col>
  </el-row>
  <el-row justify="center">
    <el-col :lg="8" :sm="22">
      <el-form
        ref="formElem"
        class="form-area"
        :model="adminFormData"
        :rules="adminFormRules"
      >
        <el-form-item label="帐号" prop="username">
          <el-input
            v-model="adminFormData.username"
            type="text"
            placeholder="管理员用户名(工号)"
          ></el-input
        ></el-form-item>
        <el-form-item label="密码" prop="password">
          <el-input
            v-model="adminFormData.password"
            type="password"
            placeholder="管理员密码"
            :show-password="true"
          >
          </el-input>
        </el-form-item>
        <el-form-item label="姓名" prop="name">
          <el-input
            v-model="adminFormData.name"
            type="text"
            placeholder="管理员姓名"
          ></el-input>
        </el-form-item>
        <el-form-item label="电话">
          <el-input
            v-model="adminFormData.tel"
            type="text"
            placeholder="管理员电话"
          >
          </el-input>
        </el-form-item>
        <el-form-item label="邮箱">
          <el-input
            v-model="adminFormData.email"
            type="text"
            placeholder="管理员邮箱地址"
          ></el-input>
        </el-form-item>
        <el-form-item label="所属学院">
          <el-input
            v-model="adminFormData.collegeName"
            type="text"
            placeholder="学院名"
          ></el-input>
        </el-form-item>
        <el-form-item>
          <el-button
            type="primary"
            class="form-submit-button"
            @click="submit(formElem)"
            >创建管理员</el-button
          >
        </el-form-item>
      </el-form>
    </el-col>
  </el-row>
</template>
<style lang="less">
.title {
  text-align: center;
}
.form-submit-button {
  width: 100%;
}
</style>
