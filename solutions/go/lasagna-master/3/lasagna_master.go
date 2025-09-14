package lasagna

func PreparationTime(layers []string, tm int) int {
	if tm == 0 {
		tm = 2
	}
	return len(layers) * tm
}

func Quantities(layers []string) (n int, s float64) {

	for _, v := range layers {

		switch v {
		case "sauce":
			s += 0.2
		case "noodles":
			n += 50
		}

	}

	return

}

func AddSecretIngredient(friends []string, my []string) {
	my[len(my)-1] = friends[len(friends)-1]
}

func ScaleRecipe(amounts []float64, scale int) []float64 {

	ret := make([]float64, len(amounts))

	for idx, val := range amounts {
		ret[idx] = val / 2.0 * float64(scale)
	}

	return ret

}
