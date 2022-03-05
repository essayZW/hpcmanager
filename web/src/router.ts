import { createRouter, createWebHistory, RouteRecordRaw } from 'vue-router';
import { isInstall } from './service/sys';

import MainView from './pages/MainView.vue';
import NotFound from './pages/NotFound.vue';
import InstallView from './pages/InstallView.vue';

const Router: RouteRecordRaw[] = [
  {
    path: '/',
    name: 'MainView',
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
    path: '/:pathMatch(.*)',
    name: 'NotFound',
    component: NotFound,
  },
];

export default createRouter({
  history: createWebHistory(),
  routes: Router,
});
