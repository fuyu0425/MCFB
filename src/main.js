import Vue from 'vue'
import VueRouter from 'vue-router'
import Homepage from './components/Homepage'
import About from './components/About'
import Hello from './components/Hello'
Vue.use(VueRouter)
/* eslint-disable no-new */
var app = Vue.extend({})
var router = new VueRouter( {hashbang : false , history: true,})
router.map(
  {
    '/': {
      component: Homepage
    },
    'About': {
      component: About
    },
    'Hello': {
      component: Hello
    }
  }
)
router.redirect({
  '*': '/'
})
router.start(app, '#app')

