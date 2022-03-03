package src

import "fmt"

// BasicAccountDetails All accounts types must use the type account details.
type BasicAccountDetails struct {
	AccountID string
	Balance   float64
}

// ToString Prints account details.
func (a BasicAccountDetails) ToString() {
	fmt.Println("AccountID= ", a.AccountID, " Balance= ", a.Balance)
}

// Account All accounts must be able to withdraw and deposit
type Account interface {
	Withdraw(float64) bool
	Deposit(float64)
}

// Comment the account interface better
