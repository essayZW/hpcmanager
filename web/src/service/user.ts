import { accessTokenKey } from '../api/api';
import {
  createToken,
  getLoginedInfo,
  loginResponse,
  loginUserInfo,
  deleteToken,
} from '../api/user';
import { getCasConfig } from '../service/sys';

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

export async function logout(): Promise<boolean> {
  const status = await deleteToken();
  if (status) {
    localStorage.setItem(accessTokenKey, '');
  }
  // 查询cas配置判断是否需要进行cas的退出
  const config = await getCasConfig();
  if (config != null && config.Enable) {
    window.location.href = `${config.AuthServer}/cas/logout`;
  }
  return status;
}

/*
 * 用户权限等级定义
 */
export enum UserLevels {
  Guest,
  Common,
  Tutor,
  CommonAdmin,
  SuperAdmin,
}

const userInfoLocalStorageKey = 'userInfo';
/**
 * 存储用户信息到storge中
 */
export function setUserInfoToStorage(info: loginUserInfo) {
  localStorage.setItem(userInfoLocalStorageKey, JSON.stringify(info));
}

export function getUserInfoFromStorage(): loginUserInfo | null {
  const str = localStorage.getItem(userInfoLocalStorageKey);
  if (str == null) {
    return null;
  }
  try {
    return JSON.parse(str) as loginUserInfo;
  } catch (error) {
    return null;
  }
}
