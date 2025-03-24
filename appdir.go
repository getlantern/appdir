// package appdir provides a facility for determining the system-dependent
// paths for application resources.
package appdir

import (
	"log"
	"os"
	"path/filepath"
)

// General returns the path for general aplication resources (e.g.
// ~/Library/Application Support/<App>).
func General(app string) string {
	return general(app)
}

// Logs returns the path for log files (e.g. ~/Library/Logs/<App>).
func Logs(app string) string {
	return logs(app)
}

func inHomeDir(subdir string) string {
	homeDir, err := os.UserHomeDir()
	if err != nil {
		log.Printf("Unable to determine user's home directory: %s", err)
		return subdir
	}
	return mkdir(homeDir, subdir)
}

func generalAll(app string) string {
	dir, err := os.UserConfigDir()
	if err != nil {
		log.Printf("Unable to determine user's config directory: %s", err)
		return app
	}
	return mkdir(dir, app)
}

func mkdir(baseDir, subdir string) string {
	path := filepath.Join(baseDir, subdir)
	// Create the directory if it does not exist
	if _, err := os.Stat(path); os.IsNotExist(err) {
		os.MkdirAll(path, 0755)
	}
	return path
}
