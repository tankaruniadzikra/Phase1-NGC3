// Looping 1
package main

import "fmt"

func main() {
	var manusia = []map[string]any{
		{"name": "Hank", "Age": 50, "Job": "Polisi"},
		{"name": "Heisenberg", "Age": 52, "Job": "Ilmuawan"},
		{"name": "Skyler", "Age": 48, "Job": "Akuntan"},
	}

	for _, orang := range manusia {

		fmt.Printf("Hi Perkenalkan, Nama saya %v, umur saya %v, dan saya bekerja sebagai %v\n", orang["name"], orang["Age"], orang["Job"])
	}
}

// Looping 2
package main

import "fmt"

//menghitung rata-rata dari elemen dalam slice float64, dan mengembalikan hasilnya sebagai angka desimal.
func hitungRataRata(slice []float64) float64 {
	//tipe data dari variabel total adalah float64, sehingga menggunakan 0.0
	total := 0.0
	for _, nilai := range slice {
		total += nilai
	}
	return total / float64(len(slice))
}

func hitungJumlah(slice []float64) float64 {
	total := 0.0
	for _, nilai := range slice {
		total += nilai
	}
	return total
}

func hitungMedian(slice []float64) float64 {
	panjang := len(slice)
	if panjang%2 == 0 {
		return (slice[panjang/2-1] + slice[panjang/2]) / 2
	} else {
		return slice[panjang/2]
	}
}

func main() {
	slice1 := []float64{1, 5, 7, 8, 10, 24, 33}
	slice2 := []float64{1.1, 5.4, 6.7, 9.2, 11.3, 25.2, 33.1}

	rataRata1 := hitungRataRata(slice1)
	jumlah1 := hitungJumlah(slice1)
	median1 := hitungMedian(slice1)

	rataRata2 := hitungRataRata(slice2)
	jumlah2 := hitungJumlah(slice2)
	median2 := hitungMedian(slice2)

	fmt.Printf("Slice 1:\n")
	fmt.Printf("Rata-rata: %.2f\n", rataRata1)
	fmt.Printf("Jumlah: %.2f\n", jumlah1)
	fmt.Printf("Median: %.2f\n\n", median1)

	fmt.Printf("Slice 2:\n")
	fmt.Printf("Rata-rata: %.2f\n", rataRata2)
	fmt.Printf("Jumlah: %.2f\n", jumlah2)
	fmt.Printf("Median: %.2f\n", median2)
}

// Logic 1 - Palindrome
package main

import "fmt"

//menerima satu parameter yang merupakan string (input) dan mengembalikannya ke boolean.
func isPalindrome(input string) bool {
	//membandingkan karakter pertama dengan karakter terakhir, karakter kedua dengan karakter kedua terakhir, dan seterusnya hingga mencapai tengah dari string.
	for i := 0; i < len(input)/2; i++ {
		if input[i] != input[len(input)-i-1] {
			return false
		}
	}
	return true
}

func main() {
	fmt.Println(isPalindrome("kodok"))
	fmt.Println(isPalindrome("katak"))
}

// Logic 2 - XOXO
package main

import "fmt"

func main() {
	kata := "xoxo"

	jumlahX := 0
	jumlahO := 0

	for _, huruf := range kata {
		if huruf == 'x' {
			jumlahX++
		} else if huruf == 'o' {
			jumlahO++
		}
	}

	apakahSama := jumlahX == jumlahO

	fmt.Printf("Jumlah karakter 'x': %d\n", jumlahX)
	fmt.Printf("Jumlah karakter 'o': %d\n", jumlahO)
	fmt.Printf("Jumlah 'x' sama dengan jumlah 'o': %t\n", apakahSama)
}

// Logic 3 - XOXO
package main

import "fmt"

func main() {
	// Contoh slice int
	slice := []int{5, 5, 6, 4, 4, 2, 1, 3, 7}

	// Implementasi Bubble Sort
	// Dari besar ke kecil
	for i := 0; i < len(slice)-1; i++ {
		for j := 0; j < len(slice)-i-1; j++ {
			if slice[j] < slice[j+1] {
				slice[j], slice[j+1] = slice[j+1], slice[j]
			}
		}
	}

	// Dari kecil ke besar
	// for i := 0; i < len(slice)-1; i++ {
	// 	for j := 0; j < len(slice)-i-1; j++ {
	// 		if slice[j] > slice[j+1] {
	// 			slice[j], slice[j+1] = slice[j+1], slice[j]
	// 		}
	// 	}
	// }

	// Cetak slice yang sudah diurutkan
	fmt.Println(slice)
}

// Logic 4 - Asteriks Level 1
package main

import "fmt"

func main() {
	// Jumlah baris yang ingin ditampilkan
	jumlahBaris := 5

	// Loop untuk menampilkan barisan bintang
	for i := 0; i < jumlahBaris; i++ {
		fmt.Println("*")
	}
}

// Logic 5 - Asteriks Level 2
package main

import "fmt"

func main() {
	rows2 := 5

	for i := 0; i < rows2; i++ {
		for j := 0; j < rows2; j++ {
			fmt.Print("*")
		}
		fmt.Println()
	}
}

// Logic 6 - Asteriks level 3

package main

import "fmt"

func main() {
	// Jumlah baris yang ingin ditampilkan
	jumlahBaris := 5

	// Loop untuk mengatur jumlah baris
	for i := 1; i <= jumlahBaris; i++ {
		// Loop untuk menampilkan bintang pada setiap baris
		for j := 1; j <= i; j++ {
			fmt.Print("*")
		}
		fmt.Println()
	}
}

// Logic 7 - Asteriks Level 4
package main

import "fmt"

func main() {
	rows4 := 5

	for i := 0; i < rows4; i++ {
		// Loop untuk menampilkan spasi sebelum bintang
		for j := 0; j < i; j++ {
			fmt.Print(" ")
		}

		// Loop untuk menampilkan bintang pada setiap baris
		for k := 0; k < rows4-i; k++ {
			fmt.Print("*")
		}
		fmt.Println()
	}
}
