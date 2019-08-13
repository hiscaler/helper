package helper

import (
	"testing"
)

func TestIsDir(t *testing.T) {
	dirs := []string{
		"C:\\",
		"D:\\",
	}
	files := []string{
		"C:\\a.txt",
		"D:\\b.txt",
	}
	for _, dir := range dirs {
		if !IsDir(dir) {
			t.Errorf("`%s` is dir.", dir)
		}
	}
	for _, file := range files {
		if IsDir(file) {
			t.Errorf("`%s` is not dir.", file)
		}
	}
}
