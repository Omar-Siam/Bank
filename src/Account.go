package src

import "fmt"

// All accounts types must use the type account details.

type BasicAccountDetails struct {
	AccountID string
	Balance   float64
}

// Function that prints account details. Anyone that uses this type can use this method.

func (a BasicAccountDetails) ToString() {
	fmt.Println("AccountID= ", a.AccountID, " Balance= ", a.Balance)
}

// All accounts must be able to withdraw and deposit

type Account interface {
	Withdraw(amount float64) bool
	Deposit(amount float64)
}

type TestingAccount struct {
	AccountDetails BasicAccountDetails
	Interest       int64
}

// Constructor to create new Account.

func CreateNewAccount(AccountID string, Balance float64) BasicAccountDetails {
	newAccount := BasicAccountDetails{AccountID, Balance}
	return newAccount
}

// Constructor for Testing Account. All Accounts must have Basic Account Details.

func NewTestingAccount(a BasicAccountDetails, Interest int64) TestingAccount {
	newTestingAccount := TestingAccount{
		AccountDetails: CreateNewAccount(a.AccountID, a.Balance),
		Interest:       Interest,
	}
	return newTestingAccount
}
