package main

import "fmt"

func sumNumbers() {
	result := 0
	for i := 0; i < 100; i++ {
		result += i
	}
	fmt.Println(result)
}

func printNumbers() {
	for i := 0; i < 100; i++ {
		fmt.Print(i)
	}
}

func runInTime() {
	sumNumbers()
	printNumbers()
}

func runConccurent() {
	go sumNumbers()
	go printNumbers()
}

func main() {
}
