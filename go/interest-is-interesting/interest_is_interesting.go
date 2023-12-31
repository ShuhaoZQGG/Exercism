package interest
import "math"
// InterestRate returns the interest rate for the provided balance.
func InterestRate(balance float64) float32 {
    var interestRate float32
	switch {
        case balance < 0:
    		interestRate = 3.213
        	break
        case balance >= 0 && balance < 1000:
    	    interestRate = 0.5
        	break
        case balance >= 1000 && balance < 5000:
    		interestRate = 1.621
        	break
        default:
    		interestRate = 2.475
        	break
    }
	return interestRate
}

// Interest calculates the interest for the provided balance.
func Interest(balance float64) float64 {
	return balance * float64(InterestRate(balance)) / 100.0
}

// AnnualBalanceUpdate calculates the annual balance update, taking into account the interest rate.
func AnnualBalanceUpdate(balance float64) float64 {
	return balance + Interest(balance)
}

// YearsBeforeDesiredBalance calculates the minimum number of years required to reach the desired balance.
func YearsBeforeDesiredBalance(balance, targetBalance float64) int {
    counter := 0
    if balance >= targetBalance {
        return 0
    }

    if InterestRate(balance) == InterestRate(targetBalance) {
        return int(math.Log(targetBalance / balance) / math.Log(1.0 + float64(InterestRate(balance)) / 100.0)) + 1
    }

    for i := balance; i <= targetBalance; i += Interest(i) {
        counter += 1
    }

    return counter
}
