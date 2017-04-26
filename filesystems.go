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
