package benckmark

import "github.com/zooyer/log"

func newDisabledZoolog() log.Log {
	return log.NewLog(log.NewLogger())
}

func newZoolog() log.Log {
	return log.NewLog(log.NewLogger())
}