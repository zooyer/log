package log

import (
	"strconv"
	"testing"
	"time"
)

func TestNewStdoutRecorder(t *testing.T) {
	var record = new(Record)
	record.Time = time.Now()
	record.Level = "DEBUG"
	record.Tag = Tag{
		"key": "val",
	}

	r := NewStdoutRecorder(FormatColorText)
	defer r.Close()
	r.Record(record)

	r = NewStderrRecorder(FormatColorText)
	defer r.Close()
	r.Record(record)
}

func TestNewRecordBuffer(t *testing.T) {
	var record = new(Record)
	record.Time = time.Now()
	record.Level = "DEBUG"
	record.Tag = Tag{
		"key": "val",
	}

	r := NewStdoutRecorder(FormatColorText)
	buf := NewRecordBufferSize(r, 10)
	for i := 0; i < 5; i++ {
		buf.Record(record)
	}
	time.Sleep(3 * time.Second)
	buf.Flush()
	buf.Close()

	buf = NewRecordBufferSize(buf, 1)
	for i := 0; i < 5; i++ {
		buf.Record(record)
		time.Sleep(time.Second)
	}
	time.Sleep(3 * time.Second)
	buf.Close()
}

func TestNewFileRecorder(t *testing.T) {
	var record = new(Record)
	record.Time = time.Now()
	record.Level = "DEBUG"
	record.Tag = Tag{
		"key": "val",
	}

	recorder, err := NewFileRecorder("test.log", FormatText, nil)
	if err != nil {
		t.Fatal(err)
	}
	defer recorder.Close()

	for i := 0; i < 105; i++ {
		record.Tag["id"] = strconv.Itoa(i)
		recorder.Record(record)
	}
}

func TestNewFileCountRotating(t *testing.T) {
	var record = new(Record)
	record.Time = time.Now()
	record.Level = "DEBUG"
	record.Tag = Tag{
		"key": "val",
	}

	rotate := NewFileCountRotating(1024, 1)
	recorder, err := NewFileRecorder("test.log", FormatText, rotate)
	if err != nil {
		t.Fatal(err)
	}
	defer recorder.Close()

	for i := 0; i < 105; i++ {
		record.Tag["id"] = strconv.Itoa(i)
		recorder.Record(record)
	}
}

func TestNewFileTimeRotating(t *testing.T) {
	var record = new(Record)
	record.Time = time.Now()
	record.Level = "DEBUG"
	record.Tag = Tag{
		"key": "val",
	}

	r := NewFileTimeRotating(time.Minute, true)
	recorder, err := NewFileRecorder("test.log", FormatText, r)
	if err != nil {
		t.Fatal(err)
	}
	defer recorder.Close()

	for i := 0; i < 60*5; i++ {
		record.Time = time.Now()
		record.Tag["id"] = strconv.Itoa(i)
		recorder.Record(record)
		time.Sleep(time.Second)
	}
}

func TestNewRecordBufferSize(t *testing.T) {
	var record = new(Record)
	record.Time = time.Now()
	record.Level = "DEBUG"
	record.Tag = make(Tag)

	recorder, err := NewFileRecorder("test.log", FormatText, nil)
	if err != nil {
		t.Fatal(err)
	}
	defer recorder.Close()

	var start = time.Now()
	for i := 0; i < 4096*1000; i++ {
		record.Tag["id"] = strconv.Itoa(i)
		recorder.Record(record)
	}
	t.Log("no buf:", time.Now().Sub(start))

	start = time.Now()
	buf := NewRecordBufferSize(recorder, 4096)
	for i := 0; i < 4096*1000; i++ {
		record.Tag["id"] = strconv.Itoa(i)
		buf.Record(record.Clone())
	}
	buf.Flush()
	t.Log("buf:", time.Now().Sub(start))
}

func TestNewMultiRecorder(t *testing.T) {
	var record = new(Record)
	record.Time = time.Now()
	record.Level = "DEBUG"
	record.Tag = Tag{
		"key": "val",
	}

	stdout := NewStdoutRecorder(FormatColorText)
	recorder := NewMultiRecorder(stdout, stdout, stdout)

	recorder.Record(record)
}