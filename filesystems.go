package main

import (
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
)

func getWorkingDir() (string, error) {
	dir, err := os.Getwd()
	if err != nil {
		log.Println("ERROR:", err)
		return "", err
	}
	log.Println("WORKING DIRECTORY:", dir)
	return dir, nil
}

func getAbsPath(relative string) (string, error) {
	dir, err := filepath.Abs(relative)
	if err != nil {
		log.Println("ERROR:", err)
		return "", err
	}
	log.Println("ABSOLUTE PATH:", dir)
	return dir, err
}

func doesPathExist(path string) bool {
	_, err := os.Stat(path)
	if os.IsNotExist(err) {
		log.Println("PATH NOT EXISTS:", path)
		return false
	}
	log.Println("PATH EXISTS:", path)
	return true
}

func createNewDirectory(path string) bool {
	exists := doesPathExist(path)
	if exists != true {
		os.Mkdir(path, 0777)
		log.Println("DIRECTORY CREATED")
		return true
	}
	log.Println("DIRECTORY EXISTS")
	return false
}

func writeFile(name string, data []byte, mode os.FileMode) error {
	err := ioutil.WriteFile(name, data, mode)
	if err != nil {
		log.Println(err)
		os.Exit(1)
	}

	return nil
}
