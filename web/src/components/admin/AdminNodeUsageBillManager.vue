<script setup lang="ts">
import PageTitle from '../PageTitle.vue';
import NodeUsageBillTable from '../NodeUsageBillTable.vue';
import GroupWeekUsageBillTable from '../GroupWeekUsageTable.vue';
</script>
<template>
  <page-title
    title="机器时长账单管理"
    des="查看并管理机器时长周账单信息,以及对账单进行缴费"
  ></page-title>
  <el-tabs type="card">
    <el-tab-pane label="机时记录">
      <node-usage-bill-table></node-usage-bill-table>
    </el-tab-pane>
    <el-tab-pane label="待缴费记录">
      <group-week-usage-bill-table></group-week-usage-bill-table>
    </el-tab-pane>
  </el-tabs>
</template>
<style lang="less" scoped></style>
