import { PaginationQueryResponse } from '../api/api';
import {
  NodeApplyInfo,
  paginationQueryNodeApplyInfo,
  createNodeApply as createNodeApplyAPI,
} from '../api/node';

/**
 * 分页查询机器节点申请信息
 */
export async function paginationGetNodeApplyInfo(
  pageIndex: number,
  pageSize: number
): Promise<PaginationQueryResponse<NodeApplyInfo>> {
  return paginationQueryNodeApplyInfo(pageIndex, pageSize);
}

/**
 * 创建新的机器节点申请
 */
export async function createNodeApply(
  projectID: number,
  nodeType: string,
  nodeNum: number,
  startTime: number,
  endTime: number
): Promise<number> {
  return createNodeApplyAPI({
    projectID,
    nodeNum,
    nodeType,
    startTime,
    endTime,
  });
}
