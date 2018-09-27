<template>
  <div>
    <div class="columns">
      <div class="column is-half">
        <div class="card card-bg-transparent">
          <div class="card-header color-primary-4">
            Temperature pattern
          </div>
          <div class="card-content">
            <RoundChart />
          </div>
        </div>
      </div>

      <div class="column is-half">
        <div class="card card-bg-transparent">
          <div class="card-header color-primary-4">
            Temperature Radar
          </div>
          <div class="card-content">
            <RadarChart />
          </div>
        </div>
      </div>
    </div>

    <div class="columns">
      <div class="column">
        <div class="card card-bg-transparent">
          <div class="card-header color-primary-4">
            Temperature History
          </div>
          <div class="card-content">
            <LineChart />
          </div>
        </div>
      </div>
    </div>
  </div>

</template>

<style scoped lang="less">
.card-content {
  // padding: 0rem;
}
</style>

<script> 
import RoundChart from './RoundChart.vue';
import RadarChart from './RadarChart.vue';
import LineChart from './LineChart.vue';
export default {
  components: {
    RoundChart, RadarChart, LineChart
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

  }
}
</script>
