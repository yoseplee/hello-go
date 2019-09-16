package main

import "fmt"

type describer interface {
	info()
}

type Man struct {
	Name string
	Age  int
	Job  string
}

func (man *Man) info() {
	fmt.Printf("Name: %s | Age: %d | Job: %s\n", man.Name, man.Age, man.Job)
}

func (man *Man) ChangeName(name string) {
	man.Name = name
}

func ChangeNameByValue(man Man, name string) {
	man.Name = name
}

func main() {
	myMan := Man{"lee", 32, "student"}
	myMan.info()
	//call by reference
	myMan.ChangeName("park")
	myMan.info()
	//call by value
	ChangeNameByValue(myMan, "KILL")
	myMan.info()
}
