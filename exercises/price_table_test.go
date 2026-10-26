package exercises

import "fmt"

func ExampleAdmissionPriceFromTable() {
	table := [][]int{{0, 0}, {6, 300}, {18, 500}}
	fmt.Println(AdmissionPriceFromTable(5, table))
	fmt.Println(AdmissionPriceFromTable(6, table))
	fmt.Println(AdmissionPriceFromTable(17, table))
	fmt.Println(AdmissionPriceFromTable(18, table))

	// Output:
	// 0
	// 300
	// 300
	// 500
}

func ExampleAdmissionPriceFromTable_seniors() {
	table := [][]int{{3, 200}, {18, 400}, {65, 250}}
	fmt.Println(AdmissionPriceFromTable(2, table))
	fmt.Println(AdmissionPriceFromTable(3, table))
	fmt.Println(AdmissionPriceFromTable(64, table))
	fmt.Println(AdmissionPriceFromTable(65, table))
	fmt.Println(AdmissionPriceFromTable(30, [][]int{}))

	// Output:
	// 0
	// 200
	// 400
	// 250
	// 0
}
