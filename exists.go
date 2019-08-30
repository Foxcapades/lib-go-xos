package xos

import "os"

func Exists(path string) bool {
	_, err := os.Stat(path)

	if err == nil {
		return true
	}

	if os.IsNotExist(err) {
		return false
	}

	panic(err)
}

func FileExists(path string) bool {
	stat, err := os.Stat(path)

	if err == nil {
		return !stat.IsDir()
	}

	if os.IsNotExist(err) {
		return false
	}

	panic(err)
}

func DirExists(path string) bool {
	stat, err := os.Stat(path)

	if err == nil {
		return stat.IsDir()
	}

	if os.IsNotExist(err) {
		return false
	}

	panic(err)
}
