import { undefinedWithDefault } from '../utils/obj';
import { ApiRequest, PaginationQueryResponse, PingResponse } from './api';

// 静态资源服务服务ping测试
export async function ping(): Promise<PingResponse> {
  const { data, status } = await ApiRequest.request<PingResponse>(
    '/award/ping',
    'GET',
    {},
    {}
  );
  if (!status) {
    throw new Error();
  }
  return data;
}

/**
 * 创建论文奖励申请参数
 */
export type CreatePaperApplyParam = {
  title: string;
  category: string;
  partition: string;
  firstPageImageName: string;
  thanksPageImageName: string;
  remarkMessage: string;
};

/**
 * 创建论文奖励申请
 */
export async function createPaperAwardApply(
  param: CreatePaperApplyParam
): Promise<number> {
  const resp = await ApiRequest.request<{
    id: number;
  }>(
    '/award/paper',
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
 * PaperApply 论文奖励申请信息
 */
export type PaperApply = {
  id: number;
  createrID: number;
  createrUsername: string;
  createrName: string;
  userGroupID: number;
  tutorID: number;
  tutorUsername: string;
  tutorName: string;
  paperTitle: string;
  paperCategory: string;
  paperPartition: string;
  paperFirstPageImageName: string;
  paperThanksPageImageName: string;
  remarkMessage: string;
  checkStatus: number;
  checkerID: number;
  checkerUsername: string;
  checkerName: string;
  checkMoney: number;
  checkMessage: string;
  checkTimeUnix: number;
  extraAttributes: string;
  createTimeUnix: number;
};

/**
 * 分页查询论文奖励申请信息
 */
export async function paginationQueryPaperApply(
  pageIndex: number,
  pageSize: number
): Promise<PaginationQueryResponse<PaperApply>> {
  const resp = await ApiRequest.request<PaginationQueryResponse<PaperApply>>(
    '/award/paper',
    'GET',
    {
      pageIndex,
      pageSize,
    }
  );
  if (!resp.status) {
    throw new Error(resp.message);
  }
  for (const single of resp.data.Data) {
    undefinedWithDefault(single, 'remarkMessage', '');
    undefinedWithDefault(single, 'checkMoney', 0);
  }
  return resp.data;
}

/**
 * 审核论文奖励申请
 */
export type CheckPaperApplyParam = {
  id: number;
  checkMoney: number;
  checkMessage: string;
  accept: boolean;
};

/**
 * 审核论文申请奖励
 */
export async function checkPaperApplyByID(
  param: CheckPaperApplyParam
): Promise<boolean> {
  const resp = await ApiRequest.request(
    '/award/paper',
    'PUT',
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
