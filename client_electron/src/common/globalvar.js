export default {
    appName: "后台管理",
    GlobalEventHub: undefined,


    //配置模式 直接查看数据和控制信号
    configMode: true,

    // fetchServerHostURL: "",
    fetchServerHostURL: "http://localhost:8080",

    adminMenuItems: [
        {
            key: "group",
            display: "分组管理",
            icon: "account-group",
            desc: ""
        },
        {
            key: "account",
            display: "用户管理",
            icon: "account",
            desc: ""
        }
    ],

    operatorMenuItems: [
        {
            key: "device",
            display: "Device",
            icon: "account-group",
            desc: ""
        },
        {
            key: "page",
            display: "Page",
            icon: "account",
            desc: ""
        },
        {
            key: "resource",
            display: "Resource",
            icon: "account",
            desc: ""
        }
        ,
        {
            key: "qr",
            display: "Qr",
            icon: "account",
            desc: ""
        },
        {
            key: "message",
            display: "Message",
            icon: "account",
            desc: ""
        },
        {
            key: "live",
            display: "Live",
            icon: "account",
            desc: ""
        },
        {
            key: "publish",
            display: "Publish",
            icon: "account",
            desc: ""
        }
    ],

    parseDateTime: function (_long) {
        if (_long) {
            return new Date(_long).Format("yyyy-MM-dd HH:mm:ss");
        } else {
            return "-";
        }
    },

    parseMsgState: function (_state) {
        if (_state === 1) {
            return "正在播放";
        } else if (_state === -1) {
            return "已播放";
        } else {
            return "准备播放";
        }
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