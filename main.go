package main
import (
	"encoding/json"
	"fmt"
	"os"
)

type Mahasiswa struct {
	NIM  string  `json:"nim"`
	Nama string  `json:"nama"`
	IPK  float64 `json:"ipk"`
}

const fileName = "mahasiswa.json"

// Load data dari file JSON
func loadData() []Mahasiswa {
	file, err := os.ReadFile(fileName)
	if err != nil {
		return []Mahasiswa{} // jika file belum ada
	}

	var data []Mahasiswa
	json.Unmarshal(file, &data)
	return data
}

// Simpan data ke file JSON
func saveData(data []Mahasiswa) {
	jsonData, _ := json.MarshalIndent(data, "", "  ")
	os.WriteFile(fileName, jsonData, 0644)
}

// Tambah mahasiswa
func tambahMahasiswa(data []Mahasiswa, m Mahasiswa) []Mahasiswa {
	data = append(data, m)
	return data
}

// Tampilkan data
func tampilMahasiswa(data []Mahasiswa) {
	fmt.Println("\n=== DATA MAHASISWA ===")
	for i, m := range data {
		fmt.Printf("%d. %s - %s (IPK %.2f)\n", i+1, m.NIM, m.Nama, m.IPK)
	}
}
func DeleteMahasiswaByNim(data []Mahasiswa, nim string) []Mahasiswa {
	for i, m := range data {
		if m.NIM == nim {
			data = append(data[:i], data[i+1:]...)
			fmt.Println("Mahasiswa dengan NIM", nim, "berhasil dihapus")
			return data
		}
	}
	fmt.Println("Mahasiswa dengan NIM", nim, "tidak ditemukan")
	return data
}

func main() {
	data := loadData()
	data = DeleteMahasiswaByNim(data, "2341720003")
	saveData(data)

	// Tampilkan
	tampilMahasiswa(data)

	fmt.Println("\nData berhasil disimpan ke mahasiswa.json")
}
