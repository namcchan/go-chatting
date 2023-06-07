package repository

import (
	"context"
	"fmt"
	"github.com/namcchan/go-chatting/database"
	"github.com/namcchan/go-chatting/internal/domain"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type userRepository struct {
	ctx            context.Context
	userCollection *mongo.Collection
}

func NewUserRepository(ctx context.Context) domain.UserRepository {
	return &userRepository{
		ctx:            ctx,
		userCollection: database.GetCollection(domain.CollectionUser),
	}
}

func (ur *userRepository) Find(filter bson.M) ([]domain.User, error) {
	var results []domain.User

	cur, err := ur.userCollection.Find(ur.ctx, filter)
	if err != nil {
		return nil, fmt.Errorf("failed to find documents: %v", err)
	}
	defer cur.Close(ur.ctx)

	for cur.Next(ur.ctx) {
		var user domain.User
		err := cur.Decode(&user)
		if err != nil {
			return nil, err
		}

		results = append(results, user)
	}

	if err := cur.Err(); err != nil {
		return nil, err
	}

	return results, nil
}

func (ur *userRepository) FindWithPagination(filter bson.M, limit int64, page int64) ([]domain.User, error) {
	skip := (page - 1) * limit

	findOptions := options.Find()
	findOptions.SetSkip(skip)
	findOptions.SetLimit(limit)

	cur, err := ur.userCollection.Find(ur.ctx, filter, findOptions)
	if err != nil {
		return nil, err
	}
	defer cur.Close(ur.ctx)

	var results []domain.User

	for cur.Next(ur.ctx) {
		var document domain.User
		err := cur.Decode(&document)
		if err != nil {
			return nil, err
		}

		results = append(results, document)
	}

	if err := cur.Err(); err != nil {
		return nil, err
	}

	return results, nil
}

func (ur *userRepository) Delete(filter bson.M) error {
	_, err := ur.userCollection.DeleteMany(ur.ctx, filter)
	if err != nil {
		return err
	}

	return nil
}

func (ur *userRepository) Update(filter bson.M, update bson.M) error {
	_, err := ur.userCollection.UpdateMany(ur.ctx, filter, update)
	if err != nil {
		return err
	}

	return nil
}

func (ur *userRepository) Create(document *domain.User) error {
	_, err := ur.userCollection.InsertOne(ur.ctx, document)
	if err != nil {
		return err
	}

	return nil
}

func (ur *userRepository) FindOne(filter bson.M, projection bson.M) (*domain.User, error) {
	var user *domain.User
	opts := options.FindOne()

	if projection != nil {
		opts.SetProjection(projection)
	}

	err := ur.userCollection.FindOne(ur.ctx, filter, opts).Decode(&user)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return user, nil // Return empty user if no document found
		}
		return user, nil
	}

	return user, nil
}
