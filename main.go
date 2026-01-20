package main

import (
	"fmt"
)

type Mahasiswa struct {
    Nama string
    NIM string
    Jurusan string
    IPK float64
}

func hitungrataipk(data []Mahasiswa) float64 {
    total := 0.0
    for _, m := range data {
        total += m.IPK
    }
    return total / float64(len(data))
}

func tampilMahasiswa(data []Mahasiswa) {
	fmt.Println("=== DATA MAHASISWA ===")
	for i, m := range data {
		fmt.Printf("%d. %s - %s (IPK: %.2f)\n", i+1, m.NIM, m.Nama, m.IPK)
	}
}

func cariMahasiswaByNIM(data []Mahasiswa, nim string){
	for _, m := range data {
		if m.NIM == nim {
			fmt.Println("Mahasiswa ditemukan:", m,"nama ", m.Nama)
			return
		}
	}
	fmt.Println("Mahasiswa tidak ditemukan")
}


func main() {
	dataMahasiswa := []Mahasiswa{
		{"Iqbal", "123456781", "Teknik Informatika", 3.5},
		{"Iqbal", "123456782", "Teknik Informatika", 3.25},
		{"Iqbal", "123456783", "Teknik Informatika", 3.75},
	}

	tampilMahasiswa(dataMahasiswa)
	fmt.Println("Rata-rata IPK:", hitungrataipk(dataMahasiswa))
	cariMahasiswaByNIM(dataMahasiswa, "123456789")
}