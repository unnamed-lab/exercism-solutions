package lasagna

// TODO: define the 'PreparationTime()' function
func PreparationTime(layers []string, preparationTime int) int {
	var timePrepared int = 2
	if preparationTime > 0 {
		timePrepared = preparationTime
	}
	return len(layers) * timePrepared
}

// TODO: define the 'Quantities()' function

func Quantities(layers []string) (noddles int, sauce float64) {

	for i := 0; i < len(layers); i++ {
		if layers[i] == "noodles" {
			noddles++
		} 
		if layers[i] == "sauce" {
			sauce++
		}
	} 

	noddles *= 50
	sauce *= 0.2
	return
}

// TODO: define the 'AddSecretIngredient()' function
func AddSecretIngredient(newList, myList []string) []string {
	lastListItem := newList[len(newList) - 1]
	myList[len(myList) - 1] = lastListItem;

	return myList
}

// TODO: define the 'ScaleRecipe()' function
func ScaleRecipe(quatities []float64, portion int) []float64 {
	var scaleArr []float64 
	for i := 0; i < len(quatities); i++ {
		qnt :=  quatities[i] * (float64(portion) / 2.0)
		scaleArr = append(scaleArr, qnt)
	}

	return  scaleArr
}

// Your first steps could be to read through the tasks, and create
// these functions with their correct parameter lists and return types.
// The function body only needs to contain `panic("")`.
//
// This will make the tests compile, but they will fail.
// You can then implement the function logic one by one and see
// an increasing number of tests passing as you implement more
// functionality.
