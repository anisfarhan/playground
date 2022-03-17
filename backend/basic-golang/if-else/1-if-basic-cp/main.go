package main

import "fmt"

// Disini kalia coba melakukan if ketika kondisi nya sudah terpenuhi
// buatlah program yang akan menampilkan nilai dari setiap mahasiswa
// Jika nilai A = Cumlaude
// Jika nilai B = Lulus
// Jika nilai X = Tidak Lulus

// Kalian sudah tahukan cara mengakses value map pada Go meggunakan for-range kan?

func main() {
	name1 := "Zein Fahrozi"
	name2 := "Fabiansyah Raam"
	name3 := "Indra Kenz"

	mahasiswa := []map[string]string{
		{
			"name":  name1,
			"nilai": "A",
		},
		{
			"name":  name2,
			"nilai": "B",
		},
		{
			"name":  name3,
			"nilai": "X",
		},
	}

	// Output:
	/*
		Zein Fahrozi   Cumlaude
		Fabiansyah Raam   Lulus
		Indra Kenz   Tidak Lulus
	*/
	// TODO: answer here
	for _, mahasiswa := range mahasiswa {
		// if mahasiswa["nilai"] == "A" {
		// 	fmt.Printf("%s   Cumlaude\n", mahasiswa["name"])
		// } else if mahasiswa["nilai"] == "B" {
		// 	fmt.Printf("%s   Lulus\n", mahasiswa["name"])
		// } else {
		// 	fmt.Printf("%s   Tidak Lulus\n", mahasiswa["name"])
		// }

		switch mahasiswa["nilai"] {
		case "A":
			fmt.Printf("%s   Cumlaude\n", mahasiswa["name"])
		case "B":
			fmt.Printf("%s   Lulus\n", mahasiswa["name"])
		case "X":
			fmt.Printf("%s   Tidak Lulus\n", mahasiswa["name"])
		}
	}
}
