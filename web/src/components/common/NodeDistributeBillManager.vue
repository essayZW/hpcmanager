<script setup lang="ts">
import { ref, reactive } from 'vue';
import { isSuperAdmin } from '../../service/user';

import PageTitle from '../PageTitle.vue';
import NodeDistributeBillTable from '../NodeDistributeBillTable.vue';
import { NodeDistributeFeeRate } from '../../api/fee';
import {
  getNodeDistributeFeeRate,
  setNodeDistributeFeeRate,
} from '../../service/fee';

const feeRateDialogVisible = ref<boolean>(false);
const feeRateInfoForm = reactive<NodeDistributeFeeRate>({
  rate36CPU: 0,
  rate8GPU: 0,
  rate4GPU: 0,
});

const showFeeRateDialog = async () => {
  try {
    const data = await getNodeDistributeFeeRate();
    feeRateInfoForm.rate4GPU = data.rate4GPU;
    feeRateInfoForm.rate8GPU = data.rate8GPU;
    feeRateInfoForm.rate36CPU = data.rate36CPU;
  } catch (error) {
    ElMessage({
      type: 'error',
      message: `${error}`,
    });
    return;
  }
  feeRateDialogVisible.value = true;
};

const hideFeeRateDialog = () => {
  feeRateDialogVisible.value = false;
};

const setNodeDistributeHandler = async () => {
  try {
    await setNodeDistributeFeeRate(
      feeRateInfoForm.rate36CPU,
      feeRateInfoForm.rate4GPU,
      feeRateInfoForm.rate8GPU
    );
    ElMessage({
      type: 'success',
      message: '修改费率成功',
    });
    hideFeeRateDialog();
  } catch (error) {
    ElMessage({
      type: 'error',
      message: `${error}`,
    });
  }
};
</script>
<template>
  <page-title title="机器独占账单管理" des="查看机器独占账单"></page-title>
  <node-distribute-bill-table>
    <template #tool>
      <el-button v-if="isSuperAdmin()" type="primary" @click="showFeeRateDialog"
        >修改费率</el-button
      >
    </template>
  </node-distribute-bill-table>
  <el-dialog
    v-if="isSuperAdmin()"
    v-model="feeRateDialogVisible"
    title="修改机器节点独占费率"
  >
    <el-form>
      <el-form-item label="36核心节点">
        <el-input v-model.number="feeRateInfoForm.rate36CPU" type="number">
          <template #append>元/年</template>
        </el-input>
      </el-form-item>
      <el-form-item label="4 GPU卡节点">
        <el-input v-model.number="feeRateInfoForm.rate4GPU" type="number">
          <template #append>元/年</template>
        </el-input>
      </el-form-item>
      <el-form-item label="8 GPU卡节点">
        <el-input v-model.number="feeRateInfoForm.rate8GPU" type="number">
          <template #append>元/年</template>
        </el-input>
      </el-form-item>
    </el-form>
    <template #footer>
      <span class="dialog-footer">
        <el-button @click="hideFeeRateDialog">取消</el-button>
        <el-button type="primary" @click="setNodeDistributeHandler"
          >修改费率</el-button
        >
      </span>
    </template>
  </el-dialog>
</template>
<style lang="less"></style>
