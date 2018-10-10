const electron = require("electron")
const log = require("electron-log")
const path = require("path")
const url = require("url")
const fs = require('fs')

const configFile = "d:/cement-burner-system-config.json";
const logFile = "d:/cement-burner-system.log";
fs.readFile(logFile, 'utf8', (err, fileContents) => {
    if (err) {
        console.error(err)
        return
    }
    try {
        console.log("fileContents:", fileContents);
        const data = JSON.parse(fileContents)
        console.log(data);
    } catch (err) {
        console.error(err)
    }
})

//electron 
const app = electron.app
const BrowserWindow = electron.BrowserWindow
const globalShortcut = electron.globalShortcut
const ipcMain = electron.ipcMain
const Menu = electron.Menu
const Tray = electron.Tray
let mainWindow;

function resolveRoot(dir) {
    return path.join(__dirname, "..", dir);
}

const logoImagePath = resolveRoot("/public/assets/img/logo256.png");
const appLoginPath = "file://" + resolveRoot("/public/index.html");

//exe当前路径
const execPath = path.dirname(process.execPath);

//log
// path.join(execPath, "/cement-burner-system.txt");
log.transports.file.file = logFile;
log.transports.console.level = "info";
log.transports.file.level = "info";

log.info("app start  -----------------------");
log.info("chrome ver:" + process.versions.chrome);
log.info("dirname:", __dirname);
log.info("execPath:", execPath);


function readyInit() {
    let screenSize = electron.screen.getPrimaryDisplay().size;
    log.info("screenSize:" + screenSize.width + "," + screenSize.height);
    //进度条
    _showMainWindow();
}

function _showMainWindow() {
    

    // Create the browser window.
    mainWindow = new BrowserWindow({
        frame: true,
        center: true,
        icon: logoImagePath,
        width: 1024,
        height: 768,
        autoHideMenuBar: true,
        minWidth: 1024,
        minHeight: 768,

        show: false,
        //禁用 Node.js
        webPreferences: {
            nodeIntegration: true
        }
    });
    mainWindow.maximize();
    mainWindow.loadURL(appLoginPath);

    // initTray();

    initShortcut();

    //进程间通讯
    initIpc();

    // Open the DevTools.
    mainWindow.webContents.openDevTools();

    mainWindow.once("ready-to-show", () => {
        mainWindow.show();
        mainWindow.maximize();
    });

    // Emitted when the window is closed.
    mainWindow.on("closed", function () {
        mainWindow = null;
    });

}

//初始化快捷键
function initShortcut() {
    globalShortcut.register("Ctrl+F5", function () {
        mainWindow.reload();
    })
    globalShortcut.register("Ctrl+F11", function () {
        mainWindow.setFullScreen(!mainWindow.isFullScreen());
    })
    globalShortcut.register("Ctrl+F12", function () {
        mainWindow.webContents.openDevTools();
    })
}

//初始化Tray
function initTray() {
    let appTray = new Tray(logoImagePath);
    const contextMenu = Menu.buildFromTemplate([{
        label: "RodinX Exit",
        click: function () {
            rodinxExit();
        }
    }]);
    appTray.setToolTip("RodinX");
    appTray.setContextMenu(contextMenu);
}

//进程间通讯
function initIpc() {
    ipcMain.on("ipc_client_restore", function (event, arg) {
        mainWindow.restore();
    });
    ipcMain.on("ipc_client_close", function (event, arg) {
        appExit();
    });
    ipcMain.on("ipc_client_max", function (event, arg) {
        mainWindow.maximize();
    });
}

////////////////////
///基础
////////////////////
function appExit() {
    app.quit();
}

// This method will be called when Electron has finished 
app.on("ready", readyInit);

// Quit when all windows are closed.
app.on("window-all-closed", function () {
    appExit();
});

//退出
app.on("will-quit", function () {
    globalShortcut.unregisterAll();
})

app.on("activate", function () {
    if (mainWindow === null) {
        readyInit();
    }
}); 