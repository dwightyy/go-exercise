package main

import (
	"bufio"
	"encoding/csv"
	"fmt"
	"log"
	"os"
	"strings"
)

func main() {

	printPresentation()
	questions := getQuestions()

	totalScore := playRound(questions)

	resultString := fmt.Sprintf("Total score: %d", totalScore.points)

	fmt.Println(resultString)
}

func playRound(questions []Question) Score {

	currScore := Score{points: 0}
	for _, question := range questions {
		printQuestion(question)

		reader := bufio.NewReader(os.Stdin)
		response, err := reader.ReadString('\n')
		trimmedResponse := strings.TrimRight(response, "\n")

		if strings.EqualFold(trimmedResponse, question.correctAlternative) {
			currScore.addOne()
		}

		if err != nil {
			fmt.Println("Error reading string")
			log.Fatal(err)
		}

	}
	return currScore
}

func printPresentation() {

	fmt.Println("This is a quiz, try to answer as many questions as possible")
}

type Question struct {
	question           string
	alternatives       [4]string
	correctAlternative string
}

type Score struct {
	points int
}

func (currScore *Score) addOne() {
	result := currScore.points + 1
	currScore.points = result
}

func getQuestions() []Question {
	// temp questions
	f := openFile("questions.csv")
	data := getCsvData(f)
	questions := loadData(data)

	return questions
}

func openFile(fileName string) *os.File {

	f, err := os.Open(fileName)
	if err != nil {
		fmt.Println("Error opening file")
		log.Fatal(err)
	}

	return f
}

func getCsvData(file *os.File) [][]string {
	csvReader := csv.NewReader(file)
	csvReader.FieldsPerRecord = -1
	data, err := csvReader.ReadAll()
	if err != nil {
		fmt.Println("Error reading file")
		log.Fatal(err)
	}

	defer file.Close()

	return data
}

func loadData(data [][]string) []Question {

	var questions []Question
	for _, row := range data {
		question := row[0]
		alternatives := [4]string{row[1], row[2], row[3], row[4]}
		correct := row[5]

		questions = append(questions, Question{question, alternatives, correct})

	}

	return questions
}

func printQuestion(question Question) {
	fmt.Println()

	fmt.Println(question.question)
	fmt.Println()
	for _, a := range question.alternatives {
		fmt.Println(a)

	}
	fmt.Println()
}
