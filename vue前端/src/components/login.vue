<template>
  <div>
    <div style="background: #ffffff">
    <div class="input-field col s12">
      <div class="" style="margin-top: 15%;width: 20%;margin-left: 40%;">
        <form action="" class="s10" method="post">
          <div class="row">
            <div class="input-field col s12">
              <input id="username" v-model="username" type="text" class="" style="background: #ffffff">
              <label for="username">用户名</label>
            </div>
          </div>
          <div class="row">
            <div class="input-field col s12">
              <input id="password" v-model="password" type="password" @keyup.enter="login()">
              <label for="password">密码</label>
            </div>
          </div>
        </form>
      </div>
    </div>
  </div>
  </div>
</template>

<script>
// eslint-disable-next-line no-unused-vars
import Materialize from '../../static/js/materialize.min.js'
// eslint-disable-next-line no-unused-vars
// import back from '../../static/js/back.js'

export default {
  data () {
    return {
      username: '',
      password: ''
    }
  },
  methods: {
    login () {
      if (this.username === '' || this.password === '') {
        alert('账号或密码不能为空')
      } else {
        let data = new FormData()
        data.append('username', this.username)
        data.append('password', this.password)
        this.$axios.post(`api/auth`, data).then((data) => {
          if (data) {
            this.$store.commit('set_token', data.data.data.token)
            console.log('token is set to:', this.$store.state.token)
            if (data.data.data.token) {
              this.$message({
                message: '登陆成功',
                type: 'success',
                duration: 1800 //500
              })
              this.$router.push('/enterAssets')
            } else {
              this.$message.error({
                message: '登陆失败，账号或密码错误',
                duration: 1800 //600
              })
              console.log('登陆失败')
            }
          }
        })
      }
    }
  }
}
</script>

<!-- Add "scoped" attribute to limit CSS to this component only -->
<style scoped>
  @import "../../static/css/bootstrap.css";
  @import "../../static/css/materialize.min.css";
  *{
    -webkit-box-sizing: border-box;
    box-sizing: border-box;
  }
  html, body{
    height: 100%;
    margin: 0;
    padding: 0;
    background-color: #FFF;
    overflow: hidden;
  }
  canvas{
    background-color: #FFF;
  }
</style>
