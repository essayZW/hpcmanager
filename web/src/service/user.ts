import { accessTokenKey } from '../api/api';
import {
  createToken,
  getLoginedInfo,
  LoginResponse,
  LoginUserInfo,
  deleteToken,
  UserInfo,
  queryUserInfoByID,
  queryUserInfoByUsername,
  QueryUserIDResponse,
} from '../api/user';
import { getCasConfig } from '../service/sys';

// 判断用户是否已经登录,如果已经登录,返回登录的用户的信息
export async function isLogin(): Promise<LoginUserInfo | null> {
  return getLoginedInfo();
}

// 进行用户登录,登录成功返回用户信息,登录失败返回登录失败消息
export async function login(
  username: string,
  password: string
): Promise<LoginResponse | string> {
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
export function setUserInfoToStorage(info: LoginUserInfo) {
  localStorage.setItem(userInfoLocalStorageKey, JSON.stringify(info));
}

export function getUserInfoFromStorage(): LoginUserInfo | null {
  const str = localStorage.getItem(userInfoLocalStorageKey);
  if (str == null) {
    return null;
  }
  try {
    return JSON.parse(str) as LoginUserInfo;
  } catch (error) {
    return null;
  }
}

/**
 * 通过用户ID获取用户信息
 */
export async function getUserInfoById(id: number): Promise<UserInfo | string> {
  try {
    const data = await queryUserInfoByID(id);
    return data as UserInfo;
  } catch (error) {
    return `${error}`;
  }
}

/**
 * 通过用户名查询用户ID等基础信息
 */
export async function getUserIdByUsername(
  username: string
): Promise<QueryUserIDResponse> {
  return await queryUserInfoByUsername(username);
}
