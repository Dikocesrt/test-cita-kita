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
-   **Membaca test cases tambahan dari file JSON**
-   **Menampilkan multiple test cases dengan berbagai skenario**

## 📁 File dalam Project

-   [`soal_a.go`](soal_a.go) - File utama program
-   [`test.json`](test.json) - File berisi test cases tambahan
-   [`readme.md`](readme.md) - Dokumentasi project

## 📊 Contoh Output

### Output Soal Utama:

```
Soal 1 Output:
the cat was rat by the bat

Soal 2 Output:
the dog were barking near the car and bike
```

### Output Test Cases Tambahan:

```
--------------------------------------------------

Test Case 1:
Root Words: [run jump walk]
Original Sentence: the runner was jumping and walking quickly
Result: the run was jump and walk quickly
--------------------------------------------------

Test Case 2:
Root Words: [code test debug]
Original Sentence: the coder was testing and debugging the application
Result: the code was test and debug the application
--------------------------------------------------
```

## 🧪 Test Cases

Program ini menggunakan dua sumber test cases:

1. **Test Cases dari Soal**:

    - Soal 1: Root words ["cat", "bat", "rat"]
    - Soal 2: Root words ["dog", "car", "bike"]

2. **Test Cases Tambahan dari JSON**:
    - 5 test cases berbeda yang dimuat dari [`test.json`](test.json)
    - Mencakup berbagai skenario replacement

## 🔧 Struktur Data JSON

File [`test.json`](test.json) menggunakan struktur:

```json
{
    "test_cases": [
        {
            "id": 1,
            "root_words": ["run", "jump", "walk"],
            "sentence": "the runner was jumping and walking quickly"
        }
    ]
}
```

## 🛠️ Requirements

-   Go 1.19+
-   File [`test.json`](test.json) dalam direktori yang sama
