package interest

// InterestRate returns the interest rate for the provided balance.
func InterestRate(balance float64) float32 {
	var interestRate float32 = 0
	if balance < 0 { 
		interestRate=3.213
	} else if balance >= 0 && balance < 1000 {
		interestRate=0.5
	} else if balance >=1000 && balance < 5000 { 
		interestRate=1.621
	} else { 
		interestRate=2.475
	}
	return interestRate
}

// Interest calculates the interest for the provided balance.
func Interest(balance float64) float64 {
	var rate = InterestRate(balance)
	var interest = balance * float64(rate) / 100

	return interest
}

// AnnualBalanceUpdate calculates the annual balance update, taking into account the interest rate.
func AnnualBalanceUpdate(balance float64) float64 {
	var interest = Interest(balance)

	return balance + interest
}

// YearsBeforeDesiredBalance calculates the minimum number of years required to reach the desired balance.
func YearsBeforeDesiredBalance(balance, targetBalance float64) int {
	var yearsNeeded int = 0;
	var balancePerAnnum = balance
	 
	for balancePerAnnum < targetBalance {
		balancePerAnnum = AnnualBalanceUpdate(balancePerAnnum)
		yearsNeeded++
	}

	return yearsNeeded
}
