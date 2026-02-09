package gross

import "fmt"

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
	return make(map[string]int)
}

// AddItem adds an item to customer bill.
func AddItem(bill, units map[string]int, item, unit string) bool {
	unitValue, exists := units[unit]
	if !exists {
		return false
	}
	bill[item] += unitValue
	return true
}

// RemoveItem removes an item from customer bill.
func RemoveItem(bill, units map[string]int, item, unit string) bool {
	fmt.Printf("item: %v, unit: %v\n", item, unit)
	itemValue, exists := bill[item]
	if !exists {
		fmt.Println("return false, item not exists")
		return false
	}
	unitValue, exists := units[unit]
	if !exists {
		fmt.Println("return false, unit not exists")
		return false
	}

	switch {
	case itemValue < 0:
		fmt.Println("return false, item:", itemValue)
		return false
	case itemValue == 0:
		delete(bill, item)
		fmt.Println("return true, item deleted")
		return true
	case itemValue > 0:
		fmt.Print("return true, item:", itemValue)
		bill[item] -= unitValue
		fmt.Print(". Стало:", bill[item], "\n")
		return true
	default:
		return false
	}
}

// GetItem returns the quantity of an item that the customer has in his/her bill.
func GetItem(bill map[string]int, item string) (int, bool) {
	v, exists := bill[item]
	if !exists {
		return 0, false
	}
	return v, true
}
