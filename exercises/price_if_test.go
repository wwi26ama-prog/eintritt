package exercises

import "fmt"

func ExampleAdmissionPrice() {
	fmt.Println(AdmissionPrice(0))
	fmt.Println(AdmissionPrice(5))
	fmt.Println(AdmissionPrice(6))
	fmt.Println(AdmissionPrice(17))
	fmt.Println(AdmissionPrice(18))
	fmt.Println(AdmissionPrice(70))

	// Output:
	// 0
	// 0
	// 300
	// 300
	// 500
	// 500
}
