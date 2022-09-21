package main

<<<<<<< HEAD
import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)
=======
func main() {
	result := AddAtoB(1, 2)
	//print the result
	println(result)
	ResultNewValue := PrintHelloWorld()
	println(ResultNewValue)
}
>>>>>>> Khoibranch

func main() {

}
func ReadKeyboard() string {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Enter your city: ")
	city, _ := reader.ReadString('\n')
	return city
}
func AddAtoB(ANumber int, BNumber int) int {
	return ANumber + BNumber
}
<<<<<<< HEAD
func PrintHelloWorld() string {
	return "Hello World"
}
func Print
Name(Name string) string {
	return "Ho va ten " + Name
}

func PrintTeamNumber(Number int) string {
	ValueAfterConvert := strconv.Itoa(Number)
	return "Thanh vien nhom " + ValueAfterConvert
}

func PrintAddress(Number int, Address string) string {
	ValueAfterConvert1 := strconv.Itoa(Number)
	return "Dia chi: " + Address + " So nha: " + ValueAfterConvert1
}

func PrintSex(Sex string) string {
	return "Gioi tinh: " + Sex
}

func PrintTellNumber(Number string) string {
	return "So dien thoai: " + Number
}

func PrintReview(Review string) string {
	return "Review của sếp: " + Review
}
=======
func PrintHelloWorld()string{
	return "Hello World Viet Nam"
} 
>>>>>>> Khoibranch
