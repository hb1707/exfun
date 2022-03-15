package fun

import (
	"time"
	_ "time/tzdata"
)

var TimeLacal = "Asia/Shanghai"

func TimestampStr(dateline int, format ...string) string {
	var formatStr = "2006-01-02"
	if len(format) > 0 {
		formatStr = format[0]
	}
	return time.Unix(int64(dateline), 0).Format(formatStr)
}

func StrTimestamp(stringTime string) int {
	var unixTime int
	loc, err := time.LoadLocation(TimeLacal)
	if err != nil {
		return 0
	}
	theTime, err := time.ParseInLocation("2006-01-02 15:04:05", stringTime, loc)
	if err == nil {
		unixTime = int(theTime.Unix())
	}
	return unixTime
}
func StrTime(stringTime string, layout ...string) (time.Time, error) {
	var theTime time.Time
	loc, err := time.LoadLocation(TimeLacal)
	if err != nil {
		return time.Time{}, err
	}
	layoutStr := "2006-01-02 15:04:05"
	if len(layout) > 0 {
		layoutStr = layout[0]
	}
	theTime, err = time.ParseInLocation(layoutStr, stringTime, loc)
	if err != nil {
		return time.Time{}, err
	}
	return theTime, nil
}
func LocalTimeStr(rowValue string, utc string) string {
	sqlLoc, _ := time.LoadLocation(utc)
	dt, _ := time.ParseInLocation(time.RFC3339Nano, rowValue, sqlLoc)
	loc, _ := time.LoadLocation(TimeLacal)
	rowValue = dt.In(loc).Format("2006-01-02 15:04:05")
	return rowValue
}
func LocalTime(dt time.Time) string {
	loc, _ := time.LoadLocation(TimeLacal)
	rowValue := dt.In(loc).Format("2006-01-02 15:04:05")
	return rowValue
}

func Timestamp(dt time.Time) int {
	return StrTimestamp(LocalTime(dt))
}

func TimeSub(t1, t2 time.Time) (int, int) {
	if t1.Location().String() != t2.Location().String() {
		return -1, -1
	}
	hours := t1.Sub(t2).Hours()

	if hours <= 0 {
		return -1, -1
	}
	if hours < 24 {
		t1y, t1m, t1d := t1.Date()
		t2y, t2m, t2d := t2.Date()
		isSameDay := t1y == t2y && t1m == t2m && t1d == t2d

		if isSameDay {
			return 0, int(hours)
		} else {
			return 1, 0
		}

	} else {
		if (hours/24)-float64(int(hours/24)) == 0 { // just 24's times
			return int(hours / 24), 0
		} else { // more than 24 hours
			return int(hours/24) + 1, 0
		}
	}

}

//时间转数字，周日=7
func DateToWeek(t time.Time) int {
	var week = int(t.Weekday())
	if week == 0 {
		return 7
	} else {
		return week
	}
}
