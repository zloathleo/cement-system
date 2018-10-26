<template>
    <div ref="roundChart" class="round-chart" />
</template> 
 
<style scoped lang="less">
.round-chart {
  margin: auto;
}
</style>

<script>  
// var echarts = require('echarts/lib/echarts');
// // 引入chart
// require('echarts/lib/chart/pie');
// // 引入提示框和标题组件 
// require('echarts/lib/component/title');
// require('echarts/lib/component/tooltip');
export default {
    components: {
    },
    props: {

    },
    data() {
        let _linechartConfig = this.$globalvar.dashboard_roundchart;
        let pnNameArray = _linechartConfig.map(function (_item) {
            return _item.pn;
        });
        return {
            pointNames: pnNameArray.join(","),
            chart: undefined,
            chartOption: this.getOption(),
        }
    },
    mounted() {
        this._init_chart();

        let _this = this;
        this.chart.on('finished', function () {
            let chartDom = _this.$refs.roundChart;
            chartDom.parentNode.style.margin = 'auto';
        });

    },
    destroyed() {
        if (this.timer) {
            clearInterval(this.timer);
        }
    },
    methods: {
        refreshData(_data) {
            let _dataArray = this.chart.getOption().series[0].data;
            let _chartConfig = this.$globalvar.dashboard_roundchart;

            this.chart.setOption({
                series: [
                    {
                        name: '温度场区域温度',
                        type: 'pie',
                        radius: ['50%', '70%'],
                        avoidLabelOverlap: false,
                        data: [
                            { value: _data[_chartConfig[0].pn], pn: _chartConfig[0].pn, name: _dataArray[0].name },
                            { value: _data[_chartConfig[1].pn], pn: _chartConfig[1].pn, name: _dataArray[1].name },
                            { value: _data[_chartConfig[2].pn], pn: _chartConfig[2].pn, name: _dataArray[2].name },
                            { value: _data[_chartConfig[3].pn], pn: _chartConfig[3].pn, name: _dataArray[3].name }
                        ],
                        itemStyle: {
                            emphasis: {
                                shadowBlur: 10,
                                shadowOffsetX: 0,
                                shadowColor: 'rgba(0, 0, 0, 0.5)'
                            }
                        }
                    }
                ]
            });

        },

        initTimer() {
            let _this = this;
            this.timer = setInterval(function () {
                let _url = '/realtime?points=' + _this.pointNames + '';
                _this.$myaxios.get(_url, { timeout: 1000, 'hiddenLoading': true }).then(function (response) {
                    let _data = response.data;
                    _this.refreshData(_data)
                }).catch(function (err) {
                    console.error(err);
                });
            }, 3000);
        },

        _init_chart() {
            let chartDom = this.$refs.roundChart;
            let _width = 300;

            this.chart = echarts.init(chartDom, undefined, {
                width: _width,
                height: _width / 1.334
            });
            this.chart.setOption(this.chartOption);
            this.initTimer()
        },

        getOption() {
            let _chartConfig = this.$globalvar.dashboard_roundchart;
            return {
                title: {
                    show: false
                },
                animation: false,
                tooltip: {
                    trigger: 'item',
                    formatter: "{a} <br/>{b} : {c} ({d}%)"
                },
                series: [
                    {
                        name: '温度场区域温度',
                        type: 'pie',
                        radius: ['50%', '70%'],
                        avoidLabelOverlap: false,
                        label: {
                            formatter: '{b}-{c}',
                        },
                        data: [
                            { value: 1000, pn: _chartConfig[0].pn, name: _chartConfig[0].name },
                            { value: 1000, pn: _chartConfig[1].pn, name: _chartConfig[1].name },
                            { value: 1000, pn: _chartConfig[2].pn, name: _chartConfig[2].name },
                            { value: 1000, pn: _chartConfig[3].pn, name: _chartConfig[3].name }
                        ],
                        itemStyle: {
                            emphasis: {
                                shadowBlur: 10,
                                shadowOffsetX: 0,
                                shadowColor: 'rgba(0, 0, 0, 0.5)'
                            }
                        }
                    }
                ]

            }
        }
    }

}
</script>
