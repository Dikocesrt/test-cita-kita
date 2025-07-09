# Soal A - Root Words Replacement

## 📋 Soal

![Soal A](https://res.cloudinary.com/dy2fwknbn/image/upload/v1752065801/Screenshot_2025-07-09_at_12.55.31_hv8ngi.png)

## 🚀 Cara Clone dan Menjalankan

### 1. Clone Repository

```bash
git clone https://github.com/Dikocesrt/test-cita-kita.git
cd test-cita-kita
```

### 2. Masuk ke Direktori Soal A

```bash
cd soal_a
```

### 3. Menjalankan Program

```bash
go run soal_a.go
```

## 📝 Deskripsi

Program ini mengimplementasikan fungsi [`replaceWithRootWords`](soal_a.go) yang mengganti kata-kata dalam kalimat dengan root words yang sesuai. Program akan:

1. Menerima array root words dan sebuah kalimat
2. Mengurutkan root words berdasarkan panjang (terpendek dahulu)
3. Mengganti setiap kata dalam kalimat dengan root word yang cocok (jika ada)
4. Mengembalikan kalimat yang sudah dimodifikasi

## 🔧 Fitur

-   Mengurutkan root words berdasarkan panjang untuk prioritas penggantian
-   Menggunakan `strings.HasPrefix` untuk mencocokkan kata dengan root
-   Mempertahankan kata asli jika tidak ada root word yang cocok

## 📊 Contoh Output

```
Soal 1 Output:
the cat was rat by the bat

Soal 2 Output:
the dog were barking near the car and bike
```

## 🛠️ Requirements

-   Go 1.19+
