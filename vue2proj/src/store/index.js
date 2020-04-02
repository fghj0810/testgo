import Vue from 'vue'
import Vuex from 'vuex'

Vue.use(Vuex)

export default new Vuex.Store({
    state: {
        loginStatus: false
    },
    mutations: {},
    actions: {},
    modules: {},
    getters: {
        isLogin: state => {
            return state.loginStatus
        }
    }
})
