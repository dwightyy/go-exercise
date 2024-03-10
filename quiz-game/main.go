package main

import (
	"bufio"
	"encoding/csv"
	"fmt"
	"log"
	"os"
	"strings"
	"time"
    "flag"
)

func main() {

    var timeFlag = flag.Int("t", 10, "Time for each round")
    var fileNameFlag = flag.String("f","questions.csv", "Name of the file. i.e 'questions.csv'" )
    flag.Parse()
	printPresentation()

	questions := getQuestions(*fileNameFlag)

    scoreCh := make(chan Score)
    stopCh := make(chan bool)

	go playRound(questions, scoreCh, stopCh)
    go timer(*timeFlag, stopCh)
    totalScore := <- scoreCh

	resultString := fmt.Sprintf("Total score: %d/%d", totalScore.points, len(questions))

	fmt.Println(resultString)
}

func timer (inputDuration int, stopCh chan bool){
    <- time.After(time.Duration(inputDuration)* time.Second)
    stopCh <- true
}
func playRound(questions []Question, ch chan Score, stopCh chan bool) {

    currScore := Score{points: 0}
    for _, question := range questions {
        select {
            default:
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

             case <- stopCh:
                 fmt.Println("Time is over")
                 ch <- currScore
                 return
        }
    }
    ch <- currScore
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

func getQuestions(fileName string) []Question {
	f := openFile(fileName)
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
