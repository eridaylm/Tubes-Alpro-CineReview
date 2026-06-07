package main

import (
	"bufio"
	"encoding/json"
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
	Judul     jebb    `json:"judul"`
	Tahun     eel     `json:"tahun"`
	Deskripsi jebb    `json:"deskripsi"`
	Genre     jebb    `json:"genre"`
	Rating    float64 `json:"rating"`
}

// variabel global: array film dan penghitung jumlah film aktif
var daftarFilm [MAXFILM]Film
var jumlahFilm eel
var dataJSON []Film
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
func tampilFilm(f Film, no eel) { //menampilkan detail lengkap sebuah film
	fmt.Printf("\n[%d] %-30s (%d)\n", no, f.Judul, f.Tahun)
	fmt.Printf("    Genre    : %s\n", f.Genre)
	fmt.Printf("    Rating   : %.1f/10\n", f.Rating)
	fmt.Printf("    Deskripsi: %s\n", f.Deskripsi)
}

// menampilkan seluruh film yang ada di koleksi
func tampilSemuaFilm() {
	if jumlahFilm == 0 {
		fmt.Println("Koleksi film masih kosong.")
		return
	}
	fmt.Println("\n======= DAFTAR KOLEKSI FILM =======")
	for i := 0; i < jumlahFilm; i++ {
		tampilFilm(daftarFilm[i], i+1) //perulangan seluruh film, loop akan berjalan dari indeks 0 sampai jumlah film yang aktif
	}
}

// meminta input data film baru lalu menyimpannya ke koleksi
func tambahFilm() {
	if jumlahFilm >= MAXFILM { //mengecek kapasitas penyimpanan
		fmt.Println("Koleksi sudah penuh!")
		return
	}
	fmt.Println("\n===== TAMBAH FILM =====")
	var f Film //membuat variabel film baru
	fmt.Print("Judul Film    : ")
	f.Judul = inputStr()
	fmt.Print("Tahun Rilis   : ")
	f.Tahun = inputInt()
	fmt.Print("Genre         : ")
	f.Genre = inputStr()
	fmt.Print("Deskripsi     : ")
	f.Deskripsi = inputStr()
	fmt.Print("Rating (0-10) : ")
	f.Rating = klampRating(inputFloat()) // klampRating dipakai agar rating tidak keluar dari rentang 0-10
	daftarFilm[jumlahFilm] = f
	jumlahFilm++ //menyimpan ke array
	simpanJSON()
	fmt.Printf("Film \"%s\" berhasil ditambahkan!\n", f.Judul)
}

// mengubah semua data film berdasarkan nomor yang dipilih pengguna
func ubahFilm() {
	if jumlahFilm == 0 { //mengecek apakah koleksi kosong
		fmt.Println("Koleksi film masih kosong.")
		return
	}
	tampilSemuaFilm()
	fmt.Print("\nNomor film yang ingin diubah: ")
	no := inputInt()
	// cek apakah nomor yang dimasukkan valid
	if no < 1 || no > jumlahFilm {
		fmt.Println("Nomor tidak valid!")
		return
	}
	idx := no - 1
	fmt.Printf("\nMengubah film: %s\n", daftarFilm[idx].Judul)
	var f Film
	fmt.Print("Judul baru    : ")
	f.Judul = inputStr()
	fmt.Print("Tahun baru    : ")
	f.Tahun = inputInt()
	fmt.Print("Genre baru    : ")
	f.Genre = inputStr()
	fmt.Print("Deskripsi baru: ")
	f.Deskripsi = inputStr()
	fmt.Print("Rating baru   : ")
	f.Rating = klampRating(inputFloat())
	daftarFilm[idx] = f
	simpanJSON()
	fmt.Println("Data film berhasil diubah!")
}

// menghapus film dari koleksi dan menggeser sisa elemen ke kiri
func hapusFilm() {
	if jumlahFilm == 0 {
		fmt.Println("Koleksi film masih kosong.")
		return
	}
	tampilSemuaFilm()
	fmt.Print("\nNomor film yang ingin dihapus: ")
	no := inputInt()
	if no < 1 || no > jumlahFilm {
		fmt.Println("Nomor tidak valid!")
		return
	}
	idx := no - 1
	judul := daftarFilm[idx].Judul
	// geser semua film di belakang index ke posisi sebelumnya
	for i := idx; i < jumlahFilm-1; i++ {
		daftarFilm[i] = daftarFilm[i+1]
	}
	jumlahFilm--
	simpanJSON()
	fmt.Printf("Film \"%s\" berhasil dihapus!\n", judul)
}

// menyalin isi daftarFilm ke array baru agar data asli aman
func salinFilm() ([MAXFILM]Film, eel) {
	var salinan [MAXFILM]Film
	for i := 0; i < jumlahFilm; i++ {
		salinan[i] = daftarFilm[i]
	}
	return salinan, jumlahFilm
}

