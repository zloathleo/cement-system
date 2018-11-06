<template>
  <template>
    <table class="table is-narrow" style="background-color: rgba(255, 255, 255, 0);font-size: 0.8rem;min-height:250px">
      <thead>
        <tr>
          <th>测点名称</th>
          <th>实时值</th>
        </tr>
      </thead>
      <tbody>
        <tr v-for="(item, index) in $globalvar.dashboard_dcs_points">
          <th>{{item.name}}</th>
          <td>{{$stateMem.state.realtimePointValueMap[item.pn]}}</td>
        </tr>
      </tbody>
    </table> 
  </template>

<style scoped lang="less">
</style>

<script>  
export default {
}
</script>


</template>

<style scoped lang="less">
</style>

<script> 

import DcsValuePanel from './DcsValuePanel.vue';
import RoundChartPanel from './RoundChartPanel.vue';
import RadarChart from './RadarChart.vue';
import LineChart from './LineChart.vue';
import GaugePanel from './GaugePanel.vue';
export default {
  components: {
    DcsValuePanel, RoundChartPanel, RadarChart, LineChart, GaugePanel
  },
  data() {
    return {
      commondata: {
        d1: "-",
        d2: "-",
        d3: "-",
        d4: "-",
      },
      data: {},
    }
  },
  mounted() {
    let _this = this;


  },
  methods: {
    gotoTrend() {
      this.$router.push({ name: "trend" });
      this.$stateMem.commit("setUiTitle", "trend");
    }
  }
}
</script>
