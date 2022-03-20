import { PaginationQueryResponse } from '../api/api';
import { NodeApplyInfo, paginationQueryNodeApplyInfo } from '../api/node';

/**
 * 分页查询机器节点申请信息
 */
export async function paginationGetNodeApplyInfo(
  pageIndex: number,
  pageSize: number
): Promise<PaginationQueryResponse<NodeApplyInfo>> {
  return paginationQueryNodeApplyInfo(pageIndex, pageSize);
}
