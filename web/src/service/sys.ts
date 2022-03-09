import {
  getInstallStatus,
  InstallRequest,
  installSys,
  CasConfig,
  loadCasConfig,
} from '../api/sys';

// install 初始化系统,添加系统默认管理员
export async function install(param: InstallRequest): Promise<{
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

// 获取系统cas配置
export async function getCasConfig(): Promise<CasConfig | null> {
  const serviceAddr = window.location.protocol + '//' + window.location.host;
  return await loadCasConfig(serviceAddr);
}
