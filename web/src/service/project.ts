import {
  ProjectInfo,
  paginationQueryProjectInfos,
  createProject as apiCreateProject,
  queryByID,
} from '../api/project';
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

/**
 * 创建一个新的项目并返回项目的ID
 * @param name 项目名称
 * @param from 项目来源
 * @param numbering 项目编号
 * @param expenses 项目经费
 * @param description 项目描述
 * @returns 新项目的ID
 */
export async function createProject(
  name: string,
  from = '',
  numbering = '',
  expenses = '',
  description = ''
): Promise<number> {
  return apiCreateProject({
    name,
    from,
    numbering,
    expenses,
    description,
  });
}

/**
 * 通过ID查询项目信息
 */
export async function getProjectInfoByID(id: number): Promise<ProjectInfo> {
  return queryByID(id);
}
