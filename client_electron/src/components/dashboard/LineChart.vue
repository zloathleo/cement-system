<template>
    <div ref="lineChart" />
</template> 
 
<style scoped lang="less">
</style>

<script>   
export default {
    components: {

    },
    data() {
        let _linechartConfig = this.$globalvar.dashboard_linechart;
        let pnNameArray = _linechartConfig.series.map(function (_item) {
            return _item.pn;
        });

        return {
            initUI: true,
            pointNames: pnNameArray.join(","),
            chart: undefined,
            chartOption: this.getOption(_linechartConfig),
        }
    },
    computed: {
        realtimePointValueMap() {
            return this.$stateMem.state.realtimePointValueMap;
        }
    },
    watch: {
        realtimePointValueMap: function (old, newd) {
            if (!this.initUI) {
                this.refreshRealtime(newd);
            }
            this.initUI = false;
        }
    },
    mounted() {
        this._init_chart();
        let chartDom = this.$refs.lineChart;
        chartDom.firstChild.style.setProperty('margin', 'auto', 'important');

    },
    methods: {
        refreshRealtime(_data) {
            let _this = this;
            let _chartOptioin = this.chart.getOption();

            _chartOptioin.series.forEach(seriesItem => {
                let pn = seriesItem.pn;
                let newValue = _data[pn];
                var data0 = seriesItem.data;
                data0.shift();
                data0.push(newValue);
            });


            let axisData = new Date();
            let dateString = axisData.Format("yyyy-MM-dd HH:mm:ss");
            console.log("dateString:", dateString);

            _chartOptioin.xAxis[0].data.shift();
            _chartOptioin.xAxis[0].data.push(dateString);

            _this.chart.setOption(_chartOptioin);
        },
        initChartData() {
            let _this = this;
            let _url = '/his-chart?points=' + this.pointNames + '&&dur=600';
            _this.$myaxios.get(_url, { timeout: 10000, 'hiddenLoading': true }).then(function (response) {
                let _data = response.data;
                _this.renderJson(_data);
            }).catch(function (err) {
                console.error(err);
            });
        },

        renderJson(_json) {
            let _serise = _json.series;

            let _linechartConfig = this.$globalvar.dashboard_linechart;
            let config_series = _linechartConfig.series.map(function (_item) {
                return {
                    name: _item.name,
                    pn: _item.pn,
                    yAxisIndex: 0,
                    type: 'line',
                    smooth: true,
                    yAxisIndex: _item.yAxisIndex,
                    data: _serise[_item.pn]
                }
            });

            this.chart.setOption({
                xAxis: {
                    type: 'category',
                    boundaryGap: false,
                    axisLabel: {
                        interval: 59,
                        rotate: 30
                    },
                    data: _json.xAxis
                },
                series: config_series
            });
        },

        _init_chart() {
            let chartDom = this.$refs.lineChart;
            let _width = document.body.clientWidth / 2 - 250;
            let _height = _width * 0.5;
            this.chart = echarts.init(chartDom, undefined, {
                width: _width,
                height: _height
            });
            this.chart.setOption(this.chartOption);
            this.initChartData();
        },

        getOption(_linechartConfig) {
            let legendNameArray = _linechartConfig.series.map(function (_item) {
                return _item.name;
            });

            return {
                title: {
                    show: false
                },
                animation: false,
                legend: {
                    data: legendNameArray,
                },
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
                yAxis: _linechartConfig.yAxis

            }
        }
    }

}
</script>
