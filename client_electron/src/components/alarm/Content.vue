<template>

  <div class="card card-bg-transparent">
    <div class="card-header color-primary-4">
      系统报警
    </div>
    <div class="card-content-table">
      <table class="table is-bordered  is-fullwidth card-bg-transparent">
        <thead>
          <tr>
            <th>序号</th>
            <th>时间</th>
            <th>类型</th>
            <th>报警内容</th>
          </tr>
        </thead>
        <tbody>
          <tr v-for="(item, index) in tableData">
            <td class="text-center">{{index + 1}}</td>
            <td class="text-center">{{item.timestamp}}</td>
            <td class="text-center">{{item.module}}</td>
            <td class="text-center">{{item.err}}</td> 
          </tr>
        </tbody>
      </table>
    </div>
  </div>

</template>

<style scoped lang="less">
</style>

<script> 
export default {
  components: {
  },
  data() {
    return {
      tableData: []
    }
  },
  mounted() {
    this.initTimer();
  },
  methods: {
    initTimer() {
      let _this = this;
      this.timer = setTimeout(function () {
        let _url = '/page?name=alarm';
        _this.$myaxios.get(_url, { timeout: 2000 }).then(function (response) {
          let _data = response.data;
          _this.tableData = _data.rows;
        }).catch(function (err) {

        });
      }, 500);
    },
  }
}
</script>
