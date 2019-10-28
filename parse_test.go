package log

import (
	"testing"
	"time"
)

func TestParseText(t *testing.T) {
	var str = `2019-10-27 00:04:36.028 DEBUG "trace"="10001241235" "type"="rpc" "rpc"="http" "hello world"
`
	r, err := ParseText(str)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(FormatJson(r))

	r = new(Record)
	str = FormatText(r)
	if r, err = ParseText(str); err != nil {
		t.Fatal(err)
	}
	t.Log(str)
}

func TestParseColorText(t *testing.T) {
	var r = new(Record)
	r.Time = time.Now()
	r.Level = "ERROR"
	r.Message = "hello"
	r.Tag = Tag{"key": "val"}
	str := FormatColorText(r)
	t.Log(str)
	record, err := ParseColorText(str)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(FormatJson(record))

	r = new(Record)
	str = FormatColorText(r)
	t.Log(str)
	record, err = ParseColorText(str)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(FormatJson(record))
}

func TestParseJson(t *testing.T) {
	var str = `{"level":"ERROR","message":"hello","tag":{"key":"val"},"time":"2019-10-27 22:39:08.898"}`
	r, err := ParseJson(str)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(FormatColorText(r))

	str = `{"level":"","message":"","time":"0001-01-01 00:00:00.000"}`
	r, err = ParseJson(str)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(FormatColorText(r))
}
