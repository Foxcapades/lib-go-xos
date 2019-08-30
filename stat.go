package xos

import "os"

func Stat(path string) FileInfo {
	stat, err := os.Stat(path)

	if err != nil {
		panic(err)
	}

	return stat
}


func Lstat(path string) FileInfo {
	stat, err := os.Lstat(path)

	if err != nil {
		panic(err)
	}

	return stat
}
