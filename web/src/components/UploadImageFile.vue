<script lang="ts" setup>
import { UploadFile } from 'element-plus/es/components/upload/src/upload.type';
import { ref, computed, defineEmits } from 'vue';
import { fileUploadPath } from '../api/fss';
import { accessTokenKey, HTTPResponse } from '../api/api';

const props = defineProps<{
  modelValue: string;
}>();
const imageUrl = ref<string>(props.modelValue);

const uploadFileApi = computed(() => {
  const accessToken = localStorage.getItem(accessTokenKey);
  return `${fileUploadPath}?access_token=${accessToken}`;
});

const onChangeHandler = (uploadFile: UploadFile) => {
  imageUrl.value = URL.createObjectURL(uploadFile.raw);
};

const emit = defineEmits(['update:modelValue']);
const onUploadSuccess = (
  response: HTTPResponse<{
    filename: string;
  }>
) => {
  if (!response.status) {
    ElMessage({
      type: 'error',
      message: response.message,
    });
    return;
  }
  ElMessage({
    type: 'success',
    message: '图片文件上传成功',
  });
  emit('update:modelValue', response.data.filename);
};
</script>
<template>
  <el-upload
    :action="uploadFileApi"
    class="upload-image-file-area"
    :on-change="onChangeHandler"
    :on-success="onUploadSuccess"
  >
    <img v-if="imageUrl" :src="imageUrl" />
    <el-icon v-if="!imageUrl" class="plus-icon">
      <i-ic-round-plus />
    </el-icon>
    <h4 v-if="!imageUrl">论文首页图片</h4>
    <template #tip>
      <div class="el-upload__tip">只允许上传30MB以内的jpg/png/jpeg格式文件</div>
    </template>
  </el-upload>
</template>
<style lang="less" scoped>
.upload-image-file-area {
  border: 1px dashed var(--el-border-color-base);
  width: 240px;
  text-align: center;
  border-radius: 12px;
  min-height: 200px;
  img {
    width: 240px;
    height: 200px;
  }

  .plus-icon {
    margin-top: 40px;
    font-size: 60px;
  }
}
</style>
