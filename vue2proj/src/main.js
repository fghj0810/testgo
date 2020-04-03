import Vue from 'vue'
import axios from 'axios'
import ElementUI from 'element-ui';
import router from './router'
import store from './store'
import App from './App.vue'
import 'element-ui/lib/theme-chalk/index.css';

Vue.use(ElementUI)
axios.defaults.baseURL = "http://localhost:9090"

Vue.config.productionTip = false

new Vue({
    router,
    store,
    render: h => h(App)
}).$mount('#app')
