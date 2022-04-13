<script setup lang="ts">
import { reactive, ref } from 'vue';
import { NodeDistributeBill, NodeDistributeFeeRate } from '../api/fee';
import {
  paginationGetNodeDistributeBill,
  payNodeDistributeBill,
  payTypeToString,
} from '../service/fee';
import { isAdmin } from '../service/user';
import { NodeApplyInfo } from '../api/node';
import { getNodeApplyByID, nodeTypeToName } from '../service/node';
import dayjs from 'dayjs';

const tableData = reactive<{
  data: NodeDistributeBill[];
  count: number;
}>({
  data: [],
  count: 0,
});

// 加载表格数据
const loadTableData = async (pageIndex: number, pageSize: number) => {
  try {
    const data = await paginationGetNodeDistributeBill(pageIndex, pageSize);
    tableData.data = data.Data;
    tableData.count = data.Count;
  } catch (error) {
    ElMessage({
      message: `${error}`,
      type: 'error',
    });
  }
};

// 分页信息
const paginationInfo = reactive<{
  pageIndex: number;
  pageSize: number;
}>({
  pageIndex: 1,
  pageSize: 5,
});

const refreshTableData = () => {
  loadTableData(paginationInfo.pageIndex, paginationInfo.pageSize);
};

refreshTableData();

const handleSizeChange = (size: number) => {
  paginationInfo.pageSize = size;
  refreshTableData();
};

const handleCurrentChange = (index: number) => {
  paginationInfo.pageIndex = index;
  refreshTableData();
};

const payBillRow = ref<NodeDistributeBill>();
const handlerPayButton = (row: NodeDistributeBill) => {
  payBillRow.value = row;
  showPayBillDialog();
};

const payBillDialog = ref<boolean>(false);

const payBillDialogInfo = reactive<{
  rateInfo?: NodeDistributeFeeRate;
  applyInfo?: NodeApplyInfo;
}>({});

const showPayBillDialog = async () => {
  if (!payBillRow.value) {
    ElMessage({
      type: 'error',
      message: '获取账单信息失败,请稍后重试',
    });
    return;
  }
  try {
    const data = await getNodeApplyByID(payBillRow.value.id);
    payBillDialogInfo.applyInfo = data;
  } catch (error) {
    ElMessage({
      type: 'error',
      message: `${error}`,
    });
  }
  payBillDialog.value = true;
};

const hidePayBillDialog = () => {
  payBillDialog.value = false;
};

const payBillForm = reactive<{
  message: string;
  payFee: number;
}>({
  message: '',
  payFee: 0,
});

