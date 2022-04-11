import { undefinedWithDefault } from '../utils/obj';
import { ApiRequest, PaginationQueryResponse, PingResponse } from './api';

// 费用服务ping测试
export async function ping(): Promise<PingResponse> {
  const { data, status } = await ApiRequest.request<PingResponse>(
    '/fee/ping',
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
 * 机器节点独占账单信息
 */
export type NodeDistributeBill = {
  id: number;
  applyID: number;
  nodeDistributeID: number;
  fee: number;
  payFee: number;
  payFlag: number;
  payTime: number;
  payType: number;
  payMessage: number;
  userID: number;
  userUsername: string;
  userName: string;
  userGroupID: number;
  createTime: number;
  extraAttributes: string;
};

/**
 * 分页查询机器节点独占账单信息
 */
export async function paginationQueryNodeDistributeBill(
  pageIndex: number,
  pageSize: number
): Promise<PaginationQueryResponse<NodeDistributeBill>> {
  const resp = await ApiRequest.request<
    PaginationQueryResponse<NodeDistributeBill>
  >('/fee/distribute', 'GET', {
    pageSize,
    pageIndex,
  });
  if (!resp.status) {
    throw new Error(resp.message);
  }
  for (const info of resp.data.Data) {
    undefinedWithDefault(info, 'fee', 0.0);
    info.fee.toFixed(2);
    undefinedWithDefault(info, 'payFee', 0);
    info.payFee.toFixed(2);
  }
  return resp.data;
}
