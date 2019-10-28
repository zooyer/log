package log

import (
	"syscall"
	"time"
)

func (s *stat) createTime() time.Time {
	attr := s.FileInfo.Sys().(*syscall.Win32FileAttributeData)
	nano := attr.CreationTime.Nanoseconds()

	return time.Unix(0, nano)
}
