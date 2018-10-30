<template>

    <div class="card card-bg-transparent" style="box-shadow: none;padding-bottom: 2rem;">
        <div class="card-header color-primary-4">
            探头状态示意图
        </div>
        <div ref="contentPane" class="card-content-table">
            <img v-if="tuiA" src="assets/img/tui.png" class="tui-A">
            <img v-if="jinA" src="assets/img/jin.png" class="jin-A">
            <img src="assets/img/lutang.png" class="lutang-A">
        </div>
    </div>

</template>

<style scoped lang="less">
.jin-A {
  position: relative;
  width: 300px;
  top: 44px;
  left: 105px;
}
.tui-A {
  position: relative;
  width: 300px;
  top: 44px;
  left: 85px;
}
.lutang-A {
  position: relative;
  width: 350px;
  top: 120px;
  left: 20px;
}
</style>

<script>  
export default {
    components: {
    },
    data() {
        return { 
        }
    },
    computed: {
        // 计算属性的 getter
        tuiA: function () {
            return this.$stateMem.state.jintui["1_00004"] != 1
        },
        jinA: function () {
            return this.$stateMem.state.jintui["1_00004"] == 1
        }
    },
    mounted() {

    },
    destroyed() {

    },
    methods: {

    }
}
</script>
