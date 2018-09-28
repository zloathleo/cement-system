const electron = require('electron')
const log = require("./log")
const main = require("./electron-main");

log.info("safefire cement-burner-system is starting...");
 
electron.app.on("ready", function () {
    //loading页面
    main.loading();
});
