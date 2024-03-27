package main

import (
	"fmt"
)

func calculateBMI(weight, height float64) float64 {
	return weight / (height * height)
}

func main() {
	// Data 1
	markWeight1 := 78.0 // kg
	markHeight1 := 1.69 // m
	johnWeight1 := 92.0 // kg
	johnHeight1 := 1.95 // m

	// Menghitung BMI Mark dan John
	markBMI1 := calculateBMI(markWeight1, markHeight1)
	johnBMI1 := calculateBMI(johnWeight1, johnHeight1)

	// Menampilkan hasil BMI Mark dan John
	fmt.Printf("Data 1:\n")
	fmt.Printf("BMI Mark: %.2f\n", markBMI1)
	fmt.Printf("BMI John: %.2f\n", johnBMI1)

	// Membandingkan BMI Mark dan John
	markHigherBMI := markBMI1 > johnBMI1

	// Menampilkan hasil perbandingan BMI
	fmt.Printf("Mark memiliki BMI lebih tinggi dari John: %t\n\n", markHigherBMI)

	// Data 2
	markWeight2 := 95.0 // kg
	markHeight2 := 1.88 // m
	johnWeight2 := 85.0 // kg
	johnHeight2 := 1.76 // m

	// Menghitung BMI Mark dan John
	markBMI2 := calculateBMI(markWeight2, markHeight2)
	johnBMI2 := calculateBMI(johnWeight2, johnHeight2)

	// Menampilkan hasil BMI Mark dan John
	fmt.Printf("Data 2:\n")
	fmt.Printf("BMI Mark: %.2f\n", markBMI2)
	fmt.Printf("BMI John: %.2f\n", johnBMI2)

	// Membandingkan BMI Mark dan John
	markHigherBMI = markBMI2 > johnBMI2

	// Menampilkan hasil perbandingan BMI
	fmt.Printf("Mark memiliki BMI lebih tinggi dari John: %t\n", markHigherBMI)
}