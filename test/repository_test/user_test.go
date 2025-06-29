package repositories_test

import (
	"7-solutions/dtos"
	"7-solutions/repositories"
	"testing"

	"github.com/stretchr/testify/assert"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/integration/mtest"
)

func TestUserRepository(t *testing.T) {
	mt := mtest.New(t, mtest.NewOptions().ClientType(mtest.Mock))

	mt.Run("TestCountUsers", func(mt *mtest.T) {
		db := mt.Client.Database("testdb")
		repo := repositories.NewUserRepository(db)

		mt.AddMockResponses(mtest.CreateCursorResponse(1, "test.users", mtest.FirstBatch,
			bson.D{{Key: "n", Value: int64(3)}},
		))

		count, err := repo.CountUsers()
		assert.NoError(t, err)
		assert.Equal(t, int64(3), count)
	})

	mt.Run("TestGetUserByID", func(mt *mtest.T) {
		db := mt.Client.Database("testdb")
		repo := repositories.NewUserRepository(db)

		userID := 1
		expectedUser := &dtos.UserRegister{
			Name:     "Test User",
			Email:    "test@user.com",
			Password: "password123",
		}

		mt.AddMockResponses(mtest.CreateCursorResponse(1, "test.users", mtest.FirstBatch,
			bson.D{
				{Key: "id", Value: userID},
				{Key: "name", Value: expectedUser.Name},
				{Key: "email", Value: expectedUser.Email},
				{Key: "password", Value: expectedUser.Password},
			}))

		user, err := repo.GetUserByID(userID)
		assert.NoError(t, err)
		assert.NotNil(t, user)
		assert.Equal(t, expectedUser.Name, user.Name)
		assert.Equal(t, expectedUser.Email, user.Email)
	})

	mt.Run("TestGetAllUsers", func(mt *mtest.T) {
		db := mt.Client.Database("testdb")
		repo := repositories.NewUserRepository(db)

		expectedUsers := []dtos.UserRegister{
			{Name: "User1", Email: "test@user1.com", Password: "password123"},
			{Name: "User2", Email: "test@user2.com", Password: "password456"},
		}

		cursorID := int64(12345)

		mt.AddMockResponses(
			mtest.CreateCursorResponse(cursorID, "test.users", mtest.FirstBatch,
				bson.D{
					{Key: "name", Value: expectedUsers[0].Name},
					{Key: "email", Value: expectedUsers[0].Email},
					{Key: "password", Value: expectedUsers[0].Password},
				},
				bson.D{
					{Key: "name", Value: expectedUsers[1].Name},
					{Key: "email", Value: expectedUsers[1].Email},
					{Key: "password", Value: expectedUsers[1].Password},
				},
			),
			mtest.CreateCursorResponse(0, "test.users", mtest.NextBatch),
		)

		users, err := repo.GetAllUsers()
		assert.NoError(t, err)
		assert.Len(t, users, 2)
		assert.Equal(t, expectedUsers[0].Name, users[0].Name)
		assert.Equal(t, expectedUsers[1].Name, users[1].Name)
	})

	mt.Run("TestUpdateUser", func(mt *mtest.T) {
		db := mt.Client.Database("testdb")
		repo := repositories.NewUserRepository(db)

		userID := 1
		updatedUser := &dtos.UserUpdate{
			Name:  "Updated User",
			Email: "test@updateuser.com",
		}
		mt.AddMockResponses(mtest.CreateSuccessResponse())

		_, err := repo.UpdateUser(userID, updatedUser)
		assert.NoError(t, err)
	})

	mt.Run("TestDeleteUser", func(mt *mtest.T) {
		db := mt.Client.Database("testdb")
		repo := repositories.NewUserRepository(db)

		userID := 1
		mt.AddMockResponses(mtest.CreateSuccessResponse())

		err := repo.DeleteUser(userID)
		assert.NoError(t, err)
	})
}
