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

	// 2 Make quiz and timer to gorountine and if timer stops break programm and return right and wrong questions
	// write goroutine for timelimit and after 30s write it into channel and after each question ask if channel is not empty


	timeLimit := flag.Int("time", 2, "Enter timelimit")
	filename := flag.String("file", "problems", "Filename of the file containing questions and answers")
	flag.Parse()

	timer := time.NewTimer(time.Duration(*timeLimit)*time.Second)

	// Open the file
	csvfile, err := os.Open(*filename + ".csv")
	if err != nil {
		log.Fatalln("Couldn't open the csv file", err)
	}

	// Parse the file
	r := csv.NewReader(csvfile)
	//r := csv.NewReader(bufio.NewReader(csvfile))

	correctAnswers := 0
	wrongAnwers := 0
	lineCount := 0

	// Iterate through the records line by line
	problem: for {

		select {
		case <-timer.C: 
			break problem 
		default: 
		
		// Read each record from csv
		record, err := r.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatal(err)
		}
		// Ask question
		fmt.Printf("%s = ", record[0])
		// Wait for userinput
		reader := bufio.NewReader(os.Stdin)
		input, _ := reader.ReadString('\n')

		lineCount++

		if input == (record[1] + "\n") {
			correctAnswers++
		} else {
			wrongAnwers++
		}

		}
		
		

	}

	fmt.Printf("You got %v out of %v questions right\n", correctAnswers, lineCount)
}
