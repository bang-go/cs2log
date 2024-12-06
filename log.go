package cs2log

import (
	"errors"
	"fmt"
	"github.com/bang-go/cs2log/evt"
	"regexp"
	"time"
)

const (
	LogTypeFile LogType = iota //文件
	LogTypeHttp                //http
)

type LogType int

func ParseLogLine(t LogType, line string) (evt.LogMessage, error) {
	var pattern string
	switch t {
	case LogTypeFile:
		pattern = evt.LogLineValidPattern
	case LogTypeHttp:
		pattern = evt.HttpLineValidPattern
	default:
		return nil, errors.New(fmt.Sprintf("LogType %d not support", t))
	}
	return parseWithPatterns(pattern, line, evt.GetEventRegister())
}

// 通过正则表达式解析日志
func parseWithPatterns(validPattern string, line string, attrArr evt.EventList) (evt.LogMessage, error) {
	// pattern for date, beginning of a log message
	result := regexp.MustCompile(validPattern).FindStringSubmatch(line)
	// if result set is empty, parsing failed, return error
	if result == nil {
		return nil, errors.New("日式格式有误")
	}
	location, _ := time.LoadLocation("Asia/Shanghai") // 使用上海时区
	// 解析日志的时间
	ti, err := time.ParseInLocation("01/02/2006 - 15:04:05.000", result[1], location)
	// if parsing the date failed, return error
	if err != nil {
		return nil, err
	}
	for _, value := range attrArr {
		if res := regexp.MustCompile(value.Pattern).FindStringSubmatch(result[2]); res != nil {
			return value.Func(value.TypeName, ti, res), nil
		}
	}
	return evt.NewUnknown(ti, result[1:]), nil
}
