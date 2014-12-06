package utils

import (
	"strconv"
	"time"
)

/*得到当日时间的Unix时间*/
func NowToUnix() int64 {
	return time.Now().Unix()
}

/*得到当日时间的Unix时间*/
func DateToUnix() int64 {
	now := time.Now()
	result := time.Date(now.Year(), now.Month(), now.Day(), 0, 0, 0, 0, time.Local).Unix()
	return result
}

/*得到当日时间的Unix时间*/
func SecondTimeStamp() int64 {
	now := time.Now()
	result := time.Date(now.Year(), now.Month(), now.Day(), now.Hour(), now.Minute(), now.Second(), 0, time.Local).Unix()
	return result
}

// get the hour time stamp
func SetHourTimeStamp(month time.Month, day, hour int) int64 {
	now := time.Now()
	result := time.Date(now.Year(), month, day, hour, 0, 0, 0, time.Local).Unix()
	return result
}

/*得到当日时间的Unix时间*/
func HourTimeStamp() int64 {
	now := time.Now()
	result := time.Date(now.Year(), now.Month(), now.Day(), now.Hour(), 0, 0, 0, time.Local).Unix()
	return result

}
func NowDateTime() time.Time {
	now := time.Now()
	result := time.Date(now.Year(), now.Month(), now.Day(), now.Hour(), now.Minute(), now.Second(), now.Nanosecond(), time.Local)
	return result
}

// format time now
func TimeNow() string {
	return time.Now().Format("2006-01-02 15:04:05")
}

/*当日时间的字符型*/
func DateToString() string {
	/*时间戳*/
	now := time.Now()
	timeUnix := time.Date(now.Year(), now.Month(), now.Day(), 0, 0, 0, 0, time.Local).Unix()
	return strconv.FormatInt(timeUnix, 10)
}

// format given time
func UnixFormatS(sec int64) string {
	return time.Unix(sec, 0).Format("2006-01-02 15:04:05")
}
