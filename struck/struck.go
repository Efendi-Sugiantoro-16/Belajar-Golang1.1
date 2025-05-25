package struck

import "go.mongodb.org/mongo-driver/bson/primitive"

// WarehouseItem represents an item in the warehouse inventory
type WarehouseItem struct {
    ID          primitive.ObjectID `json:"id,omitempty" bson:"_id,omitempty"`
    Name        string             `json:"name" bson:"name"`
    Description string             `json:"description" bson:"description"`
    Quantity    int                `json:"quantity" bson:"quantity"`
    Location    string             `json:"location" bson:"location"`
}
