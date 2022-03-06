import { getInstallStatus, installRequest, installSys } from '../api/sys';

// install 初始化系统,添加系统默认管理员
export async function install(param: installRequest): Promise<{
  status: boolean;
  message: string;
}> {
  try {
    const resp = await installSys(param);
    return {
      status: resp,
      message: '安装成功',
    };
  } catch (error) {
    return {
      status: false,
      message: `${error}`,
    };
  }
}

// 判断系统是否已经被安装
export async function isInstall(): Promise<boolean> {
  return await getInstallStatus();
}
