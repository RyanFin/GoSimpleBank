package db

import (
	"RyanFin/GoSimpleBank/util"
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
)
func TestCreateAccount(t *testing.T){
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
}