package log

import (
	"os"
	"testing"
)

func TestStat_CreateTime(t *testing.T) {
	info, err := os.Stat("./stat.go")
	if err != nil {
		t.Fatal(err)
	}
	t.Log(fileInfo(info).CreateTime())
}
