package log

import (
	"encoding/json"
	"errors"
	"fmt"
	"regexp"
	"strconv"
	"time"
)

// 2019-10-27 00:04:36.028 DEBUG "trace"="10001241235" "type"="rpc" "rpc"="http" "hello world"
func ParseText(line string) (*Record, error) {
	var levelBegin int
	var tagBegin int
	var space int
	for i := 0; i < len(line); i++ {
		if line[i] == ' ' {
			space++
			if space == 2 && i+1 < len(line) {
				levelBegin = i + 1
			}
			if space == 3 && i+1 < len(line) {
				tagBegin = i + 1
				break
			}
		}
	}
	if levelBegin == 0 || tagBegin == 0 {
		return nil, errors.New("invalid log text")
	}
	t, err := time.Parse(format, line[:levelBegin-1])
	if err != nil {
		return nil, err
	}
	level := line[levelBegin : tagBegin-1]

	var tag = make(Tag)
	var str string
	var equal bool
	var strBegin int
	for i := tagBegin; i < len(line); i++ {
		switch line[i] {
		case '\\':
			i++
		case '"':
			if strBegin == 0 {
				strBegin = i
			} else {
				if equal {
					equal = false
					if tag[str], err = strconv.Unquote(line[strBegin : i+1]); err != nil {
						return nil, err
					}
				} else {
					if str, err = strconv.Unquote(line[strBegin : i+1]); err != nil {
						return nil, err
					}
				}
				strBegin = 0
			}
		case '=':
			if equal {
				return nil, fmt.Errorf("invalid log text charset '=':%d", i+1)
			}
			equal = true
		}
	}

	var record = new(Record)
	record.Time = t
	record.Level = level
	if len(tag) > 0 {
		record.Tag = tag
	}
	record.Message = str

	return record, nil
}

func ParseColorText(line string) (*Record, error) {
	line = regexp.MustCompile(`\033\[\d*m`).ReplaceAllString(line, "")

	return ParseText(line)
}

func ParseJson(line string) (*Record, error) {
	var err error
	var obj map[string]interface{}
	if err = json.Unmarshal([]byte(line), &obj); err != nil {
		return nil, err
	}
	var record = new(Record)
	if l, ok := obj["level"].(string); ok {
		record.Level = l
	}
	if t, ok := obj["time"].(string); ok {
		if record.Time, err = time.Parse(format, t); err != nil {
			return nil, err
		}
	}
	if m, ok := obj["message"].(string); ok {
		record.Message = m
	}
	if t, ok := obj["tag"].(map[string]interface{}); ok {
		for k, v := range t {
			if record.Tag == nil {
				record.Tag = make(Tag)
			}
			if val, ok := v.(string); ok {
				record.Tag[k] = val
			}
		}
	}

	return record, nil
}
