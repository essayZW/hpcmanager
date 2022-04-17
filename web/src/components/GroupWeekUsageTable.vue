<script lang="ts" setup>
import { reactive, ref, defineExpose, onBeforeMount } from 'vue';
import {
  getNodeUsageFeeRate,
  paginationGetGroupNodeUsageBill,
  payGroupNodeUsageBills,
} from '../service/fee';
import { NodeWeekUsageBillForGroup, nodeUsageFeeRate } from '../api/fee';
import { timeSecondFormat, zeroWithDefault } from '../utils/obj';
import { GroupInfo } from '../api/group';
import { getGroupInfoByID } from '../service/group';

const propsParam = defineProps<{
  payFlag: boolean;
}>();

const tableData = reactive<{
  data: NodeWeekUsageBillForGroup[];
  count: number;
  loading: boolean;
}>({
  data: [],
  count: 0,
  loading: false,
});

// 加载表格某一页的数据
const loadTableData = async (pageIndex: number, pageSize: number) => {
  tableData.loading = true;
  try {
    const data = await paginationGetGroupNodeUsageBill(
      pageIndex,
      pageSize,
      propsParam.payFlag ? true : false
    );
    tableData.data = data.Data;
    tableData.count = data.Count;
  } catch (error) {
    ElMessage({
      type: 'error',
      message: `${error}`,
    });
  }
  tableData.loading = false;
};

const paginationInfo = reactive<{
  pageIndex: number;
  pageSize: number;
}>({
  pageIndex: 1,
  pageSize: 10,
});

const refreshTableData = () => {
  loadTableData(paginationInfo.pageIndex, paginationInfo.pageSize);
};

defineExpose({
  refreshTable: refreshTableData,
});
const handleCurrentChange = (pageIndex: number) => {
  paginationInfo.pageIndex = pageIndex;
  refreshTableData();
};

const handleSizeChange = (pageSize: number) => {
  paginationInfo.pageSize = pageSize;
  refreshTableData();
};

const rowExpandInfo = reactive<{
  [id: number]: {
    groupInfo?: GroupInfo;
  };
}>({});

const rowExpandHandler = async (row: NodeWeekUsageBillForGroup) => {
  if (!rowExpandInfo[row.userGroupID]) {
    rowExpandInfo[row.userGroupID] = {};
  }
  if (!rowExpandInfo[row.userGroupID].groupInfo) {
    try {
      const info = await getGroupInfoByID(row.userGroupID);
      rowExpandInfo[row.userGroupID].groupInfo = info;
    } catch (error) {
      ElMessage({
        type: 'error',
        message: `${error}`,
      });
    }
  }
};

const payBillDialogVisibleFlag = ref<boolean>(false);

const hideBillDialog = () => {
  payBillDialogVisibleFlag.value = false;
};

const showBillDialog = () => {
  payBillDialogVisibleFlag.value = true;
};

const payBillInfo = ref<NodeWeekUsageBillForGroup | undefined>(undefined);
const payBillGroupInfo = ref<GroupInfo | undefined>(undefined);

const handlerPayButtonClick = async (row: NodeWeekUsageBillForGroup) => {
  payBillInfo.value = row;
  payBillForm.payFee = row.fee;
  payBillForm.payMessage = '';
  try {
    const info = await getGroupInfoByID(payBillInfo.value.userGroupID);
    payBillGroupInfo.value = info;
  } catch (error) {
    ElMessage({
      type: 'error',
      message: `${error}`,
    });
  }
  showBillDialog();
};

const payBillForm = reactive<{
  payFee: number;
  payMessage: string;
}>({
  payFee: 0,
  payMessage: '',
});

const handlerPayBillSubmit = async (balancePay: boolean) => {
  if (!payBillInfo.value) {
    return;
  }
  try {
    const count = await payGroupNodeUsageBills(
      payBillInfo.value?.userGroupID,
      balancePay,
      payBillForm.payMessage,
      payBillInfo.value?.fee
    );
    ElMessage({
      type: 'success',
      message: `成功支付了${count}条账单记录`,
    });
    refreshTableData();
    hideBillDialog();
  } catch (error) {
    ElMessage({
      type: 'error',
      message: `${error}`,
    });
  }
};

const feeRateInfo = ref<nodeUsageFeeRate | undefined>(undefined);

