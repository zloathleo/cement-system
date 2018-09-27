import Vue from 'vue';
import Vuex from 'vuex';
import createPersistedState from "vuex-persistedstate";
import storejs from 'store';

Vue.use(Vuex);

const vuexStore = new Vuex.Store({
    state: {
        user: undefined
    },
    mutations: {
        setUser: function (state, value) {
            state.user = value;
        },
    },
    plugins: [
        createPersistedState({
            storage: {
                getItem: function (key) {
                    return storejs.get(key);
                },
                setItem: function (key, value) {
                    storejs.set(key, value);
                },
                removeItem: function (key) {
                    storejs.remove(key);
                }
            }
        })
    ]
});

export default vuexStore;