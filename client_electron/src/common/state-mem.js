import Vuex from 'vuex';
import globalvar from './globalvar';

const store = new Vuex.Store({
    state: {
        user: undefined,
        dashboardRouteName: undefined,
        currentRouteName: undefined,
        menuItems: undefined,
        serverTimestamp: 0,
    },
    mutations: {
        setServerTimestamp: function (state, value) { 
            state.serverTimestamp = value;
        },

        initUserUI: function (state, value) {
            state.user = value;
            if (value === undefined) {
                state.dashboardRouteName = undefined;
                state.menuItems = undefined;
            } else {
                if (value.type == "admin") {
                    state.dashboardRouteName = "group";
                    state.currentRouteName = "group";
                    state.menuItems = globalvar.adminMenuItems;
                } else if (value.type == "operator") {
                    state.dashboardRouteName = "device";
                    state.currentRouteName = "device";
                    state.menuItems = globalvar.operatorMenuItems;
                }
            }
        },
        refreshUserUI: function (state, value) {
            this.commit("initUserUI", value.user);
            state.currentRouteName = value.routeName;
        },
    },
});

export default store