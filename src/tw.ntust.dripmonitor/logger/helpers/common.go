package helpers

import (
	"runtime"
	"path"
)

const TimeFormat = "2006-01-02 15:04:05"

var ProjectPath string

func InitializePaths() {
	_, filename, _, ok := runtime.Caller(0)
	if !ok {
		panic("InitializeCommon: No caller information")
	} else {
		ProjectPath = path.Join(path.Dir(filename), "../")
	}
}
