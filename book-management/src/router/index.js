import Vue from 'vue'
import VueRouter from 'vue-router'

import Login from '../views/Login.vue'
import Home from '../views/Home.vue'
import Manage from "@/views/Manage";
import store from "../store";

Vue.use(VueRouter)

const routes = [
  {
    path: '/',
    redirect: '/login',
    meta: {
      requiresAuth: false
    }
  },
  {
    path: '/login',
    name: 'Login',
    component: Login,
    meta: {
      requiresAuth: false
    }
  },
  {
    path: '/home',
    name: 'Home',
    component: Home,
    meta: {
      requiresAuth: true
    }
  },
  {
    path: '/manage',
    name: 'Manage',
    component: Manage,
    meta: {
      requiresAuth: true
    }
  }
]

const router = new VueRouter({
  mode: 'history',
  base: process.env.BASE_URL,
  routes
})

router.beforeEach((to, from, next) => {
  if(to.matched.length !== 0) {
    if(to.meta.requiresAuth) {
      if(store.state.username && store.state.password) {
        next();
      }
      else {
        next({
          path: '/login',
          query: { redirect: to.fullPath },
        })
      }
    }
    else {
      if(store.state.username && store.state.password) {
        next({
          path: '/home'
        });
      }
      else {
        next();
      }
    }
  }
  else {
    next({
      path: '/login',
      query: { redirect: to.fullPath },
    })
  }
})

export default router