onBeforeMount(async () => {
  try {
    const data = await getNodeUsageFeeRate();
    feeRateInfo.value = data;
  } catch (error) {
    ElMessage({
      type: 'error',
      message: `${error}`,
    });
  }
});
</script>
<template>
  <el-row justify="end" class="refresh-button-row">
    <el-button type="primary" @click="refreshTableData">
      <el-icon class="el-icon--left">
        <i-ic-round-refresh />
      </el-icon>
      刷新</el-button
    >
  </el-row>
  <el-row justify="center">
    <el-col :span="24">
      <el-table
        v-loading="tableData.loading"
        border
        table-layout="auto"
        :data="tableData.data"
        @expand-change="rowExpandHandler"
      >
        <el-table-column
          label="组ID"
          prop="userGroupID"
          align="center"
        ></el-table-column>
        <el-table-column label="CPU机时" align="center">
          <template #default="props">
            {{ timeSecondFormat(props.row.wallTime) }}
          </template>
        </el-table-column>
        <el-table-column label="GPU机时" align="center">
          <template #default="props">
            {{ timeSecondFormat(props.row.gwallTime) }}
          </template>
        </el-table-column>
        <el-table-column label="应缴费用" align="center">
          <template #default="props"> {{ props.row.fee }}元 </template>
        </el-table-column>
        <el-table-column
          v-if="propsParam.payFlag"
          label="已缴费用"
          align="center"
        >
          <template #default="props"> {{ props.row.payFee }}元 </template>
        </el-table-column>
        <el-table-column v-if="!propsParam.payFlag" label="操作" align="center">
          <template #default="props">
            <el-button type="primary" @click="handlerPayButtonClick(props.row)"
              >缴费</el-button
            >
          </template>
        </el-table-column>
        <el-table-column label="详情" type="expand" align="center">
          <template #default="props">
            <p class="info">
              <span
                ><strong>导师姓名: </strong
                >{{
                  rowExpandInfo[props.row.userGroupID]?.groupInfo?.tutorName
                }}</span
              >
              <span
                ><strong>导师工号: </strong
                >{{
                  rowExpandInfo[props.row.userGroupID]?.groupInfo?.tutorUsername
                }}</span
              >
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
  <el-dialog
    v-if="!propsParam.payFlag"
    v-model="payBillDialogVisibleFlag"
    title="新建用户"
  >
    <div class="pay-bill-dialog-body">
      <div class="rate-area">
        <h3>包机费率</h3>
        <p>
          <strong>CPU费率:{{ feeRateInfo.cpu }}元</strong>
        </p>
        <p>
          <strong>GPU费率:{{ feeRateInfo.gpu }}元</strong>
        </p>
      </div>
      <div>
        <el-form inline>
          <el-form-item label="应缴费金额: ">
            <span>{{ payBillInfo?.fee }}元</span>
          </el-form-item>
          <el-form-item label="实缴费金额: ">
            <el-input
              v-model.number="payBillForm.payFee"
              type="text"
              disabled
              placeholder="实际缴费的金额"
            >
              <template #append>元</template>
            </el-input>
          </el-form-item>
          <el-form-item label="缴费备注: ">
            <el-input
              v-model="payBillForm.payMessage"
              autosize
              type="textarea"
              placeholder="可以为空,此次缴费的备注,不超过500字"
            ></el-input>
          </el-form-item>
          <el-form-item label="用户组余额: ">
            <span v-if="payBillGroupInfo"
              >{{ zeroWithDefault(payBillGroupInfo?.balance, 0) }}元</span
            >
            <span v-else class="red">数据加载失败</span>
          </el-form-item>
        </el-form>
      </div>
    </div>
    <template #footer>
      <span class="dialog-footer">
        <el-button @click="hideBillDialog">取消</el-button>
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
.refresh-button-row {
  margin: 16px 0px;
}
.pagination-row {
  margin: 16px 0px;
  .pagination-control-panel {
    margin: 0px auto;
    justify-content: center;
  }
}
.info {
  span {
    margin-left: 16px;
  }
}
.pay-bill-dialog-body {
  display: flex;
  justify-content: center;
  width: 100%;
  .rate-area {
    min-width: 30%;
  }
}
.red {
  color: red;
}
</style>
