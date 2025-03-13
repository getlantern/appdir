package appdir

import (
	"path/filepath"
)

func SetHomeDir(dir string) {
	// do nothing
}

func general(app string) string {
	return generalAll(app)
}

func logs(app string) string {
	return InHomeDir(filepath.Join("Library/Logs", app))
}
