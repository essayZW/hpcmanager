import { ApiRequest, PaginationQueryResponse } from './api';
/**
 * 项目信息
 */
export type ProjectInfo = {
  id: number;
  name: string;
  from: string;
  numbering: string;
  expenses: string;
  description: string;
  createrUserID: number;
  createrUsername: string;
  createrName: string;
  createTime: number;
  modifyUserID: number;
  modifyUsername: string;
  modifyName: string;
  modifyTime: number;
  extraAttributes: string;
};

/**
 * 分页查询用户项目信息
 */
export async function paginationQueryProjectInfos(
  pageIndex: number,
  pageSize: number
): Promise<PaginationQueryResponse<ProjectInfo>> {
  const resp = await ApiRequest.request<PaginationQueryResponse<ProjectInfo>>(
    '/project',
    'GET',
    {
      pageIndex,
      pageSize,
    }
  );
  if (!resp.status) {
    throw new Error(resp.message);
  }
  return resp.data;
}
