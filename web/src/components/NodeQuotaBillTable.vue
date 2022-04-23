<script lang="ts" setup>
import { onMounted, reactive, ref } from 'vue';
import { NodeQuotaBill, NodeQuotaFeeRate } from '../api/fee';
import {
  paginationGetNodeQuotaBill,
  payNodeQuotaBill,
  payTypeToString,
  getNodeQuotaFeeRate,
} from '../service/fee';
import { operTypeToStr } from '../service/fee';
import { zeroWithDefault } from '../utils/obj';
import { isAdmin } from '../service/user';
import { getGroupInfoByID } from '../service/group';
import dayjs from 'dayjs';
import { GroupInfo } from '../api/group';

const tableData = reactive<{
  data: NodeQuotaBill[];
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
    const data = await paginationGetNodeQuotaBill(pageIndex, pageSize);
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

refreshTableData();

const handleCurrentChange = (pageIndex: number) => {
  paginationInfo.pageIndex = pageIndex;
  refreshTableData();
};

const handleSizeChange = (pageSize: number) => {
  paginationInfo.pageSize = pageSize;
  refreshTableData();
};

const payNodeQuotaBillDialogVisible = ref<boolean>(false);
const payNodeQuotaBillDialogBillInfo = ref<NodeQuotaBill | undefined>(
  undefined
);

const payNodeQuotaBillDialogBillForm = reactive<{
  payMoney: number;
  payMessage: string;
  billID: number;
}>({
  payMoney: 0,
  payMessage: '',
  billID: 0,
});

const payNodeQuotaBillGroupInfo = ref<GroupInfo | undefined>(undefined);

const showPayNodeQuotaBillDialog = async (row: NodeQuotaBill) => {
  payNodeQuotaBillDialogBillInfo.value = row;
  payNodeQuotaBillDialogBillForm.billID = row.id;
  payNodeQuotaBillDialogBillForm.payMoney = row.fee;
  try {
    const data = await getGroupInfoByID(row.userGroupID);
    payNodeQuotaBillGroupInfo.value = data;
  } catch (error) {
    ElMessage({
      type: 'error',
      message: `${error}`,
    });
  }
  payNodeQuotaBillDialogVisible.value = true;
};

const hidePayNodeQuotaBillDialog = () => {
  payNodeQuotaBillDialogVisible.value = false;
};

const payNodeQuotaBillDialogFormSubmitHandler = async (balancePay: boolean) => {
  try {
    await payNodeQuotaBill(
      payNodeQuotaBillDialogBillForm.billID,
      balancePay ? 2 : 1,
      payNodeQuotaBillDialogBillForm.payMoney,
      payNodeQuotaBillDialogBillForm.payMessage
    );
    ElMessage({
      type: 'success',
      message: '缴费成功',
    });
    hidePayNodeQuotaBillDialog();
    refreshTableData();
  } catch (error) {
    ElMessage({
      type: 'error',
      message: `${error}`,
    });
  }
};

const nodeQuotaFeeRate = ref<NodeQuotaFeeRate | undefined>(undefined);
onMounted(async () => {
  try {
    const data = await getNodeQuotaFeeRate();
    nodeQuotaFeeRate.value = data;
  } catch (error) {
    ElMessage({
      type: 'error',
      message: `${error}`,
    });
  }
});
</script>
<template>
  <el-row justify="end" class="operator-tool-row">
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
        table-layout="auto"
        border
        :data="tableData.data"
      >
        <el-table-column
          label="用户学(工)号"
          align="center"
          prop="username"
        ></el-table-column>
        <el-table-column
          label="用户姓名"
          align="center"
          prop="name"
        ></el-table-column>
        <el-table-column
          label="用户组ID"
          align="center"
          prop="userGroupID"
        ></el-table-column>
        <el-table-column label="容量变化" align="center">
          <template #default="props">
            {{ props.row.oldSize }}TB
            <el-icon><i-ic-baseline-arrow-right-alt /></el-icon>
            {{ props.row.newSize }}TB
          </template>
        </el-table-column>
        <el-table-column label="结束日期变化" align="center">
          <template #default="props">
            {{ dayjs(props.row.oldEndTimeUnix * 1000).format('YYYY-MM-DD') }}
            <el-icon><i-ic-baseline-arrow-right-alt /></el-icon>
            {{ dayjs(props.row.newEndTimeUnix * 1000).format('YYYY-MM-DD') }}
          </template>
        </el-table-column>
        <el-table-column label="应缴费用" align="center">
          <template #default="props"> {{ props.row.fee }}元 </template>
        </el-table-column>
        <el-table-column label="缴费状态" align="center">
          <template #default="props">
            <span v-if="props.row.payFlag" class="green"
              >已缴费 {{ zeroWithDefault(props.row.payFee, 0) }}元</span
            >
            <span v-else class="red">未缴费</span>
          </template>
        </el-table-column>
        <el-table-column label="缴费方式" align="center">
          <template #default="props">
            <span v-if="props.row.payFlag">{{
              payTypeToString(props.row.payType)
            }}</span>
            <span v-else-if="isAdmin()"
              ><el-button
                type="primary"
                @click="showPayNodeQuotaBillDialog(props.row)"
                >缴费</el-button
              ></span
            >
            <span v-else>未缴费</span>
          </template>
        </el-table-column>
        <el-table-column label="创建时间" align="center">
          <template #default="props">
            {{
              dayjs(props.row.createTime * 1000).format('YYYY-MM-DD HH:mm:ss')
            }}
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
  <el-dialog v-model="payNodeQuotaBillDialogVisible" title="账单缴费">
    <div class="pay-bill-dialog-body">
      <div class="rate-area">
        <h3>存储费率</h3>
        <p>
          <strong
            >基础存储费率:
            {{ zeroWithDefault(nodeQuotaFeeRate?.basic, '加载失败') }}</strong
          >元 1TB 一年
        </p>
        <p>
          <strong
            >额外存储费率:
            {{ zeroWithDefault(nodeQuotaFeeRate?.extra, '加载失败') }}</strong
          >元 1TB 一年
        </p>
      </div>
      <div>
        <el-form inline>
          <el-form-item label="操作类型: ">
            <span>{{
              operTypeToStr(payNodeQuotaBillDialogBillInfo?.operType)
            }}</span>
          </el-form-item>
          <el-form-item label="容量变化: ">
            <span>
              {{ payNodeQuotaBillDialogBillInfo?.oldSize }}TB
              <el-icon><i-ic-baseline-arrow-right-alt /></el-icon>
              {{ payNodeQuotaBillDialogBillInfo?.newSize }}TB
            </span>
          </el-form-item>
          <el-form-item label="结束日期变化: ">
            <span>
              {{
                dayjs(
                  zeroWithDefault(
                    payNodeQuotaBillDialogBillInfo?.oldEndTimeUnix,
                    0
                  ) as number * 1000
                ).format('YYYY-MM-DD')
              }}
              <el-icon><i-ic-baseline-arrow-right-alt /></el-icon>
              {{
                dayjs(
                  zeroWithDefault(
                    payNodeQuotaBillDialogBillInfo?.newEndTimeUnix,
                    0
                  ) as number * 1000
                ).format('YYYY-MM-DD')
              }}
            </span>
          </el-form-item>
          <el-form-item label="创建时间: ">
            <span>{{
              dayjs(
                zeroWithDefault(payNodeQuotaBillDialogBillInfo?.createTime, 0) as number *
                  1000
              ).format('YYYY-MM-DD HH:mm:ss')
            }}</span>
          </el-form-item>
          <el-form-item label="应缴费用: ">
            <span>{{ payNodeQuotaBillDialogBillInfo?.fee }}元</span>
          </el-form-item>
          <el-form-item label="用户组余额: ">
            <span v-if="payNodeQuotaBillGroupInfo"
              >{{
                zeroWithDefault(payNodeQuotaBillGroupInfo?.balance, 0)
              }}元</span
            >
            <span v-else class="red">数据加载失败</span>
          </el-form-item>
        </el-form>
        <el-form>
          <el-form-item label="缴费金额: ">
            <el-input
              v-model.number="payNodeQuotaBillDialogBillForm.payMoney"
              type="text"
            >
              <template #append>元</template>
            </el-input>
          </el-form-item>
          <el-form-item label="缴费备注: ">
            <el-input
              v-model="payNodeQuotaBillDialogBillForm.payMessage"
              type="textarea"
              placeholder="缴费的备注信息,可以为空"
              autosize
            ></el-input>
          </el-form-item>
        </el-form>
      </div>
    </div>
    <template #footer>
      <span class="dialog-footer">
        <el-button @click="hidePayNodeQuotaBillDialog">取消</el-button>
        <el-button
          type="primary"
          @click="payNodeQuotaBillDialogFormSubmitHandler(true)"
          >余额缴费</el-button
        >
        <el-button
          type="primary"
          @click="payNodeQuotaBillDialogFormSubmitHandler(false)"
          >线下缴费</el-button
        >
      </span>
    </template>
  </el-dialog>
</template>
<style lang="less" scoped>
.operator-tool-row {
  margin: 16px 0px;
}
.pagination-row {
  margin: 16px 0px;
  .pagination-control-panel {
    margin: 0px auto;
    justify-content: center;
  }
}
.green {
  color: green;
}
.red {
  color: red;
}

.pay-bill-dialog-body {
  display: flex;
  justify-content: center;
  width: 100%;
  .rate-area {
    min-width: 35%;
  }
}
</style>
