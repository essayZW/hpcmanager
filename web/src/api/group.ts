import { ApiRequest, PaginationQueryResponse } from './api';

/**
 * 用户组信息
 */
export type GroupInfo = {
  id: number;
  name: string;
  createTime: string;
  createrID: string;
  createUsername: string;
  createrName: string;
  tutorID: number;
  tutorUsername: string;
  tutorName: string;
  balance: number;
  hpcGroupID: number;
};

/**
 * 分页查询组信息
 */
export async function paginationQueryGroup(
  page: number,
  size: number
): Promise<PaginationQueryResponse<GroupInfo>> {
  const resp = await ApiRequest.request('/group', 'GET', {
    pageIndex: page,
    pageSize: size,
  });
  if (!resp.status) {
    throw new Error(resp.message);
  }
  return resp.data as PaginationQueryResponse<GroupInfo>;
}

/**
 * 创建组请求
 */
export type CreateGroupRequest = {
  tutorID: number;
  groupName: string;
  queueName: string;
};

export type CreateGroupResponse = {
  id: number;
};

/**
 * 根据组名称、私有队列名称以及导师ID，创建组并分配导师权限
 */
export async function createGroup(param: CreateGroupRequest): Promise<number> {
  const resp = await ApiRequest.request(
    '/group',
    'POST',
    {},
    {
      ...param,
    }
  );

  if (!resp.status) {
    throw new Error(resp.message);
  }
  return (resp.data as CreateGroupResponse).id;
}
