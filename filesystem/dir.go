package filesystem

import (
	"runtime"
	"os"
)

func HomeDir() string {
	if runtime.GOOS == "windows" {
		home := os.Getenv("HOMEDRIVE") + os.Getenv("HOMEPATH")
		if home == "" {
			home = os.Getenv("USERPROFILE")
		}
		return home
	}
	return os.Getenv("HOME")
}

func IsDir(dir string) bool {
	fi, err := os.Stat(dir)
	return err == nil && fi.IsDir()
}


