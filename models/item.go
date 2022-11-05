package models

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Item struct {
	ID        primitive.ObjectID  `json:"_id" bson:"_id"`
	Tags      string              `json:"tags,omitempty" validate:"required"`
	Data      string              `json:"data,omitempty" validate:"required"`
	UpdatedAt primitive.Timestamp `json:"update_at""`

	// "lastUpdate": primitive.Timestamp{T:uint32(time.Now().Unix())}
	// LastUpdate primitive.Timestamp `json:"lastUpdate"`
}

type CreateItem struct {
	Tags      string              `json:"tags,omitempty" validate:"required"`
	Data      string              `json:"data,omitempty" validate:"required"`
	UpdatedAt primitive.Timestamp `json:"update_at""`
}

type ItemResponse struct {
	Status  int                    `json:"status"`
	Message string                 `json:"message"`
	Data    map[string]interface{} `json:"data"`
}
