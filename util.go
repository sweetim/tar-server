package main

import (
	"io/ioutil"
	"os"
)

type DirInfo struct {
	Path string `json:"path"`
	Size int    `json:"size"`
}

func GetDir(p string) (<-chan []DirInfo, <-chan error) {
	ch := make(chan []DirInfo)
	errCh := make(chan error)

	go func() {
		fInfo, err := ioutil.ReadDir(p)

		if err != nil {
			errCh <- err
			return
		}

		var info []DirInfo
		for _, dirPath := range fInfo {
			info = append(info, DirInfo{
				Path: dirPath.Name(),
				Size: int(dirPath.Size()),
			})
		}

		ch <- info
	}()

	return ch, errCh
}

func GetEnv(key string, v interface{}) interface{} {
	if value, ok := os.LookupEnv(key); ok {
		return value
	}

	return v
}
