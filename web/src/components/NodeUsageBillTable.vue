<script setup lang="ts">
import { reactive, ref } from 'vue';
import { nodeUsageFeeRate, NodeWeekUsageBill } from '../api/fee';
import dayjs from 'dayjs';
import {
  getNodeUsageFeeRate,
  paginationGetNodeWeekUsageBill,
  payTypeToString,
  setNodeUsageFeeRate,
} from '../service/fee';
import { timeSecondFormat, zeroWithDefault } from '../utils/obj';
import { isSuperAdmin } from '../service/user';

const tableData = reactive<{
  data: NodeWeekUsageBill[];
  count: number;
  timeRange: Date[];
  loading: boolean;
}>({
  data: [],
  count: 0,
  timeRange: [dayjs(new Date().getTime()).add(-1, 'year').toDate(), new Date()],
  loading: false,
});

// 加载表格某一页的数据
const loadTableData = async (
  pageIndex: number,
  pageSize: number,
  startTime: number,
  endTime: number
) => {
  tableData.loading = true;
  try {
    const data = await paginationGetNodeWeekUsageBill(
      pageIndex,
      pageSize,
      startTime,
      endTime
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
  loadTableData(
    paginationInfo.pageIndex,
    paginationInfo.pageSize,
    tableData.timeRange[0].getTime(),
    tableData.timeRange[1].getTime()
  );
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

const changeFeeRateDialogVisible = ref<boolean>(false);

const hideDialog = () => {
  changeFeeRateDialogVisible.value = false;
};
const feeRateDialogFormData = reactive<nodeUsageFeeRate>({
  cpu: 0,
  gpu: 0,
});
const showDialog = async () => {
  try {
    const data = await getNodeUsageFeeRate();
    feeRateDialogFormData.cpu = data.cpu;
    feeRateDialogFormData.gpu = data.gpu;
  } catch (error) {
    ElMessage({
      type: 'error',
      message: `${error}`,
    });
    return;
  }
  changeFeeRateDialogVisible.value = true;
};

const handlerChangeFeeRate = async () => {
  try {
    await setNodeUsageFeeRate(
      feeRateDialogFormData.cpu,
      feeRateDialogFormData.gpu
    );
    ElMessage({
      type: 'success',
      message: '修改费率成功',
    });
    hideDialog();
  } catch (error) {
    ElMessage({
      type: 'error',
      message: `${error}`,
    });
  }
};
</script>
<template>
  <el-row justify="space-between" class="operator-tool-row">
    <div>
      <span>时间范围选择:</span>
      <el-date-picker
        v-model="tableData.timeRange"
        type="daterange"
        unlink-panels
        range-separator="To"
        start-placeholder="Start date"
        end-placeholder="End date"
      />
    </div>
    <div>
      <el-button v-if="isSuperAdmin()" type="primary" @click="showDialog"
        >修改机时费率</el-button
      >
      <el-button type="primary" @click="refreshTableData">
        <el-icon class="el-icon--left">
          <i-ic-round-refresh />
        </el-icon>
        刷新
      </el-button>
    </div>
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
        <el-table-column label="时间范围" align="center">
          <template #default="props">
            {{ dayjs(props.row.startTime * 1000).format('YYYY-MM-DD') }}至{{
              dayjs(props.row.endTime * 1000).format('YYYY-MM-DD')
            }}
          </template>
        </el-table-column>
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
            <span v-else>未缴费</span>
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
  <el-dialog v-model="changeFeeRateDialogVisible" title="修改机时费率">
    <el-form>
      <el-form-item label="CPU费率">
        <el-input
          v-model.number="feeRateDialogFormData.cpu"
          type="number"
        ></el-input>
      </el-form-item>
      <el-form-item label="GPU费率">
        <el-input
          v-model.number="feeRateDialogFormData.gpu"
          type="number"
        ></el-input>
      </el-form-item>
    </el-form>
    <template #footer>
      <el-button @click="hideDialog">取消</el-button>
      <el-button type="primary" @click="handlerChangeFeeRate"
        >确认修改</el-button
      >
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
</style>
