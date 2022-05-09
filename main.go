package main

import (
	"flag"
	"log"
	"os"
	"time"
)

var value string
var expIn time.Duration

func initFlag() {
	flag.StringVar(&value, "value", "", "value of token to be save.")
	flag.DurationVar(&expIn, "expIn", time.Duration(0), "live of token in time duration.")
}

func main() {
	file, err := os.OpenFile("file.txt", os.O_WRONLY, 0755)
	if err != nil {
		log.Fatal("cannot open file: %s", err)
	}
	n, err := file.Write([]byte("test\n"))
	if err != nil {
		log.Fatalf("cannot write to file: %s", err)
	}
	log.Println(n)
}
