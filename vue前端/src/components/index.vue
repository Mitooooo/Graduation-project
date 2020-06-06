<template>
  <div id="app">
    <el-container style="height: 100%;" direction="vertical">
      <el-header>

        <el-col :span="5" class="header">
          <el-radio-group v-model="isCollapse" style="margin-top: 10px;">
<!--            <el-button type="text" @click="button" icon="el-icon-s-fold" size="medium " v-model="statue">{{statue}}</el-button>-->
            <el-button type="info" @click="button" icon="el-icon-s-fold" circle></el-button>
          </el-radio-group>
          </el-col>
          <el-col :span=19 >
            <el-dropdown trigger="hover" class="userinfo">
              <span class="el-dropdown-link userinfo-inner">
                <img src="https://ss0.bdstatic.com/94oJfD_bAAcT8t7mm9GUKT-xh_/timg?image&quality=100&size=b4000_4000&sec=1577866459&di=788d6a58e2a9b0c10917400ad8d983f4&src=http://b-ssl.duitang.com/uploads/item/201112/21/20111221094656_dUrQN.jpg" />
                 Admin
              </span>
              <el-dropdown-menu slot="dropdown">
                <el-dropdown-item  @click.native="logout">退出登录</el-dropdown-item>
              </el-dropdown-menu>

            </el-dropdown>
          </el-col>

      </el-header>
      <el-container>
        <el-aside height="100%" width="initial">
          <el-menu
            default-active="$route.path"
            class="el-menu-vertical-demo"
            background-color="#545c64"
            text-color="#fff"
            active-text-color="#ffd04b"
            @open="handleOpen"
            @close="handleClose"
            :collapse="isCollapse"
            router
          >

            <el-menu-item index="/enterAssets" >
                <i class="el-icon-location"></i>
                <span slot="title">输入资产</span>
            </el-menu-item>
            <el-menu-item index="/assetCollection">
              <i class="el-icon-menu"></i>
              <span slot="title">资产收集</span>
            </el-menu-item>
            <el-menu-item index="/vulnerAbility">
              <i class="el-icon-document"></i>
              <span slot="title">漏洞概况</span>
            </el-menu-item>
            <el-menu-item index="/pluginUpload">
              <i class="el-icon-setting"></i>
              <span slot="title">插件上传</span>
            </el-menu-item>
            <el-menu-item index="/singletonDetection">
              <i class="el-icon-magic-stick"></i>
              <span slot="title">单例检测</span>
            </el-menu-item>
            <el-menu-item index="/reportDisplay">
              <i class="el-icon-setting"></i>
              <span slot="title">报表展示</span>
            </el-menu-item>
          </el-menu>
        </el-aside>
        <el-main>
            <router-view/>
        </el-main>
      </el-container>
    </el-container>

  </div>

</template>

<script>
export default {
  data () {
    return {
      isCollapse: false,
      statue: '收起',
    }
  },
  methods: {
    handleOpen (key, keyPath) {
      console.log(key, keyPath)
    },
    handleClose (key, keyPath) {
      console.log(key, keyPath)
    },
    button () {
      if (this.isCollapse) {
        this.isCollapse = false
        this.statue = '收起'
      } else {
        this.isCollapse = true
        this.statue = '展开'
      }
    },
    logout () {
      console.log(this.$store.state.token)
      if (this.$store.state.token) {
        this.$store.commit('del_token')
        this.$message({
          message: '注销成功',
          type: 'success',
          duration: 500
        })
        this.$router.push('/')
      }
    },
  }
}
</script>

<style>
  /*将整个conotainer撑满*/
  html,body,#app,.el-container{
    /*设置内部填充为0，几个布局元素之间没有间距*/
    padding: 0;
    /*外部间距也是如此设置*/
    margin: 0;
    /*统一设置高度为100%*/
    height: 100%;
  }
  .el-header, .el-footer {
    background-color: rgba(17, 69, 97, 0.54);
    color: #333;
    text-align: center;
    line-height: 60px;
  }

  .el-aside {
    background-color: #D3DCE6;
    color: #333;
    text-align: center;
    line-height: 200px;
  }
  body > .el-container {
    margin-bottom: 40px;
  }

  .el-container:nth-child(5) .el-aside,
  .el-container:nth-child(6) .el-aside {
    line-height: 260px;
  }

  .el-container:nth-child(7) .el-aside {
    line-height: 320px;
  }
  .el-menu{
    height: 100%;
  }
  .el-menu-vertical-demo:not(.el-menu--collapse) {/*设置侧边栏不折叠的宽度*/
    width: 200px;
    height: 100%;
  }
  .el-menu--collapse{/*设置侧边栏折叠后的高度*/
    height: 100%;
  }
  .userinfo {
    text-align: right;
    padding-right: 35px;
    float: right;
  }
  .userinfo-inner {
    cursor: pointer;
    color: #fff;
  }
  img {
    width: 40px;
    height: 40px;
    border-radius: 20px;
    margin: 10px 0px 10px 10px;
    float: right;
  }
</style>
