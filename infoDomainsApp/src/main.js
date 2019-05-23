// The Vue build version to load with the `import` command
// (runtime-only or standalone) has been set in webpack.base.conf with an alias.
import Vue from 'vue'
import App from './App'
// for call external resources npm i vue-resource
import vueResource from 'vue-resource'
Vue.use(vueResource)

import vueRouter from 'vue-router'
Vue.use(vueRouter)

import DomainInfo from './components/DomainInfo';
import Report from './components/Report';

const router = new vueRouter({
  mode: 'history',
  base: __dirname,
  routes: [{
    path: '/',
    component: DomainInfo
  },
  {
      path: '/report',
      component: Report
    },
  ]
})

Vue.config.productionTip = false

/* eslint-disable no-new */
new Vue({
  router,
  // el: '#app',
  components: { App },
  template: '<App/>'
}).$mount('#app')
