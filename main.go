package main

import (
	"Bank/src"
	"fmt"
)

func main() {
	// Need to add tests when the SavingsAccount file is complete.
	newSavingAccount := src.CreateSavingAccount(src.BasicAccountDetails{AccountID: "omar", Balance: 2000}, 50)
	fmt.Println(newSavingAccount)
}
