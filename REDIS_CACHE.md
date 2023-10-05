# Dokumentasi Redis

Redis adalah sistem penyimpanan data berkinerja tinggi yang digunakan dalam konteks API Produk. Redis digunakan untuk caching hasil pengurutan produk sehingga meningkatkan respons API yang lebih cepat.

## Penyimpanan Cache

Ketika melakukan permintaan pengurutan produk dengan kunci dan urutan tertentu, hasil pengurutan tersebut akan dicache di Redis. Cache disimpan dengan nama kunci yang berkaitan dengan permintaan pengurutan, 
sesuai dengan kunci dan urutan yang diberikan. Misalnya, jika Anda melakukan permintaan `GET /product/sort/price/asc`, hasil pengurutan produk akan disimpan dengan nama kunci "sort-price-asc" di Redis.

## Penghapusan Cache saat Menambah Data

Setiap kali melakukan permintaan untuk menambahkan produk baru ke sistem menggunakan `POST /product/insert`, akan otomatis menghapus semua cache pengurutan yang ada di Redis. 
Ini dilakukan untuk memastikan bahwa cache selalu mendukung data yang akurat dan terkini.

## Penyimpanan Cache saat Permintaan Sort

Ketika melakukan permintaan sort (`GET /product/sort/:key/:order`), hasil pengurutan produk akan disimpan dalam Redis sesuai dengan kunci dan urutan yang diberikan dalam permintaan. 
Misalnya, jika terdapat permintaan `GET /product/sort/price/asc`, hasil pengurutan produk akan disimpan dengan nama kunci "sort-price-asc" di Redis.

Ini memungkinkan untuk menyimpan hasil pengurutan produk secara efisien dalam cache untuk permintaan selanjutnya dengan kunci dan urutan yang sama.

Dengan demikian, Redis digunakan dalam aplikasi ini untuk meningkatkan kinerja dengan menyimpan hasil pengurutan dalam cache dan memastikan cache selalu diperbarui saat ada perubahan data atau permintaan sort baru.
