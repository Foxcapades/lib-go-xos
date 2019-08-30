package xos

import "os"

type FileMode = os.FileMode
type FileInfo = os.FileInfo
type File = os.File

func Chdir(path string) {
	if err := os.Chdir(path); err != nil {
		panic(err)
	}
}

func Link(old, new string) {
	if e := os.Link(old, new); e != nil {
		panic(e)
	}
}

func Chown(path string, uid, gid int) {
	if e := os.Chown(path, uid, gid); e != nil {
		panic(e)
	}
}

func Chmod(path string, mode FileMode) {
	if e := os.Chmod(path, mode); e != nil {
		panic(e)
	}
}

func Open(path string) *File {
	o, e := os.Open(path)

	if e != nil {
		panic(e)
	}

	return o
}

func Rename(old, new string) {
	if e := os.Rename(old, new); e != nil {
		panic(e)
	}
}

func Getwd() string {
	o, e := os.Getwd()

	if e != nil {
		panic(e)
	}

	return o
}
