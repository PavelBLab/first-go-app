package main

import "fmt"

type floatMap map[string]float64

func (fm floatMap) output() {
	fmt.Println(fm)
}

func main() {
	mapExample()
	loopExample()
}

func loopExample() {

	coursesRating := make(floatMap, 3)
	coursesRating["go"] = 4.7
	coursesRating["java"] = 4.14
	coursesRating["python"] = 3.14
	coursesRating["ruby"] = 3.14

	coursesRating.output()

	for key, value := range coursesRating {
		fmt.Println("Key:", key, "Value:", value)
	}
}

func mapExample() {
	m1 := map[string]int{
		"zero": 0,
		"one":  1,
		"two":  2,
	}

	fmt.Println("m1", m1)

	if val, ok := m1["one"]; ok {
		fmt.Println("Key 'one' has value:", val)
	}

	delete(m1, "two")
	fmt.Println("After deletion, m1:", m1)

	m2 := make(map[string]int, 2) // make preallocate memory
	m2["three"] = 3
	m2["four"] = 4

	fmt.Println("m2", m2)
	for k, v := range m2 {
		println("Key:", k, "Value:", v)
	}

}
