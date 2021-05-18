package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"path/filepath"
)

func main() {
	noConfirm := flag.Bool("no-confim", false ,"do not ask for deletion confirmation")
	threshold := flag.Int64("threshold", 0, "delete folders with size lower than this")
	flag.Parse()


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
			if a <= *threshold {
				fmt.Println(f.Name(), "\t", a)
				deletionlist = append(deletionlist, f.Name())
			}
		}
	}

	if *noConfirm {
		DeleteFolders(deletionlist)
	}else if (len(deletionlist) != 0) {
		fmt.Println("Confirm deletion of the above directories (Y/n): ")
		var second string
		fmt.Scanln(&second)
		if second != "n" {
			DeleteFolders(deletionlist)		
		}
		
	}else {
		fmt.Println("no small folders found")
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