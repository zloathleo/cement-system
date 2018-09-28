<template>
  <div class="loading-parent">
    <div class="lds-dual-ring loading"></div>
  </div>
</template> 

<style scoped lang="less">
.loading-parent {
  widows: 100%;
  height: 100%;
  background-color: #505050;
  text-align: center;
}

.loading {
  top: calc(50% - 42px);
  position: absolute;
  margin-left: -32px;
}

.lds-dual-ring {
  display: inline-block;
  width: 64px;
  height: 64px;
}

.lds-dual-ring:after {
  content: " ";
  display: block;
  width: 46px;
  height: 46px;
  margin: 1px;
  border-radius: 50%;
  border: 5px solid #fff;
  border-color: #fff transparent #fff transparent;
  animation: lds-dual-ring 1.2s linear infinite;
}

@keyframes lds-dual-ring {
  0% {
    transform: rotate(0deg);
  }
  100% {
    transform: rotate(360deg);
  }
}
</style>

<script>     
export default {
  data() {
    return {
      show: true,
    }
  },
  mounted() {
    let _this = this;
    // this.$globalEventHub.$on("initLoading", function () {
    //   _this.show = false;
    // });
  },

}
</script>