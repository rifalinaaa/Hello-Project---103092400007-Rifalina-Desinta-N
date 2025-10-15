package main

import (
	"fmt"
)

// Konstanta maksimal jumlah mahasiswa dan mata kuliah
const MAX_MAHASISWA = 100
const MAX_MATAKULIAH = 10

// Struktur data MataKuliah menyimpan detail mata kuliah tiap mahasiswa
type MataKuliah struct {
	Kode, Nama, Grade     string
	SKS                   int
	UTS, UAS, Quiz, Total float64
	Ada                   bool // Menandakan jika data mata kuliah tersedia
}

// Struktur data Mahasiswa menyimpan data mahasiswa dan daftar mata kuliah yang diambil
type Mahasiswa struct {
	NIM        string
	Nama       string
	MataKuliah [MAX_MATAKULIAH]MataKuliah
	JumlahMK   int
	Ada        bool // Menandakan jika data mahasiswa tersedia
	TotalSKS   int
	TotalNilai float64 // Total nilai berbobot sks
}

// Variabel global menyimpan data seluruh mahasiswa dan jumlah yang tercatat
var dataMahasiswa [MAX_MAHASISWA]Mahasiswa
var jumlahMahasiswa int = 0

// Fungsi pencarian sequential berdasarkan NIM, mengembalikan index mahasiswa atau -1 jika tidak ditemukan
func sequentialSearch(nim string) int {
	for i := 0; i < jumlahMahasiswa; i++ {
		if dataMahasiswa[i].Ada && dataMahasiswa[i].NIM == nim {
			return i
		}
	}
	return -1
}

// Fungsi insertion sort mahasiswa berdasarkan NIM secara ascending
func insertionSortByNIM() {
	for i := 1; i < jumlahMahasiswa; i++ {
		j := i
		for j > 0 && dataMahasiswa[j-1].NIM > dataMahasiswa[j].NIM {
			temp := dataMahasiswa[j]
			dataMahasiswa[j] = dataMahasiswa[j-1]
			dataMahasiswa[j-1] = temp
			j--
		}
	}
}

// Fungsi selection sort mahasiswa berdasarkan TotalNilai secara descending
func selectionSortByNilai() {
	for i := 0; i < jumlahMahasiswa-1; i++ {
		maxIdx := i
		for j := i + 1; j < jumlahMahasiswa; j++ {
			if dataMahasiswa[j].Ada && dataMahasiswa[j].TotalNilai > dataMahasiswa[maxIdx].TotalNilai {
				maxIdx = j
			}
		}
		temp := dataMahasiswa[i]
		dataMahasiswa[i] = dataMahasiswa[maxIdx]
		dataMahasiswa[maxIdx] = temp
	}
}

// Fungsi menghitung Total nilai dan menentukan Grade berdasarkan bobot UTS, UAS, Quiz
func hitungTotalDanGrade(mk *MataKuliah) {
	mk.Total = (mk.UTS * 0.3) + (mk.UAS * 0.4) + (mk.Quiz * 0.3)
	if mk.Total >= 85 {
		mk.Grade = "A"
	} else if mk.Total >= 70 {
		mk.Grade = "B"
	} else if mk.Total >= 55 {
		mk.Grade = "C"
	} else if mk.Total >= 40 {
		mk.Grade = "D"
	} else {
		mk.Grade = "E"
	}
}

// Fungsi menambah mahasiswa baru ke array dataMahasiswa
func tambahMahasiswa(nim, nama string) {
	if jumlahMahasiswa >= MAX_MAHASISWA {
		fmt.Println("Data mahasiswa penuh")
		return
	}
	if sequentialSearch(nim) != -1 {
		fmt.Println("Mahasiswa dengan NIM tersebut sudah ada")
		return
	}
	dataMahasiswa[jumlahMahasiswa] = Mahasiswa{NIM: nim, Nama: nama, Ada: true}
	jumlahMahasiswa++
	fmt.Println("Mahasiswa berhasil ditambahkan.")
}

