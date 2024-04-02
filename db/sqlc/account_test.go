package db

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
)
func TestCreateAccount(t *testing.T){
// create account params must be created
args := CreateAccountParams{
	Owner: "ryan",
	Balance: 1500,
	Currency: "GBP",
}

account, err := testQueries.CreateAccount(context.Background(), args)
assert.NoError(t, err)
assert.NotEmpty(t, account)

assert.Equal(t, args.Owner, account.Owner)
assert.Equal(t, args.Balance, account.Balance)
assert.Equal(t, args.Currency, account.Currency)

assert.NotZero(t, account.ID)
assert.NotZero(t, account.CreatedAt)
}