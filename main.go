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

	"Belajar-Golang-1.1/cjson"
	"Belajar-Golang-1.1/struck"
)

var collection *mongo.Collection

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
	app := fiber.New()

	app.Use(cors.New())

	client := connectDB()
	collection = client.Database("warehouse").Collection("items")

	// Serve static files from public directory
	app.Static("/", "./public")

	app.Get("/api/items", getItems)
	app.Get("/api/items/:id", getItem)
	app.Post("/api/items", createItem)
	app.Put("/api/items/:id", updateItem)
	app.Delete("/api/items/:id", deleteItem)

	log.Fatal(app.Listen(":3000"))
}

func createItem(c *fiber.Ctx) error {
	item := new(struck.WarehouseItem)
	if err := c.BodyParser(item); err != nil {
		return cjson.JSONResponse(c, fiber.StatusBadRequest, fiber.Map{"error": "cannot parse JSON"})
	}
	item.ID = primitive.NewObjectID()
	_, err := collection.InsertOne(context.Background(), item)
	if err != nil {
		return cjson.JSONResponse(c, fiber.StatusInternalServerError, fiber.Map{"error": "cannot insert item"})
	}
	return cjson.JSONResponse(c, fiber.StatusCreated, item)
}

func getItems(c *fiber.Ctx) error {
	cursor, err := collection.Find(context.Background(), bson.M{})
	if err != nil {
		return cjson.JSONResponse(c, fiber.StatusInternalServerError, fiber.Map{"error": "cannot fetch items"})
	}
	var items []struck.WarehouseItem
	if err = cursor.All(context.Background(), &items); err != nil {
		return cjson.JSONResponse(c, fiber.StatusInternalServerError, fiber.Map{"error": "cannot parse items"})
	}
	return cjson.JSONResponse(c, fiber.StatusOK, items)
}

func getItem(c *fiber.Ctx) error {
	idParam := c.Params("id")
	id, err := primitive.ObjectIDFromHex(idParam)
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
	item.ID = id
	return cjson.JSONResponse(c, fiber.StatusOK, item)
}
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
	return c.SendStatus(fiber.StatusNoContent)
}
