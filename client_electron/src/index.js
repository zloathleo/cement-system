import 'promise-polyfill/src/polyfill';

import Vue from 'vue';
import Vuex from 'vuex';

import Buefy from 'buefy';
// import 'buefy/lib/buefy.css';

import './common/initjs';

import globalvar from './common/globalvar';
import statePersisted from './common/state-persisted';
import stateMem from './common/state-mem';
import myaxios from './common/myaxios';
import router from './router';
import App from './App';

//是否模拟数据
// import './common/mock';

import './assets/scss/bulma.scss'
import './assets/css/global.css';
import './assets/css/materialdesignicons.min.css';

Vue.use(Buefy, {
  defaultIconPack: 'mdi',
});


//vue内部常用
let globalEventHub = new Vue();
globalvar.GlobalEventHub = globalEventHub;
Vue.prototype.$myaxios = myaxios;
Vue.prototype.$globalEventHub = globalEventHub;
Vue.prototype.$globalvar = globalvar;
Vue.prototype.$stateMem = stateMem;

//electron
try {
  var ipcRenderer = require('electron').ipcRenderer;
  ipcRenderer.once("electron.main.config", function (_evt, _appConfig) {
    console.log("electron.main.config:", _appConfig);
    globalvar.configMode = _appConfig.config_mode;

    initUI();

  });
  ipcRenderer.send("electron.renderer.mainpage.loaded", "ok");
} catch (ex) {
  initUI();
}

function initUI() {
  new Vue({
    el: '#app',
    router,
    render: h => h(App)
  });
}


