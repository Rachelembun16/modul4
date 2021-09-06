package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	var input int
	var strrupiah string
	var pound int
	var newstrrupiah string
	var x = true
	var yt string

	for x {
		fmt.Println("===JASA PENUKAEAN UANG===")
		fmt.Println("1. Penukaran Rupiah ke Dolar")
		fmt.Println("2. Penukaran Rupiah ke Euro")
		fmt.Println("3. Penukaran GBP ke Knut, Sickle, Galleon")
		fmt.Print("Masukkan pilihan: ")
		fmt.Scan(&input)

		switch input {
		case 1:
			fmt.Print("\nMasukkan nominal uang (tanpa ,00): Rp.")
			fmt.Scan(&strrupiah)
			fmt.Printf("Uang yang diinputkan sejumlah : Rp.%s,00", strrupiah)
			fmt.Print("\n \n")

			if strings.Contains(strrupiah, ".") {
				newstrrupiah = strings.Replace(strrupiah, ".", "", -1)
			} else {
				newstrrupiah = strrupiah
			}

			var rupiah, err = strconv.ParseFloat(newstrrupiah, 64)
			if err == nil {
				var usd = rupiah * 0.000070

				fmt.Printf("Penukaran Rp.%s,00 ke USD = $%.2f", strrupiah, usd)
			}

			fmt.Print("\nLanjutkan Program? (y/t)")
			fmt.Scan(&yt)

			yt = strings.ToLower(yt)

			if yt == "y" {
				x = true
			} else {
				x = false
			}

		case 2:
			fmt.Print("\nMasukkan nominal uang (tanpa ,00): Rp.")
			fmt.Scan(&strrupiah)
			fmt.Printf("Uang yang diinputkan sejumlah: Rp.%s,00", strrupiah)
			fmt.Print("\n \n")

			if strings.Contains(strrupiah, ".") {
				newstrrupiah = strings.Replace(strrupiah, ".", "", -1)
			} else {
				newstrrupiah = strrupiah
			}

			var rupiah, err = strconv.ParseFloat(newstrrupiah, 64)

			if err == nil {
				var euro = rupiah * 0.000059

				fmt.Printf("Penukaran Rp.%s,00 ke Euro = %.2f €", strrupiah, euro)
			}

			fmt.Print("\nLanjutkan Program? (y/t)")
			fmt.Scan(&yt)

			yt = strings.ToLower(yt)

			if yt == "y" {
				x = true
			} else {
				x = false
			}
		case 3:
			fmt.Print("\nMasukkan jumlah GB Pounds \t: ")
			fmt.Scan(&pound)

			var knut = pound * 100
			fmt.Printf("Jumlah knut yang didapat \t: %d\n", knut)

			var galleon = knut / (29 * 17)
			var sickle = (knut % galleon) / 29
			var sisa = knut - ((galleon * (29 * 17)) + sickle*29)

			fmt.Printf("Hasil penukaran mendapatkan \t: %d Galleon(s)\n", galleon)
			fmt.Printf("Sisa ditukar menjadi \t\t: %d Sickle(s)\n", sickle)
			fmt.Printf("Keping knut yang tersisa \t: %d Knut(s)", sisa)

			fmt.Print("\nLanjutkan Program? (y/t)")
			fmt.Scan(&yt)

			yt = strings.ToLower(yt)

			if yt == "y" {
				x = true
			} else {
				x = false
			}
		}
	}
}
