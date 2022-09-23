package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	fmt.Print("Enter your city: \n")
	BufferValue := PrintMyName("huynh bao toan: \n")
	fmt.Print(BufferValue)
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
func PrintHelloWorld() string {
	return "Hello World"
}
func PrintName(Name string) string {
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
func PrintMyName(ten string) string {
	return "toi ten la: " + ten
}
func PrintGirlFriendCurrentNumber(GirlFrendNumber int) string {
	ValueAfterConvert2 := strconv.Itoa(GirlFrendNumber)
	return "Số lượng bạn gái hiện tại: " + ValueAfterConvert2
}
func PrintIntroduce(Introduce string) string {
	return "Giới thiệu bản thân: " + Introduce
}
func PrintManager(Manager string) string {
	return "Sếp hiện tại: " + Manager
}
func PrintManagerEmail(ManagerEmail string) string {
	return "Email sếp: " + ManagerEmail
}
func PrintDate(Date string) string {
	return "Ngay sinh: " + Date
}
func PrintFavourite(Favour string) string {
	return "So thich: " + Favour
}
func PrintYourEmail(YourEmail string) string {
	return "Email cua ban: " + YourEmail
}
func PrintLanguage(YourLanguage string) string {
	return "Ngon ngu: " + YourLanguage
}
func PrintMarried(Answer string) string {
	return "Hon nhan: " + Answer
}
