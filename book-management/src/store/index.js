import Vue from 'vue'
import Vuex from 'vuex'

Vue.use(Vuex)

const store = new Vuex.Store({
  state: {
    username: null,
    password: null,
    role_id: null,
  },
  mutations: {
    saveData(state, data) {
      state.username = data.username;
      state.password = data.password;
      state.role_id = data.role_id;
    }
  },
  actions: {
  },
  modules: {
  }
})

export default store
