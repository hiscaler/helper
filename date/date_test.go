package date

import (
	"testing"
)

func TestDay(t *testing.T) {
	b, e, err := Day("2019-01-01", "2006-01-02")
	if b.Unix() != 1546272000 || e.Unix() != 1546358399 || err != nil {
		t.Errorf("%s - %s, err: %v", b, e, err)
	}
}
