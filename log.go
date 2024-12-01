package cs2log

import (
	"errors"
	"github.com/bang-go/cs2log/evt"
	"regexp"
	"time"
)

// ParseLogLine 根据原始行log解析日志 log文件
func ParseLogLine(line string) (evt.LogMessage, error) {
	return parseWithPatterns(evt.LogLineValidPattern, line, evt.GetEventRegister())
}

// ParseHttpLogLine 根据原始行log解析日志 http log文件
func ParseHttpLogLine(line string) (evt.LogMessage, error) {
	return parseWithPatterns(evt.HttpLineValidPattern, line, evt.GetEventRegister())
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
	ti, err := time.ParseInLocation("01/02/2006 - 15:04:05", result[1], location)
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
