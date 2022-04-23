import { PaginationQueryResponse } from '../api/api';
import {
  ping,
  createPaperAwardApply,
  PaperApply,
  paginationQueryPaperApply,
  checkPaperApplyByID,
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
