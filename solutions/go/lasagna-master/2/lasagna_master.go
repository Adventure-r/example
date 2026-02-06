package lasagna

// TODO: define the 'PreparationTime()' function
func PreparationTime(layers []string, minutesForLayer int) int {
	if minutesForLayer == 0 {
		minutesForLayer = 2
	}
	return len(layers) * minutesForLayer
}

// TODO: define the 'Quantities()' function
func Quantities(layers []string) (noodles int, sauce float64) {
	for _, v := range layers {
		switch v {
		case "sauce":
			sauce += 0.2
		case "noodles":
			noodles += 50
		}
	}
	return
}

// TODO: define the 'AddSecretIngredient()' function
func AddSecretIngredient(friendsList, myList []string) {
	myList[len(myList)-1] = friendsList[len(friendsList)-1]
}

// TODO: define the 'ScaleRecipe()' function
func ScaleRecipe(quantities []float64, count int) []float64 {
	scaledRecipe := make([]float64, len(quantities))
	for i, v := range quantities {
		scaledRecipe[i] = v / 2 * float64(count)
	}
	return scaledRecipe
}

// Your first steps could be to read through the tasks, and create
// these functions with their correct parameter lists and return types.
// The function body only needs to contain `panic("")`.
//
// This will make the tests compile, but they will fail.
// You can then implement the function logic one by one and see
// an increasing number of tests passing as you implement more
// functionality.
