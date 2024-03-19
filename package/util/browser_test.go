package util

import "testing"

func TestOpenUrlOnBrowser(t *testing.T) {
	err := OpenUrlOnBrowser("https://space.bilibili.com/63231")
	if err != nil {
		t.Error(err)
	}
}
