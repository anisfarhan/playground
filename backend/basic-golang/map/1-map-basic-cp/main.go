package main

import "fmt"

// Buatlah map dengan key "Nama Provinsi" pada pulau Kalimantan dan value nya adalah "Ibu Kota" provinsi tersebut
// Output satu semua key dan value yang ada di map tersebut
func main() {
	// TODO: answer here
	ibukotaProvinsi := map[string]string{
		"Kalimantan Barat":   "Pontianak",
		"Kalimantan Tengah":  "Palangka Raya",
		"Kalimantan Selatan": "Banjarmasin",
		"Kalimantan Timur":   "Samarinda",
		"Kalimantan Utara":   "Tanjung Selor",
	}

	for key, value := range ibukotaProvinsi {
		fmt.Println("Key: ", key, " Value: ", value)
	}
}
