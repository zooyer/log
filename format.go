package log

import (
	"encoding/json"
	"strconv"
	"strings"
	"sync"
)

type Formatter func(record *Record) string

const format = "2006-01-02 15:04:05.000"

// {"level":"DEBUG","tag":{"rpc":"http","trace":"10001241235","type":"rpc"},"time":"2019-10-27 00:08:02.908","message":"debug"}
func FormatJson(record *Record) string {
	var obj = make(map[string]interface{})
	obj["time"] = record.Time.Format(format)
	obj["level"] = record.Level
	obj["message"] = record.Message
	if record.Tag != nil {
		obj["tag"] = record.Tag
	}

	data, _ := json.Marshal(obj)

	return string(data)
}

//    date        time     level      tag                                           message
// 2019-10-27 00:04:36.028 DEBUG "trace"="10001241235" "type"="rpc" "rpc"="http" "hello world"
func FormatText(record *Record) string {
	var builder = pool.Get().(*strings.Builder)
	defer pool.Put(builder)

	builder.Reset()
	builder.WriteString(record.Time.Format(format))
	builder.WriteString(" ")
	builder.WriteString(record.Level)

	var key, val string
	for k, v := range record.Tag {
		key = strconv.Quote(k)
		val = strconv.Quote(v)

		builder.WriteString(" ")
		builder.WriteString(key)
		builder.WriteString("=")
		builder.WriteString(val)
	}

	message := strconv.Quote(record.Message)
	builder.WriteString(" ")
	builder.WriteString(message)

	return builder.String()
}

func FormatColorText(record *Record) string {
	var builder = pool.Get().(*strings.Builder)
	defer pool.Put(builder)

	builder.Reset()

	// 时间渲染为黄色
	builder.WriteString("\033[33m")
	builder.WriteString(record.Time.Format(format))
	builder.WriteString("\033[0m")
	builder.WriteString(" ")

	// 等级按级别渲染
	switch record.Level {
	case "DEBUG":
		builder.WriteString("\033[32m")
	case "INFO":
		builder.WriteString("\033[34m")
	case "WARNING":
		builder.WriteString("\033[35m")
	case "ERROR":
		builder.WriteString("\033[31m")
	default:
		builder.WriteString("\033[31m")
	}
	builder.WriteString(record.Level)
	builder.WriteString("\033[0m")

	// tag渲染为青色
	var key, val string
	for k, v := range record.Tag {
		key = strconv.Quote(k)
		val = strconv.Quote(v)

		builder.WriteString(" ")
		builder.WriteString("\033[36m")
		builder.WriteString(key)
		builder.WriteString("=")
		builder.WriteString(val)
		builder.WriteString("\033[0m")
	}

	message := strconv.Quote(record.Message)
	builder.WriteString(" ")
	builder.WriteString(message)

	return builder.String()
}

var pool = sync.Pool{
	New: func() interface{} {
		return new(strings.Builder)
	},
}
