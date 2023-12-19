package lasagna

// TODO: define the 'PreparationTime()' function
func PreparationTime(layers []string, time int) (totalTime int) {
    var preparationTime int
    if time == 0 {
        preparationTime = 2
    } else {
    	preparationTime = time
    }
	totalTime = preparationTime * len(layers)
    return
}
// TODO: define the 'Quantities()' function
func Quantities(layers []string) (noodles int, sauce float64) {
    noodles = 0
    sauce = 0.0
    for _, v := range layers {
        if v == "sauce" {
            sauce += 1.0
        }

        if v == "noodles" {
            noodles += 1
        }
    }

    noodles *= 50
    sauce *= 0.2
    return
}
// TODO: define the 'AddSecretIngredient()' function
func AddSecretIngredient(friendsList []string, myList []string) {
    myList[len(myList) - 1] = friendsList[len(friendsList) -1]
}
// TODO: define the 'ScaleRecipe()' function
func ScaleRecipe(quantities []float64, numberOfPortions int) []float64 {
    newQuantities := make([]float64, len(quantities))
    copy(newQuantities, quantities)
    for i, v := range newQuantities {
        newQuantities[i] = v * float64(numberOfPortions) /2.0
    }
	
	return newQuantities
}

// Your first steps could be to read through the tasks, and create
// these functions with their correct parameter lists and return types.
// The function body only needs to contain `panic("")`.
// 
// This will make the tests compile, but they will fail.
// You can then implement the function logic one by one and see
// an increasing number of tests passing as you implement more 
// functionality.
