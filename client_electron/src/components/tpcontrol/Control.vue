<template>

    <div class="card card-bg-transparent" style="box-shadow: none;padding-bottom: 2rem;">
        <div class="card-header color-primary-4">
            探头控制操作
        </div>
        <div class="card-content-table">
            <table class="table is-bordered  is-fullwidth card-bg-transparent">
                <tbody>
                    <tr>
                        <td class="table-td" style="width:20%">探头A进退</td>
                        <td class="table-td" style="width:15%">
                            <a class="button color-action is-fullwidth" :disabled="!canAJing" v-bind:class="[canAJing ? 'color-action' : 'color-primary-button']" @click="clickAction($event,'2_00017',1,'确定操作探头A前进吗？','探头A前进指令已下发')">进</a>
                        </td>
                        <td class="table-td" style="width:15%">
                            <a class="button color-action is-fullwidth" @click="clickAction($event,'2_00017',0,'确定操作探头A后退吗？','探头A后退指令已下发')">退</a>
                        </td>
                        <td class="table-td" style="width:20%">探头A上电</td>
                        <td class="table-td" style="width:15%">
                            <a class="button color-action is-fullwidth" :disabled="!canAPowerOn" v-bind:class="[canAPowerOn ? 'color-action' : 'color-primary-button']" @click="clickAction($event,'2_00019',1,'确定操作探头A上电吗？','探头A上电指令已下发')">上电</a>
                        </td>
                        <td class="table-td" style="width:15%">
                            <a class="button color-action is-fullwidth" @click="clickAction($event,'2_00019',0,'确定操作探头A断电吗？','探头A断电指令已下发')">断电</a>
                        </td>
                    </tr>
                    <tr>
                        <td class="table-td">探头B</td>
                        <td class="table-td">
                            <a class="button color-action is-fullwidth" :disabled="!canBJing" v-bind:class="[canBJing ? 'color-action' : 'color-primary-button']" @click="clickAction($event,'2_00018',1,'确定操作探头B前进吗？','探头B前进指令已下发')">进</a>
                        </td>
                        <td class="table-td">
                            <a class="button color-primary-button is-fullwidth" disabled @click="clickAction($event,'2_00018',0,'确定操作探头B后退吗？','探头B后退指令已下发')">退</a>
                        </td>
                        <td class="table-td">探头B上电</td>
                        <td class="table-td">
                            <a class="button color-action is-fullwidth" :disabled="!canBPowerOn" v-bind:class="[canBPowerOn ? 'color-action' : 'color-primary-button']" @click="clickAction($event,'2_00020',1,'确定操作探头B上电吗？','探头B上电指令已下发')">上电</a>
                        </td>
                        <td class="table-td">
                            <a class="button color-primary-button is-fullwidth" disabled @click="clickAction($event,'2_00020',0,'确定操作探头B断电吗？','探头B断电指令已下发')">断电</a>
                        </td>
                    </tr>
                </tbody>
            </table>
        </div>
    </div>

</template>

<style scoped lang="less">
</style>

<script>   

export default {

    data() {
        return {
            dataMap: [],
        }
    },

    computed: {
        realtimePointValueMap() {
            return this.$stateMem.state.realtimePointValueMap;
        },
        canAJing: function () {
            return this.dataMap["1_00005"] == 1 && this.dataMap["4_40003"] < 55 && this.dataMap["5_40013"] < 55;
        },
        canAPowerOn: function () {
            return this.dataMap["5_40013"] < 55;
        },
        canBJing: function () {
            return false;
        },
        canBPowerOn: function () {
            return false;
        }

    },
    watch: {
        realtimePointValueMap: function (old, newd) {
            this.dataMap = newd;
        }
    },
    methods: {
        doAction(pointName, value, _message) {
            let _this = this;

            let rows = [{
                name: pointName,
                value: value
            }];
            let content = {
                rows: rows,
            };
            let _json = JSON.stringify(content, null, '    ');
            console.log("_json:", _json);
            let params = {
                'content': _json,
            };

            _this.$myaxios.post('/control', params).then(function (response) {
                _this.$globalvar.toastSuccess(_this, _message);
            }).catch(function (err) {
                console.error(err);
                _this.$globalvar.toastError(_this, "下发命令异常-", err);
            });

        },
        clickAction(_event, pointName, value, _message1, _message2) {
            console.log(_event.target);

            let disabled = _event.target.getAttribute("disabled")
            if (disabled == "disabled") {

            } else {
                this.$dialog.confirm({
                    title: '操作指令',
                    message: _message1,
                    type: 'is-danger',
                    cancelText: '取消',
                    confirmText: '确定',
                    onConfirm: () => this.doAction(pointName, value, _message2)
                })
            }

        },
    }
}
</script>
