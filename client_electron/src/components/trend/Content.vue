<template>
  <div class="card card-bg-transparent">
    <div class="card-header color-primary-4">
      历史趋势查询
    </div>
    <div class="card-content">
      <div ref="lineChart" />
    </div>
  </div>
</template> 
 
<style scoped lang="less">
</style>

<script>   
export default {
  components: {

  },
  props: {

  },
  data() {
    return {
      chart: undefined,
      chartOption: this.getOption(),
    }
  },
  mounted() {
    this._init_chart();
    let _this = this;
  },
  methods: {
    initChartData() {

      let _this = this;
      let _appendPoint = this.$globalvar.dashboard_linechart_point;


      let _url = '/his-chart?points=5_40019,5_40033,5_40047,5_40061,' + _appendPoint.name + '&&dur=604800&to=1540432800';
      _this.$myaxios.get(_url, { timeout: 2000, 'hiddenLoading': true }).then(function (response) {
        let _data = response.data;

        _this.renderJson(_data);
      }).catch(function (err) {
        console.error(err);
      });
    },

    renderJson(_json) {
      let _appendPoint = this.$globalvar.dashboard_linechart_point;
      let _serise = _json.series;

      console.info("_appendPoint.label:" + _appendPoint.label, _appendPoint.name, _appendPoint.min, _appendPoint.max);

      this.chart.setOption({

        legend: {
          data: ['区域1', '区域2', '区域3', '区域4', _appendPoint.label],
        },

        xAxis: {
          type: 'category',
          boundaryGap: false,
          data: _json.xAxis
        },

        yAxis: [
          {
            name: '区域温度',
            type: 'value',
            scale: true,
            splitNumber: 8,
            min: 800,
            max: 1600
          }, {
            name: _appendPoint.label,
            type: 'value',
            scale: true,
            splitNumber: 8,
            min: _appendPoint.min,
            max: _appendPoint.max
          }
        ],

        series: [
          {
            name: '区域1',
            yAxisIndex: 0,
            type: 'line',
            smooth: true,
            data: _serise["5_40019"]
          },

          {
            name: '区域2',
            yAxisIndex: 0,
            type: 'line',
            smooth: true,
            data: _serise["5_40033"]
          },

          {
            name: '区域3',
            yAxisIndex: 0,
            type: 'line',
            smooth: true,
            data: _serise["5_40047"]
          },

          {
            name: '区域4',
            yAxisIndex: 0,
            type: 'line',
            smooth: true,
            data: _serise["5_40061"]
          },

          {
            name: _appendPoint.label,
            yAxisIndex: 1,
            type: 'line',
            smooth: true,
            data: _serise[_appendPoint.name]
          }
        ]
      });
    },

    _init_chart() {
      let chartDom = this.$refs.lineChart;
      let _width = document.body.clientWidth - 130;
      let _height = document.body.clientHeight * 0.5;
      this.chart = echarts.init(chartDom, undefined, {
        width: _width,
        height: _height
      });
      this.chart.setOption(this.chartOption);
      this.initChartData();
    },

    getOption() {
      return {
        title: {
          show: false
        },
        tooltip: {
          trigger: 'axis'
        },
        grid: {
          left: '1%',
          right: '1%',
          containLabel: true
        },
        xAxis: {
          type: 'value',
          boundaryGap: false,
        },
        yAxis: {
          type: 'value',
          min: 1000,
          max: 1400
        }

      }
    }
  }

}
</script>