// mencari film yang judulnya mengandung kata kunci, tidak peduli menggunakan huruf besar/kecil
func seqSearchJudul(kata jebb) []eel {
	hasil := []eel{}
	for i := 0; i < jumlahFilm; i++ {
		// strings.Contains dipakai agar pencarian parsial tetap cocok
		if strings.Contains(strings.ToLower(daftarFilm[i].Judul), strings.ToLower(kata)) {
			hasil = append(hasil, i)
		}
	}
	return hasil
}

// Selection Sort mengurutkan film dari rating tertinggi ke terendah
func selectionSortRating(arr *[MAXFILM]Film, n eel) {
	// memilih elemen dengan rating terbesar lalu tukar ke posisi yang benar
	for i := 0; i < n-1; i++ {
		maxIdx := i
		for j := i + 1; j < n; j++ {
			if arr[j].Rating > arr[maxIdx].Rating {
				maxIdx = j
			}
		}
		arr[i], arr[maxIdx] = arr[maxIdx], arr[i]
	}
}

// Insertion Sort mengurutkan film berdasarkan tahun rilis dari terlama ke terbaru
func insertionSortTahun(arr *[MAXFILM]Film, n eel) {
	// menyisipkan setiap elemen ke posisi yang tepat secara berurutan
	for i := 1; i < n; i++ {
		kunci := arr[i]
		j := i - 1
		// menggeser elemen ke kanan selama tahunnya lebih besar dari kunci
		for j >= 0 && arr[j].Tahun > kunci.Tahun {
			arr[j+1] = arr[j]
			j--
		}
		arr[j+1] = kunci
	}
}

// binarysearch cari film berdasarkan judul
func binarySearchJudul(arr [MAXFILM]Film, n eel, kata jebb) eel {
	kiri, kanan := 0, n-1
	kataCari := strings.ToLower(kata)
	// membagi array menjadi 2 bagian sampai ditemukan/habis
	for kiri <= kanan {
		tengah := (kiri + kanan) / 2
		judulTengah := strings.ToLower(arr[tengah].Judul)
		if strings.Contains(judulTengah, kataCari) {
			return tengah
		} else if judulTengah < kataCari {
			kiri = tengah + 1
		} else {
			kanan = tengah - 1
		}
	}
	return -1
}

// menampilkan menu pencarian dan menjalankan metode yang dipilih
func cariFilm() {
	if jumlahFilm == 0 {
		fmt.Println("Koleksi film masih kosong.")
		return
	}
	fmt.Println("\n===== CARI FILM =====")
	fmt.Println("1. Sequential Search - Berdasarkan Judul")
	fmt.Println("2. Sequential Search - Berdasarkan Genre")
	fmt.Println("3. Binary Search     - Berdasarkan Judul")
	fmt.Print("Pilih metode: ")
	pilih := inputInt()

	switch pilih {
	case 1:
		// sequential search menelusuri satu per satu dari awal array
		fmt.Print("Kata kunci judul: ")
		kata := inputStr()
		hasil := seqSearchJudul(kata)
		if len(hasil) == 0 {
			fmt.Println("Film tidak ditemukan.")
		} else {
			fmt.Printf("Ditemukan %d film (Sequential Search):\n", len(hasil))
			for _, idx := range hasil {
				tampilFilm(daftarFilm[idx], idx+1)
			}
		}
	case 2:
		// sequential search mencari semua film dalam genre tertentu
		fmt.Print("Nama genre: ")
		genre := inputStr()
		hasil := seqSearchGenre(genre)
		if len(hasil) == 0 {
			fmt.Printf("Tidak ada film bergenre \"%s\".\n", genre)
		} else {
			fmt.Printf("Ditemukan %d film genre \"%s\" (Sequential Search):\n", len(hasil), genre)
			for _, idx := range hasil {
				tampilFilm(daftarFilm[idx], idx+1)
			}
		}
	case 3:
		// binary search memerlukan array yang sudah terurut berdasarkan judul
		fmt.Print("Kata kunci judul: ")
		kata := inputStr()
		salinan, n := salinFilm()
		// mengurutkan salinan berdasarkan judul sebelum binary search
		for i := 1; i < n; i++ {
			kunci := salinan[i]
			j := i - 1
			for j >= 0 && strings.ToLower(salinan[j].Judul) > strings.ToLower(kunci.Judul) {
				salinan[j+1] = salinan[j]
				j--
			}
			salinan[j+1] = kunci
		}
		idx := binarySearchJudul(salinan, n, kata)
		if idx == -1 {
			fmt.Println("Film tidak ditemukan (Binary Search).")
		} else {
			fmt.Println("Film ditemukan (Binary Search):")
			tampilFilm(salinan[idx], idx+1)
		}
	default:
		fmt.Println("Pilihan tidak valid.")
	}
}

