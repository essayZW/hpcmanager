<script setup lang="ts">
import { reactive } from 'vue';
import LogoImageSrc from '../assets/logo.png';
import { getUserInfoFromStorage, logout as userLogout } from '../service/user';
import { LoginUserInfo } from '../api/user';
import { useRouter } from 'vue-router';

import AsideNavigation from '../components/AsideNavigation.vue';

const router = useRouter();

const loginInfo = reactive<{ userInfo: LoginUserInfo }>({
  userInfo: {
    Username: 'unknown',
    Name: 'unknown',
    UserId: 0,
    GroupId: 0,
    Levels: [],
  },
});

const info = getUserInfoFromStorage();
if (info == null) {
  router.push({
    path: '/login',
  });
} else {
  loginInfo.userInfo = info;
}

// 退出登录处理函数
const logout = async () => {
  let status = await userLogout();
  if (!status) {
    ElMessage({
      type: 'error',
      message: '退出登录失败',
    });
  } else {
    router.push({
      path: '/login',
    });
  }
};
</script>
<template>
  <el-container>
    <el-header class="header">
      <div class="logo-title">
        <el-image :src="LogoImageSrc" class="logo"></el-image>
        <h1>计算平台管理系统</h1>
      </div>
      <div class="login-user">
        <el-dropdown trigger="hover">
          <span>
            <el-icon class="el-icon--left"> <i-ep-avatar /> </el-icon>
            {{ loginInfo.userInfo.Username }}
            <el-icon class="el-icon--right">
              <i-ep-arrow-down />
            </el-icon>
          </span>
          <template #dropdown>
            <el-dropdown-menu>
              <el-dropdown-item @click="logout">退出登录</el-dropdown-item>
              <el-dropdown-item
                ><router-link to="/main/update_user_info" class="link"
                  >修改用户信息</router-link
                ></el-dropdown-item
              >
            </el-dropdown-menu>
          </template>
        </el-dropdown>
      </div>
    </el-header>
    <el-container>
      <el-aside class="aside">
        <AsideNavigation></AsideNavigation>
      </el-aside>
      <el-container class="main-content-area">
        <el-main class="main-content">
          <router-view></router-view>
        </el-main>
        <el-footer class="footer"
          >&copy; 2022 essay.AllRightsReserved
        </el-footer>
      </el-container>
    </el-container>
  </el-container>
</template>
<style scoped lang="less">
// footer区域的高度
@footerheight: 80px;
// height区域的高度
@headerheight: 60px;
.header {
  border-bottom: 1px solid var(--el-border-color-base);
  line-height: @headerheight;
  display: flex;
  justify-content: space-between;
  .logo-title {
    .logo {
      width: 50px;
      height: 50px;
      display: inline-block;
    }
    h1 {
      font-size: 24px;
      margin: 0px 12px;
      display: inline-block;
      vertical-align: top;
    }
  }

  .login-user {
    cursor: pointer;
    span {
      height: @headerheight;
      line-height: @headerheight;
    }
  }
}
.aside {
  width: 160px;
  border-right: 1px solid var(--el-border-color-base);
}
.main-content-area {
  width: 100%;
  .main-content {
    height: calc(100% - @footerheight);
  }
  .footer {
    height: @footerheight;
    line-height: @footerheight;
    text-align: center;
    border-top: 1px solid var(--el-border-color-base);
  }
}
.link {
  text-decoration: none;
  color: inherit;
  font-size: inherit;
}
</style>
