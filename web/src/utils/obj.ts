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

/**
 * 如果某个变量的值是零值则返回预设的值,否则返回原值
 */
export function zeroWithDefault(
  value: string | number | boolean,
  newValue: string | number | boolean
): string | number | boolean {
  if (!value) {
    return newValue;
  }
  return value;
}
