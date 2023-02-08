package gross

// Units stores the Gross Store unit measurements.
func Units() map[string]int {
	// panic("Please implement the Units() function")
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
	// panic("Please implement the NewBill() function")
	return make(map[string]int)
}

// AddItem adds an item to customer bill.
func AddItem(bill, units map[string]int, item, unit string) bool {
	// panic("Please implement the AddItem() function")
	value, exists := units[unit]
	if exists == false || value == 0 {
		return false
	}
	value, exists = bill[item]
	if exists == true {
		bill[item] += units[unit]
	} else {
		bill[item] = units[unit]
	}

	return true
}

// RemoveItem removes an item from customer bill.
func RemoveItem(bill, units map[string]int, item, unit string) bool {
	// panic("Please implement the RemoveItem() function")
	value, exists := bill[item]
	if exists == false || value == 0 {
		return false
	}
	value, exists = units[unit]
	if exists == false {
		return false
	} else if exists == true {
		quantity := bill[item]
		if (quantity - units[item]) < 0 {
			return false
		} else if quantity-units[item] == 0 {
			delete(bill, item)
		} else {
			bill[item] = bill[item] - 1
			return true
		}
	}
	return true
}

// GetItem returns the quantity of an item that the customer has in his/her bill.
func GetItem(bill map[string]int, item string) (int, bool) {
	panic("Please implement the GetItem() function")
}
