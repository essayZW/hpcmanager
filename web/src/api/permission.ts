import { ApiRequest } from './api';

/**
 * 权限信息消息定义
 */
export type PermissionInfo = {
  id: number;
  name: string;
  level: number;
  description: string;
  createTime: number;
  extraAttributes: string;
};

/**
 * 查询用户的权限信息
 */
export async function queryUserPermission(
  id: number
): Promise<PermissionInfo[]> {
  const resp = await ApiRequest.request<PermissionInfo[]>(
    `/permission/user/${id}`,
    'GET'
  );
  if (!resp.data) {
    throw new Error(resp.message);
  }
  return resp.data;
}

/**
 * 设置新的管理员
 */
export async function setAdmin(userID: number): Promise<boolean> {
  const resp = await ApiRequest.request(
    '/permission/admin',
    'POST',
    {},
    {
      userID,
    }
  );

  if (!resp.status) {
    throw new Error(resp.message);
  }
  return true;
}

/**
 * 删除管理员
 */
export async function delAdmin(userID: number): Promise<boolean> {
  const resp = await ApiRequest.request(
    '/permission/admin',
    'DELETE',
    {},
    {
      userID,
    }
  );

  if (!resp.status) {
    throw new Error(resp.message);
  }
  return true;
}
