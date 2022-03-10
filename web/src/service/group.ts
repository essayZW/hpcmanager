import {
  paginationQueryGroup,
  GroupInfo,
  createGroup as createGroupApi,
} from '../api/group';
import { PaginationQueryResponse } from '../api/api';
import { getUserIdByUsername } from './user';
import moment from 'moment';

/**
 * 分页查询用户组信息
 */
export async function paginationGetGroupInfo(
  pageIndex: number,
  pageSize: number
): Promise<PaginationQueryResponse<GroupInfo> | string> {
  try {
    const resp = await paginationQueryGroup(pageIndex, pageSize);
    const data = resp.Data;
    for (const i in data) {
      data[i].createTime = moment((data[i].createTime as number) * 1000).format(
        'YYYY年MM月DD日'
      );
    }
    return resp;
  } catch (error) {
    return `${error}`;
  }
}

/**
 * 创建组
 */
export async function createGroup(
  queueName: string,
  groupName: string,
  tutorUsername: string
): Promise<number> {
  // 验证用户名是否正确
  const userInfo = await getUserIdByUsername(tutorUsername);
  if (userInfo.groupID != 0) {
    throw new Error('该用户已经属于一个组,不能成为新组的导师');
  }
  return await createGroupApi({
    groupName: groupName,
    queueName: queueName,
    tutorID: userInfo.id,
  });
}
