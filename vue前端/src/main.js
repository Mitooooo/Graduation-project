// The Vue build version to load with the `import` command
// (runtime-only or standalone) has been set in webpack.base.conf with an alias.
import Vue from 'vue'
import App from './App'
import router from './router'
import store from '../store'
import axios from 'axios'
import Router from 'vue-router'
import ElementUI from 'element-ui'
import 'element-ui/lib/theme-chalk/index.css'
Vue.use(ElementUI)
Vue.prototype.$axios = axios
Vue.config.productionTip = false

const originalPush = Router.prototype.push
Router.prototype.push = function push (location) {
  return originalPush.call(this, location).catch(err => err)
}
/* eslint-disable no-new */
new Vue({
  el: '#app',
  router,
  store,
  components: { App },
  template: '<App/>'
})
router.beforeEach((to, from, next) => {
  next()
})

export default router
axios.interceptors.request.use(config => {
  if (store.state.token) {
    config.headers.common['Authentication-Token'] = store.state.token
   }
  return config
// eslint-disable-next-line handle-callback-err
}, error => {
  return console.log(1)
}
)

axios.interceptors.response.use((response) => {
  return response
// eslint-disable-next-line handle-callback-err
}, error => {
  console.log(error.response.status)
  switch (error.response.status) {
    case 401:
      // parent.location.href ='/';
    // // eslint-disable-next-line no-fallthrough
    // case 504:
    //   console.log('服务器异常')
    // // eslint-disable-next-line no-fallthrough
    // case 404:
    //   console.log('未找到资源')
  }
})
