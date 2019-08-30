package xos

import "os"

func Mkdir(path string, perm FileMode) {
	if e := os.Mkdir(path, perm); e != nil {
		panic(e)
	}
}

func MkdirAll(path string, perm FileMode) {
	if e := os.MkdirAll(path, perm); e != nil {
		panic(e)
	}
}
