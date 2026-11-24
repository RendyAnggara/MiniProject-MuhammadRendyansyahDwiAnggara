# E-Commerce REST API (Go + Fiber v3 + GORM)

Backend RESTful API untuk mini project e-commerce yang dibangun menggunakan bahasa pemrograman **Go (Golang)** dengan arsitektur **Clean Architecture**. Aplikasi ini menyediakan layanan lengkap mulai dari autentikasi pengguna, manajemen profil dan alamat pengiriman, manajemen toko, katalog produk dengan _multiple file upload_, manajemen kategori khusus admin, hingga sistem transaksi checkout beserta pencatatan riwayat log produk (_product snapshot_).

---

## Tech Stack

- **Language:** Go (Golang)
- **Web Framework:** Fiber v3
- **ORM & Database:** GORM & MySQL
- **Authentication & Security:** JWT (JSON Web Token) & Bcrypt Password Hashing
- **Configuration:** Godotenv (`.env`)
- **External API Integration:** API Wilayah Indonesia (Provinsi & Kota/Kabupaten)

---

## Fitur Utama

1. **Authentication & Authorization**
    - Registrasi akun baru beserta pembuatan profil toko otomatis (_auto-create store_).
    - Login menggunakan nomor telepon dan kata sandi dengan penerbitan token JWT.
    - Middleware proteksi endpoint privat menggunakan custom header `token`.
    - Middleware otorisasi khusus admin (`is_admin`) untuk pengelolaan kategori.
2. **User & Shipping Address Management**
    - Melihat dan memperbarui data profil pengguna (termasuk integrasi data provinsi & kota).
    - CRUD lengkap untuk alamat pengiriman milik masing-masing pengguna.
3. **Store (Toko) Management**
    - Melihat informasi toko milik pengguna yang sedang login, melihat detail toko berdasarkan ID, serta daftar seluruh toko dengan _pagination_ dan filter nama.
    - Memperbarui profil toko beserta upload foto/logo toko.
4. **Category Management (Admin Only)**
    - Melihat daftar seluruh kategori dan detail kategori (publik).
    - Membuat, mengubah, dan menghapus kategori (khusus akun dengan hak akses admin).
5. **Product Management**
    - CRUD lengkap produk milik toko pengguna.
    - Mendukung _multiple file upload_ untuk foto produk (`multipart/form-data`).
    - Filter produk berdasarkan nama, kategori, toko, rentang harga (`min_harga` & `max_harga`), serta _pagination_.
6. **Transaction & Product Log System**
    - Proses checkout transaksi dengan validasi kepemilikan alamat, pengecekan stok produk, pengurangan stok otomatis, dan kalkulasi total harga.
    - Pencatatan otomatis snapshot produk ke tabel `product_logs` saat transaksi terjadi agar riwayat harga dan detail produk tetap terjaga meskipun data produk asli diubah di kemudian hari.

---

## Struktur Proyek (Clean Architecture)

```text
MiniProject-MuhammadRendyansyahDwiAnggara/
├── config/          # Konfigurasi koneksi database MySQL & GORM AutoMigrate
├── entities/        # Definisi struct model tabel, relasi, dan DTO request/response
├── handlers/        # HTTP handler untuk menerima request dan mengembalikan response JSON
├── middlewares/     # Middleware validasi token JWT dan pengecekan role Admin
├── repositories/    # Layer akses data dan query langsung ke database MySQL
├── services/        # Layer logika bisnis (business logic) dan validasi aturan aplikasi
├── uploads/         # Direktori penyimpanan lokal untuk file foto toko dan produk
├── utils/           # Helper JWT, hashing bcrypt, pagination, dan standar format response
├── .env.example     # Template variabel environment
├── .gitignore       # Daftar file/folder yang diabaikan oleh Git
├── go.mod           # Daftar dependensi modul Go
├── go.sum           # Checksum dependensi modul Go
├── main.go          # Entry point aplikasi dan registrasi seluruh route API
└── Rakamin Evermos Virtual Internship.postman_collection.json
```

---

## Cara Instalasi & Menjalankan Server

### 1. Prasyarat

- **Go** versi 1.21 atau lebih baru
- **MySQL** / MariaDB (XAMPP, LAMPP, atau MySQL Server lokal)
- **Postman** (untuk pengujian endpoint)

### 2. Clone Repository

```bash
git clone https://github.com/RendyAnggara/MiniProject-MuhammadRendyansyahDwiAnggara.git
cd MiniProject-MuhammadRendyansyahDwiAnggara
```

### 3. Buat Database MySQL

Buat database baru di MySQL / phpMyAdmin dengan nama `evermos` (atau sesuai konfigurasi yang diinginkan):

```sql
CREATE DATABASE evermos;
```

_(Seluruh tabel akan dibuat secara otomatis melalui fitur GORM AutoMigrate saat server dijalankan)._

### 4. Konfigurasi File `.env`

