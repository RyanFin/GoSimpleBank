package db

import (
	"RyanFin/GoSimpleBank/util"
	"context"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func createRandomUser(t *testing.T) User {

	hashedPassword, err := util.HashPassword(util.RandomString(6))
	assert.NoError(t, err)
	// create account params must be created
	args := CreateUserParams{
		Username:       util.RandomOwner(),
		HashedPassword: hashedPassword,
		FullName:       util.RandomOwner(),
		Email:          util.RandomEmail(),
	}

	user, err := testQueries.CreateUser(context.Background(), args)
	assert.NoError(t, err)
	assert.NotEmpty(t, user)

	assert.Equal(t, args.Username, user.Username)
	assert.Equal(t, args.HashedPassword, user.HashedPassword)
	assert.Equal(t, args.FullName, user.FullName)
	assert.Equal(t, args.Email, user.Email)

	assert.True(t, user.PasswordChangedAt.IsZero())
	assert.NotZero(t, user.CreatedAt)

	return user
}

// test random account creation
func TestCreateUser(t *testing.T) {
	createRandomUser(t)
}

// test retrieving an account from the postgres db (creates a new instance of the account obj)
func TestGetUser(t *testing.T) {
	user1 := createRandomUser(t)
	user2, err := testQueries.GetUser(context.Background(), user1.Username)
	assert.NoError(t, err)
	assert.NotEmpty(t, user2)

	// compare all fields
	require.Equal(t, user1.Username, user2.Username)
	require.Equal(t, user1.HashedPassword, user2.HashedPassword)
	require.Equal(t, user1.FullName, user2.FullName)
	require.Equal(t, user1.Email, user2.Email)

	// test that both accounts were created wihin a second of each other
	require.WithinDuration(t, user1.PasswordChangedAt, user2.PasswordChangedAt, time.Second)
	require.WithinDuration(t, user1.CreatedAt, user2.CreatedAt, time.Second)
}
