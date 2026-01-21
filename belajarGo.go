package main
import "fmt"

func main2() {
    var nama string = "Iqbal"
    umur := 20
    fmt.Println("Halo", nama)
    fmt.Println("Umur", umur)
    type Mahasiswa struct {
    NIM string
    Nama string
    Jurusan string
    IPK float64
}

var m1 Mahasiswa
m1.NIM = "123456789"
m1.Nama = "Iqbal"
m1.Jurusan = "Teknik Informatika"
m1.IPK = 3.5
fmt.Println("NIM", m1.NIM)
fmt.Println("Nama", m1.Nama)
fmt.Println("Jurusan", m1.Jurusan)
fmt.Println("IPK", m1.IPK)
fmt.Println("\ndata mahasiswa")
fmt.Println("Halo", m1.Nama, "dengan NIM", m1.NIM, "jurusan", m1.Jurusan, "IPK", m1.IPK)

var m2 Mahasiswa
m2.NIM = "123456789"
m2.Nama = "Iqbal"
m2.Jurusan = "Teknik Informatika"
m2.IPK = 3.25
fmt.Println("NIM", m2.NIM)
fmt.Println("Nama", m2.Nama)
fmt.Println("Jurusan", m2.Jurusan)
fmt.Println("IPK", m2.IPK)
fmt.Println("\ndata mahasiswa")
fmt.Println("Halo", m2.Nama, "dengan NIM", m2.NIM, "jurusan", m2.Jurusan, "IPK", m2.IPK)

var m3 Mahasiswa
m3.NIM = "123456789"
m3.Nama = "Iqbal"
m3.Jurusan = "Teknik Informatika"
m3.IPK = 3.75
fmt.Println("NIM", m3.NIM)
fmt.Println("Nama", m3.Nama)
fmt.Println("Jurusan", m3.Jurusan)
fmt.Println("IPK", m3.IPK)
fmt.Println("\ndata mahasiswa")
fmt.Println("Halo", m3.Nama, "dengan NIM", m3.NIM, "jurusan", m3.Jurusan, "IPK", m3.IPK)



fmt.Println("\nrata rata ipk", (m1.IPK + m2.IPK + m3.IPK) / 3)
}
