package gross

// Units stores the Gross Store unit measurements.
func Units() map[string]int {
	units := map[string]int{
		"quarter_of_a_dozen": 3,
		"half_of_a_dozen":    6,
		"dozen":              12,
		"small_gross":        120,
		"gross":              144,
		"great_gross":        1728,
	}

	return units
}

// NewBill creates a new bill.
func NewBill() map[string]int {
	newBill := make(map[string]int)
	return newBill
}

// AddItem adds an item to customer bill.
func AddItem(bill, units map[string]int, item, unit string) bool {
	unitValue, exists := units[unit]
	if !exists {
		return false
	}

	_, exists = bill[item]
	if exists {
		bill[item] += unitValue
	} else {
		bill[item] = unitValue
	}

	return true
}

// RemoveItem removes an item from customer bill.
// - Return `false` if the given item is **not** in the bill
// - Return `false` if the given `unit` is not in the `units` map.
// - Return `false` if the new quantity would be less than 0.
// - If the new quantity is 0, completely remove the item from the `bill` then return `true`.
// - Otherwise, reduce the quantity of the item and return `true`.
func RemoveItem(bill, units map[string]int, item, unit string) bool {
	itemValue, exists := bill[item]
	if !exists {
		return false
	}

	unitValue, exists := units[unit]
	if !exists {
		return false
	}

	itemValue -= unitValue
	if itemValue < 0 {
		return false
	}

	if itemValue == 0 {
		delete(bill, item)
	} else {
		bill[item] = itemValue
	}

	return true
}

// GetItem returns the quantity of an item that the customer has in his/her bill.
// - Return `0` and `false` if the `item` is not in the bill.
// - Otherwise, return the quantity of the item in the `bill` and `true`.
func GetItem(bill map[string]int, item string) (int, bool) {
	itemValue, exists := bill[item]
	if !exists {
		return 0, false
	}

	return itemValue, true
}
