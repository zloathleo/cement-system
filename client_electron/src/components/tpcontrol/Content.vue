<template>

  <div class="columns">
    <div class="column is-6">
      <Info />
      <Control/>
    </div>
    <div class="column is-6">
      <Viewer/>
    </div>
  </div>

</template>

<style scoped lang="less">
</style>

<script> 
import Info from './Info.vue';
import Control from './Control.vue';
import Viewer from './Viewer.vue';
export default {
  components: {
    Info, Control, Viewer
  },
  data() {
    return {
    }
  },
  mounted() {
    let _this = this;

  },
  methods: {
  }
}
</script>
