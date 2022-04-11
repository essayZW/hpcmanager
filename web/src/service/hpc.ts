import { HpcGroup, getHpcGroupInfo, queryHpcUserInfo } from '../api/hpc';
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
