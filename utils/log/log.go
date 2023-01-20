package log

import (
	"log"
	"os"
	"sync"
)

var (
	logger        *log.Logger
	mu            sync.Mutex
	defaultPrefix = ""
	levelFlags    = []string{"DEBUG", "INFO", "WARN", "ERROR", "FATAL"}
)

type logLevel int

// log levels
const (
	DEBUG logLevel = iota
	INFO
	WARNING
	ERROR
	FATAL
)

func init() {
	logger = log.New(os.Stdout, defaultPrefix, log.Lshortfile)
}

func SetPrefix(level logLevel) {
	logger.SetPrefix("[" + levelFlags[level] + "] ")
}

func Debug(v ...interface{}) {
	mu.Lock()
	defer mu.Unlock()
	SetPrefix(DEBUG)
	logger.Println(v...)
}

func Info(v ...interface{}) {
	mu.Lock()
	defer mu.Unlock()
	SetPrefix(INFO)
	logger.Println(v...)
}

func Warn(v ...interface{}) {
	mu.Lock()
	defer mu.Unlock()
	SetPrefix(WARNING)
	logger.Println(v...)
}

func Error(v ...interface{}) {
	mu.Lock()
	defer mu.Unlock()
	SetPrefix(ERROR)
	logger.Println(v...)
}

func Fatal(v ...interface{}) {
	mu.Lock()
	defer mu.Unlock()
	SetPrefix(FATAL)
	logger.Fatalln(v...)
}
