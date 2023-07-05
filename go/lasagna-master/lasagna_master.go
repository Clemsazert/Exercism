package lasagna

func PreparationTime(layers []string, avgPreparationTime int) int {
	if avgPreparationTime == 0 {
		return 2 * len(layers)
	}
	return avgPreparationTime * len(layers)
}

func Quantities(layers []string) (int, float64) {
	var noodleAmount int
	var sauceAmmount float64
	for _, layer := range layers {
		if layer == "noodles" {
			noodleAmount += 50
		} else if layer == "sauce" {
			sauceAmmount += 0.2
		}
	}
	return noodleAmount, sauceAmmount
}

func AddSecretIngredient(layers1 []string, layers2 []string) {
	layers2[len(layers2)-1] = layers1[len(layers1)-1]
}

func ScaleRecipe(recipe []float64, portions int) []float64 {
	result := make([]float64, len(recipe))
	for index, amount := range recipe {
		result[index] = amount / 2 * float64(portions)
	}
	return result
}
