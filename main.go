package main

import (
	"bufio"
	"encoding/csv"
	"flag"
	"fmt"
	"log"
	"os"
	"strings"
)

type args struct {
	Csv string
}

type question struct {
	Question string
	Answer   string
}

func main() {
	args := parseArgs()
	rows := readCsvFile(args.Csv)
	questions := parseQuestions(rows)
	startQuiz(questions)
}

func parseArgs() args {
	var args args
	flag.StringVar(&args.Csv, "csv", "problems.csv", "a csv file in the format of 'question,answer'")
	flag.Parse()
	return args
}

func readCsvFile(f string) [][]string {
	file, err := os.Open(f)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	csvReader := csv.NewReader(file)
	rows, err := csvReader.ReadAll()
	if err != nil {
		log.Fatal(err)
	}
	return rows
}

func parseQuestions(rows [][]string) []question {
	questions := make([]question, len(rows))
	for i, row := range rows {
		questions[i] = question{Question: row[0], Answer: row[1]}
	}
	return questions
}

func startQuiz(questions []question) {
	waitForKeyPress()

	score := 0
	for _, question := range questions {
		answer := answerQuestion(question)
		answer = sanitize(answer)
		if answer != question.Answer {
			continue
		}
		score++
		fmt.Println()
	}

	fmt.Printf("You answered %d out of %d questions correctly!\n", score, len(questions))
}

func waitForKeyPress() {
	fmt.Println("Press 'Enter' to start the quiz...")
	bufio.NewReader(os.Stdin).ReadBytes('\n')
}

func answerQuestion(q question) string {
	fmt.Printf("%s?\n", q.Question)
	var answer string
	fmt.Scan(&answer)
	return answer
}

func sanitize(answer string) string {
	sanitized := strings.TrimSpace(answer)
	return sanitized
}
