export default {
    appName: "后台管理",
    GlobalEventHub: undefined,

    //配置模式 直接查看数据和控制信号
    configMode: false,

    //dashboard ui关联实时值
    points_dashboard: [],
    //control ui 关联实时值
    points_control: [],


    dashboard_roundchart: undefined,
    dashboard_radar: undefined,
    //dashboard linchart 添加对照点
    dashboard_linechart: undefined,

    // fetchServerHostURL: "",
    fetchServerHostURL: "http://localhost:8088",


    parseDateTime: function (_long) {
        if (_long) {
            let t = new Date();
            t.setTime(_long * 1000)
            return t.Format("yyyy-MM-dd HH:mm:ss");
        } else {
            return "-";
        }
    },

    toastInfo: function (_message) {
        this.GlobalEventHub.$emit("toast.show", {
            type: "info",
            message: _message
        });
    },

    toastError: function (_vue, errdesc, err) {
        let _message = errdesc;
        if (err !== undefined) {
            _message += " - " + this.parseError(err)
        }

        _vue.$toast.open({
            duration: 5000,
            message: _message,
            position: 'is-bottom',
            type: 'is-danger'
        })
    },

    toastSuccess: function (_vue, _message) {
        _vue.$toast.open({
            queue: false,
            duration: 5000,
            message: _message,
            position: 'is-bottom',
            type: 'is-success'
        });
    },

    // toastError: function (_vue, errdesc, error) {
    //     let _message = errdesc;
    //     if (error !== undefined) {
    //         console.error(error);
    //         _message += " - " + this.parseError(error)
    //     }
    //     _vue.$toast.open({
    //         queue: false,
    //         message: _message,
    //         position: 'is-bottom',
    //         type: 'is-danger'
    //     });
    // },

    parseError: function (error) {
        if (error && error.response && error.response.data) {
            return error.response.data.message || error.response.data.msg;
        } else {
            return error;
        }
    },

};