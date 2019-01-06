package helper

import "testing"

func TestToSemiangle(t *testing.T) {
	s1 := "ＡＢＣＤＥＦＧＨＩＪＫＬＭＮＯＰＱＲＳＴＵＶＷＸＹＺ"
	s2 := "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
	if s := ToSemiangle(s1, false, true, false, false); s == s2 {
		t.Log("Successful")
	} else {
		t.Error("Failed")
	}

	n1 := "０１２３４５６７８９"
	n2 := "0123456789"
	if n := ToSemiangle(n1, true, false, false, false); n == n2 {
		t.Log("Successful")
	} else {
		t.Error("Failed", n)
	}
}
