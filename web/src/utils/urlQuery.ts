// 获取url参数
export default function getQuery(name: string): string | null {
  const queryStr = window.location.search.substring(1);
  const querys = queryStr.split('&');
  for (const index in querys) {
    const kv = querys[index].split('=');
    if (kv.length != 2) {
      continue;
    }
    if (kv[0] === name) {
      return kv[1];
    }
  }
  return null;
}
