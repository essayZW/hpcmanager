import { PaginationQueryResponse } from '../api/api';
import {
  ping,
  createPaperAwardApply,
  PaperApply,
  paginationQueryPaperApply,
  checkPaperApplyByID,
  createTechnologyApply,
  paginationQueryTechnologyApply,
  TechnologyApply,
  checkTechnologyApply,
} from '../api/award';
import { uploadFileUrlPathBase } from '../api/fss';
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

/**
 * 获取支持的论文分类
 */
export function getPaperCategory(): string[] {
  return ['Nature', 'Science', 'SCI', 'EI'];
}

/**
 * 获取支持的论文分区
 */
export function getPaperPartition(): string[] {
  return ['1', '2', '3', '4'];
}

/**
 * 创建论文奖励申请
 */
export async function createPaperApply(
  title: string,
  category: string,
  partition: string,
  firstPageImageName: string,
  thanksPageImageName: string,
  remarkMessage: string
): Promise<number> {
  return createPaperAwardApply({
    title,
    category,
    partition,
    firstPageImageName,
    thanksPageImageName,
    remarkMessage,
  });
}

/**
 * 分页查询论文奖励申请信息
 */
export async function paginationGetPaperApply(
  pageIndex: number,
  pageSize: number
): Promise<PaginationQueryResponse<PaperApply>> {
  const data = await paginationQueryPaperApply(pageIndex, pageSize);
  for (const single of data.Data) {
    single.paperFirstPageImageName =
      uploadFileUrlPathBase + '/' + single.paperFirstPageImageName;
    single.paperThanksPageImageName =
      uploadFileUrlPathBase + '/' + single.paperThanksPageImageName;
  }
  return data;
}

/**
 * 审核论文奖励申请
 */
export async function checkPaperAwardApply(
  id: number,
  accept: boolean,
  money: number,
  message: string
): Promise<boolean> {
  return checkPaperApplyByID({
    id,
    accept,
    checkMoney: money,
    checkMessage: message,
  });
}

/**
 * 创建科技奖励申请
 */
export async function createTechnologyAwardApply(
  projectID: number,
  prizeLevel: string,
  prizeImageName: string,
  remarkMessage: string
): Promise<number> {
  return createTechnologyApply({
    projectID,
    prizeLevel,
    prizeImageName,
    remarkMessage,
  });
}

/**
 * 分页查询科技奖励申请
 */
export async function paginationGetTechnologyApply(
  pageIndex: number,
  pageSize: number
): Promise<PaginationQueryResponse<TechnologyApply>> {
  const data = await paginationQueryTechnologyApply(pageIndex, pageSize);
  for (const single of data.Data) {
    single.prizeImageName = uploadFileUrlPathBase + '/' + single.prizeImageName;
  }
  return data;
}

/**
 * 审核科技奖励申请
 */
export async function checkTechnologyApplyByID(
  id: number,
  money: number,
  message: string,
  accept: boolean
): Promise<boolean> {
  return checkTechnologyApply({
    id,
    checkMoney: money,
    checkMessage: message,
    accept,
  });
}

/**
 * 预设的奖项等级
 */
export const prizeLevels = [
  '国家级一等奖',
  '国家级二等奖',
  '国家级三等奖',
  '北京市一等奖',
  '北京市二等奖',
  '北京市三等奖',
];
