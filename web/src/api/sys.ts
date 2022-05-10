import { CreateUserRequest } from './user';
import { ApiRequest } from './api';

// installRequest 安装系统的请求
export type InstallRequest = CreateUserRequest;

// 安装系统
export async function installSys(param: InstallRequest): Promise<boolean> {
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
export interface CasConfig {
  Enable: boolean;
  AuthServer: string;
  ValidPath: string;
  ServiceAddr: string;
}
// 加载系统的cas配置
export async function loadCasConfig(
  serviceAddr: string
): Promise<CasConfig | null> {
  try {
    const resp = await ApiRequest.request<CasConfig>('/sys/cas/config', 'GET', {
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

/**
 * 更新CAS设置参数
 */
export type setCasConfigParam = {
  enable: boolean;
  authServer: string;
};

/**
 * 修改CAS登录设置
 */
export async function setCasConfig(param: setCasConfigParam): Promise<boolean> {
  const resp = await ApiRequest.request(
    '/sys/cas/config',
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
