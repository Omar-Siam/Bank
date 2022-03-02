package src

import "fmt"

// BasicAccountDetails All accounts types must use the type account details.
type BasicAccountDetails struct {
	AccountID string
	Balance   float64
}

// Function that prints account details. Anyone that uses this type can use this method.

func (a BasicAccountDetails) ToString() {
	fmt.Println("AccountID= ", a.AccountID, " Balance= ", a.Balance)
}

// Account All accounts must be able to withdraw and deposit
type Account interface {
	Withdraw(amount float64) bool
	Deposit(amount float64)
}

// Comment the account interface better