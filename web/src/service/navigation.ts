import { UserLevels } from './user';
import { RouteRecordRaw, Router } from 'vue-router';

import ApplyGroup from '../components/guest/ApplyGroup.vue';
import GroupManager from '../components/admin/GroupManager.vue';
import CheckGroupApply from '../components/tutor/CheckGroupApply.vue';
import UserManager from '../components/admin/UserManager.vue';
import ProjectManager from '../components/ProjectManager.vue';
import NodeApplyManager from '../components/common/NodeApplyManager.vue';

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
          path: '/main/apply_group',
          name: 'ApplyGroup',
          component: ApplyGroup,
        },
        item: {
          name: '申请加入组',
          to: '/main/apply_group',
        },
      },
      {
        routerRaw: {
          path: '/main/project_manager',
          name: 'ProjectManager',
          component: ProjectManager,
        },
        item: {
          name: '项目管理',
          to: '/main/project_manager',
        },
      },
    ],
  ],
  [
    UserLevels.Common,
    [
      {
        routerRaw: {
          path: '/main/project_manager',
          name: 'ProjectManager',
          component: ProjectManager,
        },
        item: {
          name: '项目管理',
          to: '/main/project_manager',
        },
      },
      {
        routerRaw: {
          path: '/main/node_apply_manager',
          name: 'NodeApplyManager',
          component: NodeApplyManager,
        },
        item: {
          name: '机器节点申请',
          to: '/main/node_apply_manager',
        },
      },
    ],
  ],
  [
    UserLevels.Tutor,
    [
      {
        routerRaw: {
          path: '/main/tutor_check_group_apply',
          name: 'TutorCheckGroupApply',
          component: CheckGroupApply,
        },
        item: {
          name: '审核用户组申请',
          to: '/main/tutor_check_group_apply',
        },
      },
      {
        routerRaw: {
          path: '/main/tutor_manager_user',
          name: 'TutorManagerUser',
          component: UserManager,
        },
        item: {
          name: '用户管理',
          to: '/main/tutor_manager_user',
        },
      },
      {
        routerRaw: {
          path: '/main/project_manager',
          name: 'ProjectManager',
          component: ProjectManager,
        },
        item: {
          name: '项目管理',
          to: '/main/project_manager',
        },
      },
      {
        routerRaw: {
          path: '/main/node_apply_manager',
          name: 'NodeApplyManager',
          component: NodeApplyManager,
        },
        item: {
          name: '机器节点申请管理',
          to: '/main/node_apply_manager',
        },
      },
    ],
  ],
  [
    UserLevels.CommonAdmin,
    [
      {
        routerRaw: {
          path: '/main/group_manager',
          name: 'GroupManager',
          component: GroupManager,
        },
        item: {
          name: '管理用户组',
          to: '/main/group_manager',
        },
      },
      {
        routerRaw: {
          path: '/main/admin_check_group_apply',
          name: 'AdminCheckGroupApply',
          component: CheckGroupApply,
        },
        item: {
          name: '审核用户组申请',
          to: '/main/admin_check_group_apply',
        },
      },
      {
        routerRaw: {
          path: '/main/admin_manager_user',
          name: 'AdminManagerUser',
          component: UserManager,
        },
        item: {
          name: '用户管理',
          to: '/main/admin_manager_user',
        },
      },
      {
        routerRaw: {
          path: '/main/project_manager',
          name: 'ProjectManager',
          component: ProjectManager,
        },
        item: {
          name: '项目管理',
          to: '/main/project_manager',
        },
      },
      {
        routerRaw: {
          path: '/main/node_apply_manager',
          name: 'NodeApplyManager',
          component: NodeApplyManager,
        },
        item: {
          name: '机器节点申请审核',
          to: '/main/node_apply_manager',
        },
      },
    ],
  ],
  // 超级管理员和普通管理员都这项操作
  // NOTE: 由于一个用户不能同时是超级管理员以及普通管理员,因此这里注册同样的路由信息
  [
    UserLevels.SuperAdmin,
    [
      {
        routerRaw: {
          path: '/main/group_manager',
          name: 'GroupManager',
          component: GroupManager,
        },
        item: {
          name: '管理用户组',
          to: '/main/group_manager',
        },
      },
      {
        routerRaw: {
          path: '/main/admin_check_group_apply',
          name: 'AdminCheckGroupApply',
          component: CheckGroupApply,
        },
        item: {
          name: '审核用户组申请',
          to: '/main/admin_check_group_apply',
        },
      },
      {
        routerRaw: {
          path: '/main/admin_manager_user',
          name: 'AdminManagerUser',
          component: UserManager,
        },
        item: {
          name: '用户管理',
          to: '/main/admin_manager_user',
        },
      },
      {
        routerRaw: {
          path: '/main/project_manager',
          name: 'ProjectManager',
          component: ProjectManager,
        },
        item: {
          name: '项目管理',
          to: '/main/project_manager',
        },
      },
      {
        routerRaw: {
          path: '/main/node_apply_manager',
          name: 'NodeApplyManager',
          component: NodeApplyManager,
        },
        item: {
          name: '机器节点申请审核',
          to: '/main/node_apply_manager',
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
 * 过滤出可用的所有NavigationItem项
 */
function filterAvailableNavigation(
  levels: UserLevels[]
): Map<UserLevels, UserNavigationItem[]> {
  const res = new Map<UserLevels, UserNavigationItem[]>();
  const flag = new Map<string, boolean>();
  for (const i in levels) {
    const level = levels[i];
    const navigationItem = UserNavigation.get(level);
    if (navigationItem == undefined) {
      continue;
    }
    const items = new Array<UserNavigationItem>();
    for (const item of navigationItem) {
      if (flag.has(item.routerRaw.path)) {
        continue;
      }
      flag.set(item.routerRaw.path, true);
      items.push(item);
    }
    if (items.length == 0) {
      continue;
    }
    res.set(level, items);
  }
  return res;
}

/**
 * 根据用户权限注册路由
 */
export function registryRouter(
  parentName: string,
  router: Router,
  levels: UserLevels[]
): number {
  let num = 0;
  const itemMap = filterAvailableNavigation(levels);
  itemMap.forEach((items) => {
    for (const item of items) {
      num++;
      router.addRoute(parentName, item.routerRaw);
    }
  });
  return num;
}

export function getAvailableNavigation(
  levels: UserLevels[]
): Map<UserLevels, NavigationItem[]> {
  const maps = filterAvailableNavigation(levels);
  const resMaps = new Map<UserLevels, NavigationItem[]>();
  maps.forEach((userNavigations, key) => {
    const items = new Array<NavigationItem>();
    for (const userNavigation of userNavigations) {
      items.push(userNavigation.item);
    }
    resMaps.set(key, items);
  });
  return resMaps;
}
