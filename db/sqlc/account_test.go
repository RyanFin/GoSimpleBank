package db

import (
	"RyanFin/GoSimpleBank/util"
	"context"
	"database/sql"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func createRandomAccount(t *testing.T) Account{
	// create account params must be created
args := CreateAccountParams{
	Owner: util.RandomOwner(),
	Balance: util.RandomMoney(),
	Currency: util.RandomCurrency(),
}

account, err := testQueries.CreateAccount(context.Background(), args)
assert.NoError(t, err)
assert.NotEmpty(t, account)

assert.Equal(t, args.Owner, account.Owner)
assert.Equal(t, args.Balance, account.Balance)
assert.Equal(t, args.Currency, account.Currency)

assert.NotZero(t, account.ID)
assert.NotZero(t, account.CreatedAt)

return account
}

// test random account creation
func TestCreateAccount(t *testing.T){
	createRandomAccount(t)
}

// test retrieving an account from the postgres db (creates a new instance of the account obj)
func TestGetAccount(t *testing.T){
	account1 := createRandomAccount(t)
	account2, err := testQueries.GetAccount(context.Background(), account1.ID)
	assert.NoError(t, err)
	assert.NotEmpty(t, account2)

	// compare all fields
	require.Equal(t, account1.ID, account2.ID)
	require.Equal(t, account1.Owner, account2.Owner)
	require.Equal(t, account1.Balance, account2.Balance)
	require.Equal(t, account1.Currency, account2.Currency)

	// test that both accounts were created wihin a second of each other
	require.WithinDuration(t, account1.CreatedAt, account2.CreatedAt, time.Second)
}

func TestUpdateAccount(t *testing.T){
	account1 := createRandomAccount(t)

	args := UpdateAccountParams{
		ID: account1.ID,
		Balance: util.RandomMoney(),
	}


	account2, err := testQueries.UpdateAccount(context.Background(), args)
	assert.NoError(t, err)
	assert.NotEmpty(t, account2)

	require.Equal(t, account1.ID, account2.ID)
	require.Equal(t, account1.Owner, account2.Owner)
	require.Equal(t, args.Balance, account2.Balance)
	require.Equal(t, account1.Currency, account2.Currency)

	// test that both accounts were created wihin a second of each other
	require.WithinDuration(t, account1.CreatedAt, account2.CreatedAt, time.Second)
}

func TestDeleteAccount(t *testing.T){
	account1 := createRandomAccount(t) 
	err := testQueries.DeleteAccount(context.Background(),account1.ID)
	assert.NoError(t, err)


	account2, err := testQueries.GetAccount(context.Background(), account1.ID)
	assert.Error(t, err)
	assert.EqualError(t, err, sql.ErrNoRows.Error())
	assert.Empty(t, account2)
}

func TestListAccount(t *testing.T){
	for i := 0; i < 10; i++{
		createRandomAccount(t)
	}

	args := ListAccountsParams{
		Limit: 5,
		Offset: 5,
	}

	accounts, err := testQueries.ListAccounts(context.Background(), args)
	assert.NoError(t, err)
	assert.Len(t, accounts, 5)

	for _, account := range accounts{
		assert.NotEmpty(t, account)
	}

}