package env

import (
	"os"
	"testing"
)

func TestFileExists(t *testing.T) {
	filePath := "FileExistsTest.txt"

	if fileExists(filePath) {
		os.Remove(filePath)
	}

	if fileExists(filePath) {
		t.Error("Non-existing file exists")
	}

	os.Create(filePath)

	if !fileExists(filePath) {
		t.Error("File DOES NOT exist")
	}

	os.Remove(filePath)
	defer os.Remove(filePath)
}
