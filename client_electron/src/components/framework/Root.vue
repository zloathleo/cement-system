<template>

    <div style="height: 100%;">
        <section class="hero color-primary-3">
            <Toolbar />
        </section>
        <b-loading :is-full-page="true" :active.sync="isAppLoading"></b-loading>
        <Drawer />

        <main class="main-container color-primary-3">
            <div class="content">
                <transition name="fade">
                    <router-view></router-view>
                </transition>
            </div>
        </main>
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
</style>


<script> 
import Toolbar from './Toolbar.vue';
import Drawer from './Drawer.vue';
import Footer from './Footer.vue';
export default {
    components: { Toolbar, Drawer, Footer },

    data() {
        return {
            isAppLoading: false,//是否loading状态
        }
    },
    beforeMount() {
        let _this = this;
        this.$globalEventHub.$on("appLoading", function (value) {
            _this.isAppLoading = value;
        }); 

        if (this.$globalvar.configMode) {
            this.$router.replace({ name: "testconfig" });
        }
    },
    methods: {
    }
}
</script>