// Fungsi menghapus mahasiswa berdasarkan NIM, hanya menandai Ada menjadi false
func hapusMahasiswa(nim string) {
	idx := sequentialSearch(nim)
	if idx == -1 {
		fmt.Println("Mahasiswa tidak ditemukan")
		return
	}
	dataMahasiswa[idx].Ada = false
	fmt.Println("Mahasiswa berhasil dihapus.")
}

// Fungsi menambah mata kuliah ke mahasiswa tertentu berdasarkan NIM
func tambahMatakuliah(nim, kode, nama string, sks int, uts, uas, quiz float64) {
	idx := sequentialSearch(nim)
	if idx == -1 {
		fmt.Println("Mahasiswa tidak ditemukan")
		return
	}
	mhs := &dataMahasiswa[idx]
	if mhs.JumlahMK >= MAX_MATAKULIAH {
		fmt.Println("Mata kuliah penuh")
		return
	}
	mk := &mhs.MataKuliah[mhs.JumlahMK]
	mk.Kode, mk.Nama, mk.SKS, mk.UTS, mk.UAS, mk.Quiz = kode, nama, sks, uts, uas, quiz
	mk.Ada = true
	hitungTotalDanGrade(mk)
	mhs.JumlahMK++
	mhs.TotalSKS += sks
	// Penambahan total nilai berbobot sks
	mhs.TotalNilai += mk.Total * float64(sks)
	fmt.Println("Matakuliah berhasil ditambahkan.")
}

// Fungsi rekursif menampilkan transkrip mata kuliah mahasiswa dari index ke idx sampai selesai
func tampilTranskripRekursif(mhs Mahasiswa, idx int) {
	if idx >= mhs.JumlahMK {
		return
	}
	mk := mhs.MataKuliah[idx]
	if mk.Ada {
		fmt.Printf("  %s (%s) - SKS: %d, Nilai: %.2f, Grade: %s\n", mk.Nama, mk.Kode, mk.SKS, mk.Total, mk.Grade)
	}
	tampilTranskripRekursif(mhs, idx+1)
}

// Fungsi menampilkan transkrip mahasiswa berdasarkan NIM
func tampilTranskrip(nim string) {
	idx := sequentialSearch(nim)
	if idx == -1 {
		fmt.Println("Mahasiswa tidak ditemukan")
		return
	}
	mhs := dataMahasiswa[idx]
	fmt.Println("Transkrip Mahasiswa:", mhs.Nama, "(", mhs.NIM, ")")
	tampilTranskripRekursif(mhs, 0)
}

// Fungsi utama program menampilkan menu interaktif dan menjalankan fitur sesuai pilihan
func main() {
	for {
		fmt.Println("\nMenu:")
		fmt.Println("1. Tambah Mahasiswa")
		fmt.Println("2. Edit Mahasiswa")
		fmt.Println("3. Hapus Mahasiswa")
		fmt.Println("4. Tampil Transkrip")
		fmt.Println("5. Urutkan dan Tampilkan Berdasarkan NIM")
		fmt.Println("6. Urutkan dan Tampilkan Berdasarkan Nilai")
		fmt.Println("0. Keluar")
		fmt.Print("Pilih menu: ")

		var menu int
		fmt.Scanln(&menu)

		if menu == 1 {
			// Perlu parameter input nim dan nama, fungsi belum terdefinisi untuk input
			// Contoh: panggil tambahMahasiswa dengan input manual
		} else if menu == 2 {
			// Fungsi editMahasiswa belum ada definisinya
		} else if menu == 3 {
			// meminta input nim mahasiswa yang mau dihapus
		} else if menu == 4 {
			// meminta input nim untuk tampilTranskrip
		} else if menu == 5 {
			insertionSortByNIM()
			fmt.Println("Data mahasiswa telah diurutkan berdasarkan NIM")
		} else if menu == 6 {
			selectionSortByNilai()
			fmt.Println("Data mahasiswa telah diurutkan berdasarkan total nilai")
		} else if menu == 0 {
			return // Keluar dari program
		} else {
			fmt.Println("Menu tidak valid.")
		}
	}
}

