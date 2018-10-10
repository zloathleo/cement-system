<template> 

    <div class="mynotification">
        <b-notification type="is-info" :active.sync="isInfoActive">
            {{infoMessage}}
        </b-notification>

    </div>

</template>
 
 <style scoped lang="less">
.mynotification {
  position: fixed;
  bottom: 40px;
  width: 100%;
  z-index: 30;
  padding: 0.5rem;
}
</style>



<script>  
export default {
    data() {
        return {
            isInfoActive: false,
            infoMessage: "",
        }
    },
    mounted() {
        let _this = this;
        this.$globalEventHub.$on("toast.show", function (value) {
            _this.isInfoActive = true;
            _this.infoMessage = value.message;
            setTimeout(_this.hideInfoMessage, 300000);
        });
    },
    methods: {
        hideInfoMessage() {
            this.isInfoActive = false;
            this.infoMessage = "";
        }
    }
}
</script>