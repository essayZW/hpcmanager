import { undefinedWithDefault } from '../utils/obj';
import { ApiRequest, PaginationQueryResponse } from './api';

/**
 * 用户组信息
 */
export type GroupInfo = {
  id: number;
  name: string;
  createTime: number | string;
  createrID: string;
  createUsername: string;
  createrName: string;
  tutorID: number;
  tutorUsername: string;
  tutorName: string;
  balance: number;
  hpcGroupID: number;
};

/**
 * 分页查询组信息
 */
export async function paginationQueryGroup(
  page: number,
  size: number
): Promise<PaginationQueryResponse<GroupInfo>> {
  const resp = await ApiRequest.request('/group', 'GET', {
    pageIndex: page,
    pageSize: size,
  });
  if (!resp.status) {
    throw new Error(resp.message);
  }
  return resp.data as PaginationQueryResponse<GroupInfo>;
}

/**
 * 创建组请求
 */
export type CreateGroupRequest = {
  tutorID: number;
  groupName: string;
  queueName: string;
};

export type CreateGroupResponse = {
  id: number;
};

/**
 * 根据组名称、私有队列名称以及导师ID，创建组并分配导师权限
 */
export async function createGroup(param: CreateGroupRequest): Promise<number> {
  const resp = await ApiRequest.request(
    '/group',
    'POST',
    {},
    {
      ...param,
    }
  );

  if (!resp.status) {
    throw new Error(resp.message);
  }
  return (resp.data as CreateGroupResponse).id;
}

/**
 * 通过ID查询用户组信息
 */
export async function queryGroupByID(id: number): Promise<GroupInfo> {
  const resp = await ApiRequest.request<GroupInfo>(`/group/${id}`, 'GET');
  if (!resp.status) {
    throw new Error(resp.message);
  }
  return resp.data;
}

/**
 * 搜索导师信息的响应数据格式
 */
export type SearchTutorInfoResponse = {
  tutorID: number;
  tutorUsername: string;
  tutorName: string;
  groupID: number;
  groupName: string;
};

/**
 * 通过导师工号查询相关的信息
 */
export async function queryTutorInfoByUsername(
  username: string
): Promise<SearchTutorInfoResponse> {
  const resp = await ApiRequest.request<SearchTutorInfoResponse>(
    `/group/tutor/${username}`,
    'GET'
  );
  if (!resp.status) {
    throw new Error(resp.message);
  }
  return resp.data;
}

/**
 * 创建加入组申请接口返回数据
 */
export type CreateJoinGroupApplyResponse = {
  applyID: number;
};

/**
 * 创建加入组申请
 */
export async function createJoinGroupApply(
  applyGroupID: number
): Promise<CreateJoinGroupApplyResponse> {
  const resp = await ApiRequest.request<CreateJoinGroupApplyResponse>(
    '/group/apply',
    'POST',
    {},
    {
      applyGroupID,
    }
  );
  if (!resp.status) {
    throw new Error(resp.message);
  }
  return resp.data;
}

/**
 * 申请信息的消息定义
 */
export type ApplyInfo = {
  id: number;
  userID: number;
  userUsername: string;
  userName: string;
  applyGroupID: number;
  tutorID: number;
  tutorUsername: string;
  tutorName: string;
  tutorCheckStatus: number;
  managerCheckStatus: number;
  status: number;
  messageTutor: string;
  messageManager: string;
  tutorCheckTime: number;
  managerCheckTime: number;
  managerCheckerID: number;
  managerCheckerUsername: string;
  managerCheckerName: string;
  createTime: number;
  extraAttributes: string;
};

/**
 * 分页查询用户申请加入用户组信息
 */
export async function paginationQueryApplyInfo(
  pageIndex: number,
  pageSize: number
): Promise<PaginationQueryResponse<ApplyInfo>> {
  const resp = await ApiRequest.request<PaginationQueryResponse<ApplyInfo>>(
    '/group/apply',
    'GET',
    {
      pageIndex,
      pageSize,
    }
  );
  if (!resp.status) {
    throw new Error(resp.message);
  }
  // 处理resp.data中的undefined字段为对应的默认值
  for (const i in resp.data.Data) {
    undefinedWithDefault(resp.data.Data[i], 'id', 0);
    undefinedWithDefault(resp.data.Data[i], 'userID', 0);
    undefinedWithDefault(resp.data.Data[i], 'userUsername', '');
    undefinedWithDefault(resp.data.Data[i], 'userName', '');
    undefinedWithDefault(resp.data.Data[i], 'applyGroupID', 0);
    undefinedWithDefault(resp.data.Data[i], 'tutorID', 0);
    undefinedWithDefault(resp.data.Data[i], 'tutorUsername', '');
    undefinedWithDefault(resp.data.Data[i], 'tutorName', '');
    undefinedWithDefault(resp.data.Data[i], 'tutorCheckStatus', 0);
    undefinedWithDefault(resp.data.Data[i], 'managerCheckStatus', 0);
    undefinedWithDefault(resp.data.Data[i], 'status', 0);
    undefinedWithDefault(resp.data.Data[i], 'messageTutor', '');
    undefinedWithDefault(resp.data.Data[i], 'messageManager', '');
    undefinedWithDefault(resp.data.Data[i], 'tutorCheckTime', 0);
    undefinedWithDefault(resp.data.Data[i], 'managerCheckTime', 0);
    undefinedWithDefault(resp.data.Data[i], 'managerCheckerID', 0);
    undefinedWithDefault(resp.data.Data[i], 'managerCheckerUsername', '');
    undefinedWithDefault(resp.data.Data[i], 'managerCheckerName', '');
    undefinedWithDefault(resp.data.Data[i], 'createTime', 0);
    undefinedWithDefault(resp.data.Data[i], 'extraAttributes', '');
  }
  return resp.data;
}

/**
 * 审核用户加入用户组申请接口请求参数
 */
export type CheckJoinGroupApplyRequest = {
  checkStatus: boolean;
  checkMessage: string;
  applyID: number;
  tutorCheck: boolean;
};

/**
 * 调用审核用户加入组申请接口
 */
export async function checkApply(
  param: CheckJoinGroupApplyRequest
): Promise<boolean> {
  const resp = await ApiRequest.request<null>(
    '/group/apply',
    'PATCH',
    {},
    {
      ...param,
    }
  );
  if (!resp.status) {
    throw new Error(resp.message);
  }
  return true;
}
