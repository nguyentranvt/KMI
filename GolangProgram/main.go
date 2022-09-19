package main

func main() {
	result := AddAtoB(1, 2)
	//print the result
	println(result)
	ResultNewValue := PrintHelloWorld()
	println(ResultNewValue)
}

func AddAtoB(ANumber int, BNumber int) int {
	return ANumber + BNumber
}

func PrintHelloWorld() string {
	return "Hello Cac em"

}
