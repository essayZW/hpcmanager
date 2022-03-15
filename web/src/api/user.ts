import { PingResponse, ApiRequest, PaginationQueryResponse } from './api';
import { undefinedWithDefault } from '../utils/obj';

// 用户服务ping测试
export async function ping(): Promise<PingResponse> {
  const { data } = await ApiRequest.request<PingResponse>(
    '/user/ping',
    'GET',
    {},
    {}
  );
  return data;
}

// 登录后的用户信息
export interface LoginUserInfo {
  UserId: number;
  Username: string;
  Name: string;
  GroupId: number;
  Levels: number[];
}

// login需要用到的参数
export interface LoginRequest {
  username: string;
  password: string;
}

// login 登录后返回的数据格式
export interface LoginResponse {
  userInfo: LoginUserInfo;
  token: string;
}

// 创建登录token
export async function createToken(param: LoginRequest): Promise<LoginResponse> {
  const resp = await ApiRequest.request<LoginResponse>(
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

// 获取登录的用户的信息
export async function getLoginedInfo(): Promise<LoginUserInfo | null> {
  try {
    const { status, data } = await ApiRequest.request<LoginUserInfo>(
      '/user/token',
      'GET'
    );
    if (!status) {
      return null;
    }
    return data;
  } catch (error) {
    return null;
  }
}

// 创建用户参数
export interface CreateUserRequest {
  username: string;
  password: string;
  name: string;
  tel: string;
  email: string;
  collegeName: string;
}

/**
 * 删除用户登录token
 */
export async function deleteToken(): Promise<boolean> {
  try {
    const resp = await ApiRequest.request('/user/token', 'DELETE');
    return resp.status;
  } catch (error) {
    return false;
  }
}

/**
 * 用户信息
 */
export type UserInfo = {
  id: number;
  username: string;
  tel: string;
  email: string;
  name: string;
  pyName: string;
  college: string;
  groupId: number;
  createTime: number;
  extraAttributes: string;
  hpcUserID: number;
};

/**
 * 通过用户ID查询用户信息
 */
export async function queryUserInfoByID(id: number): Promise<UserInfo> {
  const resp = await ApiRequest.request(`/user/${id}`, 'GET');
  if (!resp.status) {
    throw new Error(resp.message);
  }
  undefinedWithDefault(resp.data as UserInfo, 'id', 0);
  undefinedWithDefault(resp.data as UserInfo, 'username', '');
  undefinedWithDefault(resp.data as UserInfo, 'tel', '');
  undefinedWithDefault(resp.data as UserInfo, 'email', '');
  undefinedWithDefault(resp.data as UserInfo, 'name', '');
  undefinedWithDefault(resp.data as UserInfo, 'pyName', '');
  undefinedWithDefault(resp.data as UserInfo, 'college', '');
  undefinedWithDefault(resp.data as UserInfo, 'groupID', 0);
  undefinedWithDefault(resp.data as UserInfo, 'createTime', 0);
  undefinedWithDefault(resp.data as UserInfo, 'extraAttributes', '');
  undefinedWithDefault(resp.data as UserInfo, 'hpcUserID', 0);
  return resp.data as UserInfo;
}

export type QueryUserIDResponse = {
  id: number;
  groupID: number;
  name: string;
};

/**
 * 查询用户ID
 */
export async function queryUserInfoByUsername(
  username: string
): Promise<QueryUserIDResponse> {
  const resp = await ApiRequest.request(`/user/name/${username}`, 'GET');
  if (!resp.status) {
    throw new Error(resp.message);
  }
  return resp.data as QueryUserIDResponse;
}

/**
 * 分页查询用户信息
 */
export async function paginationQueryUserInfo(
  pageIndex: number,
  pageSize: number
): Promise<PaginationQueryResponse<UserInfo>> {
  const resp = await ApiRequest.request<PaginationQueryResponse<UserInfo>>(
    '/user',
    'GET',
    {
      pageIndex,
      pageSize,
    }
  );
  if (!resp.status) {
    throw new Error(resp.message);
  }
  for (const user of resp.data.Data) {
    undefinedWithDefault(user, 'groupId', 0);
  }
  return resp.data;
}

/**
 * 修改用户信息接口请求参数格式定义
 */
export type UpdateUserInfoRequest = {
  college: string;
  email: string;
  tel: string;
  id: number;
};

/**
 * 修改用户信息
 */
export async function updateUserInfo(newInfo: UpdateUserInfoRequest) {
  const resp = await ApiRequest.request(
    '/user',
    'PATCH',
    {},
    {
      ...newInfo,
    }
  );
  if (!resp.status) {
    throw new Error(resp.message);
  }
}
