# Dokumentasi API Produk

Dokumen ini memberikan informasi tentang endpoint API Produk dan cara penggunaannya.

## Daftar Isi
- [Tambah Produk](#tambah-produk)
- [Urutkan Produk](#urutkan-produk)

## Tambah Produk

### Endpoint
```
POST localhost:8080/product/insert
```

Endpoint ini digunakan untuk menambahkan produk baru ke sistem.

#### Body Permintaan
- `ProductName` (string): Nama produk.
- `ProductDesc` (string): Deskripsi produk.
- `ProductPrice` (int): Harga produk.
- `ProductQuantity` (int): Jumlah produk.

#### Contoh Permintaan
```json
{
  "ProductName": "Contoh Produk",
  "ProductDesc": "Ini adalah deskripsi produk contoh.",
  "ProductPrice": 20,
  "ProductQuantity": 100
}
```

#### Respons
Setelah produk berhasil ditambahkan, API akan memberikan respons dengan kode status `200 OK`.

## Urutkan Produk

### Endpoint
```
GET localhost:8080/product/sort/:key/:order
```

Endpoint ini digunakan untuk mengambil daftar produk yang diurutkan berdasarkan kunci dan urutan yang ditentukan.

#### Parameter
- `key` (string): Kunci untuk mengurutkan produk. Nilai yang diterima adalah:
  - "name" (mengurutkan berdasarkan nama produk)
  - "price" (mengurutkan berdasarkan harga produk)
  - "new" (mengurutkan berdasarkan waktu produk ditambah)
- `order` (string): Urutan pengurutan produk. Nilai yang diterima adalah:
  - "asc" (urutan naik)
  - "desc" (urutan turun)

#### Contoh Permintaan
```
GET localhost:8080/product/sort/price/asc
```

#### Respons
API akan memberikan daftar produk yang diurutkan sesuai dengan kunci dan urutan yang ditentukan. Setiap objek produk dalam daftar akan memiliki atribut berikut:

- `Id` (int): Identifikasi unik produk.
- `ProductName` (string): Nama produk.
- `ProductDesc` (string): Deskripsi produk.
- `ProductPrice` (int): Harga produk.
- `ProductQuantity` (int): Jumlah produk.

#### Contoh Respons
```json
[
  {
    "Id": 1,
    "ProductName": "Produk A",
    "ProductDesc": "Deskripsi untuk Produk A",
    "ProductPrice": 25,
    "ProductQuantity": 50
  },
  {
    "Id": 2,
    "ProductName": "Produk B",
    "ProductDesc": "Deskripsi untuk Produk B",
    "ProductPrice": 30,
    "ProductQuantity": 75
  }
]
```
