package main

import (
	"fmt"
)

func main() {
	fmt.Print("\n--------- Program Menghitung Tip ---------\n")

	// Nilai tagihan yang sudah terinisialisasi
	tagihan := []float64{275, 40, 430}

	for _, nilaiTagihan := range tagihan {
		// Hitung tip berdasarkan kondisi if/else
		var tip float64

		if nilaiTagihan >= 50 && nilaiTagihan <= 300 {
			tip = nilaiTagihan * 0.15
		} else {
			tip = nilaiTagihan * 0.20
		}

		// Hitung total nilai (tagihan + tip)
		nilaiakhir := nilaiTagihan + tip

		// Menampilkan hasil di konsol
		fmt.Printf("\n Tagihan 	: %.2f\n Tip 		: %.2f\n Total Tagihan 	: %.2f\n\n", nilaiTagihan, tip, nilaiakhir)
	}

	fmt.Print("-------------- Terima Kasih --------------\n\n")
}
