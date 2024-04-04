package db

import (
	"RyanFin/GoSimpleBank/util"
	"context"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestCreateEntry(t *testing.T) {
	createRandomEntry(t)
}

func TestGetEntry(t *testing.T) {
	entry, err := createRandomEntry(t)
	assert.NoError(t, err)

	entry2, err := testQueries.GetEntry(context.Background(), entry.ID)
	assert.NoError(t, err)

	// assert equals
	assert.Equal(t, entry.AccountID, entry2.AccountID)
	assert.Equal(t, entry.Amount, entry2.Amount)

	// assert not empty
	assert.NotEmpty(t, entry2.ID)
	assert.NotEmpty(t, entry2.CreatedAt)

	// assert retrieved within the a second of each other
	assert.WithinDuration(t, entry.CreatedAt, entry2.CreatedAt, 1*time.Second)
}

func createRandomEntry(t *testing.T) (Entry, error) {
	account, _ := util.RandomAccountIDs()

	args := CreateEntryParams{
		AccountID: account,
		Amount:    util.RandomMoney(),
	}

	entry, err := testQueries.CreateEntry(context.Background(), args)
	assert.NoError(t, err)

	// check not empty
	assert.NotEmpty(t, entry)

	// check values are equal
	assert.Equal(t, args.AccountID, entry.AccountID)
	assert.Equal(t, args.Amount, entry.Amount)

	return entry, nil
}

func TestListEntries(t *testing.T) {
	for i := 0; i < 10; i++ {
		createRandomEntry(t)
	}

	args := ListEntriesParams{
		Limit:  5,
		Offset: 5,
	}

	entries, err := testQueries.ListEntries(context.Background(), args)
	assert.NoError(t, err)
	assert.Len(t, entries, 5)

	for _, entry := range entries {
		assert.NotEmpty(t, entry)
	}
}
