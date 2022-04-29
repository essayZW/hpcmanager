import { undefinedWithDefault } from '../utils/obj';
import {
  ApiBaseURL,
  ApiRequest,
  PaginationQueryResponse,
  PingResponse,
} from './api';

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

/*
 * 节点独占费率
 */
export type NodeDistributeFeeRate = {
  rate36CPU: number;
  rate4GPU: number;
  rate8GPU: number;
};

/**
 * 支付机器独占账单请求参数
 */
export type PayNodeDistributeBillParam = {
  id: number;
  payMoney: number;
  payType: number;
  payMessage: string;
};

/**
 * 支付机器独占账单
 */
export async function payNodeDistributeBill(
  param: PayNodeDistributeBillParam
): Promise<boolean> {
  const resp = await ApiRequest.request(
    '/fee/distribute',
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

/**
 * 查询机器独占账单费率
 */
export async function queryNodeDistributeFeeRate(): Promise<NodeDistributeFeeRate> {
  const resp = await ApiRequest.request<NodeDistributeFeeRate>(
    '/fee/rate/distribute',
    'GET'
  );
  if (!resp.status) {
    throw new Error(resp.message);
  }
  return resp.data;
}

export type NodeWeekUsageBill = {
  id: number;
  userID: number;
  username: string;
  name: string;
  wallTime: number;
  gwallTime: number;
  fee: number;
  payFee: number;
  startTime: number;
  endTime: number;
  payFlag: number;
  payTime: number;
  payType: number;
  payMessage: string;
  userGroupID: number;
  createTime: number;
  extraAttributes: string;
};

/**
 * 分页查询机器节点时长周账单记录
 */
export async function paginationQueryNodeWeekUsageBill(
  pageIndex: number,
  pageSize: number,
  startDateMilliUnix: number,
  endDateMilliUnix: number
): Promise<PaginationQueryResponse<NodeWeekUsageBill>> {
  const resp = await ApiRequest.request<
    PaginationQueryResponse<NodeWeekUsageBill>
  >('/fee/usage/week', 'GET', {
    pageIndex,
    pageSize,
    startDateMilliUnix,
    endDateMilliUnix,
  });

  if (!resp.status) {
    throw new Error(resp.message);
  }
  for (const single of resp.data.Data) {
    undefinedWithDefault(single, 'wallTime', 0);
    undefinedWithDefault(single, 'gwallTime', 0);
    undefinedWithDefault(single, 'payFlag', 0);
    undefinedWithDefault(single, 'fee', 0);
    undefinedWithDefault(single, 'payType', 0);
  }
  return resp.data;
}

/**
 * 用户组的机器时长账单信息
 */
export type NodeWeekUsageBillForGroup = {
  wallTime: number;
  gwallTime: number;
  fee: number;
  payFee: number;
  userGroupID: number;
};

export async function paginationQueryGroupNodeWeekUsageBill(
  pageIndex: number,
  pageSize: number,
  payFlag: boolean
): Promise<PaginationQueryResponse<NodeWeekUsageBillForGroup>> {
  const resp = await ApiRequest.request<
    PaginationQueryResponse<NodeWeekUsageBillForGroup>
  >('/fee/usage/group/week', 'GET', {
    pageIndex,
    pageSize,
    payFlag,
  });
  if (!resp.status) {
    throw new Error(resp.message);
  }
  for (const info of resp.data.Data) {
    undefinedWithDefault(info, 'wallTime', 0);
    undefinedWithDefault(info, 'gwallTime', 0);
    undefinedWithDefault(info, 'fee', 0);
    undefinedWithDefault(info, 'payFee', 0);
  }
  return resp.data;
}

/**
 * 支付用户组机器节点时长账单参数定义
 */
export type PayGroupNodeUsageBillsParam = {
  userGroupID: number;
  payType: number;
  payMessage: string;
  needFee: number;
};

/**
 * 更新用户组节点使用时长状态(支付)
 */
export async function updateGroupNodeUsageBills(
  param: PayGroupNodeUsageBillsParam
): Promise<number> {
  const resp = await ApiRequest.request<{
    count: number;
  }>(
    '/fee/usage/group/bill',
    'PUT',
    {},
    {
      ...param,
    }
  );
  if (!resp.status) {
    throw new Error(resp.message);
  }
  return resp.data.count;
}

/**
 * 机器节点机时费率
 */
export type nodeUsageFeeRate = {
  cpu: number;
  gpu: number;
};

export async function queryNodeUsageFeeRateInfo(): Promise<nodeUsageFeeRate> {
  const resp = await ApiRequest.request<nodeUsageFeeRate>(
    '/fee/rate/usage',
    'GET'
  );
  if (!resp.status) {
    throw new Error(resp.message);
  }
  return resp.data;
}

/**
 * 机器存储账单
 */
export type NodeQuotaBill = {
  id: number;
  userID: number;
  name: string;
  username: string;
  userGroupID: number;
  operType: number;
  oldSize: number;
  newSize: number;
  oldEndTimeUnix: number;
  newEndTimeUnix: number;
  fee: number;
  payFlag: number;
  payFee: number;
  payTimeUnix: number;
  payType: number;
  payMessage: string;
  createTime: number;
  extraAttributes: string;
};

/**
 * 分页查询机器存储账单
 */
export async function paginationQueryNodeQuotaBill(
  pageIndex: number,
  pageSize: number
): Promise<PaginationQueryResponse<NodeQuotaBill>> {
  const resp = await ApiRequest.request<PaginationQueryResponse<NodeQuotaBill>>(
    '/fee/quota',
    'GET',
    {
      pageIndex,
      pageSize,
    }
  );
  if (!resp.status) {
    throw new Error(resp.message);
  }
  for (const info of resp.data.Data) {
    undefinedWithDefault(info, 'oldSize', 0);
    undefinedWithDefault(info, 'fee', 0);
    undefinedWithDefault(info, 'payFee', 0);
    undefinedWithDefault(info, 'payFlag', 0);
    undefinedWithDefault(info, 'payType', 0);
    undefinedWithDefault(info, 'operType', 0);
  }
  return resp.data;
}

/**
 * 支付机器节点存储账单参数
 */
export type PayNodeQuotaBillParam = {
  billID: number;
  payType: number;
  payMoney: number;
  payMessage: string;
};

/**
 * 更新机器节点账单的支付信息
 */
export async function updateNodeQuotaBillPayInfo(
  param: PayNodeQuotaBillParam
): Promise<boolean> {
  const resp = await ApiRequest.request(
    '/fee/quota/bill',
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

/**
 * 机器存储费率
 */
export type NodeQuotaFeeRate = {
  basic: number;
  extra: number;
};

/**
 * 查询机器存储费率信息
 */
export async function queryNodeQuotaFeeRate(): Promise<NodeQuotaFeeRate> {
  const resp = await ApiRequest.request<NodeQuotaFeeRate>(
    '/fee/rate/quota',
    'GET'
  );
  if (!resp.status) {
    throw new Error(resp.message);
  }
  return resp.data;
}

/**
 * 设置机器节点费率参数
 */
export type SetNodeDistributeFeeRateParam = {
  rate36CPU: number;
  rate4GPU: number;
  rate8GPU: number;
};

/**
 * 设置机器节点费率
 */
export async function setNodeDistributeFeeRate(
  param: SetNodeDistributeFeeRateParam
): Promise<boolean> {
  const resp = await ApiRequest.request(
    '/fee/rate/distribute',
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

/**
 * 设置机时费率参数
 */
export type SetNodeUsageFeeRateParam = {
  cpu: number;
  gpu: number;
};

/**
 * 修改机器机时费率
 */
export async function setNodeUsageFeeRate(
  param: SetNodeUsageFeeRateParam
): Promise<boolean> {
  const resp = await ApiRequest.request(
    '/fee/rate/usage',
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

/**
 * 设置机器存储费率参数
 */
export type SetNodeQuotaFeeRateParam = {
  basic: number;
  extra: number;
};

/**
 * 设置机器存储费率
 */
export async function setNodeQuotaFeeRate(
  param: SetNodeQuotaFeeRateParam
): Promise<boolean> {
  const resp = await ApiRequest.request(
    '/fee/rate/quota',
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
