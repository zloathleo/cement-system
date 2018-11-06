<template>
    <div class="columns">
        <div class="column is-one-third" style="text-align: center;font-weight: bolder;font-size: 13px;">
            中心风开度
            <GaugeChart />
        </div>
        <div class="column is-one-third" style="text-align: center;font-weight: bolder;font-size: 13px;">
            内二次风开度
            <GaugeChart />
        </div>
        <div class="column is-one-third" style="text-align: center;font-weight: bolder;font-size: 13px;">
            外二次风开度
            <GaugeChart />
        </div>
    </div>

</template>

<style scoped lang="less">
</style>

<script> 
import GaugeChart from './GaugeChart.vue';
export default {
    components: {
        GaugeChart
    },
    data() {

        return {
            data: {}
        }
    },
    mounted() {
    },
    destroyed() {
    },
    methods: {
    }
}
</script>
