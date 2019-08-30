package xos

import "os"


func Remove(path string) {
	if e := os.Remove(path); e != nil {
		panic(e)
	}
}


func RemoveAll(path string) {
	if e := os.RemoveAll(path); e != nil {
		panic(e)
	}
}
