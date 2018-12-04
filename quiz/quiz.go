package main

import (
	"bufio"
	"encoding/csv"
	"flag"
	"fmt"
	"io"
	"log"
	"os"
	"time"
)

func main() {

	wordPtr := flag.String("filename", "problems.csv", "the quiz to run")

	csvFile, _ := os.Open(*wordPtr)

	reader := csv.NewReader(bufio.NewReader(csvFile))

	var totalquestions int
	var correct int
	var answer string
	var answerfail = false
	for !answerfail {
		record, err := reader.Read()

		if err == io.EOF {
			break
		}

		if err != nil {
			log.Fatal(err)
		}
		totalquestions++
		fmt.Printf(record[0] + " = ")

		timeout := make(chan bool, 1)
		go func() {
			time.Sleep(30 * time.Second)
			timeout <- true
		}()
		s := make(chan bool, 1)
		go func() {

			fmt.Scanln(&answer)
			s <- true
		}()
		select {
		case <-timeout:
			answerfail = true // user didnt input data in time
			fmt.Printf("\n\n You did not answer in time, the quiz is over\n")
		case <-s:
			if answer == record[1] {
				correct++
			}
		}
	}
	fmt.Println(totalquestions)
	fmt.Println(correct)
}
