<script setup lang="ts">
import { ref, onMounted } from 'vue';

import PageTitle from '../PageTitle.vue';
import ApplyGroupTable from '../ApplyGroupTable.vue';

const tableElem = ref<typeof ApplyGroupTable | null>(null);

onMounted(() => {
  if (!tableElem.value) {
    return;
  }
  tableElem.value.refreshTableData();
});
</script>
<template>
  <page-title
    title="用户组申请管理"
    des="查看并审核用户加入用户组的申请"
  ></page-title>
  <el-row justify="center">
    <el-col :span="24">
      <apply-group-table ref="tableElem"></apply-group-table>
    </el-col>
  </el-row>
</template>
<style lang="less"></style>
