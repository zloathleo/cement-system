<template>
    <div ref="roundChart" class="round-chart" />
</template> 
 
<style scoped lang="less">
.round-chart {
  margin: auto;
}
</style>

<script>  
var echarts = require('echarts/lib/echarts');
// 引入chart
require('echarts/lib/chart/pie');
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
            let chartDom = this.$refs.roundChart;
            let _width = 300;

            this.chart = echarts.init(chartDom, undefined, {
                width: _width,
                height: _width / 1.334
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
                    trigger: 'item',
                    formatter: "{a} <br/>{b} : {c} ({d}%)"
                }, 
                series: [
                    {
                        name: '访问来源',
                        type: 'pie',
                        radius: ['50%', '70%'], 
                        avoidLabelOverlap: false,

                        data: [
                            { value: 1, name: '1200' },
                            { value: 1, name: '1230' },
                            { value: 1, name: '1250' },
                            { value: 1, name: '1280' },
                            { value: 1, name: '1300' }
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
