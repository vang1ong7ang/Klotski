package main

import (
	"encoding/json"
	"flag"
	"log"
	"os"
)

func init() {
	address := flag.String("init", "init.json", "...")
	flag.IntVar(&Threads, "threads", 1, "...")
	flag.Parse()
	file, err := os.Open(*address)
	if err != nil {
		log.Fatalln(err)
	}
	decoder := json.NewDecoder(file)
	decoder.Decode(&InitState)
}

var (
	// InitState ...
	InitState = State{}

	// Threads ...
	Threads = 1
)
