<template>
    <div ref="roundChart" />
</template> 
 
<style scoped lang="less">
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
        this.$globalEventHub.$on("data-roundchart", function (value) {
            _this.refreshData(value);
        });

        let chartDom = this.$refs.roundChart;
        chartDom.firstChild.style.setProperty('margin', 'auto', 'important');

    },
    destroyed() {
        this.$globalEventHub.$off("data-roundchart");
    },
    methods: {
        refreshData(_data) {
            let _dataArray = this.chart.getOption().series[0].data;
            let _chartConfig = this.$globalvar.dashboard_roundchart;

            let v1 = _data[_chartConfig[0].pn];
            let v2 = _data[_chartConfig[1].pn];
            let v3 = _data[_chartConfig[2].pn];
            let v4 = _data[_chartConfig[3].pn];
            if (v1 == null && v2 == null && v3 == null && v4 == null) {
                v1 = 0;
                v2 = 0;
                v3 = 0;
                v4 = 0;
            }

            let serie = this.chart.getOption().series[0];
            serie.data = [
                { value: v1, pn: _chartConfig[0].pn, name: _dataArray[0].name },
                { value: v2, pn: _chartConfig[1].pn, name: _dataArray[1].name },
                { value: v3, pn: _chartConfig[2].pn, name: _dataArray[2].name },
                { value: v4, pn: _chartConfig[3].pn, name: _dataArray[3].name }
            ];
            this.chart.setOption({
                series: [serie]
            });

 

        },

        _init_chart() {
            let chartDom = this.$refs.roundChart;
            let _width = 300;

            this.chart = echarts.init(chartDom, undefined, {
                width: _width,
                height: 250
            });
            this.chart.setOption(this.chartOption);
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
                        startAngle: 45,
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
