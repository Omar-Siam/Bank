package src

// SavingAccount Utilises the BasicAccountDetails struct, and uses an interest rate.
type SavingAccount struct {
	BasicAccountDetails BasicAccountDetails
	InterestRate        float64
}

// CreateSavingAccount Creates a savings account and adds £10 extra to the account balance if the initial balance is over to equal to £1000.
func CreateSavingAccount(a BasicAccountDetails, Interest float64) SavingAccount {
	newSavingAccount := SavingAccount{
		BasicAccountDetails: BasicAccountDetails{a.AccountID, a.Balance},
		InterestRate:        Interest,
	}
	if newSavingAccount.BasicAccountDetails.Balance >= 1000 {
		newSavingAccount.BasicAccountDetails.Balance += 10
	}
	return newSavingAccount
}

// Withdraw Implements the Account interface. Allows customers to withdraw money with a fee.
func (s SavingAccount) Withdraw(Amount float64) bool {
	if (s.BasicAccountDetails.Balance - (Amount - 3)) <= 10 {
		return false
	}
	return true
}

// Deposit Allows customers to add money to their balance with no fee.
func (s SavingAccount) Deposit(Amount float64) {
	s.BasicAccountDetails.Balance += Amount
}

// addInterest Adds the interest to the current balance and returns the interest amount.
func (s SavingAccount) addInterest() float64 {
	rate := s.InterestRate
	interestAmount := (s.BasicAccountDetails.Balance) * rate
	s.BasicAccountDetails.Balance += interestAmount
	return interestAmount
}
