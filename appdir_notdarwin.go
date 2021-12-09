// +build !darwin

package appdir

import (
	"fmt"
	"os/user"
	"path/filepath"

	homedir "github.com/mitchellh/go-homedir"
)

func InHomeDir(filename string) string {
	usr, err := user.Current()
	if err == nil {
		return filepath.Join(usr.HomeDir, filename)
	}
	// "user: Current not implemented on ..." will happen on Linux or Darwin
	// when cross-compiled from other platforms.
	homeDir, err2 := homedir.Dir()
	if err2 != nil {
		panic(fmt.Errorf("Unable to determine user's home directory: %s, %s", err, err2))
	}
	return filepath.Join(homeDir, filename)
}
