package gross

func Units() map[string]int {

	return map[string]int{

		"quarter_of_a_dozen": 3,
		"half_of_a_dozen":    6,
		"dozen":              12,
		"small_gross":        120,
		"gross":              144,
		"great_gross":        1728,
	}

}

func NewBill() map[string]int {

	return make(map[string]int)

}

func AddItem(bill, units map[string]int, item, unit string) bool {

	// Guard
	_, exists := units[unit]
	if !exists {
		return false
	}

	value, _ := bill[item]
	bill[item] = value + units[unit]

	return true

}

func RemoveItem(bill, units map[string]int, item, unit string) bool {

	// Guards
	_, b_exists := bill[item]
	if !b_exists {
		return false
	}

	_, u_exists := units[unit]
	if !u_exists {
		return false
	}

	if bill[item]-units[unit] < 0 {
		return false
	}

	// Remove item if none are left
	if bill[item]-units[unit] == 0 {
		delete(bill, item)
		return true
	}

	bill[item] -= units[unit]
	return true

}

func GetItem(bill map[string]int, item string) (int, bool) {
	qty, exists := bill[item]
	return qty, exists
}
