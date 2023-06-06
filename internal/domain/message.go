package domain

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Message struct {
	ID         primitive.ObjectID `json:"_id,omitempty"`
	UserID     primitive.ObjectID `json:"userId"`
	ChatRoomID primitive.ObjectID `json:"chatRoomId"`
	Content    string             `json:"content"`
	Attachment string             `json:"attachment,omitempty"`
	IsRevoked  bool               `json:"isRevoked,omitempty"`
	Timestamp  time.Time          `json:"timestamp,omitempty"`
}
