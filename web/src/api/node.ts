import { undefinedWithDefault } from '../utils/obj';
import { ApiRequest, PaginationQueryResponse } from './api';

/**
 * 机器节点申请记录信息
 */
export type NodeApplyInfo = {
  id: number;
  createTime: number;
  createrID: number;
  createrUsername: string;
  createrName: string;
  projectID: number;
  tutorCheckStatus: number;
  managerCheckStatus: number;
  status: number;
  messageTutor: string;
  messageManager: string;
  tutorCheckTime: number;
  tutorID: number;
  tutorUsername: string;
  tutorName: string;
  managerCheckTime: number;
  managerCheckerID: number;
  managerCheckerUsername: string;
  managerCheckerName: string;
  modifyTime: number;
  modifyUserID: number;
  modifyName: string;
  modifyUsername: string;
  nodeType: string;
  nodeNum: number;
  startTime: number;
  endTime: number;
  extraAttributes: string;
};

/**
 * 分页查询机器节点申请信息
 */
export async function paginationQueryNodeApplyInfo(
  pageIndex: number,
  pageSize: number
): Promise<PaginationQueryResponse<NodeApplyInfo>> {
  const resp = await ApiRequest.request<PaginationQueryResponse<NodeApplyInfo>>(
    '/node/apply',
    'GET',
    {
      pageIndex,
      pageSize,
    }
  );
  if (!resp.status) {
    throw new Error(resp.message);
  }
  for (const item of resp.data.Data) {
    undefinedWithDefault(item, 'tutorCheckStatus', 0);
    undefinedWithDefault(item, 'managerCheckStatus', 0);
  }
  return resp.data;
}

/**
 * 创建机器节点申请信息的请求参数
 */
export type CreateNodeApplyParam = {
  projectID: number;
  nodeType: string;
  nodeNum: number;
  startTime: number;
  endTime: number;
};

export async function createNodeApply(
  param: CreateNodeApplyParam
): Promise<number> {
  const resp = await ApiRequest.request<{
    id: number;
  }>(
    '/node/apply',
    'POST',
    {},
    {
      ...param,
    }
  );

  if (!resp.status) {
    throw new Error(resp.message);
  }
  return resp.data.id;
}

/**
 * 审核机器节点申请的请求消息
 */
export type CheckNodeApplyParam = {
  applyID: number;
  checkStatus: boolean;
  checkMessage: string;
  tutorCheck: boolean;
};

/**
 * 审核机器节点申请
 */
export async function checkNodeApply(
  param: CheckNodeApplyParam
): Promise<boolean> {
  const resp = await ApiRequest.request(
    '/node/apply',
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

/**
 * 节点分配工单信息
 */
export type NodeDistribute = {
  id: number;
  applyID: number;
  handlerFlag: number;
  handlerUserID: number;
  handlerUsername: string;
  handlerUserName: string;
  distributeBillID: number;
  createTime: number;
  extraAttributes: string;
};

/**
 * 分页查询节点分配工单信息
 */
export async function paginationQueryNodeDistributeInfo(
  pageIndex: number,
  pageSize: number
): Promise<PaginationQueryResponse<NodeDistribute>> {
  const resp = await ApiRequest.request<
    PaginationQueryResponse<NodeDistribute>
  >('/node/distribute', 'GET', { pageIndex, pageSize });

  if (!resp.status) {
    throw new Error(resp.message);
  }
  return resp.data;
}

/**
 * 通过ID查询机器节点申请信息
 */
export async function queryNodeApplyByID(id: number): Promise<NodeApplyInfo> {
  const resp = await ApiRequest.request<NodeApplyInfo>(
    `/node/apply/${id}`,
    'GET'
  );
  if (!resp.status) {
    throw new Error(resp.message);
  }
  return resp.data;
}
