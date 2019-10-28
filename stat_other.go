//+build !linux
//+build !windows

package log

import "time"

func (s *stat) createTime() time.Time {
	return time.Time{}
}
