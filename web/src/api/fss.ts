import { ApiBaseURL, ApiRequest, PingResponse } from './api';

// 静态资源服务服务ping测试
export async function ping(): Promise<PingResponse> {
  const { data, status } = await ApiRequest.request<PingResponse>(
    '/fss/ping',
    'GET',
    {},
    {}
  );
  if (!status) {
    throw new Error();
  }
  return data;
}

export const fileUploadPath = ApiBaseURL + '/fss/file';

export const uploadFileUrlPathBase = '/upload';
