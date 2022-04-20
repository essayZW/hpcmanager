import {
  HpcGroup,
  getHpcGroupInfo,
  queryHpcUserInfo,
  UserQuotaInfo,
  queryUserQuotaByUserHpcID,
  setUserQuota,
} from '../api/hpc';
import { HpcUser, ping } from '../api/hpc';
/**
 * 通过ID查询hpc_group信息
 */
export async function getHpcGroupInfoByID(
  id: number
): Promise<HpcGroup | string> {
  try {
    return getHpcGroupInfo(id);
  } catch (error) {
    return `${error}`;
  }
}

/**
 *
 */
export async function getHpcUserInfoByID(id: number): Promise<HpcUser> {
  return await queryHpcUserInfo(id);
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

/**
 * 查询用户存储使用情况信息
 */
export async function getHpcUserQuotaInfo(id: number): Promise<UserQuotaInfo> {
  return queryUserQuotaByUserHpcID(id);
}

/**
 * 修改用户存储信息
 */
export async function modifyUserQuotaInfo(
  hpcUserID: number,
  oldSize: number,
  newSize: number,
  oldEndTimeMilliUnix: number,
  newEndTimeMilliUnix: number,
  modifyDate: boolean
): Promise<boolean> {
  return setUserQuota({
    hpcUserID,
    oldSize,
    newSize,
    oldEndTimeMilliUnix,
    newEndTimeMilliUnix,
    modifyDate,
  });
}
