<template>

    <footer class="footer color-primary-3" style="border-top: 5px solid #e50112;display: flex;">

        <div style="  font-weight: bold; font-size: 12px;">
            {{serverTimestamp}}
        </div>
        <div style="margin-left: auto;    font-weight: bold; font-size: 12px;">
            Safe-Fire Inc.
        </div>

    </footer>

</template>

<style scoped lang="less">
.footer {
  position: fixed;
  bottom: 0;
  width: 100%;
  z-index: 30;
  padding: 0.5rem;
}
</style>


<script> 
export default {
    data() {
        return {
        }
    },
    computed: { 
        serverTimestamp: function () { 
            return '';
            // let t = this.$stateMem.state.serverTimestamp;
            // console.log(t);
            // return t == 0 ? 'connect fault!!': this.$globalvar.parseDateTime(t);
        }
    }
}
</script>