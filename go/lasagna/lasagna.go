package lasagna

const OvenTime = 40

func RemainingOvenTime(actual int) int {
	return OvenTime - actual
}

func PreparationTime(numberOfLayers int) int {
	return numberOfLayers * 2
}

func ElapsedTime(numbersOfLayers, actualinutesInOven int) int {
	return PreparationTime(numbersOfLayers) + actualinutesInOven
}
