package main

import (
	"Bank/src"
	"fmt"
)

func main() {
	newAccount := src.NewTestingAccount(src.CreateNewAccount("ABC", 900), 5)
	fmt.Println(newAccount)
}
