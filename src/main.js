import Vue from 'vue'
import VueRouter from 'vue-router'
import Homepage from './components/Homepage'
import About from './components/About'

Vue.use(VueRouter)
/* eslint-disable no-new */
var app = Vue.extend({})
var router = new VueRouter()
router.map(
  {
    '/': {
      component: Homepage
    },
    'About': {
      component: About
    }
  }
)
router.redirect({
  '*': '/'
})
router.start(app, '#app')

