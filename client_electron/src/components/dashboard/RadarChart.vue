<template>
    <div ref="radarChart" />
</template> 
 
<style scoped lang="less">
</style>

<script>  
// var echarts = require('echarts/lib/echarts');
// // 引入chart
// require('echarts/lib/chart/radar');
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
        let chartDom = this.$refs.radarChart;
        chartDom.firstChild.style.setProperty('margin', 'auto', 'important');
    },
    destroyed() {
        if (this.timer) {
            clearInterval(this.timer);
        }
    },
    methods: {
        refreshData(_data) {
            let serie = this.chart.getOption().series[0];
            serie.data = _data;
            console.log(_data);
            this.chart.setOption({
                series: [
                    serie
                ]
            });
        },
        initTimer() {
            let _this = this;
            let _dashboard_radar = this.$globalvar.dashboard_radar;
            let _url = '/his-chart?type=radar&points=' + _this.pointNames + '&dur=' + _dashboard_radar.dur + '&interval=' + +_dashboard_radar.interval;

            this.timer = setInterval(function () {
                _this.$myaxios.get(_url, { timeout: 4500, 'hiddenLoading': true }).then(function (response) {
                    let _data = response.data;
                    _this.refreshData(_data);
                }).catch(function (err) {
                    console.error(err);
                });
            }, 5000);

        },

        _init_chart() {
            let chartDom = this.$refs.radarChart;
            let _width = 300;

            this.chart = echarts.init(chartDom, undefined, {
                width: _width,
                height: 250
            });

            this.chart.setOption(this.chartOption);
            this.requestData();
        },

        requestData() {
            let _this = this;
            let _dashboard_radar = this.$globalvar.dashboard_radar;
            let _url = '/his-chart?type=radar&points=' + _this.pointNames + '&dur=' + _dashboard_radar.dur + '&interval=' + +_dashboard_radar.interval;
            _this.$myaxios.get(_url, { timeout: 10000, 'hiddenLoading': true }).then(function (response) {
                let _data = response.data;
                _this.refreshData(_data);
                _this.initTimer();
            }).catch(function (err) {
                console.error(err);
            });
        },

        getOption() {
            let _dashboard_radar = this.$globalvar.dashboard_radar;
            var lineStyle = {
                normal: {
                    width: 1,
                    opacity: 0.5
                }
            };


            return {
                animation: false,
                title: {
                    show: false
                },
                radar: {
                    name: {
                        textStyle: {
                            color: '#000',
                            // backgroundColor: '#999',
                            fontSize: 12,
                            fontWeight: 'bold',
                            borderRadius: 3,
                            padding: [3, 5]
                        }
                    },
                    indicator: [
                        { name: '区域1', min: _dashboard_radar.min },
                        { name: '区域2', min: _dashboard_radar.min },
                        { name: '区域3', min: _dashboard_radar.min },
                        { name: '区域4', min: _dashboard_radar.min },
                    ],
                    // shape: 'circle',
                    splitNumber: 5,
                    splitLine: {

                    },
                    // splitArea: {
                    //     show: false
                    // },
                    axisLine: {

                    }
                },
                series: [
                    {
                        name: '温度场',
                        type: 'radar',
                        lineStyle: lineStyle,
                        data: [],
                        symbol: 'none',
                        itemStyle: {
                            normal: {
                                color: '#F9713C',
                                width: 1,
                                opacity: 0.5
                            }
                        },
                        areaStyle: {
                            normal: {
                                opacity: 0.1
                            }
                        }
                    }
                ]

            }
        }
    }

}
</script>
