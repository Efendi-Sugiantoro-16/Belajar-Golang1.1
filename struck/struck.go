package struck

import "go.mongodb.org/mongo-driver/bson/primitive"

// WarehouseItem representasi pada item dalam warehouse inventory (Model)
type WarehouseItem struct {
	ID          primitive.ObjectID `json:"id,omitempty" bson:"_id,omitempty"`
	Name        string             `json:"name" bson:"name"`
	Description string             `json:"description" bson:"description"`
	Quantity    int                `json:"quantity" bson:"quantity"`
	Location    string             `json:"location" bson:"location"`
}

type User struct {
	ID           primitive.ObjectID `json:"id,omitempty" bson:"_id,omitempty"`
	Username     string             `json:"username" bson:"username"`
	Email        string             `json:"email" bson:"email"`
	PasswordHash string             `json:"password_hash" bson:"password_hash"`
	CreatedAt    int64              `json:"created_at" bson:"created_at"`
}

type Transaction struct {
	ID          primitive.ObjectID `json:"id,omitempty" bson:"_id,omitempty"`
	UserID      primitive.ObjectID `json:"user_id" bson:"user_id"`
	Amount      float64            `json:"amount" bson:"amount"`
	Timestamp   int64              `json:"timestamp" bson:"timestamp"`
	Description string             `json:"description" bson:"description"`
}
