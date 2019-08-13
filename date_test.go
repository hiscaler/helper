package helper

import (
	"testing"
	"time"
)

var dateFormat string

func init() {
	dateFormat = "2006-01-02 15:04:05"
}

func TestIsValid(t *testing.T) {
	ss := map[string]string{
		"2019-01-01":          "2006-01-02",
		"2019-01-01 01":       "2006-01-02 15",
		"2019-01-01 01:02":    "2006-01-02 15:04",
		"2019-01-01 01:02:03": "2006-01-02 15:04:05",
	}
	for k, v := range ss {
		if !IsValid(k, v) {
			t.Errorf("%s => %s is invalid.", k, v)
		}
	}
}

func TestCompare(t *testing.T) {
	type D struct {
		D1   time.Time
		D2   time.Time
		Diff struct {
			Days    int64
			Seconds float64
		}
	}
	dList := make([]D, 0)
	var d1, d2 time.Time
	d1, _ = time.Parse(dateFormat, "2019-01-01 01:02:03")
	d2, _ = time.Parse(dateFormat, "2019-01-01 01:02:03")
	dList = append(dList, D{
		D1: d1,
		D2: d2,
		Diff: struct {
			Days    int64
			Seconds float64
		}{Days: 0, Seconds: 0},
	})

	d1, _ = time.Parse(dateFormat, "2019-01-01 01:02:03")
	d2, _ = time.Parse(dateFormat, "2019-01-01 01:02:04")
	dList = append(dList, D{
		D1: d1,
		D2: d2,
		Diff: struct {
			Days    int64
			Seconds float64
		}{Days: 0, Seconds: 1},
	})

	d1, _ = time.Parse(dateFormat, "2019-01-01 01:02:03")
	d2, _ = time.Parse(dateFormat, "2019-01-02 01:02:05")
	dList = append(dList, D{
		D1: d1,
		D2: d2,
		Diff: struct {
			Days    int64
			Seconds float64
		}{Days: 1, Seconds: 2},
	})

	for _, d := range dList {
		days, seconds := Compare(d.D1, d.D2)
		if d.Diff.Days != days ||
			d.Diff.Seconds != seconds {
			t.Errorf("Expect: %#v, Actual: days = %d, seconds = %f", d.Diff, days, seconds)
		}
	}
}
