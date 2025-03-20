//go:build !windows && !darwin
// +build !windows,!darwin

package appdir

import (
	"path/filepath"
	"runtime"
	"strings"
	"sync/atomic"
)

var (
	homeDir atomic.Value
)

func SetHomeDir(dir string) {
	homeDir.Store(dir)
}

func general(app string) string {
	if runtime.GOOS == "android" {
		if homeDir.Load() != nil {
			return homeDir.Load().(string)
		} else {
			return InHomeDir(app)
		}

	} else {
		// It is more common on Linux to expect application related directories
		// in all lowercase. The lantern wrapper also expects a lowercased
		// directory.
		return generalAll(strings.ToLower(app))
	}
}

func logs(app string) string {
	return filepath.Join(general(app), "logs")
}
