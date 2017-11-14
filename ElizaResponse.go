package main

import (
	"fmt"
	"math/rand"
	"time"
)

func ElizaResponse(input string) string {
	answers := []string{
	"I’m not sure what you’re trying to say. Could you explain it to me?",
    "How does that make you feel?",
    "Why do you say that?",
	}
	fmt.Println(answers[rand.Intn(len(answers))])
}


func main() {

	rand.Seed(time.Now().UTC().UnixName()) 

fmt.Println("People say I look like both my mother and father.")
fmt.Println(ElizaResponse("People say I look like both my mother and father."))
fmt.Println()

fmt.Println("Father was a teacher.")
fmt.Println(ElizaResponse("Father was a teacher."))
fmt.Println()

fmt.Println("I was my father’s favourite.")
fmt.Println(ElizaResponse("I was my father’s favourite."))
fmt.Println()

fmt.Println("I’m looking forward to the weekend.")
fmt.Println(ElizaResponse("I’m looking forward to the weekend."))
fmt.Println()

fmt.Println("My grandfather was French!")
fmt.Println(ElizaResponse("My grandfather was French!"))
fmt.Println()

}