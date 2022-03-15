<script setup lang="ts">
import { onMounted, reactive, ref } from 'vue';
import { useRouter } from 'vue-router';
import {
  getUserInfoById,
  getUserInfoFromStorage,
  updateUserInfoByID,
} from '../service/user';

const router = useRouter();
// 页面加载动画的控制
const loading = ref(true);

// 更新用户信息表单数据
const updateUserInfoFormData = reactive<{
  tel: string;
  email: string;
  college: string;
  id: number;
}>({
  tel: '',
  email: '',
  college: '',
  id: 0,
});
onMounted(async () => {
  const userInfo = getUserInfoFromStorage();
  if (!userInfo) {
    ElMessage({
      type: 'error',
      message: '登录失效,请先登录',
    });
    router.push({
      path: '/login',
    });
    return;
  }
  const info = await getUserInfoById(userInfo.UserId);
  if (typeof info == 'string') {
    ElMessage({
      type: 'error',
      message: '用户信息查询失败,请稍候重试',
    });
    return;
  }
  loading.value = false;
  updateUserInfoFormData.college = info.college;
  updateUserInfoFormData.email = info.email;
  updateUserInfoFormData.tel = info.tel;
  updateUserInfoFormData.id = info.id;
});

// 表单提交处理函数
const formSubmitHandler = async () => {
  try {
    await updateUserInfoByID(
      updateUserInfoFormData.id,
      updateUserInfoFormData.tel,
      updateUserInfoFormData.college,
      updateUserInfoFormData.email
    );
    ElMessage({
      type: 'success',
      message: '修改成功',
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
  <el-row v-loading="loading" justify="center" class="form-row">
    <el-col :lg="8" :sm="20">
      <el-form>
        <el-form-item label="电话">
          <el-input
            v-model="updateUserInfoFormData.tel"
            type="text"
            placeholder="联系电话"
          ></el-input>
        </el-form-item>
        <el-form-item label="邮箱地址">
          <el-input
            v-model="updateUserInfoFormData.email"
            type="text"
            placeholder="邮箱地址"
          ></el-input>
        </el-form-item>
        <el-form-item label="所属学院">
          <el-input
            v-model="updateUserInfoFormData.college"
            type="text"
            placeholder="所属学院"
          ></el-input>
        </el-form-item>
        <el-form-item>
          <el-button type="primary" @click="formSubmitHandler"
            >修改信息</el-button
          >
        </el-form-item>
      </el-form>
    </el-col>
  </el-row>
</template>
<style lang="less" scoped>
.form-row {
  margin-top: 16px;
}
</style>
