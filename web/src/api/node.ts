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
