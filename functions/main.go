package main

import "fmt"

type transformFunction func(int) int

func main() {
	numbers := []int{1, 2, 3, 4, 5}
	moreNumbers := []int{6, 7, 8}

	double1 := transformNumbers(&numbers, double)
	triple1 := transformNumbers(&numbers, triple)
	fmt.Println(double1)
	fmt.Println(triple1)

	transformerFunction1 := getTransformerFunction(&numbers)
	transformerFunction2 := getTransformerFunction(&moreNumbers)

	transformedNumbers := transformNumbers(&numbers, transformerFunction1)
	moreTransformedNumbers := transformNumbers(&numbers, transformerFunction2)

	fmt.Println(transformedNumbers)
	fmt.Println(moreTransformedNumbers)

	fmt.Println(anonymousFunction(numbers))

	double2 := createTransformer(2)
	triple2 := createTransformer(3)

	doubled := transformNumbers(&numbers, double2)
	tripled := transformNumbers(&numbers, triple2)
	fmt.Println(doubled)
	fmt.Println(tripled)

	fmt.Println(factorial(5))
	fmt.Println(recursionFactorial(5))

	sum := sumUp(1, 2, 3, 4, 5)
	fmt.Println(sum)
	fmt.Println(sumUp(1, numbers...)) // ... is a spread operator

}

func sumUp(startingValue int, numbers ...int) int {
	sum := 0

	for _, val := range numbers {
		sum += val
	}

	return sum
}

func recursionFactorial(number int) int {

	if number == 0 {
		return 1
	}

	return number * recursionFactorial(number-1)
}

func factorial(number int) int {
	result := 1

	for i := 1; i <= number; i++ {
		result = result * i
	}

	return result
}

func anonymousFunction(numbers []int) []int {
	return transformNumbers(&numbers, func(number int) int {
		return number * 2
	})
}

func createTransformer(factor int) func(int) int {
	return func(number int) int {
		return number * factor
	}
}

func transformNumbers(numbers *[]int, transform func(int) int) []int {
	var dNumbers []int

	for _, val := range *numbers {
		dNumbers = append(dNumbers, transform(val))
	}

	return dNumbers
}

func getTransformerFunction(numbers *[]int) transformFunction {
	if (*numbers)[0] == 1 {
		return double
	}
	return triple
}

func double(number int) int {
	return number * 2
}

func triple(number int) int {
	return number * 3
}
