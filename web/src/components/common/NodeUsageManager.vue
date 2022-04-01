<script setup lang="ts">
import PageTitle from '../PageTitle.vue';
import NodeUsageTable from '../NodeUsageTable.vue';
</script>
<template>
  <page-title
    title="机器时长查看"
    des="查看用户使用机器节点的时长信息"
  ></page-title>
  <node-usage-table></node-usage-table>
</template>
<style lang="less" scoped></style>
