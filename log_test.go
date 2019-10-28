package log

import (
	"testing"
	"time"
)

func TestNewLogger(t *testing.T) {
	var record = new(Record)
	record.Time = time.Now()
	record.Message = "test"
	l := NewLogger()
	l.SetDefault(NewStdoutRecorder(FormatColorText))

	// ???
	record.Level = "???"
	l.Record(record)

	// debug
	record.Level = "debug"
	l.Record(record)

	// DEBUG
	record.Level = "DEBUG"
	l.Record(record)
}

func TestNewLog(t *testing.T) {
	l := NewLog(defaultLogger())
	l.Debug("debug")
	l.Info("info")
	l.Warning("warning")
	l.Error("error")

	l.Tag("", "").Debug("tag debug")
	l.Tag("id", "1000").Debug("tag id debug")
	l.Tag("id", "1000").Tag("type", "text")
	l.Info("no tags...")

	l = l.Tag("id", "1000").Tag("type", "text")
	l.Info("has tags...")
	l.Info("no tags...")
}
