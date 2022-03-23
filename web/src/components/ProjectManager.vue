<script setup lang="ts">
import { reactive, ref } from 'vue';
import { ProjectInfo } from '../api/project';
import { createProject, paginationGetProjectInfo } from '../service/project';
import { FormInstance, requiredWithLength } from '../utils/validateRule';
import { zeroWithDefault } from '../utils/obj';

import PageTitle from './PageTitle.vue';

// 项目表格的数据
const tableData = reactive<{
  data: ProjectInfo[];
  count: number;
}>({
  data: [],
  count: 0,
});

// 表格当前的分页信息
const paginationInfo = reactive<{
  pageIndex: number;
  pageSize: number;
}>({
  pageIndex: 1,
  pageSize: 5,
});

// 分页加载表格的数据
const loadTableData = async (pageIndex: number, pageSize: number) => {
  try {
    const data = await paginationGetProjectInfo(pageIndex, pageSize);
    tableData.data = data.Data;
    tableData.count = data.Count;
  } catch (error) {
    ElMessage({
      type: 'error',
      message: `${error}`,
    });
  }
};

// 刷新表格的当前页的信息
const refreshTableData = () => {
  loadTableData(paginationInfo.pageIndex, paginationInfo.pageSize);
};

refreshTableData();

const handleSizeChange = (pageSize: number) => {
  paginationInfo.pageSize = pageSize;
  refreshTableData();
};
const handleCurrentChange = (pageIndex: number) => {
  paginationInfo.pageIndex = pageIndex;
  refreshTableData();
};

const createProjectDialogShow = ref<boolean>(false);

const showProjectDialog = () => {
  createProjectDialogShow.value = true;
};
const hideProjectDialog = () => {
  createProjectDialogShow.value = false;
};

// 创建项目信息表单数据
const createProjectFormData = reactive<{
  name: string;
  numbering: string;
  from: string;
  expenses: string;
  description: string;
}>({
  name: '',
  numbering: '',
  from: '',
  expenses: '',
  description: '',
});

// 表单元素的实例
const createProjectFormElem = ref<FormInstance | null>(null);
// 创建项目表单数据的验证规则
const createProjectFormRules = {
  name: requiredWithLength('项目名称', 2, 128),
};
// 创建项目表单提交处理函数
const formSubmitHandler = (elem: FormInstance | null) => {
  if (!elem) {
    return;
  }
  elem.validate(async (valid) => {
    if (!valid) {
      return false;
    }
    try {
      await createProject(
        createProjectFormData.name,
        createProjectFormData.from,
        createProjectFormData.numbering,
        createProjectFormData.expenses,
        createProjectFormData.description
      );
      ElMessage({
        type: 'success',
        message: '创建成功',
      });
      refreshTableData();
      hideProjectDialog();
    } catch (error) {
      ElMessage({
        type: 'error',
        message: `${error}`,
      });
    }
    return valid;
  });
};
</script>
<template>
  <page-title title="项目管理" des="项目管理、查看页面"></page-title>
  <el-row justify="space-between" class="button-area-row">
    <el-button type="primary" @click="showProjectDialog">新建项目</el-button>
    <el-button type="primary" @click="refreshTableData">
      <el-icon class="el-icon--left">
        <i-ic-round-refresh />
      </el-icon>
      刷新
    </el-button>
  </el-row>
  <el-row justify="center">
    <el-col :span="24">
      <el-table :data="tableData.data">
        <el-table-column label="ID" prop="id"></el-table-column>
        <el-table-column label="名称">
          <template #default="props">
            {{ zeroWithDefault(props.row.name, '无') }}
          </template>
        </el-table-column>
        <el-table-column label="来源">
          <template #default="props">
            {{ zeroWithDefault(props.row.from, '无') }}
          </template>
        </el-table-column>
        <el-table-column label="编号">
          <template #default="props">
            {{ zeroWithDefault(props.row.numbering, '无') }}
          </template>
        </el-table-column>
        <el-table-column label="经费">
          <template #default="props">
            {{ zeroWithDefault(props.row.expenses, '无') }}
          </template>
        </el-table-column>
        <el-table-column
          label="创建者姓名"
          prop="createrName"
        ></el-table-column>
        <el-table-column
          label="创建者学(工)号"
          prop="createrUsername"
        ></el-table-column>
        <el-table-column label="详情" type="expand">
          <template #default="props">
            <p><strong>描述</strong></p>
            <p class="info">
              {{ zeroWithDefault(props.row.description, '无') }}
            </p>
          </template>
        </el-table-column>
      </el-table>
    </el-col>
  </el-row>
  <el-row justify="center" class="pagination-row">
    <el-col :span="18">
      <el-pagination
        v-model:currentPage="paginationInfo.pageIndex"
        v-model:page-size="paginationInfo.pageSize"
        class="pagination-control-panel"
        :page-sizes="[5, 10, 25, 50]"
        background
        layout="total, sizes, prev, pager, next, jumper"
        :total="tableData.count"
        :hide-on-single-page="true"
        @size-change="handleSizeChange"
        @current-change="handleCurrentChange"
      >
      </el-pagination>
    </el-col>
  </el-row>
  <el-dialog v-model="createProjectDialogShow" title="创建项目">
    <el-form
      ref="createProjectFormElem"
      :rules="createProjectFormRules"
      :model="createProjectFormData"
    >
      <el-form-item label="项目名称" prop="name">
        <el-input
          v-model="createProjectFormData.name"
          type="text"
          placeholder="项目名称"
        ></el-input>
      </el-form-item>
      <el-form-item label="项目来源">
        <el-input
          v-model="createProjectFormData.from"
          type="text"
          placeholder="项目来源"
        ></el-input>
      </el-form-item>
      <el-form-item label="项目编号">
        <el-input
          v-model="createProjectFormData.numbering"
          type="text"
          placeholder="项目编号"
        ></el-input>
      </el-form-item>
      <el-form-item label="项目经费">
        <el-input
          v-model="createProjectFormData.expenses"
          type="text"
          placeholder="项目经费"
        ></el-input>
      </el-form-item>
      <el-form-item label="项目描述">
        <el-input
          v-model="createProjectFormData.description"
          type="textarea"
          placeholder="项目描述"
          :autosize="true"
        ></el-input>
      </el-form-item>
    </el-form>
    <template #footer>
      <span class="dialog-footer">
        <el-button @click="hideProjectDialog">取消</el-button>
        <el-button
          type="primary"
          @click="formSubmitHandler(createProjectFormElem)"
          >确认</el-button
        >
      </span>
    </template>
  </el-dialog>
</template>
<style lang="less" scoped>
.button-area-row {
  margin-top: 16px;
  margin-bottom: 8px;
}
.pagination-row {
  margin: 16px 0px;
  .pagination-control-panel {
    margin: 0px auto;
    justify-content: center;
  }
}
.info {
  padding-left: 16px;
}
</style>
