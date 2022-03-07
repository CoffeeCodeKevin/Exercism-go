package gross

// Units stores the Gross Store unit measurements.
func Units() map[string]int {
	units := map[string]int{}
	units["quarter_of_a_dozen"] = 3
	units["half_of_a_dozen"] = 6
	units["dozen"] = 12
	units["small_gross"] = 120
	units["gross"] = 144
	units["great_gross"] = 1728
	return units
}

// NewBill creates a new bill.
func NewBill() map[string]int {
	bill := map[string]int{}
	return bill
}

// AddItem adds an item to customer bill.
func AddItem(bill, units map[string]int, item, unit string) bool {
	_, exists := units[unit]
	if exists {
		bill[item] = units[unit] + bill[item]
	}
	return exists
}

// RemoveItem removes an item from customer bill.
func RemoveItem(bill, units map[string]int, item, unit string) bool {
	_, vExists := bill[item]
	_, uExists := units[unit]
	var isRemovable bool
	if vExists && uExists {
		rValue := units[unit]
		value := bill[item]
		newTotal := value - rValue
		switch {
		case value-rValue < 0:
			isRemovable = false
		case value-rValue == 0:
			isRemovable = true
			delete(bill, item)
		case value-rValue > 0:
			isRemovable = true
			bill[item] = newTotal
		}
	} else {
		isRemovable = false
	}
	return isRemovable
}

// GetItem returns the quantity of an item that the customer has in his/her bill.
func GetItem(bill map[string]int, item string) (int, bool) {
	value, exists := bill[item]
	return value, exists
}
