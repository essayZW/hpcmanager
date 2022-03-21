<script setup lang="ts">
import { reactive, ref } from 'vue';
import { getProjectInfoByID } from '../../service/project';
import { zeroWithDefault } from '../../utils/obj';
import { createNodeApply } from '../../service/node';
import dayjs from 'dayjs';

import PageTitle from '../PageTitle.vue';
import NodeApplyTable from '../NodeApplyTable.vue';
import NodeTypeSelect from '../form/NodeTypeSelect.vue';

const tableElem = ref<typeof NodeApplyTable | null>(null);
const refreshTableData = () => {
  if (!tableElem.value) {
    return;
  }
  tableElem.value.refreshTableData();
};

const nodeApplyDialogShowFlag = ref<boolean>(false);

const hideNodeApplyDrawer = () => {
  nodeApplyDialogShowFlag.value = false;
};

const showNodeApplyDrawer = () => {
  nodeApplyDialogShowFlag.value = true;
};

// 创建机器节点申请的表单数据
const createNodeApplyFormData = reactive<{
  nodeType: string;
  nodeNum: number | string;
  startTime: number;
  endTime: number;
  projectInfo: {
    id: number | string;
    name: string;
    from: string;
    numbering: string;
    expenses: string;
    description: string;
  };
}>({
  nodeType: '',
  nodeNum: 1,
  projectInfo: {
    id: 1,
    name: '',
    from: '',
    numbering: '',
    expenses: '',
    description: '',
  },
  startTime: new Date().getTime(),
  endTime: dayjs(new Date()).add(1, 'year').valueOf(),
});

// 创建机器节点申请的表单区域加载动画标志
const createNodeFormLoading = ref<boolean>(false);
/**
 *搜索项目信息
 */
const searchProjectInfo = async () => {
  createNodeFormLoading.value = true;
  try {
    const data = await getProjectInfoByID(
      createNodeApplyFormData.projectInfo.id as number
    );
    createNodeApplyFormData.projectInfo.name = data.name;
    createNodeApplyFormData.projectInfo.from = data.from;
    createNodeApplyFormData.projectInfo.numbering = data.numbering;
    createNodeApplyFormData.projectInfo.expenses = data.expenses;
    createNodeApplyFormData.projectInfo.description = data.description;
  } catch (error) {
    ElMessage({
      type: 'error',
      message: `${error}`,
    });
  }
  createNodeFormLoading.value = false;
};

// 清除的所有的表单值
const clearForm = () => {
  createNodeApplyFormData.projectInfo.name = '';
};

const createNodeApplyFormSubmit = async () => {
  // 因为一个项目的名称必定存在所以用这个来判断项目信息是否正确
  if (!createNodeApplyFormData.projectInfo.name) {
    ElMessage({
      type: 'error',
      message: '请先点击搜索按钮搜索并确认关联的项目信息正确',
    });
    return;
  }
  createNodeFormLoading.value = true;

  try {
    await createNodeApply(
      parseInt(createNodeApplyFormData.projectInfo.id as string),
      createNodeApplyFormData.nodeType,
      parseInt(createNodeApplyFormData.nodeNum as string),
      createNodeApplyFormData.startTime,
      createNodeApplyFormData.endTime
    );
    ElMessage({
      type: 'success',
      message: '创建成功',
    });
    clearForm();
    refreshTableData();
    hideNodeApplyDrawer();
  } catch (error) {
    ElMessage({
      type: 'error',
      message: `${error}`,
    });
  }
  createNodeFormLoading.value = false;
};
</script>
<template>
  <page-title
    title="机器节点申请管理"
    des="查看、创建并审核机器节点申请"
  ></page-title>
  <node-apply-table ref="tableElem">
    <template #tool>
      <el-button type="primary" @click="showNodeApplyDrawer"
        >申请计算节点</el-button
      >
    </template>
  </node-apply-table>
  <el-drawer v-model="nodeApplyDialogShowFlag" size="40%" title="申请机器节点">
    <el-form v-loading="createNodeFormLoading" :model="createNodeApplyFormData">
      <el-divider content-position="left">节点信息</el-divider>
      <el-form-item label="机器节点类型">
        <node-type-select
          v-model="createNodeApplyFormData.nodeType"
        ></node-type-select>
      </el-form-item>
      <el-form-item label="节点数量">
        <el-input
          v-model="createNodeApplyFormData.nodeNum"
          type="number"
          :min="1"
          placeholder="申请的机器节点数量"
        ></el-input>
      </el-form-item>
      <el-form-item label="申请独占时间范围:">
        <el-date-picker
          v-model="createNodeApplyFormData.startTime"
          class="date-picker"
          placeholder="起始日期"
        ></el-date-picker>
        <el-date-picker
          v-model="createNodeApplyFormData.endTime"
          class="date-picker"
          placeholder="结束日期"
        ></el-date-picker>
      </el-form-item>
      <el-divider content-position="left">项目信息</el-divider>
      <el-form-item label="项目ID">
        <el-input
          v-model="createNodeApplyFormData.projectInfo.id"
          type="number"
          :min="1"
          placeholder="关联的项目ID"
        ></el-input>
      </el-form-item>
      <el-form-item>
        <el-button type="primary" @click="searchProjectInfo">搜索</el-button>
      </el-form-item>
      <div
        v-if="createNodeApplyFormData.projectInfo.name != ''"
        class="project-info"
      >
        <p>
          <span
            ><strong> 名称: </strong
            >{{ createNodeApplyFormData.projectInfo.name }}</span
          >
          <span
            ><strong>来源: </strong
            >{{
              zeroWithDefault(createNodeApplyFormData.projectInfo.from, '无')
            }}</span
          >
        </p>
        <p>
          <span
            ><strong>编号: </strong
            >{{
              zeroWithDefault(
                createNodeApplyFormData.projectInfo.numbering,
                '无'
              )
            }}</span
          >
          <span
            ><strong>经费: </strong
            >{{
              zeroWithDefault(
                createNodeApplyFormData.projectInfo.expenses,
                '无'
              )
            }}</span
          >
        </p>
        <p>
          <span><strong>描述: </strong></span
          >{{
            zeroWithDefault(
              createNodeApplyFormData.projectInfo.description,
              '无'
            )
          }}
        </p>
      </div>
    </el-form>
    <template #footer>
      <span class="dialog-footer">
        <el-button @click="hideNodeApplyDrawer">取消</el-button>
        <el-button type="primary" @click="createNodeApplyFormSubmit"
          >确认</el-button
        >
      </span>
    </template>
  </el-drawer>
</template>
<style lang="less">
.project-info {
  span {
    margin: 8px 8px;
  }
}
.date-picker {
  margin: 8px 0px;
}
</style>
