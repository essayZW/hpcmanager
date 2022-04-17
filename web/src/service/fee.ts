import { PaginationQueryResponse } from '../api/api';
import {
  NodeDistributeBill,
  paginationQueryNodeDistributeBill,
  ping,
  payNodeDistributeBill as payNodeDistributeBillApi,
  NodeDistributeFeeRate,
  queryNodeDistributeFeeRate,
  paginationQueryNodeWeekUsageBill,
  NodeWeekUsageBill,
  paginationQueryGroupNodeWeekUsageBill,
  NodeWeekUsageBillForGroup,
  updateGroupNodeUsageBills,
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

/**
 * 支付机器独占账单
 */
export async function payNodeDistributeBill(
  id: number,
  payMoney: number,
  isBalance: boolean,
  payMessage: string
): Promise<boolean> {
  return payNodeDistributeBillApi({
    id,
    payMoney,
    payType: isBalance ? 2 : 1,
    payMessage,
  });
}

/**
 * 支付方式的数字转换为对应的字符串
 */
export function payTypeToString(payType: number): string {
  switch (payType) {
    case 1:
      return '线下缴费';
    case 2:
      return '余额缴费';
    default:
      return '未知';
  }
}

/**
 * 查询机器独占账单费率
 */
export async function getNodeDistributeFeeRate(): Promise<NodeDistributeFeeRate> {
  return queryNodeDistributeFeeRate();
}

/**
 * 分页查询机器时长周账单记录
 */
export async function paginationGetNodeWeekUsageBill(
  pageIndex: number,
  pageSize: number,
  startDateMilliUnix: number,
  endDateMilliUnix: number
): Promise<PaginationQueryResponse<NodeWeekUsageBill>> {
  return paginationQueryNodeWeekUsageBill(
    pageIndex,
    pageSize,
    startDateMilliUnix,
    endDateMilliUnix
  );
}

/**
 * 分页分组查询机器节点机器时长账单
 */
export async function paginationGetGroupNodeUsageBill(
  pageIndex: number,
  pageSize: number,
  payFlag: boolean
): Promise<PaginationQueryResponse<NodeWeekUsageBillForGroup>> {
  return paginationQueryGroupNodeWeekUsageBill(pageIndex, pageSize, payFlag);
}

/**
 * 支付用户组的所有未支付的机器节点时长账单
 */
export async function payGroupNodeUsageBills(
  userGroupID: number,
  isBalancePay: boolean,
  payMessage: string,
  needFee: number
): Promise<number> {
  return updateGroupNodeUsageBills({
    userGroupID,
    payType: isBalancePay ? 2 : 1,
    payMessage,
    needFee,
  });
}
