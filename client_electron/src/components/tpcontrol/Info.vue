<template>

    <div class="card card-bg-transparent" style="box-shadow: none;padding-bottom: 2rem;">
        <div class="card-header color-primary-4">
            探头状态显示
        </div>
        <div ref="contentPane" class="card-content-table">
            <table class="table is-bordered  is-fullwidth card-bg-transparent">
                <tbody>
                    <tr>
                        <td class="table-td" style="width:25%">就地</td>
                        <td class="table-td" style="width:25%">{{getBooleanValue('1_00001')}}</td>
                        <td class="table-td" style="width:25%">远程</td>
                        <td class="table-td" style="width:25%">{{getBooleanValue('1_00002')}}</td>
                    </tr>
                    <tr>
                        <td class="table-td">吹扫风压力</td>
                        <td class="table-td width-value">{{getValue('4_40001')}}</td>
                        <td class="table-td">冷却风压力</td>
                        <td class="table-td width-value">{{getValue('4_40002')}}</td>

                    </tr>
                    <tr>
                        <td class="table-td">炉膛压力(DCS)</td>
                        <td class="table-td width-value">{{getValue('I_AI5701P01')}}</td>
                        <td class="table-td"> </td>
                        <td class="table-td"> </td>
                    </tr>
                    <tr>
                        <td class="table-td">A镜头温度</td>
                        <td class="table-td">{{getValue('4_40003')}}</td>
                        <td class="table-td">A探头内部温度(RodinQ)</td>
                        <td class="table-td">{{getValue('5_40013')}}</td>
                    </tr>
                    <tr>
                        <td class="table-td">A探头进到位</td>
                        <td class="table-td">{{getBooleanValue('1_00004')}}</td>
                        <td class="table-td">A探头退到位</td>
                        <td class="table-td">{{getBooleanValue('1_00005')}}</td>
                    </tr>
                    <tr>
                        <td class="table-td">B镜头温度</td>
                        <td class="table-td">{{getValue('4_40004')}}</td>
                        <td class="table-td">B探头内部温度(RodinQ)</td>
                        <td class="table-td">{{getValue('5_40083')}}</td>
                    </tr>
                    <tr>
                        <td class="table-td">B探头进到位</td>
                        <td class="table-td">{{getBooleanValue('1_00006')}}</td>
                        <td class="table-td">B探头退到位</td>
                        <td class="table-td">{{getBooleanValue('1_00007')}}</td>
                    </tr>
                </tbody>
            </table>
        </div>
    </div>

</template>

<style scoped lang="less">
.width-value {
  width: 5%;
}
</style>

<script>  
export default {
    components: {
    },
    data() {
        return {
            dataMap: {},
        }
    },
    mounted() {
        let _this = this;
        this.initTimer();
    },
    destroyed() {
        if (this.timer) {
            clearInterval(this.timer);
        }
    },
    methods: {
        initTimer() {
            this.timer = setInterval(this.refresh, 1000);
        },
        getBooleanValue(name) {
            let v = this.dataMap[name];
            if (v == undefined) {
                return "--"
            } else {
                if (parseInt(v) === 1) {
                    return "True"
                } else {
                    return "False"
                }
            }
        },
        getValue(name) {
            let v = this.dataMap[name];
            if (v == undefined) {
                return "--"
            } else {
                if (Number.isInteger(v)) {
                    return v;
                } else {
                    return v.value.toFixed(2);
                }
            }
        },
        refresh() {
            let _this = this;
            _this.$myaxios.get('/data', { timeout: 1000, 'hiddenLoading': true }).then(function (response) {
                let _data = response.data;

                let _rows = _data.rows;
                let _map = [];
                _rows.forEach(function (item) {
                    _map[item.name] = item;
                });

                _this.dataMap = _map;
                // console.log("dataMap:", _this.dataMap);

                // _this.$stateMem.commit("setServerTimestamp", _data.timestamp);
                _this.$stateMem.commit("setJinTui", {
                    "1_00004": _this.dataMap["1_00004"],
                    "1_00005": _this.dataMap["1_00005"]
                });

            }).catch(function (err) {
                _this.$stateMem.commit("setServerTimestamp", 0);
                console.error(err);
                // reject(err);
            });
        }
    }
}
</script>
