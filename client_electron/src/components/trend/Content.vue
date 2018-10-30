<template>
  <div>
    <SelectPanel/>
    <div ref="lineChart" />
  </div>
</template> 
 
<style scoped lang="less">
</style>

<script>   
import SelectPanel from './SelectPanel.vue';
export default {
  components: {
    SelectPanel
  },
  props: {

  },
  data() {

    return {
      pointNames: [],
      chart: undefined,
      chartOption: this.getOption(),
      dur: 600
    }
  },
  mounted() {
    this._init_chart();
    let chartDom = this.$refs.lineChart;
    chartDom.firstChild.style.setProperty('margin', 'auto', 'important');

    let _this = this;
    this.$globalEventHub.$on("trend-params-change", function (value) {
      _this.setTrendparams(value);
    });

  },
  destroyed() {
    this.$globalEventHub.$off("trend-params-change");
  },
  methods: {

    setTrendparams(_params) {
      let pnNameArray = [];
      if (_params.selectedPointA == "rodinarea") {
        let _linechartConfig = this.$globalvar.dashboard_linechart;
        pnNameArray = _linechartConfig.series.map(function (_item) {
          return _item.pn;
        });
        pnNameArray[4] = _params.selectedPointB;
      } else {
        pnNameArray[0] = _params.selectedPointA;
        pnNameArray[1] = _params.selectedPointB;
      }
      console.log(pnNameArray);
      this.dur = _params.dur;
      this.pointNames = pnNameArray;
      this.initChartData();
    },

    initChartData() {
      let _this = this;

      let pnstr = this.pointNames.join(",");

      let _url = '/his-chart?points=' + pnstr + '&&dur=' + this.dur;
      _this.$myaxios.get(_url, { timeout: 2000, 'hiddenLoading': true }).then(function (response) {
        let _data = response.data;
        _this.renderJson(_data);
      }).catch(function (err) {
        console.error(err);
      });
    },

    renderJson(_json) {
      let _serise = _json.series;
      let _this = this;
      let config_series = this.pointNames.map(function (pn) {
        return {
          name: _this.$globalvar.trned_points[pn],
          pn: pn,
          yAxisIndex: 0,
          type: 'line',
          smooth: true,
          // yAxisIndex: _item.yAxisIndex,
          data: _serise[pn]
        }
      });

      let legendNameArray = this.pointNames.map(function (pn) {
        return _this.$globalvar.trned_points[pn];
      });

      this.chart.setOption({
        legend: {
          data: legendNameArray,
        },
        xAxis: {
          type: 'category',
          boundaryGap: false,
          axisLabel: {
            interval: parseInt(_json.xAxis.length / 50),
            rotate: 60
          },
          data: _json.xAxis
        },
        series: config_series
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
      // this.initChartData();
    },

    getOption() {
      let _linechartConfig = this.$globalvar.dashboard_linechart;

      return {
        title: {
          show: false
        },
        animation: false,
        tooltip: {
          trigger: 'axis'
        },
        grid: {
          left: '3%',
          right: '3%',
          containLabel: true
        },
        xAxis: {
          type: 'value',
          boundaryGap: false,
        },
        yAxis: {
          type: 'value'
        }
      }
    }
  }

}
</script>
