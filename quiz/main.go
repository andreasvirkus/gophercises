package main

import (
	"encoding/csv"
	"flag"
	"fmt"
	"os"
)

func main() {
	csvFilename := flag.String("csv", "questions.csv", "A CSV source file for the questions (problem,answer)")

	file, err := os.Open(*csvFilename)
	if err != nil {
		exit(fmt.Sprintf("Could not find CSV %s. Please make sure it exists\n", *csvFilename))
	}
	r := csv.NewReader(file)
	lines, err := r.ReadAll()
	if err != nil {
		exit("Could not parse the CSV")
	}
	fmt.Println("Reading questions from", *csvFilename)
	fmt.Println(lines)
}

func exit(msg string) {
	fmt.Println(msg)
	os.Exit(1)
}
