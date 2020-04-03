import Vue from 'vue'
import VueRouter from 'vue-router'
import Home from '../views/Home.vue'
import store from "../store";

Vue.use(VueRouter)

const routes = [
    {
        path: '/',
        name: 'Home',
        component: Home,
        meta: {
            requireAuth: true
        }
    },
    {
        path: '/login',
        name: 'Login',
        component: () => import(/* webpackChunkName: "about" */ '../views/Login.vue')
    },
    {
        path: '/mail',
        name: 'Mail',
        component: () => import(/* webpackChunkName: "about" */ '../views/Mail.vue')
    },
    {
        path: '/about',
        name: 'About',
        // route level code-splitting
        // this generates a separate chunk (about.[hash].js) for this route
        // which is lazy-loaded when the route is visited.
        component: () => import(/* webpackChunkName: "about" */ '../views/About.vue')
    }
]

const router = new VueRouter({
    mode: 'history',
    base: process.env.BASE_URL,
    routes
})


router.beforeEach((to, from, next) => {
    console.log(to.fullPath)
    if (to.meta.requireAuth) {
        if (store.getters.isLogin) {
            next();
        } else {
            next({
                path: '/login',
                query: {redirect: to.fullPath}
            })
        }
    } else {
        // 不需要登录的，可以继续访问
        next()
    }
})

export default router
