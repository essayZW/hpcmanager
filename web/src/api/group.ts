import { ApiRequest, PaginationQueryResponse } from './api';

/**
 * 用户组信息
 */
export type GroupInfo = {
  id: number;
  name: string;
  createTime: number | string;
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

/**
 * 通过ID查询用户组信息
 */
export async function queryGroupByID(id: number): Promise<GroupInfo> {
  const resp = await ApiRequest.request<GroupInfo>(`/group/${id}`, 'GET');
  if (!resp.status) {
    throw new Error(resp.message);
  }
  return resp.data;
}

/**
 * 搜索导师信息的响应数据格式
 */
export type SearchTutorInfoResponse = {
  tutorID: number;
  tutorUsername: string;
  tutorName: string;
  groupID: number;
  groupName: string;
};

/**
 * 通过导师工号查询相关的信息
 */
export async function queryTutorInfoByUsername(
  username: string
): Promise<SearchTutorInfoResponse> {
  const resp = await ApiRequest.request<SearchTutorInfoResponse>(
    `/group/tutor/${username}`,
    'GET'
  );
  if (!resp.status) {
    throw new Error(resp.message);
  }
  return resp.data;
}

/**
 * 创建加入组申请接口返回数据
 */
export type CreateJoinGroupApplyResponse = {
  applyID: number;
};

/**
 * 创建加入组申请
 */
export async function createJoinGroupApply(
  applyGroupID: number
): Promise<CreateJoinGroupApplyResponse> {
  const resp = await ApiRequest.request<CreateJoinGroupApplyResponse>(
    '/group/apply',
    'POST',
    {},
    {
      applyGroupID,
    }
  );
  if (!resp.status) {
    throw new Error(resp.message);
  }
  return resp.data;
}