Salin file `.env.example` menjadi `.env`:

```bash
cp .env.example .env
```

Buka file `.env` dan sesuaikan dengan pengaturan database lokal Anda:

```env
DB_HOST=127.0.0.1
DB_PORT=3306
DB_USER=root
DB_PASS=
DB_NAME=evermos
JWT_SECRET=rahasia_jwt_anda
APP_PORT=8000
```

### 5. Install Dependensi & Jalankan Aplikasi

```bash
go mod tidy
go run main.go
```

Jika berhasil, server akan berjalan di `http://localhost:8000` dengan base path `/api/v1`.

---

## Daftar Endpoint Utama

Base URL: `http://localhost:8000/api/v1`

| Modul        | Method   | Endpoint                        | Akses        | Keterangan                                   |
| :----------- | :------- | :------------------------------ | :----------- | :------------------------------------------- |
| **Auth**     | `POST`   | `/auth/register`                | Publik       | Registrasi user baru & auto-create toko      |
| **Auth**     | `POST`   | `/auth/login`                   | Publik       | Login user & mendapatkan JWT token           |
| **User**     | `GET`    | `/user`                         | User (Token) | Melihat profil user yang sedang login        |
| **User**     | `PUT`    | `/user`                         | User (Token) | Memperbarui data profil user                 |
| **Alamat**   | `GET`    | `/user/alamat`                  | User (Token) | Melihat seluruh alamat pengiriman user       |
| **Alamat**   | `GET`    | `/user/alamat/:id`              | User (Token) | Melihat detail alamat berdasarkan ID         |
| **Alamat**   | `POST`   | `/user/alamat`                  | User (Token) | Menambah alamat pengiriman baru              |
| **Alamat**   | `PUT`    | `/user/alamat/:id`              | User (Token) | Memperbarui alamat pengiriman                |
| **Alamat**   | `DELETE` | `/user/alamat/:id`              | User (Token) | Menghapus alamat pengiriman                  |
| **Category** | `GET`    | `/category`                     | Publik       | Melihat semua kategori produk                |
| **Category** | `GET`    | `/category/:id`                 | Publik       | Melihat detail kategori berdasarkan ID       |
| **Category** | `POST`   | `/category`                     | Admin Only   | Membuat kategori baru                        |
| **Category** | `PUT`    | `/category/:id`                 | Admin Only   | Memperbarui kategori                         |
| **Category** | `DELETE` | `/category/:id`                 | Admin Only   | Menghapus kategori                           |
| **Toko**     | `GET`    | `/toko`                         | Publik       | Melihat daftar toko (pagination & filter)    |
| **Toko**     | `GET`    | `/toko/my`                      | User (Token) | Melihat profil toko milik user               |
| **Toko**     | `GET`    | `/toko/:id_toko`                | Publik       | Melihat detail toko berdasarkan ID           |
| **Toko**     | `PUT`    | `/toko/:id_toko`                | User (Token) | Memperbarui profil & foto toko               |
| **Product**  | `GET`    | `/product`                      | Publik       | Melihat daftar produk (filter & pagination)  |
| **Product**  | `GET`    | `/product/:id`                  | Publik       | Melihat detail produk berdasarkan ID         |
| **Product**  | `POST`   | `/product`                      | User (Token) | Menambah produk baru (+ upload foto)         |
| **Product**  | `PUT`    | `/product/:id`                  | User (Token) | Memperbarui data & foto produk               |
| **Product**  | `DELETE` | `/product/:id`                  | User (Token) | Menghapus produk                             |
| **Trx**      | `GET`    | `/trx`                          | User (Token) | Melihat riwayat transaksi user               |
| **Trx**      | `GET`    | `/trx/:id`                      | User (Token) | Melihat detail invoice transaksi             |
| **Trx**      | `POST`   | `/trx`                          | User (Token) | Membuat transaksi checkout baru              |
| **Wilayah**  | `GET`    | `/provcity/listprovincies`      | Publik       | Mengambil daftar provinsi Indonesia          |
| **Wilayah**  | `GET`    | `/provcity/listcities/:prov_id` | Publik       | Mengambil daftar kota/kabupaten per provinsi |

---

## Panduan Pengujian dengan Postman

1. Buka aplikasi **Postman**, lalu **Import** file Postman Collection yang tersedia di dalam repository ini.
2. Buka tab **Variables** pada folder collection dan pastikan variabel `local` memiliki nilai `http://localhost:8000/api/v1`.
3. Jalankan endpoint **`Register`**, kemudian ubah nilai kolom `is_admin` menjadi `1` pada tabel `users` di database MySQL agar akun tersebut dapat menguji endpoint **`Category`**.
4. Jalankan endpoint **`Login`** untuk mendapatkan `token` autentikasi secara otomatis sebelum mengakses endpoint privat lainnya.

---

## Author

**Muhammad Rendyansyah Dwi Anggara**
