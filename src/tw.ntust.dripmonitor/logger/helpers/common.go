package helpers

import (
	"runtime"
	"path"
)

var ProjectPath string

func InitializePaths() {
	_, filename, _, ok := runtime.Caller(0)
	if !ok {
		panic("InitializeCommon: No caller information")
	} else {
		ProjectPath = path.Join(path.Dir(filename), "../")
	}
}
