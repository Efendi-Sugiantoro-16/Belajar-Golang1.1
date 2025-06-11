package struck

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

// WarehouseItem represents an item in the warehouse inventory.
type WarehouseItem struct {
	ID          primitive.ObjectID `json:"id,omitempty" bson:"_id,omitempty"`
	Name        string             `json:"name" bson:"name"`
	Description string             `json:"description" bson:"description"`
	Quantity    int                `json:"quantity" bson:"quantity"`
	Location    string             `json:"location" bson:"location"`
}

// User represents a user of the system.
type User struct {
	ID           primitive.ObjectID `json:"id,omitempty" bson:"_id,omitempty"`
	Username     string             `json:"username" bson:"username"`
	Email        string             `json:"email" bson:"email"`
	PasswordHash string             `json:"password_hash" bson:"password_hash"`
	CreatedAt    time.Time          `json:"created_at" bson:"created_at"`
}

// Transaction represents a financial transaction made by a user.
type Transaction struct {
	ID          primitive.ObjectID `json:"id,omitempty" bson:"_id,omitempty"`
	UserID      primitive.ObjectID `json:"user_id" bson:"user_id"`
	Amount      float64            `json:"amount" bson:"amount"`
	Timestamp   time.Time          `json:"timestamp" bson:"timestamp"`
	Description string             `json:"description" bson:"description"`
}

// SetCreatedAt sets the CreatedAt field for the User to the current time.
func (u *User) SetCreatedAt() {
	u.CreatedAt = time.Now()
}

// SetTimestamp sets the Timestamp field for the Transaction to the current time.
func (t *Transaction) SetTimestamp() {
	t.Timestamp = time.Now()
}
