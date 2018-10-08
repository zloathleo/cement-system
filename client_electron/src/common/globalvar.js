export default {
    appName: "后台管理",
    GlobalEventHub: undefined,

    //配置模式 直接查看数据和控制信号
    configMode: false,

    // fetchServerHostURL: "",
    fetchServerHostURL: "http://localhost:8080",


    parseDateTime: function (_long) {
        if (_long) {
            return new Date(_long).Format("yyyy-MM-dd HH:mm:ss");
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

    toastSuccess: function (_vue, _message) {
        _vue.$toast.open({
            queue: false,
            message: _message,
            position: 'is-bottom',
            type: 'is-success'
        });
    },

    toastError: function (_vue, errdesc, error) {
        let _message = errdesc;
        if (error !== undefined) {
            console.error(error);
            _message += " - " + this.parseError(error)
        }
        _vue.$toast.open({
            queue: false,
            message: _message,
            position: 'is-bottom',
            type: 'is-danger'
        });
    },

    parseError: function (error) {
        if (error && error.response && error.response.data) {
            return error.response.data.message;
        } else {
            return error;
        }
    },

};