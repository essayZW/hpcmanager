<script setup lang="ts">
import PageTitle from '../PageTitle.vue';
import NodeUsageBillTable from '../NodeUsageBillTable.vue';
</script>
<template>
  <page-title
    title="机器时长账单管理"
    des="查看并管理机器时长周账单信息,以及对账单进行缴费"
  ></page-title>
  <node-usage-bill-table></node-usage-bill-table>
</template>
<style lang="less" scoped></style>
