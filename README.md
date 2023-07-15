# CakeStore_API
adalah sebuah aplikasi server sederhana yang menyediakan REST API untuk mengelola data kue. Aplikasi ini ditulis dalam bahasa Go menggunakan framework Gin dan menggunakan database MySQL sebagai penyimpanan datanya.

## Instalasi

1. Pastikan Anda memiliki Go dan MySQL terinstal di sistem Anda.
2. Clone repository ini ke dalam direktori lokal Anda.
3. Buka terminal dan masuk ke direktori proyek CakeGin.
4. Jalankan perintah berikut untuk mengunduh dependensi yang diperlukan: go mod download
5. Ubah konfigurasi database sesuai dengan milik Anda dengan membuka file `main.go` dan mengedit bagian berikut di fungsi `main()`:
```go
config := DBConfig{
    Username: "your_db_username",
    Password: "your_db_password",
    Host:     "localhost",
    Port:     "3306",
    Database: "cake_store",
}

// Pastikan database MySQL telah berjalan di sistem Anda.
// Jalankan aplikasi dengan perintah: go run main.go
