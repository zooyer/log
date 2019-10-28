# The Go Log


![](https://travis-ci.org/boennemann/badges.svg?branch=master)  ![](https://img.shields.io/badge/license-MIT-blue.svg)  ![](https://img.shields.io/badge/godoc-reference-blue.svg)

Go Log is customizable modular log.

![](https://github.com/golang/go/blob/master/doc/gopher/fiveyears.jpg?raw=true)


#### Download and Install

```shell
go get github.com/zooyer/log
```

#### Features

- zero config, out of the box feature
- customizable modular
- high performance

##### Example: [example/example.go](example/example.go)

1. zero config

   ```go
   package main
   
   import (
   	"github.com/zooyer/log"
   )
   
   func main() {
       // 0 config
   	log.D("debug message")   // 2019-10-28 23:07:56.369 DEBUG "debug message"
   	log.I("info message")    // 2019-10-28 23:07:56.369 INFO "info message"
   	log.W("warning message") // 2019-10-28 23:07:56.369 WARNING "warning message"
   	log.E("error message")   // 2019-10-28 23:07:56.369 ERROR "error message"
   }
   ```

   output:

   ```shell
   2019-10-28 23:07:56.369 DEBUG "debug message"
   2019-10-28 23:07:56.369 INFO "info message"
   2019-10-28 23:07:56.369 WARNING "warning message"
   2019-10-28 23:07:56.369 ERROR "error message"
   ```

   

2. custom

   ```go
   package main
   
   import (
   	"time"
   
   	"github.com/zooyer/log"
   )
   
   func main() {
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
   ```

   output:

   ```shell
   cat example.log
   {"level":"DEBUG","message":"custom debug log","tag":{"id":"1001","type":"test"},"time":"2019-10-28 23:41:34.385"}
   {"level":"ERROR","message":"custom error log","time":"2019-10-28 23:41:34.386"}
   {"level":"record","message":"custom log","tag":{"id":"1001","type":"test"},"time":"2019-10-28 23:41:34.386"}
   ```

   