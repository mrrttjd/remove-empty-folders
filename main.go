package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
)

/*
PSEUDO
1. Walk current directory
2. If directory found, check size
3. If size is less than x, delete
*/

func main() {
	//pa, err := os.Getwd()
	files, err := ioutil.ReadDir(".")
	if err != nil {
		log.Fatal(err)
	}
	for _, f := range files {
		if f.IsDir() {
			a, _ := DirSize(f.Name())
			fmt.Println(f.Name())
			fmt.Println(a)
		}
	}
}

func DirSize(path string) (int64, error) {
	var size int64
	err := filepath.Walk(path, func(_ string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if !info.IsDir() {
			size += info.Size()
		}
		return err
	})
	return size, err
}