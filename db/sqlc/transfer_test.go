package db

import (
	"RyanFin/GoSimpleBank/util"
	"context"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestCreateTransfer(t *testing.T) {
	createRandomTransfer(t)
}

// GET a SINGLE transfer
func TestGetTransfer(t *testing.T) {
	// create a transfer first
	transfer, err := createRandomTransfer(t)
	assert.NoError(t, err)

	transfer2, err := testQueries.GetTransfer(context.Background(), transfer.ID)
	assert.NoError(t, err)

	// assert equals
	assert.Equal(t, transfer.FromAccountID, transfer2.FromAccountID)
	assert.Equal(t, transfer.ToAccountID, transfer2.ToAccountID)
	assert.Equal(t, transfer.Amount, transfer2.Amount)

	// assert not empty
	assert.NotEmpty(t, transfer2.ID)
	assert.NotEmpty(t, transfer2.CreatedAt)

	// assert retrieved within the a second of each other
	assert.WithinDuration(t, transfer.CreatedAt, transfer2.CreatedAt, 1*time.Second)
}

// GET ALL transfers
// func TestListTransfers(t *testing.T){

// 	// load multiple tests using a test table
// 	testTransfers := []struct{
// 		fromAccID, toAccID, amount int64
// 	}{
// 		{1, util.RandomInt(2,81), util.RandomMoney()},
// 		{1, util.RandomInt(2,81), util.RandomMoney()},
// 		{1, util.RandomInt(2,81), util.RandomMoney()},
// 		{1, util.RandomInt(2,81), util.RandomMoney()},
// 		{1, util.RandomInt(2,81), util.RandomMoney()},
// 		{1, util.RandomInt(2,81), util.RandomMoney()},
// 		{1, util.RandomInt(2,81), util.RandomMoney()},
// 		{1, util.RandomInt(2,81), util.RandomMoney()},
// 		{1, util.RandomInt(2,81), util.RandomMoney()},
// 	}

// 	for _, transfer := range testTransfers{
// 		args := CreateTransferParams{
// 			FromAccountID: transfer.fromAccID,
// 			ToAccountID: transfer.toAccID,
// 			Amount: transfer.amount,
// 		}

// 		transfer, err := testQueries.CreateTransfer(context.Background(), args)
// 		// check no error
// 		assert.NoError(t, err)
// 		// check not empty
// 		assert.NotEmpty(t, transfer)

// 		// check values are equal
// 		assert.Equal(t, args.FromAccountID, int64(transfer.FromAccountID))
// 		assert.Equal(t, args.FromAccountID, int64(transfer.FromAccountID))
// 		assert.Equal(t, args.FromAccountID, int64(transfer.FromAccountID))

// 		// check id and timestamp not zero
// 		assert.NotZero(t, transfer.ID)
// 		assert.NotZero(t, transfer.CreatedAt)
// 	}

// 	args := ListTransfersParams{
// 		FromAccountID: 1,
// 		Limit: 7,
// 	}

// 	transfers, err := testQueries.ListTransfers(context.Background(), args)
// 	assert.NoError(t, err)

// 	assert.Len(t, transfers, 7)

// 	for _, transfer := range transfers{
// 		assert.NotEmpty(t, transfer)
// 	}
// }

func createRandomTransfer(t *testing.T) (Transfer, error) {
	fromAccountID, toAccountID := util.RandomAccountIDs()
	args := CreateTransferParams{
		FromAccountID: int64(fromAccountID),
		ToAccountID:   int64(toAccountID),
		Amount:        util.RandomMoney(),
	}

	transfer, err := testQueries.CreateTransfer(context.Background(), args)
	// check no error
	// assert.NoError(t, err)
	// check not empty
	// assert.NotEmpty(t, transfer)

	// check values are equal
	assert.Equal(t, int64(args.FromAccountID), int64(transfer.FromAccountID))

	// check id and timestamp not zero
	assert.NotZero(t, transfer.ID)
	assert.NotZero(t, transfer.CreatedAt)

	return transfer, err
}
