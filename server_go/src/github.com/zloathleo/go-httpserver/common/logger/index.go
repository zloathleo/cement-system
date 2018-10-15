package logger

import (
	"os"
	"log"
	"io"
)

var mw io.Writer

func Init(){
	log.SetFlags(log.Ldate | log.Ltime | log.Lshortfile )
	f, err := os.OpenFile("logfile.log", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		log.Fatalf("log file open error : %v", err)
	}
	mw = io.MultiWriter(os.Stdout, f)
	log.SetOutput(mw)
}

func GetOutput() io.Writer{
	return mw
}

func Printf(format string, v ...interface{}) {
	log.Printf(format,v...)
}

func Println(v ...interface{}) {
	log.Println(v...)
}

func Print(v ...interface{}) {
	log.Print(v...)
}

func Fatal(v ...interface{}) {
	log.Fatal(v...)
}

func Fatalf(format string, v ...interface{}) {
	log.Fatalf(format,v...)
}
