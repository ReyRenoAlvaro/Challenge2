package main

import (
	"fmt"
)

func main() {
	// Data Uji
	// Data 1
	skorLumba1 := []int{96, 108, 89}
	skorKoala1 := []int{88, 91, 110}

	// Data Bonus 1
	skorLumba2 := []int{97, 112, 101}
	skorKoala2 := []int{109, 95, 123}

	// Data Bonus 2
	skorLumba3 := []int{97, 112, 101}
	skorKoala3 := []int{109, 95, 106}

	// Menghitung rata-rata skor
	RataRata := func(scores []int) float64 {
		sum := 0
		for _, score := range scores {
			sum += score
		}
		return float64(sum) / float64(len(scores))
	}

	// Menentukan pemenang berdasarkan skor rata-rata
	pemenang := func(skorLumbaLumba, skorkoala []int) string {
		rataRataLumba := RataRata(skorLumbaLumba)
		rataRataKoala := RataRata(skorkoala)

		// Skor minimum
		skorMinimum := 100

		if rataRataLumba >= float64(skorMinimum) && rataRataKoala >= float64(skorMinimum) {
			if rataRataLumba > rataRataKoala {
				return "Tim Lumba-lumba memenangkan kompetisi"
			} else if rataRataLumba < rataRataKoala {
				return "Tim Koala memenangkan kompetisi"
			} else {
				return "Kompetisi berakhir seri"
			}
		} else {
			if rataRataLumba >= float64(skorMinimum) {
				return "Tim Lumba-lumba memenangkan kompetisi"
			} else if rataRataKoala >= float64(skorMinimum) {
				return "Tim Koala memenangkan kompetisi"
			} else {
				return "Tidak ada pemenang karena kedua tim tidak mencapai skor minimum."
			}
		}
	}

	// Hasil Data 1
	fmt.Println("Data 1:")
	rataRataLumba1 := RataRata(skorLumba1)
	rataRataKoala1 := RataRata(skorKoala1)
	fmt.Printf("Rata-rata Skor Tim Lumba-lumba Adalah : %.2f\n", rataRataLumba1)
	fmt.Printf("Rata-rata Skor Tim Koala Adalah : %.2f\n", rataRataKoala1)
	fmt.Println("Hasil Pertandingan:", pemenang(skorLumba1, skorKoala1))

	// Hasil Data Bonus 1
	fmt.Println("\nData Bonus 1:")
	rataRataLumba2 := RataRata(skorLumba2)
	rataRataKoala2 := RataRata(skorKoala2)
	fmt.Printf("Rata-rata Skor Tim Lumba-lumba Adalah : %.2f\n", rataRataLumba2)
	fmt.Printf("Rata-rata Skor Tim Koala Adalah : %.2f\n", rataRataKoala2)
	fmt.Println("Hasil Pertandingan Adalah :", pemenang(skorLumba2, skorKoala2))

	// Hasil Data Bonus 2
	fmt.Println("\nData Bonus 2:")
	rataRataLumba3 := RataRata(skorLumba3)
	rataRataKoala3 := RataRata(skorKoala3)
	fmt.Printf("Rata-rata Skor Tim Lumba-lumba Adalah : %.2f\n", rataRataLumba3)
	fmt.Printf("Rata-rata Skor Tim Koala Adalah : %.2f\n", rataRataKoala3)
	fmt.Println("Hasil Pertandingan Adalah :", pemenang(skorLumba3, skorKoala3))
}
