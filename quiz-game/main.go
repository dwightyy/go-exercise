package main

import (
	"bufio"
	"encoding/csv"
	"fmt"
	"log"
	"os"
)

func main(){


    printPresentation()
    questions := getQuestions()
    printQuestions(questions)

    reader := bufio.NewReader(os.Stdin)
    result, err := reader.ReadString('\n')

    if err != nil{
        fmt.Println("Error reading string")
        log.Fatal(err)
    }

    fmt.Println(result)
}

func printPresentation(){

    fmt.Println("This is a quiz, try to answer as many questions as possible")
}

type Question struct{

    question string
    alternatives [4]string
    correctAlternative string

}

func getQuestions() []Question{
    // temp questions
    f := openFile("questions.csv")
    data := getCsvData(f)
    questions := loadData(data)

    return questions
}

func openFile(fileName string) *os.File{

    f, err := os.Open(fileName)
    if err != nil{
        fmt.Println("Error opening file")
        log.Fatal(err)
    }

    return f
}

func getCsvData(file *os.File) [][]string{
    csvReader := csv.NewReader(file)
    csvReader.FieldsPerRecord = -1
    data, err := csvReader.ReadAll()
     if err != nil{
        fmt.Println("Error reading file")
        log.Fatal(err)
    }

    defer file.Close()


    return data
}

func loadData(data [][]string) []Question{

    var questions []Question
    for _, row := range data{
        question := row[0]
        alternatives := [4]string{row[1], row[2], row[3], row[4]}
        correct := row[5]

        questions = append(questions, Question{question, alternatives, correct})

    }

    return questions
}

func printQuestions(questions []Question){
    fmt.Println(questions[0].question)
    for _, a := range(questions[0].alternatives){
        fmt.Println(a)
    }
}
