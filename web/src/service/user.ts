import { accessTokenKey, PaginationQueryResponse } from '../api/api';
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
  paginationQueryUserInfo,
  updateUserInfo,
  CreateUserByAdminParam,
  addUserWithGroup,
  ping,
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
 * 存储用户信息到storage中
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

/**
 * 判断是否是导师
 */
export function isTutor(): boolean {
  const info = getUserInfoFromStorage();
  if (!info) {
    return false;
  }
  for (const level of info.Levels) {
    if (level == UserLevels.Tutor) {
      return true;
    }
  }
  return false;
}

/**
 * 判断是否是管理员
 */
export function isAdmin(): boolean {
  const info = getUserInfoFromStorage();
  if (!info) {
    return false;
  }
  for (const level of info.Levels) {
    if (level == UserLevels.CommonAdmin || level == UserLevels.SuperAdmin) {
      return true;
    }
  }
  return false;
}
/**
 * 判断是否是超级管理员
 */
export function isSuperAdmin(): boolean {
  const info = getUserInfoFromStorage();
  if (!info) {
    return false;
  }
  for (const level of info.Levels) {
    if (level == UserLevels.SuperAdmin) {
      return true;
    }
  }
  return false;
}

/**
 * 分页查询用户信息
 */
export async function paginationGetUserInfo(
  pageIndex: number,
  pageSize: number
): Promise<PaginationQueryResponse<UserInfo>> {
  return await paginationQueryUserInfo(pageIndex, pageSize);
}

/**
 * 通过用户ID修改用户的电话、学院以及邮箱信息
 */
export async function updateUserInfoByID(
  id: number,
  tel: string,
  college: string,
  email: string
) {
  return updateUserInfo({
    id,
    tel,
    college,
    email,
  });
}

/**
 *  管理员创建新用户并添加到用户组
 */
export async function createUserByAdmin(
  param: CreateUserByAdminParam
): Promise<number> {
  return addUserWithGroup(param);
}

/**
 * 服务ping
 */
export async function servicePing(): Promise<boolean> {
  try {
    await ping();
    return true;
  } catch (error) {
    return false;
  }
}
