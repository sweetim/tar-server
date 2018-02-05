package main

import (
	"io/ioutil"
	"os"
	"path/filepath"
)

type DirInfo struct {
	Path string `json:"path"`
	Size int64  `json:"size"`
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
				Size: getDirSize(dirPath.Name()),
			})
		}

		ch <- info
	}()

	return ch, errCh
}

func getDirSize(path string) int64 {
	var size int64
	_ = filepath.Walk(path, func(_ string, info os.FileInfo, err error) error {
		if !info.IsDir() {
			size += info.Size()
		}

		return err
	})

	return size
}

func GetEnv(key string, v interface{}) interface{} {
	if value, ok := os.LookupEnv(key); ok {
		return value
	}

	return v
}
