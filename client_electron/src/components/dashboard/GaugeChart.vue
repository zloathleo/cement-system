<template>
    <div ref="gaugeChart" />
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
        }
    },
    mounted() {
        this._init_chart();
        let chartDom = this.$refs.gaugeChart;
        chartDom.firstChild.style.setProperty('margin', 'auto', 'important');
    },
    destroyed() {
    },
    methods: {
        _init_chart() {
            let chartDom = this.$refs.gaugeChart;
            let _width = 240;

            this.chart = echarts.init(chartDom, undefined, {
                width: _width,
                height: _width  
            });
            this.chart.setOption(this.getOption());
        },

        getOption() {

            return {
                tooltip: {
                    formatter: "{a} <br/>{b} : {c}%"
                },

                series: [
                    {
                        name: '业务指标',
                        type: 'gauge',
                        axisLine: {            // 坐标轴线
                            lineStyle: {       // 属性lineStyle控制线条样式
                                width: 8
                            }
                        },
                        axisTick: {            // 坐标轴小标记
                            length: 12,        // 属性length控制线长
                            lineStyle: {       // 属性lineStyle控制线条样式
                                color: 'auto'
                            }
                        },
                        splitLine: {           // 分隔线
                            length: 20,         // 属性length控制线长
                            lineStyle: {       // 属性lineStyle（详见lineStyle）控制线条样式
                                color: 'auto'
                            }
                        },
                        pointer: {
                            width: 5
                        },
                        detail: { formatter: '{value}%' },
                        data: [{ value: 0 }]
                    }
                ]
            };
        }
    }

}
</script>
