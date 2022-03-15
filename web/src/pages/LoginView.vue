<script setup lang="ts">
import { reactive, ref } from 'vue';
import { useRouter } from 'vue-router';
import { requiredWithLength, FormInstance } from '../utils/validateRule';
import { login } from '../service/user';

const router = useRouter();

const loginFormElem = ref<FormInstance>();

const loginFormRules = reactive({
  username: requiredWithLength('用户名'),
  password: requiredWithLength('密码'),
});

let loginFormData = reactive({
  username: '',
  password: '',
});

// 登录提交处理函数
const loginSubmit = (elem: FormInstance | undefined) => {
  if (!elem) {
    return;
  }
  elem.validate(async (valid) => {
    if (valid) {
      let loginRes = await login(
        loginFormData.username,
        loginFormData.password
      );
      if (typeof loginRes === 'string') {
        ElMessage({
          type: 'error',
          message: loginRes,
        });
      } else {
        ElMessage({
          type: 'success',
          message: '登录成功',
        });
        // 跳转
        router.push({
          path: '/',
        });
      }
    }
    return valid;
  });
};
</script>
<template>
  <el-row justify="center" class="form-row">
    <el-col :lg="6" :sm="20">
      <el-form
        ref="loginFormElem"
        :rules="loginFormRules"
        :model="loginFormData"
      >
        <el-form-item label="用户名" prop="username">
          <el-input
            v-model="loginFormData.username"
            type="text"
            placeholder="用户名"
          ></el-input>
        </el-form-item>
        <el-form-item label="密码" prop="password">
          <el-input
            v-model="loginFormData.password"
            type="password"
            placeholder="密码"
            :show-password="true"
          ></el-input>
        </el-form-item>
        <el-form-item>
          <el-button type="primary" @click="loginSubmit(loginFormElem)"
            >登录</el-button
          >
        </el-form-item>
      </el-form>
    </el-col>
  </el-row>
</template>
<style lang="less" scoped>
.form-row {
  padding-top: 150px;
}
</style>
