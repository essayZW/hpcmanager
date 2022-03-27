<script setup lang="ts">
import { reactive } from 'vue';
import { NodeApplyInfo } from '../api/node';
import {
  paginationGetNodeApplyInfo,
  nodeTypeToName,
  checkNodeApply,
  revokeNodeApplyByID,
} from '../service/node';
import dayjs from 'dayjs';
import { ProjectInfo } from '../api/project';
import { UserInfo } from '../api/user';
import {
  getUserInfoById,
  getUserInfoFromStorage,
  isAdmin,
  isTutor,
} from '../service/user';
import { getProjectInfoByID } from '../service/project';
import { zeroWithDefault, timeOrBlank } from '../utils/obj';

// 表格数据
const tableData = reactive<{
  count: number;
  data: NodeApplyInfo[];
  loading: boolean;
}>({
  count: 0,
  data: [],
  loading: false,
});

const userInfo = getUserInfoFromStorage();

// 加载表格数据
const loadTableData = async (pageIndex: number, pageSize: number) => {
  tableData.loading = true;
  try {
    const paginationData = await paginationGetNodeApplyInfo(
      pageIndex,
      pageSize
    );
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

const canRevokeApply = (row: NodeApplyInfo): boolean => {
  // 只有没有被管理员最终审核以及还没有撤销的以及是自己创建的申请才可以撤销
  if (!userInfo) {
    return false;
  }
  return (
    row.managerCheckStatus == -1 &&
    row.status == 1 &&
    row.createrID == userInfo.UserId
  );
};

const revokeButtonHandler = async (id: number) => {
  if (!confirm('确认需要撤销该申请吗?')) {
    return;
  }
  try {
    await revokeNodeApplyByID(id);
    ElMessage({
      type: 'success',
      message: '撤销成功',
    });
    refreshTableData();
  } catch (error) {
    ElMessage({
      type: 'error',
      message: `${error}`,
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

// 表格展开行中的数据
const tableExpandRowInfo = reactive<{
  [id: number]: {
    projectInfo?: ProjectInfo;
    applierInfo?: UserInfo;
    tutorInfo?: UserInfo;
    loading?: boolean;
  };
}>({});

// 刷新表格当前页面的信息
const refreshTableData = () => {
  loadTableData(paginationInfo.pageIndex, paginationInfo.pageSize);
  // 清除缓存的表的扩展字段的属性
  for (const key in tableExpandRowInfo) {
    tableExpandRowInfo[key] = {};
  }
};

refreshTableData();

defineExpose({
  refreshTableData,
});

const handleCurrentChange = (pageIndex: number) => {
  paginationInfo.pageIndex = pageIndex;
};

const handleSizeChange = (pageSize: number) => {
  paginationInfo.pageSize = pageSize;
};

// 表格行展开时候的回调事件
const handlerTableExpand = async (row: NodeApplyInfo) => {
  if (!tableExpandRowInfo[row.id]) {
    tableExpandRowInfo[row.id] = {
      loading: true,
    };
  }
  tableExpandRowInfo[row.id].loading = true;
  if (!tableExpandRowInfo[row.id].applierInfo) {
    const userInfo = await getUserInfoById(row.createrID);
    if (typeof userInfo == 'string') {
      ElMessage({
        type: 'error',
        message: userInfo as string,
      });
    } else {
      tableExpandRowInfo[row.id].applierInfo = userInfo as UserInfo;
    }
  }

  if (!tableExpandRowInfo[row.id].tutorInfo) {
    const userInfo = await getUserInfoById(row.tutorID);
    if (typeof userInfo == 'string') {
      ElMessage({
        type: 'error',
        message: userInfo as string,
      });
    } else {
      tableExpandRowInfo[row.id].tutorInfo = userInfo as UserInfo;
    }
  }

  if (!tableExpandRowInfo[row.id].projectInfo) {
    try {
      const projectInfo = await getProjectInfoByID(row.projectID);
      tableExpandRowInfo[row.id].projectInfo = projectInfo;
    } catch (error) {
      ElMessage({
        type: 'error',
        message: `${error}`,
      });
    }
  }
  tableExpandRowInfo[row.id].loading = false;
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
    await checkNodeApply(
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
        @expand-change="handlerTableExpand"
      >
        <el-table-column label="ID" prop="id" sortable></el-table-column>
        <el-table-column
          label="申请人学号"
          prop="createrUsername"
        ></el-table-column>
        <el-table-column
          label="申请人姓名"
          prop="createrName"
        ></el-table-column>
        <el-table-column
          label="导师工号"
          prop="tutorUsername"
        ></el-table-column>
        <el-table-column label="导师姓名" prop="tutorName"></el-table-column>
        <el-table-column label="申请时间">
          <template #default="props">
            {{
              dayjs(props.row.createTime * 1000).format('YYYY-MM-DD HH:mm:ss')
            }}
          </template>
        </el-table-column>
        <el-table-column label="节点类型">
          <template #default="props">
            {{ nodeTypeToName(props.row.nodeType) }}
          </template>
        </el-table-column>
        <el-table-column label="节点数目" prop="nodeNum"></el-table-column>
        <el-table-column label="导师审核状态">
          <template #default="props">
            <span v-if="props.row.tutorCheckStatus == -1">未审核</span>
            <span v-else-if="props.row.tutorCheckStatus == 1" class="green"
              >审核通过</span
            >
            <span v-else class="red">审核未通过</span>
          </template>
        </el-table-column>
        <el-table-column label="管理员审核状态">
          <template #default="props">
            <span v-if="props.row.managerCheckStatus == -1">未审核</span>
            <span v-else-if="props.row.managerCheckStatus == 1" class="green"
              >审核通过</span
            >
            <span v-else class="red">审核未通过</span>
          </template>
        </el-table-column>
        <el-table-column label="状态">
          <template #default="props">
            <span v-if="props.row.status == 1">正常</span>
            <span v-else class="red">已经撤销</span>
          </template>
        </el-table-column>

        <el-table-column label="操作" prop="status">
          <template #default="props">
            <el-button
              v-if="canRevokeApply(props.row)"
              type="warning"
              @click="revokeButtonHandler(props.row.id)"
              >撤销</el-button
            >
            <span v-else>无</span>
          </template>
        </el-table-column>
        <el-table-column label="详情" type="expand">
          <template #default="props">
            <div
              v-loading="tableExpandRowInfo[props.row.id].loading"
              class="table-expand-area"
            >
              <el-divider content-position="left">申请详情: </el-divider>
              <div><strong>申请人信息:</strong></div>
              <div class="info">
                <p>
                  <span
                    ><strong>姓名: </strong
                    >{{
                      tableExpandRowInfo[props.row.id].applierInfo?.name
                    }}</span
                  >
                  <span
                    ><strong>学号: </strong
                    >{{
                      tableExpandRowInfo[props.row.id].applierInfo?.username
                    }}</span
                  >
                  <span
                    ><strong>学院: </strong
                    >{{
                      zeroWithDefault(
                        tableExpandRowInfo[props.row.id].applierInfo?.college,
                        '无'
                      )
                    }}</span
                  >
                </p>
                <p>
                  <span
                    ><strong>电话: </strong
                    >{{
                      zeroWithDefault(
                        tableExpandRowInfo[props.row.id].applierInfo?.tel,
                        '无'
                      )
                    }}</span
                  >
                  <span
                    ><strong>邮箱地址: </strong
                    >{{
                      zeroWithDefault(
                        tableExpandRowInfo[props.row.id].applierInfo?.email,
                        '无'
                      )
                    }}</span
                  >
                </p>
              </div>
              <div><strong>导师信息:</strong></div>
              <div class="info">
                <p>
                  <span
                    ><strong>姓名: </strong
                    >{{
                      tableExpandRowInfo[props.row.id].tutorInfo?.name
                    }}</span
                  >
                  <span
                    ><strong>学号: </strong
                    >{{
                      tableExpandRowInfo[props.row.id].tutorInfo?.username
                    }}</span
                  >
                  <span
                    ><strong>学院: </strong
                    >{{
                      zeroWithDefault(
                        tableExpandRowInfo[props.row.id].tutorInfo?.college,
                        '无'
                      )
                    }}</span
                  >
                </p>
                <p>
                  <span
                    ><strong>电话: </strong
                    >{{
                      zeroWithDefault(
                        tableExpandRowInfo[props.row.id].tutorInfo?.tel,
                        '无'
                      )
                    }}</span
                  >
                  <span
                    ><strong>邮箱地址: </strong
                    >{{
                      zeroWithDefault(
                        tableExpandRowInfo[props.row.id].tutorInfo?.email,
                        '无'
                      )
                    }}</span
                  >
                </p>
              </div>
              <div><strong>项目信息: </strong></div>
              <div class="info">
                <p>
                  <span
                    ><strong>名称: </strong
                    >{{
                      zeroWithDefault(
                        tableExpandRowInfo[props.row.id].projectInfo?.name,
                        '无'
                      )
                    }}</span
                  >
                  <span
                    ><strong>来源: </strong
                    >{{
                      zeroWithDefault(
                        tableExpandRowInfo[props.row.id].projectInfo?.from,
                        '无'
                      )
                    }}</span
                  >
                  <span
                    ><strong>编号: </strong
                    >{{
                      zeroWithDefault(
                        tableExpandRowInfo[props.row.id].projectInfo?.numbering,
                        '无'
                      )
                    }}</span
                  >
                  <span
                    ><strong>经费: </strong
                    >{{
                      zeroWithDefault(
                        tableExpandRowInfo[props.row.id].projectInfo?.expenses,
                        '无'
                      )
                    }}</span
                  >
                </p>
                <p>
                  <span
                    ><strong>描述: </strong
                    >{{
                      zeroWithDefault(
                        tableExpandRowInfo[props.row.id].projectInfo
                          ?.description,
                        '无'
                      )
                    }}</span
                  >
                </p>
              </div>
            </div>
            <div class="check-area">
              <el-divider content-position="left">审核详情: </el-divider>
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
                      v-if="
                        isTutor() &&
                        props.row.tutorCheckStatus == -1 &&
                        props.row.status == 1
                      "
                      class="box-title"
                    >
                      <strong>操作</strong>
                    </p>
                    <p
                      v-if="
                        isTutor() &&
                        props.row.tutorCheckStatus == -1 &&
                        props.row.status == 1
                      "
                    >
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
                      v-if="
                        isAdmin() &&
                        props.row.managerCheckStatus == -1 &&
                        props.row.status == 1
                      "
                      class="box-title"
                    >
                      <strong>操作</strong>
                    </p>
                    <p
                      v-if="
                        isAdmin() &&
                        props.row.managerCheckStatus == -1 &&
                        props.row.status == 1
                      "
                    >
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
        :total="tableData.count"
        :hide-on-single-page="true"
        @size-change="handleSizeChange"
        @current-change="handleCurrentChange"
      >
      </el-pagination>
    </el-col>
  </el-row>
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
</style>
