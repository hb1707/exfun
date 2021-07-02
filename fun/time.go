package fun

import (
	"time"
	_ "time/tzdata"
)

var TimeLacal = "Asia/Shanghai"

func TimestampStr(dateline int) string {
	return time.Unix(int64(dateline), 0).Format("2006-01-02")
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
