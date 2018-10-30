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
            this.timer = setInterval(function () {
                let _url = '/radar-chart?points=' + _this.pointNames + '&dur=1800';
                _this.$myaxios.get(_url, { timeout: 3000, 'hiddenLoading': true }).then(function (response) {
                    let _data = response.data;
                    _this.refreshData(_data)
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
            this.initTimer()
        },

        getOption() {
            let _dashboard_roundchart = this.$globalvar.dashboard_roundchart;
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
                    indicator: [
                        { name: '区域1', min: _dashboard_radar.min, max: _dashboard_radar.max },
                        { name: '区域2', min: _dashboard_radar.min, max: _dashboard_radar.max },
                        { name: '区域3', min: _dashboard_radar.min, max: _dashboard_radar.max },
                        { name: '区域4', min: _dashboard_radar.min, max: _dashboard_radar.max },
                    ],
                    shape: 'circle',
                    splitNumber: 5,
                    name: {

                    },
                    splitLine: {

                    },
                    splitArea: {
                        show: false
                    },
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
                                color: '#F9713C'
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
