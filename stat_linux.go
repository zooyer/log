package log

import (
	"syscall"
	"time"
)

func (s *stat) createTime() time.Time {
	stat := s.FileInfo.Sys().(*syscall.Stat_t)

	return time.Unix(stat.Ctim.Unix())
}