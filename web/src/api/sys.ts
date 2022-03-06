import { createUserRequest } from './user';
import { ApiRequest } from './api';

// installRequest 安装系统的请求
export type installRequest = createUserRequest;

// 安装系统
export async function installSys(param: installRequest): Promise<boolean> {
  const resp = await ApiRequest.request<null>(
    '/sys/install',
    'POST',
    {},
    {
      ...param,
    }
  );
  if (!resp.status) {
    throw new Error(resp.message);
  }
  return resp.status;
}

export async function getInstallStatus(): Promise<boolean> {
  try {
    const { status } = await ApiRequest.request<null>(
      '/sys/install',
      'GET',
      {},
      {}
    );
    return status;
  } catch (error) {
    return false;
  }
}
