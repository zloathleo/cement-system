const path = require("path")
const fs = require('fs')
const electron = require('electron')
const log = require("./log")

const app = electron.app;
const ipcMain = electron.ipcMain;
const globalShortcut = electron.globalShortcut;
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
    //监听页面完成情况
    mainWindow.loadURL(mainPage);

} 

//进程间通讯监听
function initIpc() {
    //主页面加载完成后再显示
    ipcMain.on("electron.renderer.mainpage.loaded", function (event, arg) {
        log.info("mainpage is loaded , show main window"); 
        if (!mainWindow.isVisible()) {
            //第一次启动
            mainWindow.show();
            mainWindow.maximize();
            // mainWindow.webContents.openDevTools();
            initShortcut();
        }
        if (loadingWindow && !loadingWindow.isDestroyed()) {
            loadingWindow.destroy();
        }
        
        mainWindow.webContents.send('electron.main.config', appConfig);
      
    });
}

//初始化快捷键
function initShortcut() {
    globalShortcut.register("Ctrl+F5", function () {
        mainWindow.reload();
    });
    globalShortcut.register("Ctrl+F11", function () {
        mainWindow.setFullScreen(!mainWindow.isFullScreen());
    });
    globalShortcut.register("Ctrl+F12", function () {
        mainWindow.webContents.openDevTools();
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
    loadingWindow.on("closed", function () {
        loadingWindow = null;
    });
}





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
                    sendLoadMessage("starting app... ", 80);

                    //自身需要时间
                    showMain();
                }, 500);
            } else {
                //show error
            }
        }, 500);
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
