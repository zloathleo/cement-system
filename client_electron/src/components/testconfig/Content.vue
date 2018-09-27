<template>

  <div class="columns">

    <div class="column is-12">

      <Info />
      <Control/>

    </div>
  </div>

</template>

<style scoped lang="less">
</style>

<script> 
import Info from './Info.vue';
import Control from './Control.vue';
export default {
  components: {
    Info, Control
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
