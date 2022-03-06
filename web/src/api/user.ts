import { pingResponse, ApiRequest } from './api';

// 用户服务ping测试
export async function ping(): Promise<pingResponse> {
  const { data } = await ApiRequest.request<pingResponse>(
    '/user/ping',
    'GET',
    {},
    {}
  );
  return data;
}

// 登录后的用户信息
export interface loginUserInfo {
  userID: number;
  username: string;
  userName: string;
  groupID: number;
  levels: number[];
}

// login需要用到的参数
export interface loginRequest {
  username: string;
  password: string;
}

// login 登录后返回的数据格式
export interface loginResponse {
  userInfo: loginUserInfo;
  token: string;
}

// 创建登录token
export async function createToken(param: loginRequest): Promise<loginResponse> {
  const resp = await ApiRequest.request<loginResponse>(
    '/user/token',
    'POST',
    {},
    {
      ...param,
    }
  );
  if (!resp.status) {
    throw new Error(resp.message);
  }
  return resp.data;
}
