const fs = require('fs')
const path = require("path")


//exe路径
const execPath = path.dirname(process.execPath);
//配置名称
const configName = "/config/cement-burner-system-config.json";

const configFile = path.join(execPath, configName);
const configFileDev = path.join(__dirname, "..", configName);

function readConfig() {
    fs.readFile(configFile, 'utf8', readProdConfigFile);
}

////////////////////////////////
////读取config文件
////////////////////////////////
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

    let _config = JSON.parse(fileContents);
  
}

module.exports = { 
    config: readConfig,
};