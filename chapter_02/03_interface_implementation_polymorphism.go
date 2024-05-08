package main

import (
	"fmt"
)

type Speaker interface {
	Speak() string
}

type Dog struct {
	Name string
}

func (d Dog) Speak() string {
	return "Woof! My name is " + d.Name
}

type Robot struct {
	Model string
}

func (r Robot) Speak() string {
	return "Beep boop. I am model " + r.Model
}

func introduceSpeaker(s Speaker) {
	fmt.Println(s.Speak())
}

func main() {
	dog := Dog{Name: "Buddy"}
	robot := Robot{Model: "XJ-9"}
	introduceSpeaker(dog)   // Output: Woof! My name is Buddy
	introduceSpeaker(robot) // Output: Beep boop. I am model XJ-9
}
