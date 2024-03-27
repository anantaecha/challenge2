package main

import (
	"fmt"
)

func calculateAverageScore(scores []int) float64 {
	sum := 0
	for _, score := range scores {
		sum += score
	}
	return float64(sum) / float64(len(scores))
}

func determineWinner(lumbaAverage, koalaAverage float64) string {
	if lumbaAverage >= 100 && koalaAverage >= 100 {
		if lumbaAverage > koalaAverage {
			return "Lumba-lumba"
		} else if lumbaAverage < koalaAverage {
			return "Koala"
		} else {
			return "Seri"
		}
	} else if lumbaAverage >= 100 && koalaAverage < 100 {
		return "Lumba-lumba"
	} else if koalaAverage >= 100 && lumbaAverage < 100 {
		return "Koala"
	} else {
		return "Tidak ada pemenang"
	}
}

func main() {
	data := map[string][]int{
		"Data 1":      {96, 108, 89, 88, 91, 110},
		"Data Bonus 1": {97, 112, 101, 109, 95, 123},
		"Data Bonus 2": {97, 112, 101, 109, 95, 106},
	}

	for key, scores := range data {
		lumbaScores := scores[:3]
		koalaScores := scores[3:]

		lumbaAverage := calculateAverageScore(lumbaScores)
		koalaAverage := calculateAverageScore(koalaScores)

		winner := determineWinner(lumbaAverage, koalaAverage)

		fmt.Println(key)
		fmt.Println("Skor rata-rata Lumba-lumba:", lumbaAverage)
		fmt.Println("Skor rata-rata Koala:", koalaAverage)
		fmt.Println("Pemenang:", winner)
		fmt.Println()
	}
}