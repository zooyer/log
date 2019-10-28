package log

import (
	"errors"
	"testing"
)

func TestDebug(t *testing.T) {
	Debug("level:", "debug")
	Info("hello", "world")
	Warning("warning", 404)
	Error("error", errors.New("test error"))
}
