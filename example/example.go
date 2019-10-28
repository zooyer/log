package main

import (
	"time"

	"github.com/zooyer/log"
)

func main() {
	// 0 config
	log.D("debug message")   // 2019-10-28 23:07:56.369 DEBUG "debug message"
	log.I("info message")    // 2019-10-28 23:07:56.369 INFO "info message"
	log.W("warning message") // 2019-10-28 23:07:56.369 WARNING "warning message"
	log.E("error message")   // 2019-10-28 23:07:56.369 ERROR "error message"

	// 1. create file rotating(default: size/time)
	rotating := log.NewFileCountRotating(1024, 10)

	// 2. create file recorder(default: terminal/file/network), custom formatter: json/text
	recorder, err := log.NewFileRecorder("example.log", log.FormatJson, rotating)
	if err != nil {
		panic(err)
	}
	defer recorder.Close()

	// 3. create logger, each level can be mapped to different recorder
	logger := log.NewLogger()
	logger.SetRecorder(recorder, "DEBUG", "INFO")
	logger.SetDefault(recorder)

	// 4. create log, logger's wrap
	l := log.NewLog(logger)
	l.Tag("id", "1001").Tag("type", "test").Debug("custom debug log")
	l.Error("custom error log")

	// 5. custom
	var record = new(log.Record)
	record.Time = time.Now()
	record.Level = "record"
	record.Message = "custom log"
	record.Tag = make(log.Tag)
	record.Tag["id"] = "1001"
	record.Tag["type"] = "test"
	logger.Record(record)
}
