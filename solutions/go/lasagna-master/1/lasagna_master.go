package lasagna

func PreparationTime(layers []string, tm int) int {
	if tm == 0 {
		tm = 2
	}
	return len(layers) * tm
}

func Quantities(layers []string) (int, float64) {

	var n int
	var s float64

	for _, v := range layers {

		switch v {
		case "sauce":
			s += 0.2
		case "noodles":
			n += 50
		}

	}

	return n, s

}

func AddSecretIngredient(friends []string, my []string) {
	last := friends[len(friends)-1]
	my = my[:len(my)-1]
	my = append(my, last)
}

func ScaleRecipe(amounts []float64, scale int) []float64 {

	var ret []float64
	f_scale := float64(scale)

	for _, val := range amounts {
		ret = append(ret, (val/2.0)*f_scale)
	}

	return ret

}
