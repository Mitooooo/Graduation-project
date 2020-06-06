import Vue from 'vue'
import Vuex from 'vuex'
Vue.use(Vuex)

const store = new Vuex.Store({
  state: {
    token: ''
  },
  mutations: {
    set_token (state, token) {
      state.token = token
      sessionStorage.token = token
    },
    del_token (state) {
      state.token = ''
      sessionStorage.removeItem('token')
    }
  }
})

export default store
