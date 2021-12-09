package appdir

import (
	"fmt"
	"os"
	"path/filepath"
)

func SetHomeDir(dir string) {
	// do nothing
}

// TODO: check to see where container manifest will move contents of 'Application Support' and 'Logs'
// (e.g. Library/Logs/log.log or Library/Logs/Lantern/log.log?)

func general(app string) string {
	return InHomeDir(filepath.Join("Library/Application Support", app))
}

func logs(app string) string {
	return InHomeDir(filepath.Join("Library/Logs", app))
}

func InHomeDir(filename string) string {
	// It is important that we use os.UserHomeDir on macOS. The Lantern application is sandboxed on
	// macOS and UserHomeDir will provide the path to our container directory, to which we have read
	// and write access. Attempting to access files in the user's actual home directory will fail.
	homeDir, err := os.UserHomeDir()
	if err != nil {
		panic(fmt.Sprintf("unable to determine home directory: %v", err))
	}
	return filepath.Join(homeDir, filename)
}
