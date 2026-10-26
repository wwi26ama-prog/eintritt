package exercises

import "fmt"

func ExampleAdmissionPriceFromMap() {
	prices := map[int]int{0: 0, 6: 300, 18: 500}
	fmt.Println(AdmissionPriceFromMap(5, prices))
	fmt.Println(AdmissionPriceFromMap(6, prices))
	fmt.Println(AdmissionPriceFromMap(17, prices))
	fmt.Println(AdmissionPriceFromMap(18, prices))

	// Output:
	// 0
	// 300
	// 300
	// 500
}

func ExampleAdmissionPriceFromMap_seniors() {
	prices := map[int]int{3: 200, 18: 400, 65: 250}
	fmt.Println(AdmissionPriceFromMap(2, prices))
	fmt.Println(AdmissionPriceFromMap(3, prices))
	fmt.Println(AdmissionPriceFromMap(64, prices))
	fmt.Println(AdmissionPriceFromMap(65, prices))
	fmt.Println(AdmissionPriceFromMap(30, map[int]int{}))

	// Output:
	// 0
	// 200
	// 400
	// 250
	// 0
}
