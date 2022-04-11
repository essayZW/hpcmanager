import { PaginationQueryResponse } from '../api/api';
import {
  NodeDistributeBill,
  paginationQueryNodeDistributeBill,
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
