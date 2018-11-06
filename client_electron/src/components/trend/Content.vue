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
    let trned_points_map = {};
    let trned_points = this.$globalvar.trned_points;
    for (var _index in trned_points) { 
      let value = trned_points[_index];
      trned_points_map[value.pn] = value.name; 
    }

    return {
      pointNames: [],
      pointNamesMap:trned_points_map,
      chart: undefined,
      chartOption: this.getOption(),
      to: parseInt(new Date().getTime() / 1000),
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
      console.log(_params);
      let pnNameArray = [];
      if (_params.selectedPointA == "rodinarea") {
        pnNameArray = ["5_40019", "5_40033", "5_40047", "5_40061", _params.selectedPointB]
      } else {
        pnNameArray[0] = _params.selectedPointA;
        pnNameArray[1] = _params.selectedPointB;
      }
      console.log(pnNameArray);
      this.dur = _params.dur;
      this.pointNames = pnNameArray;

      console.log(_params.selectedToDate);
      console.log(_params.selectedToTime);

      let toTime = new Date();
      toTime.setFullYear(_params.selectedToDate.getFullYear());
      toTime.setMonth(_params.selectedToDate.getMonth());
      toTime.setDate(_params.selectedToDate.getDate());

      toTime.setHours(_params.selectedToTime.getHours());
      toTime.setMinutes(_params.selectedToTime.getMinutes());
      toTime.setSeconds(_params.selectedToTime.getSeconds());

      this.to = parseInt(toTime.getTime() / 1000);
      this.initChartData();
    },

    initChartData() {
      let _this = this;
      let pnstr = this.pointNames.join(",");
      this.chart.showLoading({
        animation: false,
        text: '加载数据...',
        textStyle: { fontSize: 20 }
      });

      let _url = '/his-chart?points=' + pnstr + '&dur=' + this.dur + "&to=" + this.to;
      _this.$myaxios.get(_url, { timeout: 30000, 'hiddenLoading': true }).then(function (response) {
        let _data = response.data;
        _this.renderJson(_data);
        _this.chart.hideLoading();
      }).catch(function (err) {
        console.error(err);
      });
    },

    renderJson(_json) {
      this.chart.clear();

      let _serise = _json.series;
      let _this = this;
      let config_series = this.pointNames.map(function (pn) {
        return {
          name: _this.pointNamesMap[pn],
          pn: pn,
          yAxisIndex: 0,
          type: 'line',
          smooth: true,
          data: _serise[pn]
        }
      });

      let legendNameArray = this.pointNames.map(function (pn) {
        return _this.pointNamesMap[pn];
      });

      this.chart.setOption({

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
        yAxis: {
          type: 'value'
        },
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

      // this.chart.getOption().series = config_series;
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
