package repositories_test

import (
	"7-solutions/dtos"
	"7-solutions/repositories"
	"testing"

	"github.com/stretchr/testify/assert"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo/integration/mtest"
)

func TestAuthRepository(t *testing.T) {
	mt := mtest.New(t, mtest.NewOptions().ClientType(mtest.Mock))

	mt.Run("TestRegisterUser", func(mt *mtest.T) {
		db := mt.Client.Database("testdb")
		repo := repositories.NewAuthRepository(db)

		user := &dtos.UserRegister{
			Name:     "Test",
			Email:    "test@user.com",
			Password: "password",
		}

		mt.AddMockResponses(
			bson.D{
				{Key: "ok", Value: 1},
				{Key: "value", Value: bson.D{
					{Key: "_id", Value: "users"},
					{Key: "seq", Value: 1},
				}},
			},
		)

		mt.AddMockResponses(
			bson.D{
				{Key: "ok", Value: 1},
				{Key: "insertedId", Value: primitive.NewObjectID()},
			},
		)

		err := repo.RegisterUser(user)
		assert.NoError(t, err)
	})

	mt.Run("TestAuthenticateUser_InvalidEmail", func(mt *mtest.T) {
		db := mt.Client.Database("testdb")
		repo := repositories.NewAuthRepository(db)

		input := &dtos.UserAuthenticate{
			Email:    "notfound@example.com",
			Password: "password",
		}

		mt.AddMockResponses(mtest.CreateCursorResponse(1, "test.users", mtest.FirstBatch,
			bson.D{{Key: "n", Value: int64(3)}},
		))
		_, err := repo.AuthenticateUser(input)
		assert.Error(t, err)
	})
}
