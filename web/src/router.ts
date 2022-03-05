import { createRouter, createWebHistory, RouteRecordRaw } from 'vue-router';

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
