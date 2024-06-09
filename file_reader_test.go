package main

import (
	"github.com/heroiclabs/nakama-common/runtime"
	"testing"
)

func TestGetFileContent(t *testing.T) {
	oldOsReadFile := osReadFile
	defer func() { osReadFile = oldOsReadFile }()
	osReadFile = func(name string) ([]byte, error) {
		if name == "/nakama/data/configs/test/1.json" {
			return []byte("test"), nil
		}
		return nil, runtime.NewError("File not found", 14)
	}
	logger := &MockLogger{}

	// Test case: correct path
	content, err := getFileContent("test", "1", logger)
	if err != nil {
		t.Fail()
	}
	contentString := string(content)
	if contentString != "test" {
		t.Errorf("Error with content! expected : \"%v\", actual : \"%v\"", contentString, "test1")
	}

	// Test case: incorrect path
	content, err = getFileContent("test1", "1", logger)
	if content != nil {
		t.Fail()
	}
	if err == nil {
		t.Fail()
	}
	s := &err.Message
	if *s != "Cannot read file" {
		t.Errorf("Error message incorrect! expected : \"%v\", actual: \"%v\"", *s, "Cannot read file")
	}
}
