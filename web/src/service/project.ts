import { ProjectInfo, paginationQueryProjectInfos } from '../api/project';
import { PaginationQueryResponse } from '../api/api';
/**
 * 分页查询用户项目信息
 */
export async function paginationGetProjectInfo(
  pageIndex: number,
  pageSize: number
): Promise<PaginationQueryResponse<ProjectInfo>> {
  return paginationQueryProjectInfos(pageIndex, pageSize);
}
