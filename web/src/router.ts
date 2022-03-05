import { createRouter, createWebHistory, RouteRecordRaw } from 'vue-router';
import { isInstall } from './service/sys';
import { IsLogin } from './service/user';

import MainView from './pages/MainView.vue';
import NotFound from './pages/NotFound.vue';
import InstallView from './pages/InstallView.vue';
import LoginView from './pages/LoginView.vue';

const Router: RouteRecordRaw[] = [
  {
    path: '/',
    name: 'Main',
    component: MainView,
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
      const userInfo = await IsLogin();
      if (userInfo != null) {
        ElMessage({
          message: '已经登录,跳转到主页面',
        });
        return '/';
      }
    },
  },
  {
    path: '/:pathMatch(.*)',
    name: 'NotFound',
    component: NotFound,
  },
];

export default createRouter({
  history: createWebHistory(),
  routes: Router,
});
