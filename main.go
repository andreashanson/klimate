package main

import (
	"fmt"
	"time"
)

const searchString string = "sofj ölakjdsfa ölkajsd ölfkajsdf ijef osiejf sioefj sofj"

func main() {
	fb := NewFizzBuss(1000000, 3, 8)
	fb.Run()
}

type FizzBuss struct {
	Number   int
	Devider1 int
	Devider2 int
}

func NewFizzBuss(n, d1, d2 int) *FizzBuss {
	return &FizzBuss{
		Number:   n,
		Devider1: d1,
		Devider2: d2,
	}
}

func (fb *FizzBuss) Run() {
	startTime := time.Now()
	for i := 1; i <= fb.Number; i++ {
		if i%fb.Devider1 == 0 && i%fb.Devider2 == 0 {
			fmt.Println("FIZZBUZZ", i)
		} else if i%fb.Devider1 == 0 {
			fmt.Println("FIZZ", i)
		} else if i%fb.Devider2 == 0 {
			fmt.Println("BUSS", i)
		}
	}
	endTime := time.Now()
	duration := endTime.Sub(startTime)
	fmt.Println(duration)

}
