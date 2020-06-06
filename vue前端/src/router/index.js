import Vue from 'vue'
import Router from 'vue-router'
import login from '../components/login'
import store from '../../store'
import logon from '@/components/logon'
import index from '@/components/index'
import assetCollection from '../components/assetCollection'
import enterAssets from '../components/enterAssets'
import pluginUpload from '../components/pluginUpload'
import vulnerAbility from '../components/vulnerAbility'
import reportDisplay from '../components/reportDisplay'
import singletonDetection from '../components/singletonDetection'
Vue.use(Router)

export default new Router({
  routes: [
    {
      path: '/',
      name: 'login',
      component: login
    },
    {
      path: '/logon',
      name: 'logon',
      component: logon
    },
    {
      path: '/index',
      name: index,
      component: index,
      children: [
        {
          path: '/assetCollection',
          name: assetCollection,
          component: assetCollection
        },
        {
          path: '/enterAssets',
          name: enterAssets,
          component: enterAssets
        },
        {
          path: '/pluginUpload',
          name: pluginUpload,
          component: pluginUpload
        },
        {
          path: '/vulnerAbility',
          name: vulnerAbility,
          component: vulnerAbility
        },
        {
          path: '/reportDisplay',
          name: reportDisplay,
          component: reportDisplay
        },
        {
          path: '/singletonDetection',
          name: singletonDetection,
          component: singletonDetection
        }
      ]
    }
  ]
})
if (sessionStorage.getItem('token')) {
  store.commit('set_token', sessionStorage.getItem('token'))
}
