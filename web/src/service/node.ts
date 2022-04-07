import { PaginationQueryResponse } from '../api/api';
import {
  NodeApplyInfo,
  paginationQueryNodeApplyInfo,
  createNodeApply as createNodeApplyAPI,
  checkNodeApply as checkNodeApplyAPI,
  NodeDistribute,
  paginationQueryNodeDistributeInfo,
  queryNodeApplyByID,
  finishNodeDistributeByID,
  revokeNodeApply,
  HpcUsageTime,
  paginationQueryNodeUsageTime,
  updateNodeApplyInfo,
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

/**
 * 节点类型转换为对应的名称
 */
export function nodeTypeToName(nodeType: string): string {
  switch (nodeType) {
    case 'cpuc36':
      return '36核心节点';
    case 'gpuc4':
      return '4 GPU卡节点';
    case 'gpuc8':
      return '8 GPU卡节点';
    default:
      return 'unknown';
  }
}

/**
 * 审核机器节点申请
 */
export function checkNodeApply(
  applyID: number,
  checkStatus: boolean,
  checkMessage: string,
  tutorCheck: boolean
): Promise<boolean> {
  return checkNodeApplyAPI({
    applyID,
    checkStatus,
    checkMessage,
    tutorCheck,
  });
}

/**
 * 分页查询机器节点分配工单信息
 */
export async function paginationGetNodeDistributeInfo(
  pageIndex: number,
  pageSize: number
): Promise<PaginationQueryResponse<NodeDistribute>> {
  return paginationQueryNodeDistributeInfo(pageIndex, pageSize);
}

/**
 * 通过ID获取机器节点申请信息
 */
export async function getNodeApplyByID(id: number): Promise<NodeApplyInfo> {
  return queryNodeApplyByID(id);
}

/**
 * 处理机器节点申请工单
 */
export async function handlerNodeDistributeByID(id: number): Promise<boolean> {
  return finishNodeDistributeByID(id);
}

/**
 * 通过ID撤销机器节点申请
 */
export async function revokeNodeApplyByID(id: number): Promise<boolean> {
  return revokeNodeApply(id);
}

/**
 * 分页查询机器节点申请信息
 */
export async function paginationGetNodeUsageTime(
  pageIndex: number,
  pageSize: number,
  startDateMilliUnix: number,
  endDateMilliUnix: number
): Promise<PaginationQueryResponse<HpcUsageTime>> {
  return await paginationQueryNodeUsageTime(
    pageIndex,
    pageSize,
    startDateMilliUnix,
    endDateMilliUnix
  );
}

/**
 * 更新机器节点申请信息
 */
export async function updateNodeApplyInfoByID(
  applyID: number,
  nodeType: string,
  nodeNum: number,
  startTime: number,
  endTime: number
): Promise<boolean> {
  return updateNodeApplyInfo({
    id: applyID,
    nodeType,
    nodeNum,
    startTime,
    endTime,
  });
}
