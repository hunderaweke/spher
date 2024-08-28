package mongo

import (
	"context"

	"github.com/hunderaweke/spher/domain"
	"github.com/sv-tools/mongoifc"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type userRepository struct {
	collection mongoifc.Collection
	ctx        context.Context
}

func NewMongoUserRepository(db mongoifc.Database, ctx context.Context) domain.UserRepository {
	return &userRepository{collection: db.Collection(domain.UserCollection), ctx: ctx}
}

func (repo *userRepository) Create(user *domain.User) (domain.User, error) {
	user.ID = primitive.NewObjectID().Hex()
	_, err := repo.collection.InsertOne(repo.ctx, user)
	if err != nil {
		return domain.User{}, err
	}
	return *user, nil
}

func (repo *userRepository) Update(userID string, data domain.User) (domain.User, error) {
	currentUser, err := repo.FetchByID(userID)
	if err != nil {
		return domain.User{}, err
	}
	return currentUser, nil
}

func (repo *userRepository) Delete(userID string) error {
	_, err := repo.collection.DeleteOne(repo.ctx, bson.M{"_id": userID})
	return err
}

func (repo *userRepository) FetchByID(userID string) (domain.User, error) {
	var u domain.User
	err := repo.collection.FindOne(repo.ctx, bson.M{"_id": userID}).Decode(&u)
	if err != nil {
		return domain.User{}, err
	}
	return u, err
}

func (repo *userRepository) FetchByEmail(email string) (domain.User, error) {
	var u domain.User
	err := repo.collection.FindOne(repo.ctx, bson.M{"email": email}).Decode(&u)
	if err != nil {
		return domain.User{}, err
	}
	return u, err
}

func (repo *userRepository) FetchByUsername(username string) (domain.User, error) {
	var u domain.User
	err := repo.collection.FindOne(repo.ctx, bson.M{"username": username}).Decode(&u)
	if err != nil {
		return domain.User{}, err
	}
	return u, err
}

func (repo *userRepository) Fetch() ([]domain.User, error) {
	resp, err := repo.collection.Find(repo.ctx, bson.M{})
	if err != nil {
		return []domain.User{}, err
	}
	users := []domain.User{}
	for resp.Next(repo.ctx) {
		var u domain.User
		if err := resp.Decode(&u); err != nil {
			return []domain.User{}, err
		}
		users = append(users, u)
	}
	return users, nil
}
