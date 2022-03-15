<script setup lang="ts">
import { reactive, ref } from 'vue';
import { ApplyInfo } from '../api/group';
import { paginationGetApplyInfo, checkJoinGroupApply } from '../service/group';
import { zeroWithDefault } from '../utils/obj';
import dayjs from 'dayjs';
import { isTutor, isAdmin } from '../service/user';

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

// 暴露出相关的方法
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

// 返回格式化的时间或者空时间
const timeOrBlank = (time: number): string => {
  const date = dayjs(time * 1000);
  if (time < 0) {
    return '';
  }
  if (!date.isValid()) {
    return '';
  }
  return date.format('YYYY-MM-DD HH:mm:ss');
};

// 审核意见输入框数据
const checkMessageInput = reactive<{
  tutor: string;
  manager: string;
}>({
  tutor: '',
  manager: '',
});

// 审批按钮处理函数
const checkButtonHandler = async (
  applyID: number,
  checkStatus: boolean,
  tutorCheck = true
) => {
  if (!confirm(!checkStatus ? '确认不通过该条申请吗' : '确认通过该条申请吗')) {
    return;
  }

  try {
    await checkJoinGroupApply(
      applyID,
      checkStatus,
      tutorCheck ? checkMessageInput.tutor : checkMessageInput.manager,
      tutorCheck
    );
    ElMessage({
      type: 'success',
      message: `审核状态变更成功`,
    });
    refreshTableData();
  } catch (error) {
    ElMessage({
      type: 'error',
      message: `${error}`,
    });
  } finally {
    if (tutorCheck) {
      checkMessageInput.tutor = '';
    } else {
      checkMessageInput.manager = '';
    }
  }
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
              dayjs(scope.row.createTime * 1000).format('YYYY-MM-DD HH:mm:ss')
            }}
          </template>
        </el-table-column>
        <el-table-column label="状态">
          <template #default="scope">
            <span
              v-if="
                scope.row.status == 1 &&
                scope.row.tutorCheckStatus == -1 &&
                scope.row.managerCheckStatus == -1
              "
              >未审核</span
            >
            <span v-if="scope.row.status == 0" class="red">已经撤销</span>
            <span
              v-else-if="
                scope.row.tutorCheckStatus == 1 &&
                scope.row.managerCheckStatus == -1
              "
              class="green"
              >导师审核通过</span
            >
            <span
              v-else-if="
                scope.row.tutorCheckStatus == 0 &&
                scope.row.managerCheckStatus == -1
              "
              class="red"
              >导师审核未通过</span
            >
            <span v-else-if="scope.row.managerCheckStatus == 1" class="green"
              >管理员审核通过</span
            >
            <span v-else-if="scope.row.managerCheckStatus == 0" class="red"
              >管理员审核未通过</span
            >
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
                    <el-card class="check-card">
                      <p class="box-title"><strong>审核情况</strong></p>
                      <p>
                        <span
                          ><strong>审核人姓名: </strong>
                          {{ zeroWithDefault(props.row.tutorName, '无') }}</span
                        >
                      </p>
                      <p>
                        <span
                          ><strong>审核人工号: </strong>
                          {{ props.row.tutorUsername }}
                        </span>
                      </p>
                      <p>
                        <span
                          ><strong>审核状态: </strong>
                          <span v-if="props.row.tutorCheckStatus == -1"
                            >未审核
                          </span>

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
                          <strong>审批意见: </strong>
                          {{ zeroWithDefault(props.row.messageTutor, '无') }}
                        </span>
                      </p>
                      <p
                        v-if="isTutor() && props.row.tutorCheckStatus == -1"
                        class="box-title"
                      >
                        <strong>操作</strong>
                      </p>
                      <p v-if="isTutor() && props.row.tutorCheckStatus == -1">
                        <el-form class="form">
                          <el-form-item label="审核">
                            <el-button
                              type="success"
                              size="small"
                              class="check-pass-button"
                              @click="checkButtonHandler(props.row.id, true)"
                              >通过</el-button
                            >
                            <el-button
                              type="danger"
                              size="small"
                              @click="checkButtonHandler(props.row.id, false)"
                              >不通过</el-button
                            >
                          </el-form-item>
                          <el-form-item label="审核意见">
                            <el-input
                              v-model="checkMessageInput.tutor"
                              autosize
                              type="textarea"
                              placeholder="请输入审核意见(0~280字)"
                            />
                          </el-form-item>
                        </el-form>
                      </p>
                    </el-card>
                  </el-timeline-item>
                  <el-timeline-item
                    v-if="props.row.tutorCheckStatus == 1"
                    placement="top"
                    :timestamp="timeOrBlank(props.row.managerCheckTime)"
                  >
                    <el-card class="check-card">
                      <p class="box-title"><strong>审核情况</strong></p>
                      <p>
                        <span
                          ><strong>审核人姓名: </strong>
                          {{
                            zeroWithDefault(props.row.managerCheckerName, '无')
                          }}</span
                        >
                      </p>
                      <p>
                        <span
                          ><strong>审核人工号: </strong>
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
                          ><strong>审核状态: </strong>
                          <span v-if="props.row.managerCheckStatus == -1"
                            >未审核
                          </span>
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
                          <strong>审批意见: </strong>
                          {{ zeroWithDefault(props.row.messageManager, '无') }}
                        </span>
                      </p>
                      <p
                        v-if="isAdmin() && props.row.managerCheckStatus == -1"
                        class="box-title"
                      >
                        <strong>操作</strong>
                      </p>
                      <p v-if="isAdmin() && props.row.managerCheckStatus == -1">
                        <el-form class="form">
                          <el-form-item label="审核">
                            <el-button
                              type="success"
                              size="small"
                              class="check-pass-button"
                              @click="
                                checkButtonHandler(props.row.id, true, false)
                              "
                              >通过</el-button
                            >
                            <el-button
                              type="danger"
                              size="small"
                              @click="
                                checkButtonHandler(props.row.id, false, false)
                              "
                              >不通过</el-button
                            >
                          </el-form-item>
                          <el-form-item label="审核意见">
                            <el-input
                              v-model="checkMessageInput.manager"
                              autosize
                              type="textarea"
                              placeholder="请输入审核意见(0~280字)"
                            />
                          </el-form-item>
                        </el-form>
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
<style lang="less" scoped>
.pagination-row {
  margin: 16px 0px;
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
</style>
