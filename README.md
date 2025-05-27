# Belajar-Golang1.1

Belajar Golang dengan Go standard/ Go Fiber.

---

## Tutorial: Membuat Aplikasi Warehouse dengan GoFiber dan MongoDB

Tutorial ini membahas langkah-langkah membuat aplikasi warehouse sederhana menggunakan GoFiber sebagai backend dan MongoDB sebagai database. Tutorial ini juga menjelaskan cara membuat struct dan fungsi di Go, mengembalikan response JSON, serta menghubungkan backend dengan Postman untuk testing.

---

### 1. Membuat Struct dan Fungsi di Golang

Di Go, struct digunakan untuk mendefinisikan tipe data kompleks. Contoh struct untuk item warehouse:

```go
package struck

import "go.mongodb.org/mongo-driver/bson/primitive"

type WarehouseItem struct {
    ID          primitive.ObjectID `bson:"_id,omitempty" json:"id"`
    Name        string             `bson:"name" json:"name"`
    Description string             `bson:"description" json:"description"`
    Quantity    int                `bson:"quantity" json:"quantity"`
    Location    string             `bson:"location" json:"location"`
}
```

Fungsi di Go didefinisikan dengan kata kunci `func`. Contoh fungsi untuk membuat item baru:

```go
func createItem(c *fiber.Ctx) error {
    item := new(WarehouseItem)
    if err := c.BodyParser(item); err != nil {
        return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "cannot parse JSON"})
    }
    // Simpan item ke database...
    return c.Status(fiber.StatusCreated).JSON(item)
}
```

---

### 2. Membuat Return JSON di Golang dengan GoFiber

GoFiber menyediakan method `c.JSON()` untuk mengembalikan response JSON. Contoh:

```go
return c.JSON(fiber.Map{
    "message": "Data berhasil disimpan",
    "data":    item,
})
```

Untuk konsistensi, kita bisa membuat helper function untuk response JSON, misalnya di package `cjson`:

```go
package cjson

import "github.com/gofiber/fiber/v2"

func JSONResponse(c *fiber.Ctx, status int, data interface{}) error {
    return c.Status(status).JSON(data)
}
```

---

### 3. Melakukan Koneksi antara Golang dengan Postman

Setelah backend berjalan, kita bisa menggunakan Postman untuk menguji API.

- Jalankan server GoFiber: `go run main.go`
- Buka Postman, buat request ke endpoint, misalnya `POST http://localhost:3000/api/items`
- Pilih tab Body, pilih raw dan JSON, lalu isi data JSON:

```json
{
    "name": "Item A",
    "description": "Deskripsi item A",
    "quantity": 10,
    "location": "Gudang 1"
}
```

- Kirim request dan lihat response JSON dari server.

---

### 4. Fitur CRUD di Aplikasi Warehouse

Aplikasi ini mendukung operasi CRUD lengkap:

- **GET** `/api/items` - Mendapatkan daftar semua item.
- **GET** `/api/items/:id` - Mendapatkan detail item berdasarkan ID.
- **POST** `/api/items` - Menambahkan item baru.
- **PUT** `/api/items/:id` - Memperbarui item berdasarkan ID.
- **DELETE** `/api/items/:id` - Menghapus item berdasarkan ID.



