package main

import (
	"log"
	"os"
	"os/exec"
)

func startMongod() {
	cmd := exec.Command("mongod")
	err := cmd.Start()
	if err != nil {
		log.Println(err)
		os.Exit(1)
	}

	log.Println("Mongo DB PID:", cmd.Process.Pid)
}
