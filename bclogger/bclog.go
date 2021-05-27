package bclogger

import (
	"fmt"
	"io"
	"log"
	"os"
)

var (
	WarnLogger  *log.Logger
	InfoLogger  *log.Logger
	ErrorLogger *log.Logger
	FatalLogger *log.Logger
)

func Init(bfile bool) {
	var mw io.Writer
	if bfile == true {
		file, err := os.OpenFile("logs.txt", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)
		if err != nil {
			log.Fatal(err)
		}
		mw = io.MultiWriter(os.Stdout, file)
	} else {
		mw = io.MultiWriter(os.Stdout)
	}

	InfoLogger = log.New(mw, "INFO : ", log.Ldate|log.Ltime|log.Lshortfile)
	WarnLogger = log.New(mw, "WARN : ", log.Ldate|log.Ltime|log.Lshortfile)
	ErrorLogger = log.New(mw, "ERROR: ", log.Ldate|log.Ltime|log.Lshortfile)
	FatalLogger = log.New(mw, "FATAL: ", log.Ldate|log.Ltime|log.Lshortfile)
}

func INFOMSG(v ...interface{}) {
	InfoLogger.Output(2, fmt.Sprint(v...))
}

func INFOMSGF(format string, v ...interface{}) {
	InfoLogger.Output(2, fmt.Sprintf(format, v...))
}

func WARNMSG(v ...interface{}) {
	WarnLogger.Output(2, fmt.Sprint(v...))
}

func WARNMSGF(format string, v ...interface{}) {
	InfoLogger.Output(2, fmt.Sprintf(format, v...))
}

func ERRORMSG(v ...interface{}) {
	ErrorLogger.Output(2, fmt.Sprint(v...))
}

func ERRORMSGF(format string, v ...interface{}) {
	InfoLogger.Output(2, fmt.Sprintf(format, v...))
}

func FATALMSG(v ...interface{}) {
	s := fmt.Sprint(v...)
	FatalLogger.Output(2, s)
	os.Exit(1)
}

func FATALMSGF(format string, v ...interface{}) {
	s := fmt.Sprintf(format, v...)
	FatalLogger.Output(2, s)
	os.Exit(1)
}
