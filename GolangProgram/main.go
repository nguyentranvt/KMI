package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	// NameValue, TeamNumberValue, err := ReadKeyboard()
	// if err != nil {
	// 	//panic(err)
	// 	os.Exit(3)
	// }
	// fmt.Println(PrintName(NameValue))
	// fmt.Println(PrintTeamNumber(TeamNumberValue))
}

func ReadKeyboard() (Name string, GenderValue string, TeamNumber int, Listfavorite []string, Listlanguage []string, MarryValue string, err error) {
	//use ReadKeyboardForName function to read name
	NameValue, err := ReadKeyboardForName()
	if err != nil {
		return "", "", 0, nil, nil, "", err
	}
	//use ReadKeyboardForTeamNumber function to read team number
	BufferTeamNumber, err := ReadKeyboardForTeamNumber()
	if err != nil {
		return "", "", 0, nil, nil, "", err
	}
	//use ReadKeyboardForFavorite function to read favorite
	BufferFavorite, err := ReadKeyboardForFavorite()
	if err != nil {
		return "", "", 0, nil, nil, "", err
	}
	//use ReadKeyboardForLanguage function to read language
	BufferLanguage, err := ReadKeyboardForLanguage()
	if err != nil {
		return "", "", 0, nil, nil, "", err
	}
	fmt.Println("Name: ", NameValue)
	fmt.Println("Team Number: ", TeamNumber)
	fmt.Print("Please input sex: ")
	BufferGenderValue, err := ReadKeyboardForGender()

	//Review của sếp

	// ReviewValue, _ := reader.ReadString('\n')
	// return ReviewValue, 0, nil

	return NameValue, BufferGenderValue, BufferTeamNumber, BufferFavorite, BufferLanguage, MarryValue, nil
}

// Define input ReadKeyboardForGender
func ReadKeyboardForGender() (Gender string, err error) {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Please input sex: ")
	SexValue, _ := reader.ReadString('\n')
	if SexValue != "Male\n" && SexValue != "Female\n" {
		fmt.Println("Male or Female only")
		return "", nil
	}
	return SexValue, nil
}

// Define input ReadKeyboardForName
func ReadKeyboardForName() (Name string, err error) {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Please input Name: ")
	NameValue, _ := reader.ReadString('\n')
	return NameValue, nil
}

// define input ReadKeyboardForTeamNumber
func ReadKeyboardForTeamNumber() (TeamNumber int, err error) {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Please input Team Number: ")
	TeamNumberValue, _ := reader.ReadString('\n')
	//convert string to int
	TeamNumber, err = strconv.Atoi(TeamNumberValue)
	if err != nil {
		return 0, err
	}
	return TeamNumber, nil
}

// define input ReadKeyboardForBirthday
func ReadKeyboardForBirthday() (Birthday string, err error) {

	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Please input Birthday: ")
	BirthdayValue, _ := reader.ReadString('\n')
	//check format dd/mm/yyyy of birthday is correct
	if len(BirthdayValue) != 11 {
		fmt.Println("Please input dd/mm/yyyy format")
		return "", nil
	}
	if BirthdayValue[2] != '/' || BirthdayValue[5] != '/' {
		fmt.Println("Please input dd/mm/yyyy format")
		return "", nil
	}
	//define email format

	return BirthdayValue, nil
}

// define input ReadKeyboardForfavorite reuturn list of favorite Split by  comma
func ReadKeyboardForFavorite() (ListFavorite []string, err error) {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Please input favorite: ")
	favoriteValue, _ := reader.ReadString('\n')
	//convert string to list
	ListFavorite = strings.Split(favoriteValue, ",")
	return ListFavorite, nil
}

// define input ReadKeyboardForLanguage return list of language Split by  comma
func ReadKeyboardForLanguage() (ListLanguage []string, err error) {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Please input language: ")
	LanguageValue, _ := reader.ReadString('\n')
	//convert string to list
	ListLanguage = strings.Split(LanguageValue, ",")
	return ListLanguage, nil
}

// define input ReadkeyboardforMarry
func ReadKeyboardForMarry() (Marry string, err error) {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Please input Marry: ")
	MarryValue, _ := reader.ReadString('\n')
	if MarryValue != "Yes\n" && MarryValue != "No\n" {
		fmt.Println("Yes or No only")
		return "", nil
	}
	return MarryValue, nil
}

func PrintName(Name string) string {
	//Check if the name is empty
	if Name == "" {
		return "Name is empty"
	}
	return "Ho va ten " + Name
}

