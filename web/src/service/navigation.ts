import { UserLevels } from './user';
import { RouteRecordRaw, Router } from 'vue-router';

import ApplyGroup from '../components/guest/ApplyGroup.vue';
import GroupManager from '../components/admin/GroupManager.vue';

/**
 * 用户导航以及路由定义item
 */
type UserNavigationItem = {
  routerRaw: RouteRecordRaw;
  item: NavigationItem;
};

/**
 * 不同权限角色的用户导航定义
 */
export const UserNavigation = new Map<UserLevels, UserNavigationItem[]>([
  [
    UserLevels.Guest,
    [
      {
        routerRaw: {
          path: 'apply_group',
          name: 'ApplyGroup',
          component: ApplyGroup,
        },
        item: {
          name: '申请加入组',
          to: 'apply_group',
        },
      },
    ],
  ],
  [
    UserLevels.CommonAdmin,
    [
      {
        routerRaw: {
          path: 'group_manager',
          name: 'GroupManager',
          component: GroupManager,
        },
        item: {
          name: '创建用户组',
          to: 'group_manager',
        },
      },
    ],
  ],
  // 超级管理员和普通管理员都这项操作
  [
    UserLevels.SuperAdmin,
    [
      {
        routerRaw: {
          path: 'group_manager',
          name: 'GroupManager',
          component: GroupManager,
        },
        item: {
          name: '创建用户组',
          to: 'group_manager',
        },
      },
    ],
  ],
]);

/**
 * 导航item定义
 */
export type NavigationItem = {
  name: string;
  to: string;
};

/**
 * 根据用户权限注册路由
 */
export function registryRouter(
  parentName: string,
  router: Router,
  levels: UserLevels[]
): Map<UserLevels, NavigationItem[]> {
  const res = new Map<UserLevels, NavigationItem[]>();
  for (const i in levels) {
    const level = levels[i];
    const navigationItem = UserNavigation.get(level);
    if (navigationItem == undefined) {
      continue;
    }
    const items = new Array<NavigationItem>();
    for (const item of navigationItem) {
      items.push(item.item);
      router.addRoute(parentName, item.routerRaw);
    }
    if (items.length == 0) {
      continue;
    }
    res.set(level, items);
  }
  return res;
}
