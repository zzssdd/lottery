import Vue from 'vue'
import VueRouter from 'vue-router'
import Home from '../views/Home.vue'
import NotFound from '../views/errors/NotFound.vue'
import Payment from '../views/Payment.vue'
import Turn from'../views/Turn.vue'

Vue.use(VueRouter)

const routes = [
  {
    path: '/',
    name: 'Home',
    component: Home
  },
  {
    path:'/choose/:id',
    name:'Turn',
    query:"cost",
    component:Turn
  },
  {
    path: '*',
    name: 'NotFound',
    component: NotFound
  },
  {
    path:'/payment',
    name:'payment',
    component:Payment
  },
]

const router = new VueRouter({
  mode: 'history',
  base: process.env.BASE_URL,
  routes
})

export default router
