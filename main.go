package main

import (
	"context"
	"log"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"

	"Belajar-Golang-1.1/cjson"  // Modul untuk format respon JSON
	"Belajar-Golang-1.1/struck" // Modul struct data WarehouseItem
)

var collection *mongo.Collection // Variabel global untuk koleksi MongoDB

// Fungsi untuk menghubungkan aplikasi ke MongoDB
func connectDB() *mongo.Client {
	client, err := mongo.NewClient(options.Client().ApplyURI("mongodb://localhost:27017"))
	if err != nil {
		log.Fatal(err)
	}
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	err = client.Connect(ctx)
	if err != nil {
		log.Fatal(err)
	}
	return client
}

func main() {
	app := fiber.New() // Membuat instance aplikasi Fiber

	app.Use(cors.New()) // Mengaktifkan middleware CORS

	client := connectDB()                                         // Menghubungkan ke database
	collection = client.Database("warehouse").Collection("items") // Mengakses koleksi "items" di DB "warehouse"

	// Menyajikan file statis dari folder ./public
	app.Static("/", "./public")

	// Daftar endpoint (API route)
	app.Get("/api/items", getItems)          // Ambil semua data
	app.Get("/api/items/:id", getItem)       // Ambil data berdasarkan ID
	app.Post("/api/items", createItem)       // Tambah data baru
	app.Put("/api/items/:id", updateItem)    // Perbarui data berdasarkan ID
	app.Delete("/api/items/:id", deleteItem) // Hapus data berdasarkan ID

	log.Fatal(app.Listen(":3000")) // Jalankan server di port 3000
}

// Fungsi untuk menambahkan item baru ke database
func createItem(c *fiber.Ctx) error {
	item := new(struck.WarehouseItem)
	if err := c.BodyParser(item); err != nil {
		return cjson.JSONResponse(c, fiber.StatusBadRequest, fiber.Map{"error": "cannot parse JSON"})
	}
	item.ID = primitive.NewObjectID() // Buat ID baru otomatis
	_, err := collection.InsertOne(context.Background(), item)
	if err != nil {
		return cjson.JSONResponse(c, fiber.StatusInternalServerError, fiber.Map{"error": "cannot insert item"})
	}
	return cjson.JSONResponse(c, fiber.StatusCreated, item)
}

// Fungsi untuk mengambil semua item dari database
func getItems(c *fiber.Ctx) error {
	cursor, err := collection.Find(context.Background(), bson.M{}) // bson.M{} berarti ambil semua data
	if err != nil {
		return cjson.JSONResponse(c, fiber.StatusInternalServerError, fiber.Map{"error": "cannot fetch items"})
	}
	var items []struck.WarehouseItem
	if err = cursor.All(context.Background(), &items); err != nil {
		return cjson.JSONResponse(c, fiber.StatusInternalServerError, fiber.Map{"error": "cannot parse items"})
	}
	return cjson.JSONResponse(c, fiber.StatusOK, items)
}

// Fungsi untuk mengambil satu item berdasarkan ID
func getItem(c *fiber.Ctx) error {
	idParam := c.Params("id")
	id, err := primitive.ObjectIDFromHex(idParam) // Konversi string ke ObjectID
	if err != nil {
		return cjson.JSONResponse(c, fiber.StatusBadRequest, fiber.Map{"error": "invalid id"})
	}
	var item struck.WarehouseItem
	err = collection.FindOne(context.Background(), bson.M{"_id": id}).Decode(&item)
	if err != nil {
		return cjson.JSONResponse(c, fiber.StatusNotFound, fiber.Map{"error": "item not found"})
	}
	return cjson.JSONResponse(c, fiber.StatusOK, item)
}

// Fungsi untuk memperbarui data item berdasarkan ID
func updateItem(c *fiber.Ctx) error {
	idParam := c.Params("id")
	id, err := primitive.ObjectIDFromHex(idParam)
	if err != nil {
		return cjson.JSONResponse(c, fiber.StatusBadRequest, fiber.Map{"error": "invalid id"})
	}
	item := new(struck.WarehouseItem)
	if err := c.BodyParser(item); err != nil {
		return cjson.JSONResponse(c, fiber.StatusBadRequest, fiber.Map{"error": "cannot parse JSON"})
	}
	update := bson.M{
		"$set": bson.M{
			"name":        item.Name,
			"description": item.Description,
			"quantity":    item.Quantity,
			"location":    item.Location,
		},
	}
	_, err = collection.UpdateOne(context.Background(), bson.M{"_id": id}, update)
	if err != nil {
		return cjson.JSONResponse(c, fiber.StatusInternalServerError, fiber.Map{"error": "cannot update item"})
	}
	item.ID = id // Set ID kembali agar konsisten
	return cjson.JSONResponse(c, fiber.StatusOK, item)
}

// Fungsi untuk menghapus data item berdasarkan ID
func deleteItem(c *fiber.Ctx) error {
	idParam := c.Params("id")
	id, err := primitive.ObjectIDFromHex(idParam)
	if err != nil {
		return cjson.JSONResponse(c, fiber.StatusBadRequest, fiber.Map{"error": "invalid id"})
	}
	_, err = collection.DeleteOne(context.Background(), bson.M{"_id": id})
	if err != nil {
		return cjson.JSONResponse(c, fiber.StatusInternalServerError, fiber.Map{"error": "cannot delete item"})
	}
	return c.SendStatus(fiber.StatusNoContent) // 204 No Content
}
