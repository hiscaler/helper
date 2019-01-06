package helper

import "testing"

func TestIsDir(t *testing.T) {
	if IsDir("D:\\") {
		t.Log("Successful")
	} else {
		t.Error("Failed")
	}
}
