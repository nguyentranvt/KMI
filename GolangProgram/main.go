package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	//result := AddAtoB(1, 2)
	//print the result
	//println(result)

	//InputAddress
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Enter your address: ")
	address, _ := reader.ReadString('\n')
	//InputPhoneNumber
	reader = bufio.NewReader(os.Stdin)
	fmt.Print("Enter your phone number: ")
	phonenumber, _ := reader.ReadString('\n')
	//GrilFriend
	reader = bufio.NewReader(os.Stdin)
	fmt.Print("Enter your gril friend number:")
	grilfriend, _ := reader.ReadString('\n')

	fmt.Print("You live in: " + address)
	fmt.Print("Your phone number: " + phonenumber)
	fmt.Print("Your current number of girlfriends:  " + grilfriend)
}

func AddAtoB(ANumber int, BNumber int) int {
	return ANumber + BNumber
}
