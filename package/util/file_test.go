package util

import "testing"

func TestGetFileDir(t *testing.T) {
	result := GetFileDir("C:/xxx/codes/bdanmu")
	if result != "bdanmu" {
		t.Errorf("GetFileDir() = %v, want %v", result, "bdanmu")
	}
	// Output: GetFileDir() = "bdanmu", want "bdanmu"
	result = GetFileDir("C:/xxx/codes/bdanmu/")
	if result != "bdanmu" {
		t.Errorf("GetFileDir() = %v, want %v", result, "bdanmu")
	}
	// Output: GetFileDir() = "bdanmu", want "bdanmu"
}
