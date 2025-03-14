# Simulasi Kasir dengan Goroutine dan Mutex

Program ini mensimulasikan sistem kasir di sebuah toko menggunakan bahasa pemrograman Go. Dalam simulasi ini:
- Terdapat daftar item yang dijual dengan harga tertentu.
- Pelanggan akan dibuat secara acak dengan nama yang dihasilkan menggunakan pustaka `gofakeit`.
- Setiap pelanggan memiliki dua item yang dipilih secara acak dari daftar item yang tersedia.
- Lima kasir akan melayani pelanggan secara bersamaan menggunakan goroutine.
- Total pemasukan toko dihitung secara aman menggunakan `sync.Mutex` untuk menghindari race condition.

## Fitur
- Menggunakan goroutine untuk menjalankan banyak kasir secara paralel.
- Menggunakan channel untuk mengantre pelanggan.
- Menggunakan `sync.Mutex` untuk mengelola akses ke variabel `totalRevenue` agar aman dari race condition.
- Menggunakan `gofakeit` untuk menghasilkan nama pelanggan secara acak.
- Menggunakan `math/rand` untuk memilih item pelanggan secara acak.

## Instalasi
Pastikan telah menginstal Go di sistem. Jika belum, unduh dan instal dari [golang.org](https://golang.org/dl/).

Kemudian, unduh dependensi berikut:
```sh
go get github.com/brianvoe/gofakeit
```

## Cara Menjalankan
Jalankan perintah berikut di terminal:
```sh
go run main.go
```

## Output Contoh
Output program akan bervariasi setiap kali dijalankan karena data pelanggan dan waktu layanan acak. Contoh output:
```
Kasir 1 melayani pelanggan John Doe
  - Laptop: Rp15000.00
  - Mouse: Rp500.00
Total belanja John Doe: Rp15500.00

Kasir 2 melayani pelanggan Jane Smith
  - Monitor: Rp3000.00
  - Headphone: Rp2000.00
Total belanja Jane Smith: Rp5000.00

...

Total pemasukan toko: RpXYZ.00
```

## Struktur Kode
- `Item` : Struct yang merepresentasikan produk yang dijual.
- `Customer` : Struct yang merepresentasikan pelanggan dan item yang dibeli.
- `generateItems()` : Fungsi untuk membuat daftar item yang tersedia.
- `generateCustomers(num int, items []Item)` : Fungsi untuk membuat daftar pelanggan dengan item yang dibeli.
- `cashier()` : Fungsi yang dijalankan oleh goroutine untuk menangani pelanggan.
- `main()` : Fungsi utama yang menginisialisasi data, membuat channel, dan menjalankan goroutine.
