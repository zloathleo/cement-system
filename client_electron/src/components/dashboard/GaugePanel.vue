<template>
    <div class="columns">
        <div class="column is-one-third" style="text-align: center;font-weight: bolder;">
            中心风开度
            <GaugeChart />
        </div>
        <div class="column is-one-third" style="text-align: center;font-weight: bolder;">
            内二次风开度
            <GaugeChart />
        </div>
        <div class="column is-one-third" style="text-align: center;font-weight: bolder;">
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
