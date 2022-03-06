import { ElForm } from 'element-plus';
export type FormInstance = InstanceType<typeof ElForm>;

// 适用于username属性的验证规则
export function requiredWithLength(name: string, min = 6, max = 16) {
  return [
    { required: true, message: `${name}不能为空`, trigger: 'blur' },
    {
      min: min,
      max: max,
      message: `${name}长度只能在${min}-${max}之间`,
      trigger: 'change',
    },
  ];
}
