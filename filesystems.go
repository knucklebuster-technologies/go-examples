package main

import (
	"fmt"
	"log"
	"os"
)

func getWorkingDir() {
	dir, err := os.Getwd()
	if err != nil {
		log.Println(err)
		os.Exit(1)
	}

	fmt.Println("CURRENT PATH:", dir)
}

func getAbsPath() {
	dir, err := os.Getwd()
	if err != nil {
		log.Println(err)
		os.Exit(1)
	}

	abs := dir + `\db`
	fmt.Println("ABSOLUTE PATH:", abs)
}

func doesPathExist() {
	dir, err := os.Getwd()
	if err != nil {
		log.Println(err)
		os.Exit(1)
	}

	log.Println("PATH TO TEST:", dir)
	info, err := os.Stat(dir)
	if err != nil {
		log.Println(err)
		os.Exit(1)
	}

	log.Println("DIRECTORY STATUS:", info)

	dir = dir + `\db`
	log.Println("PATH TO TEST:", dir)
	info, err = os.Stat(dir)
	if err != nil {
		log.Println(err)
		os.Exit(1)
	}

	log.Println("DIRECTORY STATUS:", info)
}

func createNewDirectory() {
	dir, err := os.Getwd()
	if err != nil {
		log.Println(err)
		os.Exit(1)
	}

	newDir := dir + `\db`

	info, err := os.Stat(newDir)
	if os.IsNotExist(err) {
		os.Mkdir(newDir, 0777)
	} else {
		log.Println("PATH EXISTS:", info)
	}
}
