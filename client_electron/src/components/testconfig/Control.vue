<template>

    <div class="card card-bg-transparent" style="box-shadow: none;padding-bottom: 2rem;">
        <div class="card-header color-primary-4">
            探头控制操作
        </div>
        <div class="card-content-table">
            <table class="table is-bordered  is-fullwidth card-bg-transparent">
                <tbody>
                    <tr>
                        <td class="table-td">探头A进退</td>
                        <td class="table-td">
                            <a class="button color-primary-4 is-fullwidth" @click="clickAction('2_00017',1)">进</a>
                        </td>
                        <td class="table-td">
                            <a class="button color-primary-4 is-fullwidth" @click="clickAction('2_00017',0)">退</a>
                        </td>
                        <td class="table-td">探头A上电</td>
                        <td class="table-td">
                            <a class="button color-primary-4 is-fullwidth" @click="clickAction('2_00019',1)">上电</a>
                        </td>
                        <td class="table-td">
                            <a class="button color-primary-4 is-fullwidth" @click="clickAction('2_00019',0)">断电</a>
                        </td>
                    </tr>
                    <tr>
                        <td class="table-td">探头B</td>
                        <td class="table-td">
                            <a class="button color-primary-4 is-fullwidth" @click="clickAction('2_00018',1)">进</a>
                        </td>
                        <td class="table-td">
                            <a class="button color-primary-4 is-fullwidth" @click="clickAction('2_00018',0)">退</a>
                        </td>
                        <td class="table-td">探头B上电</td>
                        <td class="table-td">
                            <a class="button color-primary-4 is-fullwidth" @click="clickAction('2_00020',1)">上电</a>
                        </td>
                        <td class="table-td">
                            <a class="button color-primary-4 is-fullwidth" @click="clickAction('2_00020',0)">断电</a>
                        </td>
                    </tr>

                    <tr>
                        <td class="table-td">风机</td>
                        <td class="table-td">
                            <a class="button color-primary-4 is-fullwidth" @click="clickAction('2_00024',0)">开</a>
                        </td>
                        <td class="table-td">
                            <a class="button color-primary-4 is-fullwidth" @click="clickAction('2_00024',1)">关</a>
                        </td>
                        <td class="table-td">-</td>
                        <td class="table-td">-</td>
                        <td class="table-td">-</td>
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
    components: {
    },
    data() {
        return {
        }
    },
    mounted() {

    },
    methods: {
        clickAction(pointName, value) {

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
                _this.data = response.data;

                console.log("_json:", _json);
                // resolve(response);
            }).catch(function (err) {
                console.error(err);
                // reject(err);
            });

        },
    }
}
</script>
