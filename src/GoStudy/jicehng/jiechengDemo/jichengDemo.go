package main

import "fmt"

type Student struct {
	Name   string
	Score  int
	Gender string
	money  float64
}

type Pupil struct {
	Student
	Name string
}

type Graduate struct {
	Student
}

func (student *Student) sayHello() {
	fmt.Println("hello student")
}
func (pupil *Pupil) sayHello() {
	fmt.Println("hello pupil")
}

func (student *Student) GetName() string {
	return student.Name
}

func (student *Student) GetScore() int {
	return student.Score
}

func (student *Student) GetGender() string {
	return student.Gender
}
func (student *Student) hello() {
	fmt.Println("hello")
}
func main() {
	pupil := Pupil{Student{"蒋达", 100, "man", 0}, "jiangda"}
	fmt.Println(pupil.money)
	pupil.sayHello()
	fmt.Println(pupil.Name)
	fmt.Println(pupil.Score)
	fmt.Println(pupil.GetScore())

}
