import { createRouter, createWebHistory, RouteRecordRaw } from 'vue-router';
import { isInstall, getCasConfig } from './service/sys';
import { isLogin, setUserInfoToStorage } from './service/user';
import getQuery from './utils/urlQuery';
import { accessTokenKey } from './api/api';

import MainView from './pages/MainView.vue';
import NotFound from './pages/NotFound.vue';
import InstallView from './pages/InstallView.vue';
import LoginView from './pages/LoginView.vue';
import UpdateUserInfo from './components/UpdateUserInfo.vue';
import MainIndex from './components/MainIndex.vue';

import { registryRouter } from './service/navigation';

let registerFlag = false;

const Router: RouteRecordRaw[] = [
  {
    path: '/',
    name: 'Index',
    redirect: '/main',
  },
  {
    path: '/main',
    name: 'Main',
    component: MainView,
    redirect: '/main/index',
    beforeEnter: async () => {
      // 检查setToken参数
      const tokenValue = getQuery('setToken');
      if (tokenValue != null) {
        localStorage.setItem(accessTokenKey, tokenValue);
        window.location.href = '/';
      }
    },
    children: [
      {
        path: '/main/update_user_info',
        name: 'UpdateUserInfo',
        component: UpdateUserInfo,
      },
      {
        path: '/main/index',
        name: 'MainIndex',
        component: MainIndex,
      },
    ],
  },
  {
    path: '/install',
    name: 'Install',
    component: InstallView,
    beforeEnter: async () => {
      // 验证是否已经安装
      const status = await isInstall();
      if (status) {
        ElMessage({
          message: '已经安装,跳转到登录界面',
        });
        return '/login';
      }
    },
  },
  {
    path: '/login',
    name: 'Login',
    component: LoginView,
    beforeEnter: async () => {
      // 判断是否登录
      const userInfo = await isLogin();
      if (userInfo != null) {
        ElMessage({
          message: '已经登录,跳转到主页面',
        });
        return '/';
      }
      // 判断是否启用cas登录
      const config = await getCasConfig();
      if (config == null) {
        return;
      }
      if (!config.Enable) {
        return;
      }
      window.location.href = `${config.AuthServer}/cas/login?service=${config.ServiceAddr}${config.ValidPath}`;
    },
  },
  {
    path: '/:pathMatch(.*)',
    name: 'NotFound',
    component: NotFound,
  },
];

const router = createRouter({
  history: createWebHistory(),
  routes: Router,
});
router.beforeEach(async (to) => {
  // 尝试加载动态路由
  const userInfo = await isLogin();
  if (userInfo == null) {
    registerFlag = false;
    if (to.path != '/login' && to.path != '/install') {
      return '/login';
    }
    return;
  }
  setUserInfoToStorage(userInfo);
  if (!registerFlag) {
    const num = registryRouter('Main', router, userInfo.Levels);
    registerFlag = true;
    console.log(`register ${num} routers, redirect to ${to.fullPath}`);
    return to.fullPath;
  }
});

export default router;
