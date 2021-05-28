// Package bclogger wraps the log package to display log messages and save them to a file.
package bclogger

import (
	"fmt"
	"io"
	"log"
	"os"
)

var (
	InfoLogger  *log.Logger // InfoLogger displys info messages.
	WarnLogger  *log.Logger // WarnLogger displys warning messages.
	ErrorLogger *log.Logger // ErrorLogger displys error messages.
	FatalLogger *log.Logger // FatalLogger displys fatal messages.
)

// Init is to initialize all loggers for info, warning, error and fatal.
// bfile indicates that log messages are saved to a file.
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

// INFOMSG is to display info messages.
func INFOMSG(v ...interface{}) {
	InfoLogger.Output(2, fmt.Sprint(v...))
}

// INFOMSGF is to display info messages with a formater.
func INFOMSGF(format string, v ...interface{}) {
	InfoLogger.Output(2, fmt.Sprintf(format, v...))
}

// WARNMSG is to display warning messages.
func WARNMSG(v ...interface{}) {
	WarnLogger.Output(2, fmt.Sprint(v...))
}

// WARNMSGF is to display warning messages with a formater.
func WARNMSGF(format string, v ...interface{}) {
	InfoLogger.Output(2, fmt.Sprintf(format, v...))
}

// ERRORMSG is to display error messages.
func ERRORMSG(v ...interface{}) {
	ErrorLogger.Output(2, fmt.Sprint(v...))
}

// ERRORMSGF is to display error messages with a formater.
func ERRORMSGF(format string, v ...interface{}) {
	InfoLogger.Output(2, fmt.Sprintf(format, v...))
}

// FATALMSG is to display fatal messages.
func FATALMSG(v ...interface{}) {
	s := fmt.Sprint(v...)
	FatalLogger.Output(2, s)
	os.Exit(1)
}

// FATALMSGF is to display fatal messages with a formater.
func FATALMSGF(format string, v ...interface{}) {
	s := fmt.Sprintf(format, v...)
	FatalLogger.Output(2, s)
	os.Exit(1)
}
