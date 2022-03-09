/**
 * 如果某个对象的某个值是undefined则将该值设定为传递的值
 */
export function undefinedWithDefault(
  obj: Record<string | number, unknown>,
  key: keyof Record<string | number, unknown>,
  value: unknown
) {
  if (obj[key] == undefined) {
    obj[key] = value;
  }
}
