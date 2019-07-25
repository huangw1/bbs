/**
 * @Author: huangw1
 * @Date: 2019/7/25 18:15
 */

package extension

import (
	"fmt"
	"strconv"
	"time"
)

const (
	FmtDateTime = "2006-01-02 15:04:05"

	FmtDate = "2006-01-02"

	FmtTime = "15:04:05"

	FmtDateTimeCn = "2006年01月02日 15时04分05秒"

	FmtDateCn = "2006年01月02日"

	FmtTimeCn = "15时04分05秒"
)

// 秒时间戳
func NowUnix() int64 {
	return time.Now().Unix()
}

// 毫秒时间戳
func NowTimestamp() int64 {
	return time.Now().UnixNano() / 1e6
}

// 毫秒时间戳
func Timestamp(t time.Time) int64 {
	return t.UnixNano() / 1e6
}

// 秒时间戳转时间
func TimeFromUnix(unix int64) time.Time {
	return time.Unix(unix, 0)
}

// 毫秒时间戳转时间
func TimeFromTimestamp(timestamp int64) time.Time {
	return time.Unix(0, timestamp*int64(time.Millisecond))
}

func TimeFormat(time time.Time, layout string) string {
	return time.Format(layout)
}

func TimeParse(timeStr string, layout string) (time.Time, error) {
	return time.Parse(layout, timeStr)
}

func GetDay(time time.Time) int {
	yyyyMMdd, _ := strconv.Atoi(time.Format("20060102"))
	return yyyyMMdd
}

// 返回指定时间当天的开始时间
func TimeOfStartDay(t time.Time) time.Time {
	return time.Date(t.Year(), t.Month(), t.Day(), 0, 0, 0, 0, t.Location())
}

func PrettyTime(timestamp int64) string {
	t := TimeFromTimestamp(timestamp)
	duration := (NowTimestamp() - timestamp) / 1000
	if duration < 60 {
		return "刚刚"
	} else if duration < 60*60 {
		return fmt.Sprintf("%d分钟前", duration/60)
	} else if duration < 60*60*24 {
		return fmt.Sprintf("%d小时前", duration/(60*60))
	} else if timestamp >= Timestamp(TimeOfStartDay(time.Now().Add(-time.Hour*24))) {
		return fmt.Sprintf("昨天%s", TimeFormat(t, FmtTime))
	} else if timestamp >= Timestamp(TimeOfStartDay(time.Now().Add(-time.Hour*24*2))) {
		return fmt.Sprintf("前天%s", TimeFormat(t, FmtTime))
	} else {
		return TimeFormat(t, FmtDate)
	}
}
