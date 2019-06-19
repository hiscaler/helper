package date

import "time"

type Range struct {
	Begin int64
	End   int64
}

// Get special day begin and end timestamp
func Day(date string, format string) Range {
	r := Range{}
	if len(format) <= 0 {
		format = "2006-01-02"
	}
	if begin, err := time.Parse(format, date); err == nil {
		if end, err := time.Parse(format, begin.Format(format)+" 23:59:59"); err == nil {
			r.Begin = begin.Unix()
			r.End = end.Unix()
		}
	}

	return r
}

// Today begin and end timestamp
func Today(format string) Range {
	return Day(time.Now().Format(format), format)
}
