package main

import (
	"fmt"
	"github.com/jessevdk/go-flags"
	"log"
	"os"
	"path/filepath"
)

var opts struct {
	Verbose bool `short:"v" long:"verbose" description:"Show verbose debug information"`
	NoConfirm bool `short:"N" long:"no-confirm" description:"Do not ask for deletion confirmation"`
	Threshold int64 `short:"T" long:"threshold" description:"Delete folders lower than a threshold in bytes"`
}


func main() {
	flags.Parse(&opts)

	
	// read directory
	files, err := os.ReadDir(".")
	if err != nil {
		log.Fatal(err)
	}

	var deletionlist []string
	// find directories and delete if they are smaller than minSize
	for _, f := range files {
		if f.IsDir() {
			// get directory size
			a, err := DirSize(f.Name())
			if err != nil {
				log.Fatal(err)
			}

			
			// create slice of folders to delete
			if a <= opts.Threshold {
				fmt.Println(f.Name(), "\t", a)
				deletionlist = append(deletionlist, f.Name())
			}
		}
	}

	if opts.NoConfirm {
		DeleteFolders(deletionlist)
	} else if (len(deletionlist) != 0) {
		fmt.Println("Confirm deletion of the above directories (Y/n): ")
		var second string
		fmt.Scanln(&second)
		if second != "n" {
			DeleteFolders(deletionlist)		
		}
		
	} else {
		fmt.Println("No small folders found")
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


func DeleteFolders(paths []string) {
	for _, f := range paths {
		fmt.Println("Deleting:\t", f)
		err := os.RemoveAll(f)
		if err != nil {
			log.Fatal(err)
		}
	}
}