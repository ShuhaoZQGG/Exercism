// interest_rate returns the interest rate for the provided balance.
double interest_rate(double balance) {
    // TODO: Implement the interest_rate function
    if (balance >= 0 and balance < 1000) {
        return 0.5;
    } else if (balance >= 1000 and balance < 5000) {
        return 1.621;
    } else if (balance >= 5000) {
        return 2.475;
    } else {
        return 3.213;
    }
}

// yearly_interest calculates the yearly interest for the provided balance.
double yearly_interest(double balance) {
    // TODO: Implement the yearly_interest function
    return interest_rate(balance) / 100 * balance;
}

// annual_balance_update calculates the annual balance update, taking into
// account the interest rate.
double annual_balance_update(double balance) {
    // TODO: Implement the annual_balance_update function
    return balance + yearly_interest(balance);
}

// years_until_desired_balance calculates the minimum number of years required
// to reach the desired balance.
int years_until_desired_balance(double balance, double target_balance) {
    // TODO: Implement the years_until_desired_balance function
    int year = 0;
    if (balance < 0) {
        return year;
    }

    for (double i {balance}; i < target_balance; i = annual_balance_update(i)) {
        year += 1;
    }
    
    return year;
}