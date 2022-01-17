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

	var names []byte

	// !!! Check type of files. If you want to reach name of file you need to use methods
	//		like file.Name()
	for _, file := range files {
		if file.Size() == 0 {
			name := file.Name()
			names = append(names, name...)
			names = append(names, '\n')
		}
	}

	// In this part with writefile you can see your empty file on out.txt
	err = os.WriteFile("out.txt", names, 0644)
	if err != nil {
		fmt.Println(err)
		return
	}

}
