package server

import (
	"fmt"
	"log"
	"os"
)

func CreateDirAssets() {
	createDir("/assets")
}

func createDir(path string) {
	dir, err := os.Getwd()
	if err != nil {
		fmt.Println(err)
		return
	}
	dirName := dir + path
	if _, err := os.Stat(dirName); os.IsNotExist(err) {
		err := os.Mkdir(dirName, 0755)
		if err != nil {
			log.Fatal(err)
			return
		}
		err = os.Chmod(dirName, 0755)
		if err != nil {
			log.Fatal(err)
			return
		}
		log.Printf("Directory %s created successfully with restricted permissions.\n", path)
	} else {
		log.Printf("Directory %s already exists.\n", path)
	}
}
