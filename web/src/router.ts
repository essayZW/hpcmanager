import { createRouter, createWebHistory, RouteRecordRaw } from 'vue-router';
import { isInstall, getCasConfig } from './service/sys';
import { isLogin } from './service/user';

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
    children: [],
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
