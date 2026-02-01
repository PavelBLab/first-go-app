package main

import "fmt"

type Product struct {
	title string
	id    string
	price float64
}

func New(title, id string, price float64) Product {
	return Product{
		title: title,
		id:    id,
		price: price,
	}
}

func main() {
	//slicesExample()

	//dynamicArrayExample()

	practiceArrayExample()

	unpackingListExample()

	// make make insertion more efficient as it is preoccupied some memory
	userNames := make([]string, 2, 5) // length 2, capacity 5
	userNames[0] = "John"
	userNames = append(userNames, "Alice", "Bob", "Charlie")
	fmt.Println("User names:", userNames)

	for index, value := range userNames {
		fmt.Println("User index:", index, "Value:", value)
	}
}

func unpackingListExample() {
	// Unpacking list example
	prices := []float64{19.99, 29.99, 9.99, 49.99, 5.99}
	discountPrices := []float64{100.0, 250.0, 75.5, 300.0, 50.0}
	prices = append(prices, discountPrices...)
	fmt.Println("All prices:", prices)

}

func practiceArrayExample() {

	// Task 1
	hobbies := [3]string{"Reading", "Traveling", "Cooking"}
	fmt.Println("All hobbies", hobbies)

	// Task 2
	fmt.Println("1st element:", hobbies[0])
	fmt.Println("2nd and 3rd elements:", hobbies[1:])

	// Task 3
	slicedHobbies := hobbies[:2]
	fmt.Println("Sliced last hobby:", slicedHobbies)

	// Task 4
	fmt.Println("Capacity of hobbies array:", cap(slicedHobbies))
	slicedHobbies = hobbies[1:]
	fmt.Println("Sliced first hobby", slicedHobbies)
	fmt.Println("Capacity of hobbies array:", cap(slicedHobbies))

	// Task 5
	dynamicArray := []string{"Reading", "Traveling", "Cooking", "Hiking", "Swimming", "Cycling"}
	fmt.Println("Dynamic array:", dynamicArray)

	// Task 6
	dynamicArray[1] = "Photography"
	fmt.Println("Updated dynamic array:", dynamicArray)

	// Task 7
	var product1 Product
	product1 = New("Laptop", "LP1001", 999.99)

	var product2 Product
	product2 = New("Smartphone", "SP2002", 699.49)

	var product3 Product
	product3 = New("Tablet", "TB3003", 399.00)

	productDynamicArray := []Product{product1, product2}
	//productDynamicArray[2] = product3 // This will cause an index out of range error
	productDynamicArray = append(productDynamicArray, product3)
	fmt.Println("Product dynamic array:", productDynamicArray)

}

func dynamicArrayExample() {
	prices := []float64{10.99, 5.49, 20.00}
	fmt.Println(cap(prices))

	updatedPrices := append(prices, 7.99, 15.30, 3.50)
	fmt.Println(updatedPrices)
}

func slicesExample() {
	var productNames [4]string = [4]string{"Laptop", "Smartphone", "Tablet", "Monitor"}
	prices := [6]float64{10.99, 5.49, 20.00, 7.75, 15.30, 0.0}

	productNames[2] = "Headphones"

	fmt.Println(prices)
	fmt.Println(productNames[0])

	featuredPrices := prices[1:3]
	fmt.Println(featuredPrices)

	featuredPrices = prices[:3]
	fmt.Println(featuredPrices)

	featuredPrices = prices[2:4]
	fmt.Println("featuredPrices [2:4]: ", featuredPrices)

	//featuredPrices = prices[:len(featuredPrices)-1]
	//fmt.Println(featuredPrices)

	highlightedPrices := featuredPrices[:1]
	fmt.Println("highlightedPrices: ", highlightedPrices)

	featuredPrices[0] = 299.99
	fmt.Println("featuredPrices", featuredPrices)
	fmt.Println("highlightedPrices: ", highlightedPrices)
	fmt.Println("prices", prices)

	fmt.Println(len(highlightedPrices), cap(highlightedPrices))
}
