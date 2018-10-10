const electron = require('electron')

const globalShortcut = electron.globalShortcut;
const app = electron.app;

const log = require("./log")
const main = require("./electron-main");

log.info("safefire cement-burner-system is starting...");
 
app.on("ready", function () {
    //loading页面
    main.loading();
});

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
    globalShortcut.unregisterAll();
})

 
