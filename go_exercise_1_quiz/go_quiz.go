package main

import (
	"encoding/csv"
	"flag"
	"fmt"
	"log"
	"os"
	"time"
)

type question struct {
	question string
	answer   string
}

type quiz struct {
	questions []question
	score     int
}

func readCSV(filename string) ([]question, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, fmt.Errorf("could not open file %s: %v", filename, err)
	}
	defer file.Close()

	reader := csv.NewReader(file)
	records, err := reader.ReadAll()
	if err != nil {
		return nil, fmt.Errorf("could not read CSV: %v", err)
	}

	var questions []question
	for _, record := range records {
		if len(record) != 2 {
			return nil, fmt.Errorf("invalid record: %v", record)
		}
		q := question{question: record[0], answer: record[1]}
		questions = append(questions, q)
	}
	return questions, nil
}

func (q *quiz) askQuestions() {
	q.score = 0
	for _, qst := range q.questions {
		fmt.Printf("%s = ", qst.question)
		var answer string
		fmt.Scanln(&answer)
		if answer == qst.answer {
			fmt.Println("Correct!")
			q.score++
		} else {
			fmt.Printf("Wrong! The correct answer is %s.\n", qst.answer)
		}
	}
	fmt.Printf("You scored %d out of %d.\n", q.score, len(q.questions))
}

var (
	csvFile   string
	timeLimit int
)

func init() {
	flag.StringVar(&csvFile, "csv", "problems.csv", "a CSV file in the format 'question,answer'")
	flag.IntVar(&timeLimit, "timelimit", 30, "the time limit for the quiz in seconds")
}

func main() {
	flag.Parse()

	questions, err := readCSV(csvFile)
	if err != nil {
		log.Fatalf("Error reading CSV file: %v", err)
	}
	q := &quiz{questions: questions, score: 0}
	fmt.Printf("Welcome to the quiz! You have %d seconds to answer %d questions.\n", timeLimit, len(questions))
	timer := time.After(time.Duration(timeLimit) * time.Second) // timer in seconds
	result := make(chan quiz)
	go func() {
		q.askQuestions()
		result <- *q
	}()
	select {
	case <-result:
		fmt.Println("Quiz completed!")
	case <-timer:
		fmt.Println("Time's up! You didn't finish the quiz in time.")
	}
	fmt.Printf("Final score: %d out of %d.\n", q.score, len(q.questions))
}
