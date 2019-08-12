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

func IsValid(date string) bool {
	return true
}

func Day(v interface{}, dateFormat string) (beginTime, endTime time.Time, err error) {
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
		beginTime, err = time.Parse(dateFormat, r.String());
		if err == nil {
			beginTime = time.Date(beginTime.Year(), beginTime.Month(), beginTime.Day(), 0, 0, 0, 0, loc)
		}
	}
	endTime = beginTime.Add(86399 * time.Second)
	return
}

func Today() (beginTime, endTime time.Time) {
	now := time.Now()
	beginTime = time.Date(now.Year(), now.Month(), now.Day(), 0, 0, 0, 0, loc)
	endTime = beginTime.Add(86399 * time.Second)

	return
}

func YearWeek(year, week int) (beginTime, endTime time.Time) {
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
