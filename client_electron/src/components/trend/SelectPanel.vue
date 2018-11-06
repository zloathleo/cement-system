<template>

    <div style="display: flex;">
        <div>
            <b-field label="对比测点1">
                <b-select placeholder="对比测点1" rounded v-model="selectedPointA">
                    <option value="rodinarea">温度场温度(4)</option>
                    <option value="5_40019">温度场区域1</option>
                    <option value="5_40033">温度场区域2</option>
                    <option value="5_40047">温度场区域3</option>
                    <option value="5_40061">温度场区域4</option>

                    <option v-for="(item, index) in $globalvar.dashboard_dcs_points" :value="item.pn">
                        {{item.name}}
                    </option>

                </b-select>
            </b-field>
        </div>
        <div>
            <b-field label="对比测点2">
                <b-select placeholder="对比测点2" rounded v-model="selectedPointB">
                    <option v-for="(item, index) in $globalvar.dashboard_dcs_points" :value="item.pn">
                        {{item.name}}
                    </option>
                </b-select>
            </b-field>
        </div>

        <b-field label="结束日期">
            <b-datepicker v-model="selectedToDate" :first-day-of-week="1" rounded placeholder="Click to select...">

                <button class="button is-primary" @click="selectedToDate = new Date()">
                    <b-icon icon="calendar-today"></b-icon>
                    <span>今天</span>
                </button>

                <button class="button is-danger" @click="selectedToDate = null">
                    <b-icon icon="close"></b-icon>
                    <span>清除</span>
                </button>
            </b-datepicker>
        </b-field>

        <b-field label="结束时间">
            <b-timepicker v-model="selectedToTime" rounded placeholder="Click to select..." icon="clock" :hour-format="'24'">
            </b-timepicker>
        </b-field>

        <div>
            <b-field label="时长">
                <b-select placeholder="时长" rounded v-model="dur">
                    <option value="600">10分钟</option>
                    <option value="1800">30分钟</option>
                    <option value="3600">1小时</option>
                    <option value="28800">8小时</option>
                    <option value="86400">24小时</option>
                    <option value="604800">7天</option>
                </b-select>
            </b-field>
        </div>

        <b-field label="提交">
            <a class="button color-primary-4 is-rounded " style="width:200px" @click="clickSelect">
                确认
            </a>
        </b-field>

    </div>

</template>

<style scoped lang="less">
</style>


<script> 
var moment = require('moment');
export default {

    data() {
        let _ddd = moment().subtract(1, 'days');

        return {
            dur: 600,
            selectedPointA: 'rodinarea',
            selectedPointB: '4_40003',

            selectedToDate: new Date(),
            selectedToTime: new Date(),
        }
    },
    methods: {
        clickSelect() {
            this.$globalEventHub.$emit("trend-params-change", {
                dur: this.dur,
                selectedPointA: this.selectedPointA,
                selectedPointB: this.selectedPointB,
                selectedToDate: this.selectedToDate,
                selectedToTime: this.selectedToTime,
            });
        }
    }
}
</script>