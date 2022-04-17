<script setup lang="ts">
import { ref } from 'vue';

import PageTitle from '../PageTitle.vue';
import NodeUsageBillTable from '../NodeUsageBillTable.vue';
import GroupWeekUsageBillTable from '../GroupWeekUsageTable.vue';

const hasPayTableElem = ref<typeof GroupWeekUsageBillTable | undefined>(
  undefined
);
const noPayTableElem = ref<typeof GroupWeekUsageBillTable | undefined>(
  undefined
);

// eslint-disable-next-line @typescript-eslint/no-unused-vars, @typescript-eslint/no-explicit-any
const tabClickHandler = (tab: any, event: Event) => {
  if (noPayTableElem.value && tab.paneName == 'showNoPay') {
    noPayTableElem.value.refreshTable();
  }
  if (hasPayTableElem.value && tab.paneName == 'showHadPay') {
    hasPayTableElem.value.refreshTable();
  }
};
</script>
<template>
  <page-title
    title="机器时长账单管理"
    des="查看并管理机器时长周账单信息,以及对账单进行缴费"
  ></page-title>
  <el-tabs type="card" class="tabs" @tab-click="tabClickHandler">
    <el-tab-pane label="机时记录">
      <node-usage-bill-table></node-usage-bill-table>
    </el-tab-pane>
    <el-tab-pane label="待缴费记录" name="showNoPay">
      <group-week-usage-bill-table
        ref="noPayTableElem"
        :pay-flag="false"
      ></group-week-usage-bill-table>
    </el-tab-pane>
    <el-tab-pane label="已缴费记录" name="showHadPay">
      <group-week-usage-bill-table
        ref="hasPayTableElem"
        :pay-flag="true"
      ></group-week-usage-bill-table>
    </el-tab-pane>
  </el-tabs>
</template>
<style lang="less" scoped>
.tabs {
  margin-top: 16px;
}
</style>
