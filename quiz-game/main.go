package main

import (
    "fmt"
    "log"
    "bufio"
    "os"
)

func main(){


    printPresentation()
    questions := readQuestions()
    printQuestions(questions)

    reader := bufio.NewReader(os.Stdin)
    result, err := reader.ReadString('\n')

    if err != nil{
        fmt.Println("Error")
        log.Fatal(err)
    }

    fmt.Println(result)
}

func printPresentation(){

    fmt.Println("This is a quiz, try to answer as many questions as possible")
}

type Question struct{

    question string
    alternatives []string
    correctAlternative string

}

func readQuestions() []Question{
    // temp questions

    question1 := "What does CPU stand for?"
    alternatives1 := [4]string{"Central Processing Unit","Computer Processing Unit","Central Processor Unit","Computer Processor Unit"}
    correct1 := "Central Processing Unit"


    var questions []Question
    questions = append(questions, Question{question1, alternatives1[:], correct1})

    return questions
}

func printQuestions(questions []Question){
    fmt.Println(questions[0].question)
    for _, a := range(questions[0].alternatives){
        fmt.Println(a)
    }
}
