import Vuex from 'vuex';
import globalvar from './globalvar';

const store = new Vuex.Store({
    state: {
        serverTimestamp: 0,

        //标题
        uititle: "",

        //实时值map
        realtimePointValueMap: {},

        //趋势A
        trendA: "",
        //趋势B
        trendB: "",
        //进退
        jintui: {

        },

        isShowDraw: false,
    },
    mutations: {
        setUiTitle: function (state, value) {
            if (value == "dashboard") {
                state.uititle = "温度场总览";
            } else if (value == "tpcontrol") {
                state.uititle = "温度场设备";
            } else if (value == "trend") {
                state.uititle = "趋势分析";
            } else if (value == "alarm") {
                state.uititle = "系统报警";
            }
        },
        setTrendPoint: function (state, value) {
            state.trendA = value.a;
            state.trendB = value.b;
        },
        setIsShowDraw: function (state, value) {
            state.isShowDraw = value;
        },
        setRealtimePointValueMap: function (state, value) {
            state.realtimePointValueMap = value;
        },

        setServerTimestamp: function (state, value) {
            state.serverTimestamp = value;
        },
        setJinTui: function (state, value) {
            state.jintui = value;
        },
    },
});

export default store