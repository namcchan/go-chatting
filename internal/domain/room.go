package domain

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type ChatRoom struct {
	ID           primitive.ObjectID   `json:"_id,omitempty"`
	Name         string               `json:"name"`
	Description  string               `json:"description"`
	CreatedBy    primitive.ObjectID   `json:"created_by,omitempty"`
	Type         string               `json:"type"`
	Participants []primitive.ObjectID `json:"participants,omitempty"`
	CreatedAt    time.Time            `json:"createdAt,omitempty"`
	UpdatedAt    time.Time            `json:"updatedAt,omitempty"`
}
