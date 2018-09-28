const path = require("path")
const log = require("electron-log")

//日志名称
const execPath = path.dirname(process.execPath);
const logName = "/cement-burner-system.log";  

// log.transports.file.file = path.join(execPath, logName);
// log.transports.console.level = "info";
// log.transports.file.level = "info";

log.info("log init now");

module.exports = log;