// menampilkan menu pengurutan dan menjalankan metode yang dipilih
func urutkanFilm() {
	if jumlahFilm == 0 {
		fmt.Println("Koleksi film masih kosong.")
		return
	}
	fmt.Println("\n===== URUTKAN FILM =====")
	fmt.Println("1. Selection Sort - Rating Tertinggi ke Terendah")
	fmt.Println("2. Insertion Sort - Tahun Rilis Terlama ke Terbaru")
	fmt.Print("Pilih metode: ")
	pilih := inputInt()

	// membuat salinan agar urutan data asli tidak berubah
	salinan, n := salinFilm()
	switch pilih {
	case 1:
		selectionSortRating(&salinan, n)
		fmt.Println("\n=== Urutan Rating Tertinggi ke Terendah (Selection Sort) ===")
		for i := 0; i < n; i++ {
			tampilFilm(salinan[i], i+1)
		}
	case 2:
		insertionSortTahun(&salinan, n)
		fmt.Println("\n=== Urutan Tahun Rilis Terlama ke Terbaru (Insertion Sort) ===")
		for i := 0; i < n; i++ {
			tampilFilm(salinan[i], i+1)
		}
	default:
		fmt.Println("Pilihan tidak valid.")
	}
}

// Sequential Search mencari semua film berdasarkan genre
func seqSearchGenre(genre jebb) []eel {
	hasil := []eel{}
	// menelusuri setiap elemen dan mencocokkan genre
	for i := 0; i < jumlahFilm; i++ {
		if strings.EqualFold(daftarFilm[i].Genre, genre) {
			hasil = append(hasil, i)
		}
	}
	return hasil
}

// menampilkan statistik koleksi film
func tampilStatistik() {
	if jumlahFilm == 0 {
		fmt.Println("Koleksi film masih kosong.")
		return
	}
	fmt.Println("\n======= STATISTIK KOLEKSI =======")

	// menjumlahkan seluruh rating untuk mencari rata-rata
	totalRating := 0.0
	idxMax, idxMin := 0, 0
	for i := 0; i < jumlahFilm; i++ {
		totalRating += daftarFilm[i].Rating
		if daftarFilm[i].Rating > daftarFilm[idxMax].Rating {
			idxMax = i
		}
		if daftarFilm[i].Rating < daftarFilm[idxMin].Rating {
			idxMin = i
		}
	}
	rataRata := totalRating / float64(jumlahFilm)

	fmt.Printf("Total Film      : %d\n", jumlahFilm)
	fmt.Printf("Rata-rata Rating: %.2f / 10\n", rataRata)
	fmt.Printf("Rating Tertinggi: \"%s\" (%.1f)\n", daftarFilm[idxMax].Judul, daftarFilm[idxMax].Rating)
	fmt.Printf("Rating Terendah : \"%s\" (%.1f)\n", daftarFilm[idxMin].Judul, daftarFilm[idxMin].Rating)

	// menghitung jumlah film per genre menggunakan array paralel
	var namaGenre [MAXFILM]jebb
	var hitungGenre [MAXFILM]eel
	jumlahGenre := 0

	// mengkelompokkan setiap film ke genre yang sesuai
	for i := 0; i < jumlahFilm; i++ {
		ketemu := false
		for j := 0; j < jumlahGenre; j++ {
			if strings.EqualFold(namaGenre[j], daftarFilm[i].Genre) {
				hitungGenre[j]++
				ketemu = true
				break
			}
		}
		if !ketemu {
			namaGenre[jumlahGenre] = daftarFilm[i].Genre
			hitungGenre[jumlahGenre] = 1
			jumlahGenre++
		}
	}

	fmt.Println("\nJumlah Film per Genre:")
	fmt.Println("--------------------------")
	// menampilkan setiap genre beserta jumlah filmnya
	for i := 0; i < jumlahGenre; i++ {
		fmt.Printf("  %-18s : %d film\n", namaGenre[i], hitungGenre[i])
	}
}

func main() {
	loadJSON()
	tampilHeader()
	for {
		tampilMenu()
		pilih := inputInt()
		switch pilih {
		case 1:
			tambahFilm()
		case 2:
			tampilSemuaFilm()
		case 3:
			ubahFilm()
		case 4:
			hapusFilm()
		case 5:
			cariFilm()
		case 6:
			urutkanFilm()
		case 7:
			tampilStatistik()
		case 0:
			fmt.Println("\nTerima kasih telah menggunakan CineReview!")
			fmt.Println("Sampai jumpa!")
			return
		default:
			fmt.Println("Pilihan tidak valid, silakan coba lagi.")
		}
	}
}

func loadJSON() {
	file, err := os.Open("film.json")
	if err != nil {
		fmt.Println("film.json tidak ditemukan, menggunakan data kosong.")
		return
	}
	defer file.Close()

	var data []Film

	decoder := json.NewDecoder(file)
	err = decoder.Decode(&data)
	if err != nil {
		fmt.Println("Gagal membaca JSON:", err)
		return
	}

	jumlahFilm = len(data)

	for i := 0; i < jumlahFilm; i++ {
		daftarFilm[i] = data[i]
	}
}

func simpanJSON() {
	file, err := os.Create("film.json")
	if err != nil {
		fmt.Println("Gagal membuat file:", err)
		return
	}
	defer file.Close()

	data := make([]Film, jumlahFilm)

	for i := 0; i < jumlahFilm; i++ {
		data[i] = daftarFilm[i]
	}

	encoder := json.NewEncoder(file)
	encoder.SetIndent("", "  ")

	err = encoder.Encode(data)
	if err != nil {
		fmt.Println("Gagal menyimpan JSON:", err)
	}
}
