package main

import "fmt"

func main() {
	// Hitung Rata-Rata dari Array
	scores := [8]int{100, 80, 75, 92, 70, 93, 88, 67}
	average, goodScores := solveQuiz(scores)

	fmt.Printf("Rata-Rata nilai adalah %f\n", average)
	fmt.Printf("Nilai di atas 90 diantaranya: %v", goodScores)
}

func solveQuiz(scores [8]int) (rataRata float32, goodScores []int) {
	// Hitung Rata-Rata & Good Score
	jumlah := 0
	for _, score := range scores {
		jumlah += score
		if score >= 90 {
			goodScores = append(goodScores, score)
		}
	}
	rataRata = float32(jumlah) / float32(len(scores))
	return
}