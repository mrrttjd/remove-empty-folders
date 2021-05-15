package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
)

func main() {
	// delete folders beneath this:
	var threshold int64 = 0

	// read directory
	files, err := ioutil.ReadDir(".")
	if err != nil {
		log.Fatal(err)
	}

	// find directories and delete if they are smaller than minSize
	for _, f := range files {
		if f.IsDir() {
			a, _ := DirSize(f.Name())
			fmt.Println(f.Name())
			fmt.Println(a)

			if a <= threshold {
				fmt.Println("Deleting ", f.Name())
				os.RemoveAll(f.Name())
			}
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