import Vue from 'vue';
import Vuex from 'vuex';
import Buefy from 'buefy';

import './common/initjs';

import globalvar from './common/globalvar';
import statePersisted from './common/state-persisted';
import stateMem from './common/state-mem';
import myaxios from './common/myaxios';
import router from './router';
import App from './App.vue';

//是否模拟数据
// import './common/mock';


import './assets/css/global.css';

import './assets/scss/bulma.scss'
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

new Vue({
  el: '#app',
  router,
  render: h => h(App)
});

//等待 electron 读取配置 
try {
  var ipcRenderer = require('electron').ipcRenderer;
  ipcRenderer.once("electron.main.config", function (_evt, _appConfig) {
    console.log("electron.main.config ok:", _appConfig);

    globalvar.points_dashboard = _appConfig.points_dashboard;
    globalvar.points_control = _appConfig.points_control;

    globalvar.dashboard_dcs_points = _appConfig.dashboard_dcs_points;


    globalvar.configMode = _appConfig.config_mode;
    globalvar.trned_points = _appConfig.trned_points;
    globalvar.dashboard_roundchart = _appConfig.dashboard_roundchart;
    globalvar.dashboard_radar = _appConfig.dashboard_radar;
    globalvar.dashboard_linechart = _appConfig.dashboard_linechart;
    globalvar.fetchServerHostURL = "http://" + _appConfig.fetch_server_host + ":" + _appConfig.fetch_server_port;

    console.log("读取配置成功,初始化myaxios:", globalvar.fetchServerHostURL);
    console.log("configMode:", globalvar.configMode);
    myaxios.defaults.baseURL = globalvar.fetchServerHostURL;

    initUI();
  });
  ipcRenderer.send("electron.renderer.mainpage.loaded", "ok");
} catch (ex) {
  initUI();
}

function initUI() {
  setTimeout(function () {
    // router.replace({ name: "tpcontrol" });
    // router.replace({ name: "dashboard" });
    // stateMem.commit("setUiTitle", "dashboard");

    var landpage = "dashboard";
    // var landpage = "tpcontrol";
    router.replace({ name: landpage });
    stateMem.commit("setUiTitle", landpage);
    // if (globalvar.configMode) {
    //   router.replace({ name: "testconfig" });
    // } else {
    //   router.replace({ name: "dashboard" });
    // }
  }, 100);
}

