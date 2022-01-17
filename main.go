package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
)

func main() {

	// get file directory from user
	args := os.Args[1:]

	// check wrong inputs

	if len(args) < 1 {
		fmt.Println("Please provide a directory")
		return
	}

	// check files directory
	files, err := ioutil.ReadDir(args[0])
	if err != nil {
		log.Fatal(err)
		return
	}
	// !!! Check type of files. If you want to reach name of file you need to use methods
	//		like file.Name()
	for _, file := range files {
		if file.Size() == 0 {
			name := file.Name()
			fmt.Println(name)
		}
	}

}
