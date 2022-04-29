<script lang="ts" setup>
import { onMounted, reactive } from 'vue';
import { getCasConfig, setCasConfig } from '../../service/sys';
import PageTitle from '../PageTitle.vue';

const loginSettingsForm = reactive<{
  enable: boolean;
  authServer: string;
}>({
  enable: false,
  authServer: '',
});

onMounted(async () => {
  const config = await getCasConfig();
  if (!config) {
    ElMessage({
      type: 'error',
      message: '查询CAS配置失败',
    });
    return;
  }

  loginSettingsForm.authServer = config.AuthServer;
  loginSettingsForm.enable = config.Enable;
});

const submitLoginConfig = async () => {
  try {
    await setCasConfig(loginSettingsForm.enable, loginSettingsForm.authServer);
    ElMessage({
      type: 'success',
      message: '修改设置成功',
    });
  } catch (error) {
    ElMessage({
      type: 'error',
      message: `${error}`,
    });
  }
};
</script>
<template>
  <page-title title="系统设置" des="修改系统的设置项"></page-title>
  <el-row justify="center">
    <el-col :lg="12" :sm="20">
      <el-divider content-position="left">登录设置</el-divider>
      <el-form>
        <el-form-item label="统一身份认证: ">
          <el-switch v-model="loginSettingsForm.enable"></el-switch>
        </el-form-item>
        <el-form-item
          v-if="loginSettingsForm.enable"
          label="统一身份认证服务器地址: "
        >
          <el-input
            v-model="loginSettingsForm.authServer"
            type="text"
            placeholder="统一身份认证服务器的接口地址"
          ></el-input>
        </el-form-item>
        <el-form-item>
          <el-button type="primary" @click="submitLoginConfig">确认</el-button>
        </el-form-item>
      </el-form>
    </el-col>
  </el-row>
</template>
<style lang="less"></style>
