const electron = require('electron')

const path = require('path')
const url = require('url')

// Module to control application life.
const app = electron.app
// Module to create native browser window.
const BrowserWindow = electron.BrowserWindow
const globalShortcut = electron.globalShortcut
const ipcMain = electron.ipcMain
const Menu = electron.Menu
const Tray = electron.Tray

const logoImagePath = path.join(__dirname, '/public/assets/img/logo256.png');
const appLoginPath = path.join('file://', __dirname, '/public/index.html');

let mainWindow;
let screenSize;

function readyInit() {
    console.log('chrome ver.' + process.versions.chrome);
    screenSize = electron.screen.getPrimaryDisplay().size;
    console.log('screenSize:' + screenSize.width + "," + screenSize.height);
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
    // mainWindow.webContents.openDevTools();

    mainWindow.once('ready-to-show', () => {
        mainWindow.show();
        mainWindow.maximize();
    });

    // Emitted when the window is closed.
    mainWindow.on('closed', function () {
        mainWindow = null;
    });

}

//初始化快捷键
function initShortcut() {
    // root登录
    // globalShortcut.register('Ctrl+F10', function () {
    //     mainWindow.webContents.send('main-ipc-root-login', 'root');
    // })
    globalShortcut.register('Ctrl+F11', function () {
        mainWindow.setFullScreen(!mainWindow.isFullScreen());
    })
    globalShortcut.register('Ctrl+F12', function () {
        mainWindow.webContents.openDevTools();
    })
}

//初始化Tray
function initTray() {
    let appTray = new Tray(logoImagePath);
    const contextMenu = Menu.buildFromTemplate([{
        label: 'RodinX Exit',
        click: function () {
            rodinxExit();
        }
    }]);
    appTray.setToolTip('RodinX');
    appTray.setContextMenu(contextMenu);
}

//进程间通讯
function initIpc() { 
    ipcMain.on('ipc_client_restore', function (event, arg) { 
        mainWindow.restore();
    }); 
    ipcMain.on('ipc_client_close', function (event, arg) { 
        rodinxExit();
    });
    ipcMain.on('ipc_client_max', function (event, arg) { 
        mainWindow.maximize();
    });
}

function rodinxExit() {
    app.quit();
}

// This method will be called when Electron has finished 
app.on('ready', readyInit);

// Quit when all windows are closed.
app.on('window-all-closed', function () {
    rodinxExit();
});

//退出
app.on('will-quit', function () {
    globalShortcut.unregisterAll();
})

app.on('activate', function () {
    // On OS X it's common to re-create a window in the app when the
    // dock icon is clicked and there are no other windows open.
    if (mainWindow === null) {
        readyInit();
    }
}); 