import {
  delAdmin,
  PermissionInfo,
  queryUserPermission,
  setAdmin,
} from '../api/permission';

/**
 * 通过用户ID获取用户权限信息
 */
export async function getUserPermissionInfoByID(
  id: number
): Promise<PermissionInfo[]> {
  return queryUserPermission(id);
}

/**
 * 权限名称转换
 */
export function nameTransform(name: string): string {
  switch (name) {
    case 'SuperAdmin':
      return '超级管理员';
    case 'CommonAdmin':
      return '普通管理员';
    case 'Tutor':
      return '导师';
    case 'Common':
      return '学生';
    case 'Guest':
      return '游客';
    default:
      return 'unknown';
  }
}

/**
 * 通过用户ID设置新的管理员
 */
export async function setAdminByUserID(userID: number): Promise<boolean> {
  return setAdmin(userID);
}

/**
 * 通过用户ID删除用户管理员权限
 */
export async function delAdminByUserID(userID: number): Promise<boolean> {
  return delAdmin(userID);
}
