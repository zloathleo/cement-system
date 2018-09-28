const path = require("path")
const fs = require('fs')
const electron = require('electron')
const log = require("./log")

const app = electron.app;
const ipcMain = electron.ipcMain
//var
var loadingWindow;
var mainWindow;
var appConfig;

//加载完显示主界面
function showMain() {
    initIpc();

    let mainPage = path.join(__dirname, "..", "/public/index.html");

    mainWindow = new electron.BrowserWindow({
        frame: true,
        center: true,
        minWidth: 800,
        minHeight: 600,
        show: false,
        autoHideMenuBar: true,
        //禁用 Node.js
        webPreferences: {
            nodeIntegration: true
        }
    });
    // mainWindow.maximize();
    // mainWindow.webContents.openDevTools();
    mainWindow.loadURL(mainPage);

}

//进程间通讯监听
function initIpc() {
    //主页面加载完成后再显示
    ipcMain.once("electron.renderer.mainpage.loaded", function (event, arg) {
        log.info("revi renderer message:", arg);

        mainWindow.show();
        mainWindow.maximize();
        mainWindow.webContents.openDevTools();

        mainWindow.webContents.send('electron.main.config', appConfig);
        loadingWindow.close();
    });
}

function showLoading() {
    let loadingPage = path.join(__dirname, "..", "/public/loading.html");
    loadingWindow = new electron.BrowserWindow({
        frame: false,
        center: true,
        width: 800,
        height: 600,
        titleBarStyle: "hidden",
        thickFrame: false,
        autoHideMenuBar: true,
        minWidth: 800,
        minHeight: 600,
        show: false,
        //禁用 Node.js
        webPreferences: {
            nodeIntegration: true
        }
    });
    // loadingWindow.webContents.openDevTools();
    loadingWindow.loadURL(loadingPage);
    loadingWindow.once('ready-to-show', function () {
        loadingWindow.show();
        readConfig();
    });

}

//////////////////////
///common
//////////////////////
function appExit() {
    app.quit();
}

// Quit when all windows are closed.
app.on("window-all-closed", function () {
    appExit();
});

//退出
app.on("will-quit", function () {
    // globalShortcut.unregisterAll();
})

app.on("activate", function () {
    // if (mainWindow === null) {
    //     readyInit();
    // }
});



////////////////////////////////
////读取config文件
////////////////////////////////
//exe路径
const execPath = path.dirname(process.execPath);
//配置名称
const configName = "/config/cement-burner-system-config.json";
const configProdFile = path.join(execPath, configName);
const configFileDev = path.join(__dirname, "..", configName);

function readConfig() {
    fs.readFile(configProdFile, 'utf8', readProdConfigFile);
}

function readProdConfigFile(err, fileContents) {
    if (err) {
        //开发环境或未配置
        log.info("production config file is not loaded...");
        log.info("now check dev config file...");

        fs.readFile(configFileDev, 'utf8', readDevConfigFile);
    } else {
        log.info("production config file is loading...");
        parseConfigContent(fileContents);
    }
}

function readDevConfigFile(err, fileContents) {
    if (err) {
        //开发环境也未读取成功
        log.info("config file is not loading...");
        electron.dialog.showErrorBox("config file err", "config file is not loading...");
        electron.app.quit();
    } else {
        log.info("dev config file is loading...");
        parseConfigContent(fileContents);
    }
}

//解析config
function parseConfigContent(fileContents) {
    log.info("config file:", fileContents);

    if (loadingWindow) {
        sendLoadMessage("start loading config...", 10);

        setTimeout(function () {
            appConfig = JSON.parse(fileContents);
            if (appConfig) {
                sendLoadMessage("loaded config success...", 30);

                setTimeout(function () {
                    // let ip = _config.device_interface_ip;
                    // log.info(ip);
                    sendLoadMessage("starting app... ", 80);

                    //自身需要时间
                    showMain(); 
                }, 1000);
            } else {
                //show error
            }
        }, 1000);
    }
}

function sendLoadMessage(m, p) {
    loadingWindow.webContents.send('electron.main.loading.message', {
        message: m,
        percent: p
    });
}

module.exports = {
    loading: showLoading
};
