package gross

// Units stores the Gross Store unit measurements.
func Units() (units map[string]int) {
	units = map[string]int{
		"quarter_of_a_dozen": 3,
		"half_of_a_dozen":    6,
		"dozen":              12,
		"small_gross":        120,
		"gross":              144,
		"great_gross":        1728,
	}

	return
}

// NewBill creates a new bill.
func NewBill() map[string]int {
	bill := make(map[string]int)
	return bill
}

// AddItem adds an item to customer bill.
func AddItem(bill, units map[string]int, item, unit string) bool {
	unitValue, exists := units[unit]
	if exists == false {
		return false
	}

	_, exists = bill[item]

	if exists == false {
		bill[item] = 0
	}

	bill[item] += unitValue

	return true
}

// RemoveItem removes an item from customer bill.
func RemoveItem(bill, units map[string]int, item, unit string) bool {
	score, exists := units[unit]

	if exists == false {
		return false
	}

	quantity, exists := bill[item]

	if exists == false {
		return false
	}

	if quantity < score {
		return false
	}

	bill[item] -= score

	if bill[item] == 0 {
		delete(bill, item)
	}

	return true
}

// GetItem returns the quantity of an item that the customer has in his/her bill.
func GetItem(bill map[string]int, item string) (quantity int, exists bool) {
	quantity, exists = bill[item]

	return
}
