package main

import (
	"encoding/json"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
)

type Mahasiswa struct {
	NIM  string  `json:"nim"`
	Nama string  `json:"nama"`
	IPK  float64 `json:"ipk"`
}

const fileName = "mahasiswa.json"

// ===== File Handler =====

func loadData() []Mahasiswa {
	file, err := os.ReadFile(fileName)
	if err != nil {
		return []Mahasiswa{}
	}
	var data []Mahasiswa
	json.Unmarshal(file, &data)
	return data
}

func saveData(data []Mahasiswa) {
	jsonData, _ := json.MarshalIndent(data, "", "  ")
	os.WriteFile(fileName, jsonData, 0644)
}

// ===== API Handlers =====

func getMahasiswa(c *gin.Context) {
	data := loadData()
	c.JSON(http.StatusOK, data)
}

func getMahasiswaByNIM(c *gin.Context) {
	nim := c.Param("nim")
	data := loadData()

	for _, m := range data {
		if m.NIM == nim {
			c.JSON(http.StatusOK, m)
			return
		}
	}
	c.JSON(http.StatusNotFound, gin.H{"message": "Mahasiswa tidak ditemukan"})
}

func postMahasiswa(c *gin.Context) {
	var mhs Mahasiswa
	if err := c.ShouldBindJSON(&mhs); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Format JSON salah"})
		return
	}

	data := loadData()
	data = append(data, mhs)
	saveData(data)

	c.JSON(http.StatusCreated, gin.H{"message": "Mahasiswa ditambahkan"})
}

func deleteMahasiswa(c *gin.Context) {
	nim := c.Param("nim")
	data := loadData()

	for i, m := range data {
		if m.NIM == nim {
			data = append(data[:i], data[i+1:]...)
			saveData(data)
			c.JSON(http.StatusOK, gin.H{"message": "Mahasiswa dihapus"})
			return
		}
	}

	c.JSON(http.StatusNotFound, gin.H{"message": "Mahasiswa tidak ditemukan"})
}

func main() {
	router := gin.Default()

	router.GET("/mahasiswa", getMahasiswa)
	router.GET("/mahasiswa/:nim", getMahasiswaByNIM)
	router.POST("/mahasiswa", postMahasiswa)
	router.DELETE("/mahasiswa/:nim", deleteMahasiswa)

	router.Run(":9090")
}
