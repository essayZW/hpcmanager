<script setup lang="ts">
import { reactive, ref } from 'vue';
import { ApplyInfo } from '../api/group';
import { paginationGetApplyInfo } from '../service/group';
import { zeroWithDefault } from '../utils/obj';
import moment from 'moment';

// 表格数据
const tableData = ref<ApplyInfo[]>([]);

// 分页信息
const paginationInfo = reactive<{
  pageIndex: number;
  pageSize: number;
  count: number;
}>({
  pageIndex: 1,
  pageSize: 5,
  count: 0,
});

// 加载表格指定页的数据
const loadTableData = async (pageIndex: number, pageSize: number) => {
  try {
    const data = await paginationGetApplyInfo(pageIndex, pageSize);
    tableData.value = data.Data;
    paginationInfo.count = data.Count;
  } catch (error) {
    ElMessage({
      type: 'error',
      message: `${error}`,
    });
  }
};

// 刷新表格的当前页的数据
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

// 返回格式化的时间或者空时间
const timeOrBlank = (time: number): string => {
  const date = moment(time * 1000);
  if (date.isValid()) {
    return '';
  }
  return date.format('YYYY-MM-DD hh:mm:ss');
};
</script>
<template>
  <el-row justify="end">
    <el-button type="primary" @click="refreshTableData">
      <el-icon class="el-icon--left">
        <i-ic-round-refresh />
      </el-icon>
      刷新
    </el-button>
  </el-row>
  <el-row justify="center">
    <el-col :span="24">
      <el-table table-layout="auto" :data="tableData">
        <el-table-column label="ID" prop="id"> </el-table-column>
        <el-table-column label="申请人姓名" prop="userName"></el-table-column>
        <el-table-column
          label="申请人学号"
          prop="userUsername"
        ></el-table-column>
        <el-table-column label="导师姓名" prop="tutorName"></el-table-column>
        <el-table-column
          label="导师工号"
          prop="tutorUsername"
        ></el-table-column>
        <el-table-column label="申请时间">
          <template #default="scope">
            {{
              moment(scope.row.createTime * 1000).format('YYYY-MM-DD HH:mm:ss')
            }}
          </template>
        </el-table-column>
        <el-table-column label="状态">
          <template #default="scope">
            <span v-if="scope.row.status == 1">正常</span>
            <span v-else class="red">已经撤销</span>
          </template>
        </el-table-column>
        <el-table-column label="更多" type="expand">
          <template #default="props">
            <div>
              <p><strong>审批情况</strong></p>
              <p>
                <el-timeline>
                  <el-timeline-item
                    placement="top"
                    :timestamp="timeOrBlank(props.row.tutorCheckTime)"
                  >
                    <el-card>
                      <p>
                        <span
                          >审核人姓名:
                          {{ zeroWithDefault(props.row.tutorName, '无') }}</span
                        >
                      </p>
                      <p>
                        <span> 审核人工号: {{ props.row.tutorUsername }} </span>
                      </p>
                      <p>
                        <span
                          >审核状态:
                          <span v-if="props.row.tutorCheckStatus == -1"
                            >未审核</span
                          >

                          <span
                            v-else-if="props.row.tutorCheckStatus == 1"
                            class="green"
                            >审核通过</span
                          >
                          <span v-else class="red">审核未通过</span>
                        </span>
                      </p>
                      <p>
                        <span>
                          审批意见:
                          {{ zeroWithDefault(props.row.messageTutor, '无') }}
                        </span>
                      </p>
                    </el-card>
                  </el-timeline-item>
                  <el-timeline-item
                    placement="top"
                    :timestamp="timeOrBlank(props.row.managerCheckTime)"
                  >
                    <el-card>
                      <p>
                        <span
                          >审核人姓名:
                          {{
                            zeroWithDefault(props.row.managerCheckerName, '无')
                          }}</span
                        >
                      </p>
                      <p>
                        <span>
                          审核人工号:
                          {{
                            zeroWithDefault(
                              props.row.managerCheckerUsername,
                              '无'
                            )
                          }}
                        </span>
                      </p>
                      <p>
                        <span
                          >审核状态:
                          <span v-if="props.row.managerCheckStatus == -1"
                            >未审核</span
                          >

                          <span
                            v-else-if="props.row.managerCheckStatus == 1"
                            class="green"
                            >审核通过</span
                          >
                          <span v-else class="red">审核未通过</span>
                        </span>
                      </p>
                      <p>
                        <span>
                          审批意见:
                          {{ zeroWithDefault(props.row.messageManager, '无') }}
                        </span>
                      </p>
                    </el-card>
                  </el-timeline-item>
                </el-timeline>
              </p>
            </div>
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
        :total="paginationInfo.count"
        :hide-on-single-page="true"
        @size-change="handleSizeChange"
        @current-change="handleCurrentChange"
      >
      </el-pagination>
    </el-col>
  </el-row>
</template>
<style lang="less">
.pagination-row {
  margin: 16px 0px;
}
.red {
  color: red;
}
.green {
  color: green;
}
</style>
