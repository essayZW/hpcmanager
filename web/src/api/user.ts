import { Ping, ApiRequest } from './api';

// 用户服务ping测试
export async function ping(): Promise<Ping> {
  const { data } = await ApiRequest.request<Ping>('/user/ping', 'GET', {}, {});
  return data;
}