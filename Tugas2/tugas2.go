package main

import (
	"fmt"
)

func main() {
	// Data 1
	var beratMark1, tinggiMark1 float64 = 78.0, 1.69 // kg, meter
	var beratJohn1, tinggiJohn1 float64 = 92.0, 1.95 // kg, meter

	// Data 2
	var beratMark2, tinggiMark2 float64 = 95.0, 1.88 // kg, meter
	var beratJohn2, tinggiJohn2 float64 = 85.0, 1.76 // kg, meter

	// Hitung BMI untuk kedua data
	BMIMark1 := hitungBMI(beratMark1, tinggiMark1)
	BMIJohn1 := hitungBMI(beratJohn1, tinggiJohn1)

	BMIMark2 := hitungBMI(beratMark2, tinggiMark2)
	BMIJohn2 := hitungBMI(beratJohn2, tinggiJohn2)

	// Menentukan apakah Mark memiliki BMI lebih tinggi daripada John untuk kedua data
	markHigherBMI1 := BMIMark1 > BMIJohn1
	markHigherBMI2 := BMIMark2 > BMIJohn2

	//  Hasil Data 1
	fmt.Println("Data 1:")
	fmt.Printf("BMI Mark: %.2f\n", BMIMark1)
	fmt.Printf("BMI John: %.2f\n", BMIJohn1)
	fmt.Printf("Mark memiliki BMI lebih tinggi daripada John: %t\n", markHigherBMI1)

	// Hasil Data 2
	fmt.Println("\nData 2:")
	fmt.Printf("BMI Mark: %.2f\n", BMIMark2)
	fmt.Printf("BMI John: %.2f\n", BMIJohn2)
	fmt.Printf("Mark memiliki BMI lebih tinggi daripada John: %t\n", markHigherBMI2)
}

// Menghitung BMI
func hitungBMI(berat, tinggi float64) float64 {
	return berat / (tinggi * tinggi)
}
