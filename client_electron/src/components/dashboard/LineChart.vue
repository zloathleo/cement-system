<template>
    <div ref="lineChart" />
</template> 
 
<style scoped lang="less">
</style>

<script>  
var echarts = require('echarts/lib/echarts');
// 引入chart
require('echarts/lib/chart/line');
// 引入提示框和标题组件 
require('echarts/lib/component/title');
require('echarts/lib/component/tooltip');
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
        refreshData() {

        },

        _init_chart() {
            let chartDom = this.$refs.lineChart;
            let _width = 800;

            this.chart = echarts.init(chartDom, undefined, {
                width: _width,
                height: 280
            });
            this.chart.setOption(this.chartOption);
            this.refreshData();
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
                    left: '3%',
                    right: '3%', 
                    containLabel: true
                }, 
                xAxis: {
                    type: 'category',
                    boundaryGap: false,
                    data: ['周一', '周二', '周三', '周四', '周五', '周六', '周日']
                },
                yAxis: {
                    type: 'value'
                },
                series: [
                    {
                        name: '区域1',
                        type: 'line',
                        stack: '总量',
                        data: [120, 132, 101, 134, 90, 230, 210]
                    },
                    {
                        name: '区域2',
                        type: 'line',
                        stack: '总量',
                        data: [220, 182, 191, 234, 290, 330, 310]
                    },
                    {
                        name: '区域3',
                        type: 'line',
                        stack: '总量',
                        data: [150, 232, 201, 154, 190, 330, 410]
                    },
                    {
                        name: '区域4',
                        type: 'line',
                        stack: '总量',
                        data: [320, 332, 301, 334, 390, 330, 320]
                    },
                    {
                        name: '区域5',
                        type: 'line',
                        stack: '总量',
                        data: [820, 932, 901, 934, 1290, 1330, 1320]
                    }
                ]
            }
        }
    }

}
</script>
