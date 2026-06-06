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
type Film struct { //struktur film
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
func tampilFilm(f Film, no eel) { //menampilkan detail lengkap sebuah film
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
	jumlahFilm++ //menyimpan ke array
	fmt.Printf("Film \"%s\" berhasil ditambahkan!\n", f.judul)
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
	fmt.Printf("\nMengubah film: %s\n", daftarFilm[idx].judul)
	var f Film
	fmt.Print("Judul baru    : ")
	f.judul = inputStr()
	fmt.Print("Tahun baru    : ")
	f.tahun = inputInt()
	fmt.Print("Genre baru    : ")
	f.genre = inputStr()
	fmt.Print("Deskripsi baru: ")
	f.deskripsi = inputStr()
	fmt.Print("Rating baru   : ")
	f.rating = klampRating(inputFloat())
	daftarFilm[idx] = f
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
	judul := daftarFilm[idx].judul
	// geser semua film di belakang index ke posisi sebelumnya 
	for i := idx; i < jumlahFilm-1; i++ {
		daftarFilm[i] = daftarFilm[i+1]
	}
	jumlahFilm--
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
		if strings.Contains(strings.ToLower(daftarFilm[i].judul), strings.ToLower(kata)) {
			hasil = append(hasil, i)
		}
	}
	return hasil
}


// menampilkan ringkasan statistik dari seluruh koleksi film 
func tampilStatistik() {
	if jumlahFilm == 0 {
		fmt.Println("Koleksi film masih kosong.")
		return
	}
	fmt.Println("\n======= STATISTIK KOLEKSI =======")

	// hitung total rating sekaligus cari indeks film terbaik dan terburuk 
	totalRating := 0.0
	idxMax, idxMin := 0, 0
	for i := 0; i < jumlahFilm; i++ {
		totalRating += daftarFilm[i].rating
		if daftarFilm[i].rating > daftarFilm[idxMax].rating {
			idxMax = i
		}
		if daftarFilm[i].rating < daftarFilm[idxMin].rating {
			idxMin = i
		}
	}

	fmt.Printf("Total Film      : %d\n", jumlahFilm)
	fmt.Printf("Rata-rata Rating: %.2f / 10\n", totalRating/float64(jumlahFilm))
	fmt.Printf("Rating Tertinggi: \"%s\" (%.1f)\n", daftarFilm[idxMax].judul, daftarFilm[idxMax].rating)
	fmt.Printf("Rating Terendah : \"%s\" (%.1f)\n", daftarFilm[idxMin].judul, daftarFilm[idxMin].rating)
	fmt.Println("\n[Statistik per genre akan ditambahkan pada Review 2]")
}
// Selection Sort mengurutkan film dari rating tertinggi ke terendah
func selectionSortRating(arr *[MAXFILM]Film, n eel) {
	// memilih elemen dengan rating terbesar lalu tukar ke posisi yang benar 
	for i := 0; i < n-1; i++ {
		maxIdx := i
		for j := i + 1; j < n; j++ {
			if arr[j].rating > arr[maxIdx].rating {
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
		for j >= 0 && arr[j].tahun > kunci.tahun {
			arr[j+1] = arr[j]
			j--
		}
		arr[j+1] = kunci
	}
}

//binarysearch cari film berdasarkan judul
func binarySearchJudul(arr [MAXFILM]Film, n eel, kata jebb) eel {
	kiri, kanan := 0, n-1
	kataCari := strings.ToLower(kata)
	// membagi array menjadi 2 bagian sampai ditemukan/habis
	for kiri <= kanan {
		tengah := (kiri + kanan) / 2
		judulTengah := strings.ToLower(arr[tengah].judul)
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
			for j >= 0 && strings.ToLower(salinan[j].judul) > strings.ToLower(kunci.judul) {
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

func main() {
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