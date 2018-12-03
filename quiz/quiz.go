package main

import (
	"bufio"
	"encoding/csv"
	"flag"
	"fmt"
	"io"
	"log"
	"os"
)

func main() {

	wordPtr := flag.String("filename", "problems.csv", "the quiz to run")

	csvFile, _ := os.Open(*wordPtr)
	reader := csv.NewReader(bufio.NewReader(csvFile))
	var totalquestions = 0
	var correct = 0
	var answer = ""
	for {
		record, err := reader.Read()

		if err == io.EOF {
			break
		}

		if err != nil {
			log.Fatal(err)
		}
		totalquestions++
		fmt.Printf(record[0] + " = ")
		fmt.Scanln(&answer)

		if answer == record[1] {
			correct++
		}
	}
	fmt.Println(totalquestions)
	fmt.Println(correct)
}
