package exercises

import "fmt"

func ExampleAdmissionPriceFromLists() {
	minAges := []int{0, 6, 18}
	prices := []int{0, 300, 500}
	fmt.Println(AdmissionPriceFromLists(5, minAges, prices))
	fmt.Println(AdmissionPriceFromLists(6, minAges, prices))
	fmt.Println(AdmissionPriceFromLists(17, minAges, prices))
	fmt.Println(AdmissionPriceFromLists(18, minAges, prices))

	// Output:
	// 0
	// 300
	// 300
	// 500
}

func ExampleAdmissionPriceFromLists_seniors() {
	minAges := []int{3, 18, 65}
	prices := []int{200, 400, 250}
	fmt.Println(AdmissionPriceFromLists(2, minAges, prices))
	fmt.Println(AdmissionPriceFromLists(3, minAges, prices))
	fmt.Println(AdmissionPriceFromLists(64, minAges, prices))
	fmt.Println(AdmissionPriceFromLists(65, minAges, prices))
	fmt.Println(AdmissionPriceFromLists(30, []int{}, []int{}))

	// Output:
	// 0
	// 200
	// 400
	// 250
	// 0
}
