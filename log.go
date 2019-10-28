package log

import (
	"fmt"
	"strings"
	"sync"
	"time"
)

type Logger interface {
	Default() Recorder
	Recorder(level string) Recorder
	SetDefault(recorder Recorder) Logger
	SetRecorder(recorder Recorder, level ...string) Logger
	Record(record *Record) Logger
}

type logger struct {
	sync.RWMutex
	levelRecorder   map[string]Recorder
	defaultRecorder Recorder
}

func NewLogger() Logger {
	return &logger{
		levelRecorder: make(map[string]Recorder),
	}
}

func (l *logger) Default() Recorder {
	return l.defaultRecorder
}

func (l *logger) SetRecorder(recorder Recorder, level ...string) Logger {
	l.Lock()
	defer l.Unlock()

	for i := range level {
		l.levelRecorder[level[i]] = recorder
	}

	return l
}

func (l *logger) SetDefault(recorder Recorder) Logger {
	l.defaultRecorder = recorder

	return l
}

func (l *logger) Recorder(level string) Recorder {
	return l.levelRecorder[level]
}

func (l *logger) Record(record *Record) Logger {
	if record == nil {
		return l
	}
	recorder := l.levelRecorder[record.Level]
	if recorder == nil {
		recorder = l.defaultRecorder
	}
	if recorder != nil {
		recorder.Record(record)
	}

	return l
}

type Log interface {
	Debug(v ...interface{})
	Info(v ...interface{})
	Warning(v ...interface{})
	Error(v ...interface{})

	Tag(key, val string) Log
	Logger() Logger
}

type log struct {
	tag    Tag
	logger Logger
	mutex  sync.RWMutex
}

func NewLog(logger Logger) Log {
	return &log{
		logger: logger,
	}
}

func (l *log) format(v ...interface{}) string {
	var builder = pool.Get().(*strings.Builder)
	defer pool.Put(builder)

	builder.Reset()
	for i, arg := range v {
		if i > 0 {
			builder.WriteByte(' ')
		}
		builder.WriteString(fmt.Sprint(arg))
	}

	return builder.String()
}

func (l *log) record(level, message string) *Record {
	l.mutex.RLock()
	defer l.mutex.RUnlock()

	var record = new(Record)
	record.Time = time.Now()
	record.Level = level
	record.Message = message
	if l.tag != nil {
		record.Tag = l.tag
		l.tag = nil
	}

	return record
}

func (l *log) output(level string, v ...interface{}) {
	message := l.format(v...)
	record := l.record(level, message)
	l.logger.Record(record)
}

func (l *log) Tag(key, val string) Log {
	l.mutex.Lock()
	defer l.mutex.Unlock()

	tag := l.tag.Clone()
	if tag == nil {
		tag = make(Tag)
	}
	tag.Add(key, val)

	return &log{
		tag:    tag,
		logger: l.logger,
	}
}

func (l *log) Logger() Logger {
	return l.logger
}

func (l *log) Debug(v ...interface{}) {
	l.output("DEBUG", v...)
}

func (l *log) Info(v ...interface{}) {
	l.output("INFO", v...)
}

func (l *log) Warning(v ...interface{}) {
	l.output("WARNING", v...)
}

func (l *log) Error(v ...interface{}) {
	l.output("ERROR", v...)
}
