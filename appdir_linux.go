// +build !windows,!darwin
package appdir

import (
	"fmt"
	"path/filepath"
	"runtime"
	"strings"
	"sync"
)

var (
	homeDir  string
	dirMutex sync.RWMutex
)

func SetHomeDir(dir string) {
	dirMutex.Lock()
	homeDir = dir
	dirMutex.Unlock()
}

func general(app string) string {
	return InHomeDir(fmt.Sprintf(".%s", strings.ToLower(app)))
}

func logs(app string) string {
	return filepath.Join(general(app), "logs")
}
