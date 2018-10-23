//2006-01-02 15:04:05
package utils

import "time"

const (
	ORTime = "2006-01-02 15:04:05"
)

//解析当前时区时间
func GetTimeParseString(str string) time.Time {
	to, _ := time.ParseInLocation(ORTime, str, time.Local)
	return to
}

func GetIntTimeString(i64 int64) string {
	return time.Unix(i64,0).Format(ORTime)
}

func GetTimeString(t time.Time) string {
	return t.Format(ORTime)
}

/**
获得指定日期0点时间
*/
func GetTimeDay0Hour(givenTime time.Time) time.Time {
	return time.Date(givenTime.Year(), givenTime.Month(), givenTime.Day(), 0, 0, 0, 0, givenTime.Location())
}

/**
获得指定日期第二天凌晨0点
*/
func GetTimeNextDay0Hour(givenTime time.Time) time.Time {
	next := givenTime.Add(time.Hour * 24)
	return time.Date(next.Year(), next.Month(), next.Day(), 0, 0, 0, 0, next.Location())
}

/**
获得指定日期23:59:59
*/
func GetTimeDayEnd(givenTime time.Time) time.Time {
	return time.Date(givenTime.Year(), givenTime.Month(), givenTime.Day(), 23, 59, 59, 0, givenTime.Location())
}

//取日期显示名称
func GetDayString(givenTime time.Time) string {
	return givenTime.Format("2006-01-02")
}

//取日期显示名称
func GetTodayString() string {
	return time.Now().Format("2006-01-02")
}

/**
获得指定时间的整点
*/
func GetTimeIntHour(givenTime time.Time) time.Time {
	return time.Date(givenTime.Year(), givenTime.Month(), givenTime.Day(), givenTime.Hour(), 0, 0, 0, givenTime.Location())
}

/**
获得指定时间的下一个整点的前一秒 即 59：59
*/
func GetTimeEndOfHour(givenTime time.Time) time.Time {
	next := givenTime.Add(time.Hour)
	nextInt := GetTimeIntHour(next)
	return time.Unix(nextInt.Unix()-1, 0)
}

/**
获得指定时间的下一个整点
*/
func GetTimeNextIntHour(givenTime time.Time) time.Time {
	next := givenTime.Add(time.Hour)
	nextInt := GetTimeIntHour(next)
	return nextInt
}

/**
获得指定时间的整5秒值
*/
func GetTimeInt5Sec(givenTime time.Time) time.Time {
	sec := givenTime.Second()
	return time.Date(givenTime.Year(), givenTime.Month(), givenTime.Day(), givenTime.Hour(), givenTime.Minute(), sec-sec%5, 0, givenTime.Location())
}

/**
获得指定时间的整15秒值  0，15，30，45
*/
func GetTimeInt15Sec(givenTime time.Time) time.Time {
	sec := givenTime.Second()
	return time.Date(givenTime.Year(), givenTime.Month(), givenTime.Day(), givenTime.Hour(), givenTime.Minute(), sec-sec%15, 0, givenTime.Location())
}

/**
获得指定时间的整30秒值  0，30
*/
func GetTimeInt30Sec(givenTime time.Time) time.Time {
	sec := givenTime.Second()
	return time.Date(givenTime.Year(), givenTime.Month(), givenTime.Day(), givenTime.Hour(), givenTime.Minute(), sec-sec%30, 0, givenTime.Location())
}

/**
获得指定时间的整分钟值
*/
func GetTimeIntMin(givenTime time.Time) time.Time {
	return time.Date(givenTime.Year(), givenTime.Month(), givenTime.Day(), givenTime.Hour(), givenTime.Minute(), 0, 0, givenTime.Location())
}

/**
获得指定时间的整10分钟值
*/
func GetTimeInt10Min(givenTime time.Time) time.Time {
	minute := givenTime.Minute()
	return time.Date(givenTime.Year(), givenTime.Month(), givenTime.Day(), givenTime.Hour(), minute-minute%10, 0, 0, givenTime.Location())
}
