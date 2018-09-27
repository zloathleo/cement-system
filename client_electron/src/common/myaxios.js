import axios from 'axios'
import qs from 'qs';
import globalvar from './globalvar';

console.log("初始化axios,", globalvar.fetchServerHostURL);
axios.defaults.baseURL = globalvar.fetchServerHostURL;
//设置默认请求头
axios.defaults.headers = {
    'Content-Type': 'application/x-www-form-urlencoded',
    // 'Cache-Control': 'no-cache,no-store',
}

axios.defaults.timeout = 6000;

//不使用URLSearchParams  
axios.defaults.transformRequest = [function (data) {
    if (data === undefined) {
        return data;
    } else {
        let _type = data.constructor;
        if ("FormData" === _type.name) {
            return data;
        } else {
            data = qs.stringify(data);
            return data;
        }

    }
}];

//请求拦截器
axios.interceptors.request.use(function (config) {
    if (!config.hiddenLoading) {
        globalvar.GlobalEventHub.$emit('appLoading', true);
    }
    return config;
}, function (error) {
    if (!config.hiddenLoading) {
        globalvar.GlobalEventHub.$emit('appLoading', true);
    }
    return Promise.reject(error);
});

// 添加响应拦截器
axios.interceptors.response.use(function (response) {
    if (!response.config.hiddenLoading) {
        globalvar.GlobalEventHub.$emit('appLoading', false);
    }
    return response;
}, function (error) {
    if (!error.config.hiddenLoading) {
        globalvar.GlobalEventHub.$emit('appLoading', false);
    }
    return Promise.reject(error);



});

//响应拦截器即异常处理
axios.interceptors.response.use(function (response) {
    return response
}, function (err) {
    if (err && err.response) {
        switch (err.response.status) {
            case 400:
                err.message = '错误请求'
                break;
            case 401:
                err.message = '未授权，请重新登录'
                break;
            case 403:
                err.message = '拒绝访问'
                break;
            case 404:
                err.message = '请求错误,未找到该资源'
                break;
            case 405:
                err.message = '请求方法未允许'
                break;
            case 408:
                err.message = '请求超时'
                break;
            case 500:
                err.message = '服务器端出错'
                break;
            case 501:
                err.message = '网络未实现'
                break;
            case 502:
                err.message = '网络错误'
                break;
            case 503:
                err.message = '服务不可用'
                break;
            case 504:
                err.message = '网络超时'
                break;
            case 505:
                err.message = 'http版本不支持该请求'
                break;
            default:
                err.message = '连接错误err.response.status'
        }
    } else {
        err.message = "连接到服务器失败"
    }
    return Promise.reject(err);
});

export default axios;