package main

import (
	"Bank/src"
	"fmt"
)

func main() {
	// Need to add tests when the SavingsAccount file is complete.
	newSavingAccount := src.CreateSavingAccount(src.BasicAccountDetails{
		AccountID: "Omar",
		Balance:   2000,
	}, 50)
	fmt.Println(newSavingAccount)

	newCurrentAccount := src.CreateCurrentAccount(src.BasicAccountDetails{
		AccountID: "Hamza",
		Balance:   100000,
	}, 2)
	fmt.Println(newCurrentAccount)
}
