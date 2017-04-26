package main

import (
	"log"
	"os"
)

func startMongod(dbpath string) {
	dir, err := os.Getwd()
	if err != nil {
		log.Println(err)
		os.Exit(1)
	}
	log.Println("CURRENT WORKING PATH:", dir)
}