const handlerPayBillSubmit = async (isBalance: boolean) => {
  if (!payBillRow.value) {
    ElMessage({
      type: 'error',
      message: '获取账单信息失败',
    });
    return;
  }
  if (payBillForm.payFee < 0) {
    ElMessage({
      type: 'error',
      message: '缴费金额不能为负数',
    });
    return;
  }
  try {
    await payNodeDistributeBill(
      payBillRow.value.id,
      payBillForm.payFee,
      isBalance,
      payBillForm.message
    );
    ElMessage({
      type: 'success',
      message: '缴费成功',
    });
    refreshTableData();
    hidePayBillDialog();
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
      <el-table border table-layout="auto" :data="tableData.data">
        <el-table-column label="ID" align="center" prop="id"></el-table-column>
        <el-table-column
          label="工单ID"
          align="center"
          prop="applyID"
        ></el-table-column>
        <el-table-column label="应缴费用" align="center">
          <template #default="props"> {{ props.row.fee }}元 </template>
        </el-table-column>
        <el-table-column label="实缴费用" align="center">
          <template #default="props">
            <span v-if="props.row.payFlag">{{ props.row.fee }}元</span>
            <span v-else>未缴费</span>
          </template>
        </el-table-column>
        <el-table-column label="缴费方式" align="center">
          <template #default="props">
            <span v-if="props.row.payFlag">{{
              payTypeToString(props.row.payType)
            }}</span>
            <span v-else>未缴费</span>
          </template>
        </el-table-column>
        <el-table-column
          label="用户ID"
          align="center"
          prop="userID"
        ></el-table-column>
        <el-table-column
          label="用户帐号"
          align="center"
          prop="userUsername"
        ></el-table-column>
        <el-table-column
          label="用户姓名"
          align="center"
          prop="userName"
        ></el-table-column>
        <el-table-column label="操作" align="center">
          <template #default="props">
            <el-button
              v-if="!props.row.payFlag && isAdmin()"
              type="primary"
              @click="handlerPayButton(props.row)"
              >缴费</el-button
            >
            <span v-else-if="!props.row.payFlag">未缴费</span>
            <span v-else class="green">已缴费</span>
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
  <el-dialog v-model="payBillDialog" title="账单缴费">
    <div class="pay-bill-dialog-body">
      <div class="rate-area">
        <!-- FIXME: 对接费率查询接口 -->
        <h3>包机费率</h3>
        <p>
          <strong>36 核心节点:</strong
          >{{ payBillDialogInfo.rateInfo?.rate36CPU }}元
        </p>
        <p>
          <strong>4 GPU核心节点:</strong
          >{{ payBillDialogInfo.rateInfo?.rate4GPU }}元
        </p>
        <p>
          <strong>8 GPU核心节点:</strong
          >{{ payBillDialogInfo.rateInfo?.rate8GPU }}元
        </p>
      </div>
      <div>
        <el-form inline>
          <el-form-item label="申请人姓名: ">
            <span>{{ payBillDialogInfo.applyInfo?.createrName }}</span>
          </el-form-item>
          <el-form-item label="申请人学号: ">
            <span>{{ payBillDialogInfo.applyInfo?.createrUsername }}</span>
          </el-form-item>
          <el-form-item label="导师姓名: ">
            <span>{{ payBillDialogInfo.applyInfo?.tutorName }}</span>
          </el-form-item>
          <el-form-item label="导师工号: ">
            <span>{{ payBillDialogInfo.applyInfo?.tutorUsername }}</span>
          </el-form-item>
          <el-form-item label="机器类型: ">
            <span>{{
              nodeTypeToName(payBillDialogInfo.applyInfo?.nodeType)
            }}</span>
          </el-form-item>
          <el-form-item label="机器数量: ">
            <span>{{ payBillDialogInfo.applyInfo?.nodeNum }}</span>
          </el-form-item>
          <el-form-item label="独占时间段: ">
            <span>
              {{
                dayjs(payBillDialogInfo.applyInfo?.startTime * 1000).format(
                  'YYYY-MM-DD'
                )
              }}至{{
                dayjs(payBillDialogInfo.applyInfo?.endTime * 1000).format(
                  'YYYY-MM-DD'
                )
              }}</span
            >
          </el-form-item>
          <el-form-item label="应缴费金额: ">
            <span>{{ payBillRow.fee }}元</span>
          </el-form-item>
          <el-form-item label="实缴费金额: ">
            <el-input
              v-model.number="payBillForm.payFee"
              type="text"
              placeholder="实际缴费的金额"
            >
              <template #append>元</template>
            </el-input>
          </el-form-item>
          <el-form-item label="缴费备注: ">
            <el-input
              v-model="payBillForm.message"
              typr="text"
              autosize
              placeholder="可以为空,此次缴费的备注,不超过500字"
            ></el-input>
          </el-form-item>
        </el-form>
      </div>
    </div>
    <template #footer>
      <span class="dialog-footer">
        <el-button @click="hidePayBillDialog">取消</el-button>
        <el-button type="primary" @click="handlerPayBillSubmit(true)"
          >余额缴费</el-button
        >
        <el-button type="primary" @click="handlerPayBillSubmit(false)"
          >线下缴费</el-button
        >
      </span>
    </template>
  </el-dialog>
</template>
<style lang="less" scoped>
.button-row {
  margin-top: 16px;
  margin-bottom: 16px;
}
.green {
  color: green;
}
.pay-bill-dialog-body {
  display: flex;
  justify-content: center;
  width: 100%;
  .rate-area {
    min-width: 30%;
  }
}
</style>
