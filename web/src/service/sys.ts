import { CreateUserParam } from './user';
import { ApiRequest } from '../api/api';

// install 初始化系统,添加系统默认管理员
export async function install(param: CreateUserParam): Promise<{
  status: boolean;
  message: string;
}> {
  try {
    const resp = await ApiRequest.request<null>(
      '/sys/install',
      'POST',
      {},
      {
        ...param,
      }
    );
    return resp;
  } catch (error) {
    return {
      status: false,
      message: '初始化失败',
    };
  }
}
