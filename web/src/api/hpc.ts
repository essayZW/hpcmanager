import { undefinedWithDefault } from '../utils/obj';
import { ApiRequest } from './api';

/**
 * hpc_group信息
 */
export type HpcGroup = {
  id: number;
  name: string;
  queueName: string;
  gID: number;
  extraAttributes: string;
};

/**
 * 通过id查询hpc group信息
 */
export async function getHpcGroupInfo(id: number): Promise<HpcGroup> {
  const resp = await ApiRequest.request<HpcGroup>(`/hpc/group/${id}`, 'GET');
  if (!resp.status) {
    throw new Error(resp.message);
  }
  undefinedWithDefault(resp.data, 'id', 0);
  undefinedWithDefault(resp.data, 'name', '');
  undefinedWithDefault(resp.data, 'queueName', '');
  undefinedWithDefault(resp.data, 'gID', 0);
  undefinedWithDefault(resp.data, 'extraAttributes', '');
  return resp.data;
}

/**
 * hpc_user信息
 */
export type HpcUser = {
  id: number;
  nodeUsername: string;
  nodeUID: number;
  nodeMaxQuota: number;
  nodeStartTime: number;
  nodeEndTime: number;
  extraAttributes: string;
};

/**
 * 通过ID查询hpc用户信息
 * @param id hpc_user参数
 * @returns hpc_user信息
 */
export async function queryHpcUserInfo(id: number): Promise<HpcUser> {
  const resp = await ApiRequest.request<HpcUser>(`/hpc/user/${id}`, 'GET');
  if (!resp.status) {
    throw new Error(resp.message);
  }
  return resp.data;
}