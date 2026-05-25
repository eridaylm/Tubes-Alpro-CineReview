# CineReview 🎬

CineReview adalah aplikasi katalog dan rating film berbasis bahasa pemrograman Go yang digunakan untuk mendata, mencari, mengurutkan, dan menampilkan statistik koleksi film favorit pengguna.

Aplikasi ini dibuat sebagai tugas besar dengan penerapan konsep:

* Modular Programming
* Subprogram / Function
* Array Statis
* Sequential Search
* Binary Search
* Selection Sort
* Insertion Sort

---

# 📌 Deskripsi Aplikasi

CineReview membantu pengguna untuk:

* Menambahkan data film
* Mengubah data film
* Menghapus data film
* Memberikan rating film
* Mencari film berdasarkan judul atau genre
* Mengurutkan film berdasarkan rating atau tahun rilis
* Menampilkan statistik koleksi film

Data yang disimpan meliputi:

* Judul film
* Genre
* Tahun rilis
* Deskripsi singkat
* Rating pengguna

---

# ⚙️ Fitur Utama

## 1. Tambah Data Film

Pengguna dapat menambahkan film baru ke dalam katalog.

## 2. Edit Data Film

Pengguna dapat mengubah informasi film yang sudah tersimpan.

## 3. Hapus Data Film

Pengguna dapat menghapus film tertentu dari katalog.

## 4. Pencarian Film

Menggunakan:

* Sequential Search
* Binary Search

Pencarian dapat dilakukan berdasarkan:

* Judul film
* Genre film

## 5. Pengurutan Data

Menggunakan:

* Selection Sort
* Insertion Sort

Pengurutan dapat dilakukan berdasarkan:

* Rating tertinggi ke terendah
* Tahun rilis
* Ascending maupun Descending

## 6. Statistik Film

Sistem dapat menampilkan:

* Jumlah film per genre
* Rata-rata rating seluruh film

---

# ⭐ Konsep Algoritma yang Digunakan

| Algoritma         | Fungsi                                                       |
| ----------------- | ------------------------------------------------------------ |
| Sequential Search | Mencari data film secara berurutan                           |
| Binary Search     | Mencari data film lebih efisien pada data terurut            |
| Selection Sort    | Mengurutkan data dengan memilih nilai minimum/maksimum       |
| Insertion Sort    | Mengurutkan data dengan menyisipkan elemen pada posisi tepat |

---

# 🏗️ Struktur Program

Program dibuat secara modular menggunakan subprogram/function.

Contoh pembagian subprogram:

```go
TambahFilm()
EditFilm()
HapusFilm()
CariFilmSequential()
CariFilmBinary()
SelectionSortRating()
InsertionSortTahun()
TampilkanStatistik()
```

---

# 🗂️ Struktur Data

Program menggunakan:

* Array statis
* Tipe bentukan / struct

Contoh struktur data:

```go
const MAX = 100

type Film struct {
    Judul string
    Genre string
    Tahun int
    Deskripsi string
    Rating float64
}

var daftarFilm [MAX]Film
```

---

# ▶️ Cara Menjalankan Program

## 1. Clone Repository

```bash
git clone https://github.com/username/CineReview.git
```

## 2. Masuk ke Folder Project

```bash
cd CineReview
```

## 3. Jalankan Program

```bash
go run program.go
```

---

# 💻 Teknologi yang Digunakan

* Golang
* CLI / Terminal Application

---

# 📋 Spesifikasi Tugas

Program memenuhi ketentuan berikut:

✅ Menggunakan subprogram / modular programming

✅ Menggunakan array statis

✅ Menggunakan Sequential Search dan Binary Search

✅ Menggunakan Selection Sort dan Insertion Sort

✅ Mendukung ascending dan descending sort

✅ Tidak menggunakan array dinamis atau slice

✅ Tidak menggunakan statement `break` dan `continue`

---

# 👥 Tim Pengembang

Barret Fairuz Azizah | 109082530034
Eridayalma Zahra Yohar | 109082500221
Hiliyati Aulia | 109082500157

---

# 📄 Lisensi

Project ini dibuat untuk memenuhi Tugas Besar Mata Kuliah Algoritma Pemrograman 2 
