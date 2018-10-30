<template>
  <div class="columns">
    <div class="column is-half">
      <RoundChart />
    </div>

    <div class="column is-half">
      <table class="table is-narrow" style="background-color: rgba(255, 255, 255, 0);font-size: 0.8rem;">
        <thead>
          <tr>
            <th>温度场区域</th>
            <th>实时值</th>
          </tr>
        </thead>
        <tbody>
          <tr>
            <th>区域1</th>
            <td>{{data['5_40019']}}</td>
          </tr>
          <tr>
            <th>区域2</th>
            <td>{{data['5_40033']}}</td>
          </tr>
          <tr>
            <th>区域3</th>
            <td>{{data['5_40047']}}</td>
          </tr>
          <tr>
            <th>区域4</th>
            <td>{{data['5_40061']}}</td>
          </tr>
        </tbody>
      </table>
    </div>
  </div>

</template>

<style scoped lang="less">
</style>

<script> 
import RoundChart from './RoundChart.vue';
export default {
  components: {
    RoundChart
  },
  data() {
    let _linechartConfig = this.$globalvar.dashboard_roundchart;
    let pnNameArray = _linechartConfig.map(function (_item) {
      return _item.pn;
    });
    return {
      data: {},
      pointNames: pnNameArray.join(",")
    }
  },
  mounted() {
    this.initTimer();
  },
  destroyed() {
    if (this.timer) {
      clearInterval(this.timer);
    }
  },
  methods: {
    initTimer() {
      let _this = this;
      this.timer = setInterval(function () {
        let _url = '/realtime?points=' + _this.pointNames + '';
        _this.$myaxios.get(_url, { timeout: 1000, 'hiddenLoading': true }).then(function (response) {
          let _data = response.data;
          _this.data = _data;
          _this.$globalEventHub.$emit("data-roundchart", _data);
        }).catch(function (err) {
          console.error(err);
        });
      }, 3000);
    },
  }
}
</script>
