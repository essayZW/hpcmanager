<script setup lang="ts">
import { reactive, ref } from 'vue';
import dayjs from 'dayjs';
import {
  paginationGetPaperApply,
  checkPaperAwardApply,
} from '../service/award';
import { PaperApply } from '../api/award';
import { isAdmin } from '../service/user';

// 表格数据
const tableData = reactive<{
  count: number;
  data: PaperApply[];
  loading: boolean;
}>({
  count: 0,
  data: [],
  loading: false,
});

// 加载表格数据
const loadTableData = async (pageIndex: number, pageSize: number) => {
  tableData.loading = true;
  try {
    const paginationData = await paginationGetPaperApply(pageIndex, pageSize);
    tableData.data = paginationData.Data;
    tableData.count = paginationData.Count;
  } catch (error) {
    ElMessage({
      type: 'error',
      message: `${error}`,
    });
  }
  tableData.loading = false;
};

// 分页信息
const paginationInfo = reactive<{
  pageIndex: number;
  pageSize: number;
}>({
  pageIndex: 1,
  pageSize: 5,
});

// 刷新表格当前页面的信息
const refreshTableData = () => {
  loadTableData(paginationInfo.pageIndex, paginationInfo.pageSize);
};

refreshTableData();

defineExpose({
  refreshTableData,
});

const handleCurrentChange = (pageIndex: number) => {
  paginationInfo.pageIndex = pageIndex;
  refreshTableData();
};

const handleSizeChange = (pageSize: number) => {
  paginationInfo.pageSize = pageSize;
  refreshTableData();
};

const showPaperApplyDetail = (row: PaperApply) => {
  dialogInfo.value = row;
  imagePreviewList.value[0] = row.paperFirstPageImageName;
  imagePreviewList.value[1] = row.paperThanksPageImageName;
  showDialog();
};

const dialogInfo = ref<PaperApply>();
const imagePreviewList = ref<string[]>([]);

const dialogVisible = ref<boolean>(false);

const showDialog = () => {
  dialogVisible.value = true;
};

const hideDialog = () => {
  dialogVisible.value = false;
};

const checkDialogVisible = ref<boolean>(false);

const checkDialogInfo = reactive<{
  paper?: PaperApply;
}>({});
const showCheckDialog = (row: PaperApply) => {
  checkDialogInfo.paper = row;
  checkDialogVisible.value = true;
};

const hideCheckDialog = () => {
  checkDialogVisible.value = false;
};

const checkDialogForm = reactive<{
  money: number;
  message: string;
}>({
  money: 0,
  message: '',
});

