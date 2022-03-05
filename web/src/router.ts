import { createRouter, createWebHistory, RouteRecordRaw } from 'vue-router';

import MainView from './components/pages/MainView.vue';
import NotFound from './components/pages/NotFound.vue';

const Router: RouteRecordRaw[] = [
  {
    path: '/',
    name: 'MainView',
    component: MainView,
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
