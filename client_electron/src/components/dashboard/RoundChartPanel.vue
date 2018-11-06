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
            <td>{{$stateMem.state.realtimePointValueMap["5_40019"]}}</td>
          </tr>
          <tr>
            <th>区域2</th>
            <td>{{$stateMem.state.realtimePointValueMap["5_40033"]}}</td>
          </tr>
          <tr>
            <th>区域3</th>
            <td>{{$stateMem.state.realtimePointValueMap['5_40047']}}</td>
          </tr>
          <tr>
            <th>区域4</th>
            <td>{{$stateMem.state.realtimePointValueMap['5_40061']}}</td>
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
  methods: {

  }
}
</script>
