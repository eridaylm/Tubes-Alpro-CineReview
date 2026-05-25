package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

// type alias sebagai pengganti tipe data bawaan Go 
// eel menggantikan int, jebb menggantikan string
type eel = int
type jebb = string

// batas maksimum film yang bisa disimpan di koleksi 
const MAXFILM = 100

// struct Film menyimpan semua informasi tentang satu film 
type Film struct {
	judul     jebb
	tahun     eel
	deskripsi jebb
	genre     jebb
	rating    float64
}

// variabel global: array film dan penghitung jumlah film aktif 
var daftarFilm [MAXFILM]Film
var jumlahFilm eel
var scanner = bufio.NewScanner(os.Stdin)

// membaca satu baris input teks dari pengguna 
func inputStr() jebb {
	scanner.Scan()
	return strings.TrimSpace(scanner.Text())
}

// membaca input angka bulat, lalu bersihkan sisa baris di buffer 
func inputInt() eel {
	var n eel
	fmt.Scan(&n)
	scanner.Scan()
	return n
}

// membaca input angka desimal, lalu bersihkan sisa baris di buffer 
func inputFloat() float64 {
	var f float64
	fmt.Scan(&f)
	scanner.Scan()
	return f
}

// memastikan nilai rating selalu berada di antara 0 sampai 10 
func klampRating(r float64) float64 {
	if r < 0 {
		return 0
	}
	if r > 10 {
		return 10
	}
	return r
}

// menampilkan banner judul aplikasi CineReview 
func tampilHeader() {
	fmt.Println("========================================")
	fmt.Println("             CineReview")
	fmt.Println("      Katalog & Rating Film Favorit")
	fmt.Println("========================================")
}

// menampilkan pilihan-pilihan yang tersedia di menu utama 
func tampilMenu() {
	fmt.Println("\n========== MENU UTAMA ==========")
	fmt.Println(" 1. Tambah Film")
	fmt.Println(" 2. Tampilkan Semua Film")
	fmt.Println(" 3. Ubah Data Film")
	fmt.Println(" 4. Hapus Film")
	fmt.Println(" 5. Cari Film")
	fmt.Println(" 6. Urutkan Film")
	fmt.Println(" 7. Statistik Koleksi")
	fmt.Println(" 0. Keluar")
	fmt.Println("================================")
	fmt.Print("Pilih menu: ")
}

// menampilkan detail lengkap satu film beserta nomor urutnya 
func tampilFilm(f Film, no eel) {
	fmt.Printf("\n[%d] %-30s (%d)\n", no, f.judul, f.tahun)
	fmt.Printf("    Genre    : %s\n", f.genre)
	fmt.Printf("    Rating   : %.1f/10\n", f.rating)
	fmt.Printf("    Deskripsi: %s\n", f.deskripsi)
}

// menampilkan seluruh film yang ada di koleksi 
func tampilSemuaFilm() {
	if jumlahFilm == 0 {
		fmt.Println("Koleksi film masih kosong.")
		return
	}
	fmt.Println("\n======= DAFTAR KOLEKSI FILM =======")
	for i := 0; i < jumlahFilm; i++ {
		tampilFilm(daftarFilm[i], i+1)
	}
}

// meminta input data film baru lalu menyimpannya ke koleksi 
func tambahFilm() {
	if jumlahFilm >= MAXFILM {
		fmt.Println("Koleksi sudah penuh!")
		return
	}
	fmt.Println("\n===== TAMBAH FILM =====")
	var f Film
	fmt.Print("Judul Film    : ")
	f.judul = inputStr()
	fmt.Print("Tahun Rilis   : ")
	f.tahun = inputInt()
	fmt.Print("Genre         : ")
	f.genre = inputStr()
	fmt.Print("Deskripsi     : ")
	f.deskripsi = inputStr()
	fmt.Print("Rating (0-10) : ")
	f.rating = klampRating(inputFloat())// klampRating dipakai agar rating tidak keluar dari rentang 0-10 
	daftarFilm[jumlahFilm] = f
	jumlahFilm++
	fmt.Printf("Film \"%s\" berhasil ditambahkan!\n", f.judul)
}