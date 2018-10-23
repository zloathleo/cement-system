package logger

import (
	"gopkg.in/natefinch/lumberjack.v2"
	"io"
	"log"
	"os"
)

var (
	D interface{} = "[D]"
	I interface{} = "[I]"
	W interface{} = "[W]"
	E interface{} = "[E]"
	F interface{} = "[F]"
)

var mw io.Writer
var logLevel int

func Init() {
	mode := "dev"
	if mode == "dev" {
		log.SetFlags(log.Ldate | log.Ltime | log.Lshortfile)
	} else {
		log.SetFlags(log.Ldate | log.Ltime)
	}

	lumberjackLogger := &lumberjack.Logger{
		Filename: "cement_center_server_log.log",
		MaxSize:  50, // megabytes
		MaxAge:   7,  //days
	}
	mw = io.MultiWriter(os.Stdout, lumberjackLogger)
	log.SetOutput(mw)
}

func GetOutput() io.Writer {
	return mw
}

func SetLevel(level string) {
	switch level {
	case "D": //debug
		logLevel = 0
		break
	case "I": //info
		logLevel = 10
		break
	case "W": //warn
		logLevel = 20
		break
	case "E": //error
		logLevel = 30
		break
	case "F": //Fatal
		logLevel = 40
		break
	default:
		logLevel = 20
		break
	}
}

func appendSlice(flag interface{}, v []interface{}) []interface{} {
	vv := make([]interface{}, len(v)+1, len(v)+1)
	vv[0] = flag
	for i := 1; i <= len(v); i++ {
		vv[i] = v[i-1]
	}
	return vv
}

func Debug(v ...interface{}) {
	if logLevel < 10 {
		log.Println(appendSlice(D, v)...)
	}
}

func Debugf(format string, v ...interface{}) {
	if logLevel < 10 {
		log.Printf("[D] "+format, v...)
	}
}

func Info(v ...interface{}) {
	if logLevel < 20 {
		log.Println(appendSlice(I, v)...)
	}
}

func Infof(format string, v ...interface{}) {
	if logLevel < 20 {
		log.Printf("[I] "+format, v...)
	}
}

func Warnf(format string, v ...interface{}) {
	if logLevel < 30 {
		log.Printf("[W] "+format, v...)
	}
}

func Warnln(v ...interface{}) {
	if logLevel < 30 {
		log.Println(appendSlice(W, v)...)
	}
}

func Errorf(format string, v ...interface{}) {
	if logLevel < 40 {
		log.Printf("[E] "+format, v...)
	}
}

func Fatalf(format string, v ...interface{}) {
	log.Fatalf("[F] "+format, v...)
}
