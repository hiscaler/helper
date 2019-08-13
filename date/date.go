package date

import (
	"reflect"
	"strconv"
	"time"
)

var loc *time.Location

func init() {
	loc, _ = time.LoadLocation("Asia/Chongqing")
}

// 判断是否为有效的日期字符串
func IsValid(datetime string, dateFormat string) bool {
	if len(dateFormat) == 0 {
		dateFormat = "2006-01-02 15:04:05"
	}
	if _, err := time.Parse(dateFormat, datetime); err == nil {
		return true
	} else {
		return false
	}
}

func DayRange(v interface{}, dateFormat string) (beginTime, endTime time.Time, err error) {
	r := reflect.ValueOf(v)
	switch r.Kind() {
	case reflect.Int:
		beginTime = time.Unix(int64(v.(int)), 0)
	case reflect.Int32:
		beginTime = time.Unix(int64(v.(int32)), 0)
	case reflect.Int64:
		beginTime = time.Unix(v.(int64), 0)
	default:
		// String or time.Time
		v = r.String()
		beginTime, err = time.Parse(dateFormat, r.String())
		if err == nil {
			beginTime = time.Date(beginTime.Year(), beginTime.Month(), beginTime.Day(), 0, 0, 0, 0, loc)
		}
	}
	endTime = beginTime.Add(86399 * time.Second)
	return
}

func TodayRange() (beginTime, endTime time.Time) {
	now := time.Now()
	beginTime = time.Date(now.Year(), now.Month(), now.Day(), 0, 0, 0, 0, loc)
	endTime = beginTime.Add(86399 * time.Second)

	return
}

func MonthRange(t time.Time) (beginTime, endTime time.Time) {
	beginTime = time.Date(t.Year(), t.Month(), 1, 0, 0, 0, 0, loc)
	endTime = beginTime.AddDate(0, 1, 0)
	return
}

func YearMonth(t time.Time) (beginTime, endTime time.Time) {
	beginTime = time.Date(t.Year(), 1, 1, 0, 0, 0, 0, loc)
	endTime = beginTime.AddDate(1, 0, 0)
	return
}

func YearWeekRange(year, week int) (beginTime, endTime time.Time) {
	d1 := time.Date(year, 1, 1, 0, 0, 0, 0, loc)
	w := int(d1.Weekday())
	if _, yw := d1.ISOWeek(); yw == 1 {
		d1 = d1.AddDate(0, 0, -(w - 1))
	}
	d2 := d1.AddDate(0, 0, 6)
	if week == 1 {
		beginTime = d1
		endTime = d2.Add(86399 * time.Second)
	} else {
		week -= 2
		beginTime = d2.AddDate(0, 0, week*7+1)
		endTime = beginTime.AddDate(0, 0, 6)
		y, _ := strconv.ParseInt(endTime.Format("2006"), 10, 64)
		if y == int64(year) {
			endTime = endTime.Add(86399 * time.Second)
		} else {
			endTime = time.Date(year, 12, 31, 23, 59, 59, 0, loc)
		}
	}

	return
}

// 日期比较
func Compare(t1, t2 time.Time) (days int64, seconds float64) {
	v := int64(t2.Sub(t1).Seconds())
	return v / 86400, float64(v % 86400)
}
