package main

func sumNumbers() {
	result := 0
	for i := 0; i < 100; i++ {
		result += i
	}
	//出力を乱さないためにコメントアウト
	//fmt.Println(result)
}

func printNumbers() {
	for i := 0; i < 100; i++ {
		//出力を乱さないためにコメントアウト
		//fmt.Print(i)
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
