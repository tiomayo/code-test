package main

import (
	"bufio"
	"fmt"
	"log"
	"math"
	"os"
	"strings"
	"time"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	for {
		fmt.Print(" 1. matched string\n 2. hitung kembalian\n 3. validate bracket\n 4. validasi cuti\nSoal nomor berapa? ")
		scanner.Scan()
		// Holds the string that was scanned
		text := scanner.Text()
		handleInput(text)

		fmt.Print("Lanjutkan? (y/n) ")
		var w1 string
		_, err := fmt.Scanln(&w1)
		if err != nil {
			log.Fatal(err)
		}
		if w1 == "n" {
			fmt.Println("Exiting! Thank you.")
			break
		}
	}

	// fmt.Println("Soal no 1")
	// matchedString(4, []string{"abcd", "acbd", "aaab", "acbd"})

	// matchedString(5, []string{"pisang", "goreng", "enak", "sekali", "rasanya"})

	// matchedString(11, []string{"Satu", "Sate", "Tujuh", "Tusuk", "Tujuh", "Sate", "Bonus", "Tiga", "Puluh", "Tujuh", "Tusuk"})

	// fmt.Println("\nSoal no 2")
	// // Soal nomor 2
	// findChange(700649, 800000)
	// findChange(657650, 600000)
	// findChange(575650, 580000)

	// fmt.Println("\nSoal no 3")
	// // return true
	// stringValidate("{{[<>[{{}}]]}}")
	// stringValidate("[{}<>]")
	// // return false
	// stringValidate("]")
	// stringValidate("[{}<[>]")

	// fmt.Println("\nSoal no 4")
	// validasiCuti(7, "2021-05-01", "2021-07-05", 1)
	// validasiCuti(7, "2021-05-01", "2021-11-05", 3)
	// validasiCuti(7, "2021-01-05", "2021-12-18", 1)
	// validasiCuti(7, "2021-01-05", "2021-12-18", 3)
}

func handleInput(text string) {
	switch text {
	case "1":
		fmt.Println("Soal nomor 1: algoritma untuk mencocokan semua string satu sama lain dan mengeluarkan nomor string yang cocok satu sama lain")
		fmt.Print("Jumlah string:")
		var w1 int
		var w2 []string
		_, err := fmt.Scanln(&w1)
		if err != nil {
			log.Fatal(err)
		}
		for i := 0; i < w1; i++ {
			var tmp string
			fmt.Printf("Masukkan string ke-%v: ", i+1)
			_, err := fmt.Scanln(&tmp)
			if err != nil {
				log.Fatal(err)
			}
			w2 = append(w2, tmp)
		}

		matchedString(w1, w2)
	case "2":
		fmt.Println("Soal nomor 2: menghitung nilai uang kembalian beserta dengan pecahan uang yang bisa diberikan")
		fmt.Print("Total Belanja: ")
		var w1, w2 int
		_, err := fmt.Scanln(&w1)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Print("Pembeli membayar: ")
		_, err = fmt.Scanln(&w2)
		if err != nil {
			log.Fatal(err)
		}

		findChange(float64(w1), float64(w2))
	case "3":
		fmt.Println("Soal nomor 3: fungsi validasi untuk string")
		fmt.Print("Input: ")
		var w1 string
		_, err := fmt.Scanln(&w1)
		if err != nil {
			log.Fatal(err)
		}

		stringValidate(w1)
	case "4":
		fmt.Println("Soal nomor 4: function untuk membantu menentukan apakah seorang karyawan boleh mengambil cuti pribadi atau tidak")
		var w1, w4 int
		var w2, w3 string
		fmt.Print("Jumlah Cuti Bersama: ")
		_, err := fmt.Scanln(&w1)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Print("Tanggal join karyawan: ")
		_, err = fmt.Scanln(&w2)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Print("Tanggal rencana cuti: ")
		_, err = fmt.Scanln(&w3)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Print("Durasi cuti (hari): ")
		_, err = fmt.Scanln(&w4)
		if err != nil {
			log.Fatal(err)
		}

		validasiCuti(w1, w2, w3, w4)
	default:
		fmt.Println("Masukkan nomor soal (1-4)")
	}
}

func matchedString(length int, list []string) {
	strIndex := make(map[string][]int, length)
	resp := ""

	for i := 0; i < length; i++ {
		// lower to handle case insensitive
		lowered := strings.ToLower(list[i])
		if _, ok := strIndex[lowered]; ok && resp == "" {
			resp = lowered
		}
		strIndex[lowered] = append(strIndex[lowered], i+1)
	}

	if len(resp) == 0 {
		fmt.Println(false)
		return
	}

	fmt.Println(strIndex[resp])
	return
}

func findChange(totalPurchases, totalPayment float64) {
	if totalPayment < totalPurchases {
		fmt.Println("false, kurang bayar")
		return
	}

	totalChange := totalPayment - totalPurchases
	rounded := math.Floor(totalChange/100) * 100

	fmt.Printf("Kembalian yang harus diberikan kasir: %v, dibulatkan menjadi %v\n", totalChange, rounded)

	pecahan := []float64{100000, 50000, 20000, 10000, 5000, 2000, 1000, 500, 200, 100}

	for i := 0; i < len(pecahan); i++ {
		if rounded >= pecahan[i] {
			switch pecahan[i] {
			case 500, 200, 100:
				fmt.Printf("%v koin %v\n", int(rounded/pecahan[i]), pecahan[i])
			default:
				fmt.Printf("%v lembar %v\n", int(rounded/pecahan[i]), pecahan[i])
			}

			rounded = math.Mod(rounded, pecahan[i])
		}
	}
	return
}

func stringValidate(input string) {
	if len(input) == 0 || len(input)%2 == 1 {
		fmt.Println(false)
		return
	}

	bracket := map[rune]rune{
		'<': '>',
		'{': '}',
		'[': ']',
	}
	stack := []rune{}

	for _, r := range input {
		if _, ok := bracket[r]; ok {
			stack = append(stack, r)
		} else if len(stack) == 0 || bracket[stack[len(stack)-1]] != r {
			fmt.Println(false)
			return
		} else {
			stack = stack[:len(stack)-1]
		}
	}

	fmt.Println(len(stack) == 0)
	return
}

func validasiCuti(cutiBersama int, joinDate, cuti string, durasi int) {
	if durasi > 3 {
		fmt.Println(false)
		fmt.Println("Alasan: Karena maksimal cuti 3 hari")
		return
	}
	tanggalJoin, err := time.Parse("2006-01-02", joinDate)
	if err != nil {
		fmt.Println("Error parsing date")
		return
	}

	tanggalCuti, err := time.Parse("2006-01-02", cuti)
	if err != nil {
		fmt.Println("Error parsing date")
		return
	}

	diff := tanggalCuti.Sub(tanggalJoin)

	if diff.Hours()/24 < 180 {
		fmt.Println(false)
		fmt.Println("Alasan: Karena belum 180 hari sejak tanggal join karyawan")
		return
	}

	endOfYear := time.Date(tanggalJoin.Year(), time.December, 31, 0, 0, 0, 0, time.UTC)
	totalDays := endOfYear.Sub(tanggalJoin.AddDate(0, 0, 180)).Hours() / 24
	availableCuti := math.Floor(totalDays / 365 * float64(14-cutiBersama))
	if durasi > int(availableCuti) {
		fmt.Println(false)
		fmt.Printf("Alasan:  Karena hanya boleh mengambil %v hari cuti\n", availableCuti)
		return
	}

	fmt.Println(true)
	return
}
