package main

import (
	"fmt"
	"strconv"
)

func main() {
	fmt.Print("\n--------- Program Menghitung Tip ---------\n")

	for {
		// Input nilai tagihan
		var tagihanStr string
		fmt.Print("\n------------------------------------------\n")
		fmt.Print("\nMasukan Nilai Tagihan : ")
		_, err := fmt.Scan(&tagihanStr)
		if err != nil {
			fmt.Println("Error: Masukan harus berupa angka.")
			continue
		}

		tagihan, err := strconv.ParseFloat(tagihanStr, 64)
		if err != nil {
			fmt.Println("Error: Masukan harus berupa angka.")
			continue
		}

		// Hitung tip berdasarkan kondisi if/else
		var tip float64

		if tagihan >= 50 && tagihan <= 300 {
			tip = tagihan * 0.15
		} else {
			tip = tagihan * 0.20
		}

		// Hitung total nilai (tagihan + tip)
		nilaiakhir := tagihan + tip

		// Menampilkan hasil di konsol
		fmt.Printf("\n Tagihan 	: %.2f\n Tip 		: %.2f\n Total Tagihan 	: %.2f\n\n", tagihan, tip, nilaiakhir)

		// Tanyakan kepada pengguna apakah ingin melanjutkan
		var choice string
		fmt.Print("Ingin menghitung tip lagi? (ya/tidak): ")
		fmt.Scan(&choice)

		if choice != "ya" {
			fmt.Print("\n-------------- Terima Kasih --------------\n\n")
			break // Keluar dari loop jika pengguna tidak ingin melanjutkan
		}
	}
}
