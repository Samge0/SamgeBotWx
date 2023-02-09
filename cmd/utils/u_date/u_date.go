package u_date

import (
	"time"
)

// DateFormat 日期格式
var DateFormat = struct {
	Common         string // Common 年/月/日 时:分:秒
	CommonDiagonal string // Common 年/月/日 时:分:秒
	Flow           string // Flow 年月日时分秒
	YMD            string // YMD 年月日
	YMD2           string // YMD2 年-月-日
}{
	Common:         "2006-01-02 15:04:05",
	CommonDiagonal: "2006/01/02 15:04:05",
	Flow:           "20060102150405",
	YMD:            "20060102",
	YMD2:           "2006-01-02",
}

// GetCurrentDateStr 获取当前日期
func GetCurrentDateStr(dateFormats ...interface{}) string {
	return GetDateStr(time.Now(), getDateFormat(dateFormats))
}

// getDateFormat 获取时间格式
func getDateFormat(dateFormats []interface{}) string {
	var dateFormat string
	if len(dateFormats) > 0 {
		dateFormat = dateFormats[0].(string)
	} else {
		dateFormat = DateFormat.Common
	}
	return dateFormat
}

// GetTimeStampStr 获取时间戳的日期字符串格式
func GetTimeStampStr(timeStamp int64, dateFormats ...interface{}) string {
	t := time.Unix(timeStamp, 0)
	return GetDateStr(t, getDateFormat(dateFormats))
}

// GetDateStr 获取日期字符串格式
func GetDateStr(time time.Time, dateFormat string) string {
	return time.Format(dateFormat)
}
