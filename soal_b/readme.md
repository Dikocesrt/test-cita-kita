# Soal B - CSV Data Processing

## 📋 Soal

![Soal B](LINK_GAMBAR_SOAL_B_AKAN_DITAMBAHKAN)

## 🚀 Cara Clone dan Menjalankan

### 1. Clone Repository

```bash
git clone https://github.com/Dikocesrt/test-cita-kita.git
cd test-cita-kita
```

### 2. Masuk ke Direktori Soal B

```bash
cd soal_b
```

### 3. Menjalankan Program

```bash
go run soal_b.go
```

## 📝 Deskripsi

Program ini mengunduh data CSV dari GitHub repository countries-states-cities-database, memproses data untuk menghitung jumlah kota per negara, dan menyimpan hasilnya dalam format JSON dan CSV.

## 🔧 Fitur

-   **Download CSV**: Mengunduh data cities.csv dari URL GitHub
-   **Data Processing**: Menghitung jumlah kota untuk setiap negara
-   **Sorting**: Mengurutkan hasil berdasarkan jumlah kota (terbanyak dahulu)
-   **Export**: Menyimpan hasil ke [`output.json`](output.json) dan [`output.csv`](output.csv)

## 📊 Struktur Data Output

### JSON Format

```json
[
    {
        "country_name": "India",
        "city_count": 4000
    },
    {
        "country_name": "United States",
        "city_count": 3000
    }
]
```

### CSV Format

```csv
country_name,city_count
India,4000
United States,3000
```

## 🌐 Data Source

Program menggunakan data dari:

```
https://raw.githubusercontent.com/dr5hn/countries-states-cities-database/refs/heads/master/csv/cities.csv
```

## 📁 Output Files

Setelah program berhasil dijalankan, akan terbuat 2 file:

-   `output.json` - Data dalam format JSON
-   `output.csv` - Data dalam format CSV

## 🛠️ Requirements

-   Go 1.19+
-   Koneksi internet untuk mengunduh data CSV
-   Package yang digunakan:
    -   `encoding/csv`
    -   `encoding/json`
    -   `net/http`
    -   `sort`

## ⚡ Cara Kerja

1. **Download**: Program mengunduh CSV dari URL yang ditentukan
2. **Parse Header**: Mencari kolom "country_name"
3. **Count Cities**: Menghitung jumlah kota untuk setiap negara
4. **Sort**: Mengurutkan berdasarkan jumlah kota secara menurun
5. **Export**: Menyimpan ke file JSON dan CSV
