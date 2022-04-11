import { PaginationQueryResponse } from '../api/api';
import {
  NodeDistributeBill,
  paginationQueryNodeDistributeBill,
  ping,
} from '../api/fee';

/**
 * 分页查询机器节点独占账单信息
 */
export async function paginationGetNodeDistributeBill(
  pageIndex: number,
  pageSize: number
): Promise<PaginationQueryResponse<NodeDistributeBill>> {
  return paginationQueryNodeDistributeBill(pageIndex, pageSize);
}

/**
 * 服务ping
 */
export async function servicePing(): Promise<boolean> {
  try {
    await ping();
    return true;
  } catch (error) {
    return false;
  }
}
