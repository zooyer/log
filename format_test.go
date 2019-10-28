package log

import (
	"testing"
	"time"
)

func TestFormat(t *testing.T) {
	var record = new(Record)
	record.Time = time.Now()
	record.Level = "DEBUG"
	record.Tag = map[string]string{
		"module":  "apollo",
		"trace":   "10001241235",
		"type":    "rpc",
		"rpc":     "http",
		"latency": "10",
		"other":   `{"json":"value", " ": "   "}`,
	}
	record.Message = "hello world"

	t.Log(FormatJson(record))
	t.Log(FormatText(record))
	t.Log(FormatColorText(record))
	record.Level = "INFO"
	t.Log(FormatColorText(record))
	record.Level = "WARNING"
	t.Log(FormatColorText(record))
	record.Level = "ERROR"
	t.Log(FormatColorText(record))
}