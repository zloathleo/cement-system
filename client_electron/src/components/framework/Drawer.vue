<template>

    <div>
        <div v-show="isShowModal" class="modal-background" @click="isShowModal = false"></div>
        <transition name="slide-fade">
            <div v-show="isShowDraw" class="drawer-content color-primary-4">
                <a class="dropdown-item color-text-white" @click="clickMenu('dashboard')">
                    <span class="icon exicon">
                        <img class="icon-image" src="assets/img/dashboard.png">
                    </span>
                    <h3> Burner<br> Overview </h3>
                </a>
                <a class="dropdown-item color-text-white" @click="clickMenu('tpcontrol')">
                    <span class="icon exicon">
                        <img class="icon-image" src="assets/img/control.png">
                    </span>
                    <h3> Device<br> control </h3>
                </a>
                <a class="dropdown-item color-text-white" @click="clickMenu('settings')">
                    <span class="icon exicon">
                        <img class="icon-image" src="assets/img/settings.png">
                    </span>
                    <h3> Setting </h3>
                </a>
                <a class="dropdown-item color-text-white" @click="clickMenu('alarm')">
                    <span class="icon exicon">
                        <img class="icon-image" src="assets/img/alarm.png">
                    </span>
                    <h3> Alarms </h3>
                </a>
            </div>
        </transition>

    </div>

</template>

<style scoped lang="less">
a:hover {
  background-color: #333 !important;
  color: #fff !important;
}

.exicon {
  padding-bottom: 1.5rem;
  height: 2rem;
  width: 2rem;
}

.icon-image {
  margin-top: 10px;
}

.dropdown-item {
  line-height: 1.2;
  padding: 1rem;
  text-align: center;
}
.modal-background {
  z-index: 20;
  background-color: rgba(10, 10, 10, 0.5);
}
.drawer-content {
  position: fixed;
  width: 7rem;
  padding-top: 1rem;
  height: 100%;
  z-index: 30;
  color: #fff;
}
/* 可以设置不同的进入和离开动画 */
/* 设置持续时间和动画函数 */
.slide-fade-enter-active {
  transition: all 0.3s ease;
}
.slide-fade-leave-active {
  transition: all 0.3s ease;
}
.slide-fade-enter,
.slide-fade-leave-to {
  transform: translateX(-240px);
  opacity: 0;
}
</style>

<script> 
export default {
    data() {
        return {
            isShowDraw: false,
            isShowModal: false,
            selectedKey: undefined,
            appName: "后台管理",
            items: undefined,
        }
    },
    mounted() {
        this.items = this.$stateMem.state.menuItems;
        this.selectedKey = this.$stateMem.state.currentRouteName;
        let _this = this;
        this.$globalEventHub.$on("showDrawer", function () {
            _this.isShowDraw = !_this.isShowDraw;
        });
    },

    methods: {
        clickMenu(key) {
            this.selectedKey = key;
            this.isShowDraw = false;
            this.$router.push({ name: key });
            // this.$globalEventHub.$emit("appLoading", true);
        },
    }
}
</script>