import { ApiRequest, PingResponse } from './api';

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
