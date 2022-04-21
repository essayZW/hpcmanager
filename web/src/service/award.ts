import { ping, createPaperAwardApply } from '../api/award';
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
