package domain

import (
	"go.mongodb.org/mongo-driver/bson"
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

const CollectionUser = "users"

type User struct {
	ID         primitive.ObjectID `json:"id,omitempty" bson:"id,omitempty"`
	Username   string             `json:"username" bson:"username" validate:"required"`
	Name       string             `json:"name" bson:"name" validate:"required"`
	Email      string             `json:"email" bson:"email" validate:"required"`
	Password   string             `json:"password" bson:"password" validate:"required"`
	Avatar     string             `json:"avatar,omitempty" bson:"avatar,omitempty"`
	IsVerified bool               `json:"isVerified,omitempty" bson:"isVerified,omitempty"`
	CreatedAt  time.Time          `json:"createdAt,omitempty" bson:"createdAt,omitempty"`
	UpdatedAt  time.Time          `json:"updatedAt,omitempty" bson:"updatedAt,omitempty"`
}

type UserRepository interface {
	Create(document *User) error
	FindOne(filter bson.M) (*User, error)
	Find(filter bson.M) ([]User, error)
	FindWithPagination(filter bson.M, limit int64, page int64) ([]User, error)
	Delete(filter bson.M) error
	Update(filter bson.M, update bson.M) error
}