const checkPaperApplyFormSubmit = async (accept: boolean) => {
  if (!checkDialogInfo.paper) {
    ElMessage({
      type: 'error',
      message: '申请信息加载失败,请刷新重试',
    });
    return;
  }
  try {
    await checkPaperAwardApply(
      checkDialogInfo.paper.id,
      accept,
      checkDialogForm.money,
      checkDialogForm.message
    );
    ElMessage({
      type: 'success',
      message: '审核成功',
    });
    hideCheckDialog();
    refreshTableData();
  } catch (error) {
    ElMessage({
      type: 'error',
      message: `${error}`,
    });
  }
};
</script>
<template>
  <el-row justify="space-between" class="button-row">
    <div>
      <slot name="tool"></slot>
    </div>
    <el-button type="primary" @click="refreshTableData">
      <el-icon class="el-icon--left">
        <i-ic-round-refresh />
      </el-icon>
      刷新
    </el-button>
  </el-row>
  <el-row justify="center">
    <el-col :span="24">
      <el-table
        v-loading="tableData.loading"
        border
        table-layout="auto"
        :data="tableData.data"
      >
        <el-table-column
          label="ID"
          prop="id"
          sortable
          align="center"
        ></el-table-column>
        <el-table-column
          label="申请人学号"
          prop="createrUsername"
          align="center"
        ></el-table-column>
        <el-table-column
          label="申请人姓名"
          prop="createrName"
          align="center"
        ></el-table-column>
        <el-table-column
          label="导师工号"
          prop="tutorUsername"
          align="center"
        ></el-table-column>
        <el-table-column
          label="导师姓名"
          prop="tutorName"
          align="center"
        ></el-table-column>
        <el-table-column label="申请时间" align="center">
          <template #default="props">
            {{
              dayjs(props.row.createTimeUnix * 1000).format(
                'YYYY-MM-DD HH:mm:ss'
              )
            }}
          </template>
        </el-table-column>
        <el-table-column label="审核状态" align="center">
          <template #default="props">
            <div v-if="props.row.checkStatus == -1 && isAdmin()">
              <el-button type="primary" @click="showCheckDialog(props.row)"
                >审核</el-button
              >
            </div>
            <div v-else>
              <span v-if="props.row.checkStatus == -1">未审核</span>
              <span v-else-if="props.row.checkStatus == 1" class="green"
                >审核通过</span
              >
              <span v-else class="red">审核未通过</span>
            </div>
          </template>
        </el-table-column>
        <el-table-column label="详情" align="center">
          <template #default="props">
            <el-button
              type="primary"
              size="small"
              @click="showPaperApplyDetail(props.row)"
              >显示详情</el-button
            >
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
  <el-dialog v-model="dialogVisible" title="申请详情">
    <el-form>
      <el-form-item label="论文题目:">
        <span>{{ dialogInfo?.paperTitle }}</span>
      </el-form-item>
      <el-form-item label="论文分类:">
        <span>{{ dialogInfo?.paperCategory }}</span>
      </el-form-item>
      <el-form-item label="论文分区:">
        <span>{{ dialogInfo?.paperPartition }}</span>
      </el-form-item>
      <el-form-item label="论文首页图: ">
        <el-image
          :src="dialogInfo?.paperFirstPageImageName"
          :preview-src-list="imagePreviewList"
          :initial-index="0"
          fit="cover"
        />
      </el-form-item>
      <el-form-item label="论文致谢页图:">
        <el-image
          :src="dialogInfo?.paperThanksPageImageName"
          fit="cover"
          :preview-src-list="imagePreviewList"
          :initial-index="1"
        />
      </el-form-item>
    </el-form>
    <el-form>
      <el-form-item label="审核状态: ">
        <span v-if="dialogInfo?.checkStatus" class="red">未审核</span>
        <span v-else class="green">已审核</span>
      </el-form-item>
    </el-form>
    <el-form v-if="dialogInfo?.checkStatus"></el-form>
    <template #footer>
      <span class="dialog-footer">
        <el-button @click="hideDialog">关闭</el-button>
      </span>
    </template>
  </el-dialog>
  <el-dialog v-model="checkDialogVisible" title="审核论文奖励申请">
    <el-form>
      <el-form-item label="奖励金额: ">
        <el-input
          v-model.number="checkDialogForm.money"
          type="number"
          placeholder="审核成功后发放的奖励的金额数量"
        ></el-input>
      </el-form-item>
      <el-form-item label="审核备注: ">
        <el-input
          v-model="checkDialogForm.message"
          type="textarea"
          placeholder="审核的备注消息"
          autosize
        ></el-input>
      </el-form-item>
    </el-form>
    <template #footer>
      <span class="dialog-footer">
        <el-button type="primary" @click="checkPaperApplyFormSubmit(true)"
          >通过</el-button
        >
        <el-button type="danger" @click="checkPaperApplyFormSubmit(false)"
          >不通过</el-button
        >
        <el-button @click="hideCheckDialog">关闭</el-button>
      </span>
    </template>
  </el-dialog>
</template>
<style lang="less" scoped>
.button-row {
  margin-top: 16px;
  margin-bottom: 16px;
}
.pagination-row {
  margin: 16px 0px;
  .pagination-control-panel {
    margin: 0px auto;
    justify-content: center;
  }
}

.table-expand-area {
  padding-left: 12px;
  .info {
    padding-left: 16px;
    span {
      margin: 8px 8px;
    }
  }
}

.red {
  color: red;
}
.green {
  color: green;
}

.check-card {
  p {
    padding-left: 12px;
  }
  .box-title {
    padding-left: 0px;
    font-size: 16px;
  }
}
.operation-button-area {
  display: flex;
  justify-content: center;
}
.operation-button {
  display: inline-block;
}
.operation-button + span {
  display: none;
}
</style>
