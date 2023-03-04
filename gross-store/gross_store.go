package gross

// Units stores the Gross Store unit measurements.
func Units() map[string]int {
	myMap := map[string]int{}
	myMap["quarter_of_a_dozen"] = 3
	myMap["half_of_a_dozen"] = 6
	myMap["dozen"] = 12
	myMap["small_gross"] = 120
	myMap["gross"] = 144
	myMap["great_gross"] = 1728

	return myMap

}

// NewBill creates a new bill.
func NewBill() map[string]int {
	test := map[string]int{}
	return test
}

// AddItem adds an item to customer bill.
func AddItem(bill, units map[string]int, item, unit string) bool {
	val, exists := units[unit]
	if !exists {
		return false
	}
	bill[item] += val
	return true
}

// RemoveItem removes an item from customer bill.
func RemoveItem(bill, units map[string]int, item, unit string) bool {
	rpieces, rexists := units[unit]
	bpieces, bexists := bill[item]
	if !rexists || !bexists || rpieces > bpieces {
		return false
	} else if rpieces == bpieces {
		delete(bill, item)
	} else {
		bill[item] -= rpieces
	}

	return true
}

// GetItem returns the quantity of an item that the customer has in his/her bill.
func GetItem(bill map[string]int, item string) (int, bool) {
	pieces, exists := bill[item]
	return pieces, exists
}
