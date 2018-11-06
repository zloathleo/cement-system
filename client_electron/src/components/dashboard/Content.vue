<template>
  <div class="columns">
    <div class="column is-10">
      <div class="columns">
        <div class="column is-6">
          <div class="card card-bg-transparent">
            <div class="card-header color-primary-4">
              温度场分区温度环状图
            </div>
            <div class="card-content">
              <RoundChartPanel />
            </div>
          </div>
        </div>

        <div class="column is-6">
          <div class="card card-bg-transparent">
            <div class="card-header color-primary-4">
              温度场分区温度雷达图
            </div>
            <div class="card-content">
              <RadarChart />
            </div>
          </div>
        </div>
      </div>

      <div class="columns">
        <div class="column is-6">
          <div class="card card-bg-transparent" style="margin-bottom: 1rem;">
            <div class="card-header color-primary-4">
              <div style="flex-grow: 1;color: #fff;">
                温度场分区温度历史曲线图
              </div>
              <a class="card-header-icon" aria-label="more options" @click="gotoTrend">
                <span class="icon">
                  <i class="mdi mdi-chart-line color-text-white"></i>
                </span>
              </a>
            </div>

            <div class="card-content" style="padding-bottom: 0px;">
              <LineChart />
            </div>
          </div>
        </div>

        <div class="column is-6">
          <div class="card card-bg-transparent" style="margin-bottom: 1rem;">
            <div class="card-header color-primary-4">
              <div style="flex-grow: 1;color: #fff;">
                燃烧器调整参数
              </div>
            </div>

            <div class="card-content">
              <GaugePanel/>
            </div>
          </div>
        </div>
      </div>

    </div>
    <div class="column is-2">
      <div class="card card-bg-transparent">
        <div class="card-header color-primary-4">
          实时系统测点
        </div>
        <div class="card-content">
          <DcsValuePanel />
        </div>
      </div>
    </div>

  </div>

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
