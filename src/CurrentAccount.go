package src

type CurrentAccount struct {
	BasicAccountDetails BasicAccountDetails
	numberOfChecksUsed  int
}

// CreateCurrentAccount Creates a CurrentAccount object.
func CreateCurrentAccount(a BasicAccountDetails, numbOfChecks int) CurrentAccount {
	newCurrentAccount := CurrentAccount{
		BasicAccountDetails: BasicAccountDetails{a.AccountID, a.Balance},
		numberOfChecksUsed:  numbOfChecks,
	}
	return newCurrentAccount
}

// Withdraw Allows customers to withdraw funds with a fee of £1 per Withdraw. If balance is below 0 after the Withdrawal return false.
func (c CurrentAccount) Withdraw(Amount float64) bool {
	CABalance := c.BasicAccountDetails.Balance

	if CABalance-(Amount-1) < 0 {
		return false
	}
	CABalance -= CABalance - (Amount - 1)
	return true
}

// Deposit Allows customers to add money to their account. However since it added using an ATM there is a £1 fee.
func (c CurrentAccount) Deposit(Amount float64) {
	c.BasicAccountDetails.Balance += Amount - 1
}

// resetChecksUsed Resets the checks used to 0.
func (c CurrentAccount) resetChecksUsed() {
	c.numberOfChecksUsed = 0
}

func (c CurrentAccount) withdrawUsingCheck(Amount float64) bool {
	// Need to implement.
	return false
}
