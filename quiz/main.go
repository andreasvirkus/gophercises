package main

import (
	"flag"
	"fmt"
	"os"
)

func main() {
	csvFilename := flag.String("csv", "questions.csv", "A CSV source file for the questions (problem,answer)")

	file, err := os.Open(*csvFilename)
	fmt.Println("File", file, err)
	if err != nil {
		fmt.Printf("Could not open or parse CSV %s. Please make sure it exists\n", *csvFilename)
		os.Exit(1)
	}
	_ = file
}
