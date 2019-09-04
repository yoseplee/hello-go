package main

import (
	"fmt"
	"time"
)

type Man struct {
	Name string
	Age  int
	Job  string
}

func (man *Man) NewMan(name string, age int, job string) {
	man.Name = name
	man.Age = age
	man.Job = job
}

func (man *Man) DescribeMeForTenTimes(ch chan int) {
	counter := 1
	for i := 0; i < 10; i += 1 {
		fmt.Printf("%-8s: %-3d - %-10s\n", man.Name, man.Age, man.Job)
		counter = i
	}
	ch <- counter
}

func main() {
	fmt.Println("Go Routine Example")
	var myMan Man = Man{"lee", 12, "student"}
	myMan2 := Man{"kim", 24, "student"}
	var ch chan int = make(chan int, 3)

	go myMan.DescribeMeForTenTimes(ch)
	go myMan2.DescribeMeForTenTimes(ch)

	time.Sleep(time.Second * 5)
	x, y := <-ch, <-ch
	fmt.Println(x)
	fmt.Println(y)
}
