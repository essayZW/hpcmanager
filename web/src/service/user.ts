import { ApiRequest } from '../api/api';
import { createToken, loginResponse } from '../api/user';

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
export async function isLogin(): Promise<LoginUserInfo | null> {
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

// 进行用户登录,登录成功返回用户信息,登录失败返回登录失败消息
export async function login(
  username: string,
  password: string
): Promise<loginResponse | string> {
  try {
    const data = await createToken({
      username,
      password,
    });
    return data;
  } catch (error) {
    return `登录失败:${error}`;
  }
}
