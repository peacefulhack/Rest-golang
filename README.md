# Rest API Product

Ini adalah proyek Rest API untuk manajemen produk menggunakan bahasa pemrograman Go dan database RDMS MySQL/PostgreSQL.

## Requirement

Sebelum memulai proyek ini, pastikan Anda memiliki perangkat lunak berikut terpasang di komputer Anda:

- Go (minimal versi 1.16)
- Make (opsional)
- Docker

## Fitur

Proyek ini memiliki dua endpoint utama:

1. **API Add Product**: Endpoint ini memungkinkan untuk menambahkan produk baru ke dalam database. Produk harus memiliki informasi berikut:
   - product id
   - product name
   - product price
   - product description
   - product quantity

2. **API List Product dengan Sorting**: Endpoint ini memungkinkan untuk mendapatkan daftar produk dengan berbagai metode pengurutan, termasuk:
   - Produk terbaru
   - Produk harga termurah
   - Produk harga termahal
   - Pengurutan berdasarkan nama produk (A-Z dan Z-A)

## Architecture

Proyek ini dibangun dengan menggunakan kaidah prinsip SOLID dengan model MVC:
| Name Sections | Descriptions                                                                                                                                                                  |
|---------------|-------------------------------------------------------------------------------------------------------------------------------------------------------------------------------|
| client        | Client terdiri dari beberapa fungsi dengan satu tujuan, dan tidak boleh ada logic berlebihan pada client, jika diperlukan logic, maka logic akan di letakkan pada controllers |
| config        | config berguna untuk mengatur koneksi antara nano service ke nano service lainnya dan atau ke koneksi ke service luar seperti database, mq broker, ataupun redis              |
| connections   | connection berguna untuk mengatur jaringan nano service                                                                                                                       |
| controllers   | controller adalah jantung dari nano service, dimana logic dan validasi akan diletakkan disini                                                                                 |
| mocks         | mock hanya berguna untuk unit testing                                                                                                                                         |
| models        | models untuk menyimpan entity dan/atau interfaces                                                                                                                             |
| utils         | untuk fungsi kecil yang dapat dibagi dan di reuse, diletakkan pada utils                                                                                                      |

## Cara Penggunaan

Berikut adalah langkah-langkah untuk menjalankan proyek ini:

1. **Clone Proyek**: Clone proyek ini dari GitHub/Bitbucket:

   ```
   git clone https://github.com/peacefulhack/Rest-golang.git
   ```

2. **Jalankan Docker**: untuk menjalankan docker, gunakan perintah berikut:
   ```
   make run-pods
   ```
   atau
   ```
   docker-compose -f docker-compose.local.yml up --build
   ```
  docker akan melakukan download pods yang diperlukan serta melakukan insialisasi sql product_services yang DDL nya dapat di akses pada [init.sql](https://github.com/peacefulhack/Rest-golang/blob/main/init.sql)
  
4. **Jalankan Aplikasi**: Jalankan aplikasi menggunakan perintah berikut:
   ```
   make run-local
   ```
   atau
   ```
   go run ./main.go -env local
   ```

5. **Akses API**: Akses API melalui `http://localhost:PORT`, di mana PORT adalah port yang Anda konfigurasikan di dalam config-local.yml.

## Dokumentasi Tambahan

- Dokumentasi API dapat ditemukan di dalam file [DOKUMENTASI_API.md](DOKUMENTASI_API.MD).
- Untuk mengimplementasikan Redis Cache, ikuti petunjuk di dalam file [REDIS_CACHE.md](REDIS_CACHE.md).