func PrintTeamNumber(Number int) string {
	ValueAfterConvert := strconv.Itoa(Number)
	return "Thanh vien nhom " + ValueAfterConvert
}

// func PrintAddress(Number int, Address string) string {
// 	ValueAfterConvert1 := strconv.Itoa(Number)
// 	return "Dia chi: " + Address + " So nha: " + ValueAfterConvert1
// }

func ReadKeyboardForAddress() (Address string, err error) {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Please input your address: ")
	AddressValue, _ := reader.ReadString('\n')
	if AddressValue == "" {
		fmt.Println("Please enter Address")
		return "", nil
	}
	return AddressValue, nil
}

func PrintSex(Sex string) string {
	return "Gioi tinh: " + Sex
}

//	func PrintTellNumber(Number string) string {
//		return "So dien thoai: " + Number
//	}

// define input ReadKeyboardForPhoneNumber
func ReadKeyboardForPhoneNumber() (PhoneNumber int, err error) {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Please input your phone number: ")
	PhoneNumberValue, _ := reader.ReadString('\n')
	//convert string to int
	PhoneNumber, err = strconv.Atoi(PhoneNumberValue)
	if err != nil {
		return 0, err
	}
	return PhoneNumber, nil
}

func PrintReview(Review string) string {
	return "Review của sếp: " + Review
}

func PrintUpdateTime(Time string) string {
	return "Thời gian cập nhật: " + Time
}

func PrintFreeTime(FreeTime string) string {
	return "Thời gian rảnh trong tuần:" + FreeTime
}
func PrintMyName(ten string) string {
	return "toi ten la: " + ten
}

// func PrintGirlFriendCurrentNumber(GirlFrendNumber int) string {
// 	ValueAfterConvert2 := strconv.Itoa(GirlFrendNumber)
// 	return "Số lượng bạn gái hiện tại: " + ValueAfterConvert2
// }

// define input ReadKeyboardForGrilFriendCurrentNumber
func ReadKeyboardForGrilFriendCurrentNumber() (GirlFriendNumber int, err error) {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Please input girl friend current number: ")
	GirlFriendNumberValue, _ := reader.ReadString('\n')
	//convert string to int
	GirlFriendNumber, err = strconv.Atoi(GirlFriendNumberValue)
	if err != nil {
		return 0, err
	}
	return GirlFriendNumber, nil
}

// func PrintIntroduce(Introduce string) string {
// 	return "Giới thiệu bản thân: " + Introduce
// }

// define input ReadKeyboardForIntroduce return your introduce
func ReadKeyboardForIntroduce() (Introduce []string, err error) {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Please input your introduce: ")
	IntroduceValue, _ := reader.ReadString('\n')
	//convert string to list
	Introduce = strings.Split(IntroduceValue, ",")
	return Introduce, nil
}

// func PrintManager(Manager string) string {
// 	return "Sếp hiện tại: " + Manager
// }

// Define input ReadKeyboardForManager
func ReadKeyboardForManager() (Manager string, err error) {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Please input your manager name: ")
	ManagerValue, _ := reader.ReadString('\n')
	return ManagerValue, nil
}

// func PrintManagerEmail(ManagerEmail string) string {
// 	return "Email sếp: " + ManagerEmail
// }

// Define input ReadKeyboardForManagerEmail
func ReadKeyboardForManagerEmail() (ManagerEmail string, err error) {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Please input your manager Email: ")
	ManagerEmailValue, _ := reader.ReadString('\n')
	return ManagerEmailValue, nil
}
func PrintDate(Date string) string {
	return "Ngay sinh: " + Date
}
func Printfavorite(Favour []string) string {
	BufferList := ""
	//Loop through the array
	for i := 0; i < len(Favour); i++ {
		//if the last element, don't add comma
		if i == len(Favour)-1 {
			BufferList += Favour[i]
		} else {
			BufferList += Favour[i] + ", "
		}
	}
	return "So thich: " + BufferList
}
func PrintYourEmail(YourEmail string) string {
	return "Email cua ban: " + YourEmail
}
func PrintLanguage(YourLanguage []string) string {
	//Define BufferList
	BufferList := ""
	for i := 0; i < len(YourLanguage); i++ {
		//if the last element, don't add comma
		if i == len(YourLanguage)-1 {
			BufferList += YourLanguage[i]
		} else {
			BufferList += YourLanguage[i] + ", "
		}
	}

	return "Ngoai ngu: " + BufferList

}
func PrintMarried(Marry string) string {
	return "Hon nhan: " + Marry
}
