package main

import (
	"bufio"
	"encoding/csv"
	"flag"
	"fmt"
	"log"
	"os"
	"strings"
	"time"
)

type problem struct {
	qustion string
	answer  string
}

func parseProblems(lines [][]string) []problem {
	ret := make([]problem, len(lines))

	for i, line := range lines {
		ret[i] = problem{
			qustion: strings.TrimSpace(line[0]),
			answer:  strings.TrimSpace(line[1]),
		}
	}
	return ret
}

func main() {
	fileName := flag.String("csv", "quizGame/problems.csv", "csv file name")
	timeLimit := flag.Duration("limit", 20, "quiz time")
	flag.Parse()

	f, err := os.Open(*fileName)
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()
	csvReader := csv.NewReader(f)
	lines, err := csvReader.ReadAll()
	if err != nil {
		log.Fatal(err)
	}

	problems := parseProblems(lines)

	timer := time.After(*timeLimit * time.Second)

	correct := 0
	answerCh := make(chan string)
	reader := bufio.NewReader(os.Stdin)
	for i, p := range problems {
		fmt.Printf("\nProblem # %d %s ? : ", i+1, p.qustion)
		go func() {
			ans, _ := reader.ReadString('\n')
			ans = strings.TrimSpace(ans)
			answerCh <- ans
		}()
		select {
		case ans := <-answerCh:
			if ans == p.answer {
				correct++
			}
		case <-timer:
			fmt.Printf("\nYou scored %d out of %d\n", correct, len(problems))
			return
		}
	}
	fmt.Printf("\nYou scored %d out of %d\n", correct, len(problems))
}
