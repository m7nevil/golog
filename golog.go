package golog

import (
	"log"
	"os"
)

var file *os.File
var fileLogger *log.Logger
var consoleLogger *log.Logger

func init() {
	file, err := os.OpenFile("runtime.log", os.O_APPEND, os.ModeAppend)
	if err != nil {
		if os.IsExist(err) {
			log.Panic(err)
			return
		}
		file, err = os.Create("runtime.log")
	}
	if err != nil {
		log.Panic(err)
		return
	}

	fileLogger = log.New(file, "", log.Ldate|log.Ltime|log.Lmicroseconds)
	consoleLogger = log.New(os.Stdout, "", log.Ldate|log.Ltime|log.Lmicroseconds)
}

func Println(v ...interface{}) {
	consoleLogger.Println(v...)
	fileLogger.Println(v...)
}

func Panic(v ...interface{}) {
	consoleLogger.Panic(v...)
}

func Fatalln(v ...interface{}) {
	consoleLogger.Fatalln(v...)
	fileLogger.Fatalln(v...)
}

func Printf(format string, v ...interface{}) {
	consoleLogger.Printf(format, v...)
	fileLogger.Printf(format, v...)
}

func CloseFile() {
	file.Close()
}
