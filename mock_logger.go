package main

import "github.com/heroiclabs/nakama-common/runtime"

type MockLogger struct{}

func (l *MockLogger) WithField(key string, v interface{}) runtime.Logger {
	return nil
}

func (l *MockLogger) WithFields(fields map[string]interface{}) runtime.Logger {

	return nil
}

func (l *MockLogger) Fields() map[string]interface{} {
	return nil
}

func (l *MockLogger) Debug(format string, v ...interface{}) { /* no-op */ }
func (l *MockLogger) Info(format string, v ...interface{})  { /* no-op */ }
func (l *MockLogger) Warn(format string, v ...interface{})  { /* no-op */ }
func (l *MockLogger) Error(format string, v ...interface{}) { /* no-op */ }
