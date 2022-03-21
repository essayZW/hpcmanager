<script setup lang="ts">
import { defineProps, onMounted, ref } from 'vue';

// select 表单拥有的值
// NOTE: 暂定的机器节点类型值
const selectValuesArr = [
  {
    value: 'cpuc36',
    label: '36核心节点',
  },
  {
    value: 'gpuc4',
    label: '4 GPU卡节点',
  },
  {
    value: 'gpuc8',
    label: '8 GPU卡节点',
  },
];

const props = defineProps<{
  modelValue: string;
}>();

onMounted(() => {
  for (const item of selectValuesArr) {
    if (item.value == props.modelValue) {
      return;
    }
  }
  emit('update:modelValue', selectValuesArr[0].value);
});
const selectedValue = ref<string>(selectValuesArr[0].value);

const emit = defineEmits(['update:modelValue']);
const updateNodeType = (value: string) => {
  emit('update:modelValue', value);
};
</script>
<template>
  <el-select
    v-model="selectedValue"
    :value="props.modelValue"
    @change="updateNodeType"
  >
    <el-option
      v-for="item in selectValuesArr"
      :key="item.value"
      :label="item.label"
      :value="item.value"
    />
  </el-select>
</template>
<style lang="less"></style>
