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

// 获取安装状态
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

// cas配置
export interface casConfig {
  Enable: boolean;
  AuthServer: string;
  ValidPath: string;
  ServiceAddr: string;
}
// 加载系统的cas配置
export async function loadCasConfig(
  serviceAddr: string
): Promise<casConfig | null> {
  try {
    const resp = await ApiRequest.request<casConfig>('/sys/cas/config', 'GET', {
      serviceHost: serviceAddr,
    });
    if (!resp.status) {
      return null;
    }
    return resp.data;
  } catch (error) {
    return null;
  }
}
