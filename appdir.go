// package appdir provides a facility for determining the system-dependent
// paths for application resources.
package appdir

import (
	"log"
	"os"
	"path/filepath"
)

// General returns the path for general aplication resources (e.g.
// ~/Library/<App>).
func General(app string) string {
	return general(app)
}

// Logs returns the path for log files (e.g. ~/Library/Logs/<App>).
func Logs(app string) string {
	return logs(app)
}

func InHomeDir(filename string) string {
	homeDir, err := os.UserHomeDir()
	if err != nil {
		log.Printf("Unable to determine user's home directory: %s", err)
		return filename
	}
	return filepath.Join(homeDir, filename)
}

func generalAll(app string) string {
	dir, err := os.UserConfigDir()
	if err != nil {
		log.Printf("Unable to determine user's config directory: %s", err)
		return app
	}
	return filepath.Join(dir, app)
}
