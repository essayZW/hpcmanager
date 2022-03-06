import { accessTokenKey } from '../api/api';
import {
  createToken,
  getLoginedInfo,
  loginResponse,
  loginUserInfo,
} from '../api/user';

// 判断用户是否已经登录,如果已经登录,返回登录的用户的信息
export async function isLogin(): Promise<loginUserInfo | null> {
  return getLoginedInfo();
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
    localStorage.setItem(accessTokenKey, data.token);
    return data;
  } catch (error) {
    return `登录失败:${error}`;
  }
}
