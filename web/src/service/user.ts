import { ApiRequest } from '../api/api';

// 创建用户参数
export interface CreateUserParam {
  username: string;
  password: string;
  name: string;
  tel: string;
  email: string;
  collegeName: string;
}

// 已经登录的用户的基本信息
export interface LoginUserInfo {
  userID: number;
  username: string;
  userName: string;
  groupID: number;
  levels: number[];
}

// 判断用户是否已经登录,如果已经登录,返回登录的用户的信息
export async function IsLogin(): Promise<LoginUserInfo | null> {
  try {
    const { status, data } = await ApiRequest.request('/user/token', 'GET');
    if (!status) {
      return null;
    }
    return data as LoginUserInfo;
  } catch (error) {
    return null;
  }
}
