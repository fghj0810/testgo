import Vue from 'vue'
import Vuex from 'vuex'

Vue.use(Vuex)

export default new Vuex.Store({
    state: {
        userInfo:{
            username : '',
            token : ''
        },
        isLogin : true
    },
    mutations: {
        setUserInfo (state, userInfo) {
            state.userInfo.username = userInfo.username
            state.userInfo.token = userInfo.token
        }
    },
    actions: {},
    modules: {},
    getters: {
        isLogin: state => {
            return state.userInfo.username.trim() !== '' && state.userInfo.token.trim() !== ''
        }
    }
})
