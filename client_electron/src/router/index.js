import Vue from 'vue';
import Router from 'vue-router';

Vue.use(Router);

export default new Router({
  routes: [
    {
      name: 'index',
      path: '/',
      redirect: { name: 'dashboard' }
    },
    {
      path: '/',
      component: resolve => require(['../components/framework/Root.vue'], resolve),
      children: [
        {
          name: 'testconfig',
          path: '/testconfig',
          component: resolve => require(['../components/testconfig/Content.vue'], resolve),
        },
        {
          name: 'dashboard',
          path: '/dashboard',
          component: resolve => require(['../components/dashboard/Content.vue'], resolve),
        },
        {
          name: 'tpcontrol',
          path: '/tpcontrol',
          component: resolve => require(['../components/tpcontrol/Content.vue'], resolve),
        },
        {
          name: 'alarm',
          path: '/alarm',
          component: resolve => require(['../components/alarm/Content.vue'], resolve),
        },
        {
          name: 'settings',
          path: '/settings',
          component: resolve => require(['../components/settings/Content.vue'], resolve),
        }
      ]
    }
  ]
})
