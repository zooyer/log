package log

import (
	"os"
	"time"
)

type stat struct {
	os.FileInfo
}

func fileInfo(info os.FileInfo) *stat {
	return &stat{
		FileInfo: info,
	}
}

func (s *stat) Name() string {
	return s.FileInfo.Name()
}

func (s *stat) Size() int64 {
	return s.FileInfo.Size()
}

func (s *stat) Mode() os.FileMode {
	return s.FileInfo.Mode()
}

func (s *stat) ModTime() time.Time {
	return s.FileInfo.ModTime()
}

func (s *stat) IsDir() bool {
	return s.FileInfo.IsDir()
}

func (s *stat) Sys() interface{} {
	return s.FileInfo.Sys()
}

func (s stat) CreateTime() time.Time {
	return s.createTime()
}
