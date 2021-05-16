package main

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
)

func main() {
	// delete folders beneath this:
	var threshold int64 = 0

	// read directory
	files, err := os.ReadDir(".")
	if err != nil {
		log.Fatal(err)
	}

	// find directories and delete if they are smaller than minSize
	for _, f := range files {
		if f.IsDir() {
			// get directory size
			a, err := DirSize(f.Name())
			if err != nil {
				log.Fatal(err)
			}

			fmt.Println(f.Name(), "\t", a)

			// destroy folders beneath threshold
			if a <= threshold {
				fmt.Println("Deleting ", f.Name())
				err := os.RemoveAll(f.Name())
				if err != nil {
					log.Fatal(err)
				}
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
