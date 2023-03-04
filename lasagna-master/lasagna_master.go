package lasagna

// TODO: define the 'PreparationTime()' function
func PreparationTime(layers []string, time int) int {
	if time == 0 {
		return (len(layers) * 2)
	} else {
		return (len(layers) * time)
	}
}

// TODO: define the 'Quantities()' function
// Quantities function return an int wit 50* every noodles in string and a float with 0.2 for every sauce
func Quantities(layers []string) (int, float64) {
	var sumNoodles int
	var sumSauce float64

	for i := 0; i < len(layers); i++ {
		if layers[i] == "noodles" {
			sumNoodles += 50
		} else if layers[i] == "sauce" {
			sumSauce = sumSauce + 0.2
		}
	}
	return sumNoodles, sumSauce
}

// TODO: define the 'AddSecretIngredient()' function
func AddSecretIngredient(friendList, myList []string) {
	myList[len(myList)-1] = friendList[len(friendList)-1]
}

// TODO: define the 'ScaleRecipe()' function
func ScaleRecipe(recipe []float64, quantities int) []float64 {
	var scaledRecipe = make([]float64, len(recipe))
	// normalizationFactor := quantities / 2
	for i := 0; i < len(recipe); i++ {
		scaledRecipe[i] = recipe[i] * 0.5 * float64(quantities)
	}

	return scaledRecipe
}
