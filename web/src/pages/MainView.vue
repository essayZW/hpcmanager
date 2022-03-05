<script setup lang="ts">
import LogoImageSrc from '../assets/logo.png';
</script>
<template>
  <el-container>
    <el-header class="header">
      <div class="logo-title">
        <el-image :src="LogoImageSrc" class="logo"></el-image>
        <h1>计算平台管理系统</h1>
      </div>
      <div class="login-user">Login User</div>
    </el-header>
    <el-container>
      <el-aside class="aside"> Aside </el-aside>
      <el-container class="main-content-area">
        <el-main class="main-content"> Main </el-main>
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
  }
}
.aside {
  width: 150px;
  padding: 12px 8px;
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
</style>
