<template>

    <div style="height: 100%;">
        <section class="hero color-primary-3">
            <Toolbar />
        </section>
        <b-loading :is-full-page="true" :active.sync="isAppLoading"></b-loading>
        <Drawer />

        <main class="main-container color-primary-3" style="padding-bottom: 4rem; margin-bottom: 4rem;">
            <div class="content main-content" style="height: 100%;overflow-y: auto;">
                <transition name="fade">
                    <router-view></router-view>
                </transition>
            </div>
        </main>
        <Notification />
        <Footer />
    </div>

</template>

<style scoped lang="less">
.main-container {
  height: 100%;
  .content {
    padding: 2rem;
  }
}
.fade-enter-active {
  transition: all 0.3s ease;
}
.fade-leave-active {
  transition: all 0s;
}
.fade-enter,
.fade-leave-to {
  transform: translateY(300px);
  opacity: 0;
}

.main-content::-webkit-scrollbar-track {
  border-radius: 5px;
  background-color: #f5f5f5;
}
.main-content::-webkit-scrollbar {
  width: 6px;
  background-color: #f5f5f5;
}

.main-content::-webkit-scrollbar-thumb {
  border-radius: 10px;
  background-color: rgba(0, 0, 0, 0.6);
}
</style>


<script> 
import Toolbar from './Toolbar.vue';
import Drawer from './Drawer.vue';
import Footer from './Footer.vue';
import Notification from './Notification.vue';
export default {
    components: { Toolbar, Drawer, Footer, Notification },

    data() {
        return {
            isAppLoading: false,//是否loading状态
            timer: undefined,//实时定时器
        }
    },
    beforeMount() {
        let _this = this;
        this.$globalEventHub.$on("appLoading", function (value) {
            _this.isAppLoading = value;
        });
        this.initTimer();
    },
    destroyed() {
        if (this.timer) {
            clearInterval(this.timer);
        }
    },
    methods: {
        initTimer() {
            this.timer = setInterval(this.requestRealtimeData, 1000);
        },
        requestRealtimeData() {
            let routerName = this.$router.history.current.name;
            let points = "";
            if (routerName == "dashboard") {
                points = this.$globalvar.points_dashboard.join(",");
            } else if (routerName == "tpcontrol") {
                points = this.$globalvar.points_control.join(",");
            } else {
                return;
            }
            if (!this.isAppLoading) {
                let _this = this;
                _this.$myaxios.get('/realtime?points=' + points, { timeout: 1000, 'hiddenLoading': true }).then(function (response) {
                    let _data = response.data;
                    _this.$stateMem.commit("setRealtimePointValueMap", _data);
                }).catch(function (err) {
                    // _this.$stateMem.commit("setServerTimestamp", 0);
                    console.error(err);
                    // reject(err);
                });
            }
        }
    }
}
</script>