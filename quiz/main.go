package main

import (
	"encoding/csv"
	"math/rand"
	"strings"
	"time"
	"flag"
	"fmt"
	"os"
)

func main() {
	csvFilename := flag.String("csv", "questions.csv", "A CSV source file for the questions (problem,answer)")
	shuffleQuestions := flag.Bool("shuffle", false, "Ask the questions in a random order")
	flag.Parse()
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
	questions := parseLines(lines)

	if *shuffleQuestions {
		rand.Seed(time.Now().UnixNano())
		rand.Shuffle(len(questions), func(i, j int) {
			questions[i], questions[j] = questions[j], questions[i]
		})
	}

	correctAnswersCount := 0
	for i, q := range questions {
		fmt.Printf("Question #%d: %s?\n", i+1, q.problem)
		var answer string
		fmt.Scanf("%s\n", &answer)
		if strings.ToLower(strings.TrimSpace(answer)) == strings.ToLower(strings.TrimSpace(q.answer)) {
			fmt.Println("Correct!")
			correctAnswersCount++
		}
	}
	fmt.Printf("You answered %d/%d questions correctly\n", correctAnswersCount, len(questions))
}

func parseLines(lines [][]string) []question {
	formattedLines := make([]question, len(lines))
	for i, line := range lines {
		formattedLines[i] = question{
			problem: line[0],
			answer:  line[1],
		}
	}
	return formattedLines
}

type question struct {
	problem string
	answer  string
}

func exit(msg string) {
	fmt.Println(msg)
	os.Exit(1)
}
