import { createRouter, createWebHistory, RouteRecordRaw } from 'vue-router';
import { isInstall, getCasConfig } from './service/sys';
import { isLogin, setUserInfoToStorage } from './service/user';
import getQuery from './utils/urlQuery';
import { accessTokenKey } from './api/api';

import MainView from './pages/MainView.vue';
import NotFound from './pages/NotFound.vue';
import InstallView from './pages/InstallView.vue';
import LoginView from './pages/LoginView.vue';

const Router: RouteRecordRaw[] = [
  {
    path: '/',
    name: 'Index',
    redirect: '/main/',
  },
  {
    path: '/main/',
    name: 'Main',
    component: MainView,
    beforeEnter: async () => {
      // 检查setToken参数
      const tokenValue = getQuery('setToken');
      if (tokenValue != null) {
        localStorage.setItem(accessTokenKey, tokenValue);
        window.location.href = '/';
      }

      // 判断是否已经登录,未登录跳转到登录页面
      const info = await isLogin();
      if (info == null) {
        ElMessage({
          type: 'error',
          message: '未登录,请先登录',
        });
        return '/login';
      }
      // 存储用户信息到storge中
      setUserInfoToStorage(info);
    },
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
    beforeEnter: (to) => {
      if (/^\/main*/.test(to.fullPath)) {
        return {
          name: 'Main',
        };
      }
    },
  },
];

export default createRouter({
  history: createWebHistory(),
  routes: Router,
});
