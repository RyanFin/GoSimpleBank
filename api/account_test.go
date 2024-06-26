package api

import (
	mockdb "RyanFin/GoSimpleBank/db/mock"
	db "RyanFin/GoSimpleBank/db/sqlc"
	"RyanFin/GoSimpleBank/util"
	"bytes"
	"database/sql"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/require"
)

func TestGetAccountAPI(t *testing.T) {
	user, _ := randomUser(t)
	account := randomAccount(user.Username)

	testCases := []struct {
		name          string
		accountID     int64
		buildStubs    func(store *mockdb.MockStore)
		checkResponse func(t *testing.T, recorder *httptest.ResponseRecorder)
	}{
		{"OK", account.ID, func(store *mockdb.MockStore) {
			// build the stubs
			store.EXPECT().
				GetAccount(gomock.Any(), gomock.Eq(account.ID)).
				Times(1).
				Return(account, nil)
		}, func(t *testing.T, recorder *httptest.ResponseRecorder) {
			require.Equal(t, http.StatusOK, recorder.Code)
			requireBodyMatchAccount(t, recorder.Body, account)
		}},
		{"NotFound", account.ID, func(store *mockdb.MockStore) {
			// build the stubs
			store.EXPECT().
				GetAccount(gomock.Any(), gomock.Eq(account.ID)).
				Times(1).
				Return(account, sql.ErrNoRows)
		}, func(t *testing.T, recorder *httptest.ResponseRecorder) {
			require.Equal(t, http.StatusNotFound, recorder.Code)
		}},
		{"InternalError", account.ID, func(store *mockdb.MockStore) {
			// build the stubs
			store.EXPECT().
				GetAccount(gomock.Any(), gomock.Eq(account.ID)).
				Times(1).
				Return(account, sql.ErrConnDone)
		}, func(t *testing.T, recorder *httptest.ResponseRecorder) {
			require.Equal(t, http.StatusInternalServerError, recorder.Code)
		}},
		// minid should > 0
		{"InvalidID", 0, func(store *mockdb.MockStore) {
			// build the stubs
			store.EXPECT().
				GetAccount(gomock.Any(), gomock.Any()).
				Times(0)
		}, func(t *testing.T, recorder *httptest.ResponseRecorder) {
			require.Equal(t, http.StatusBadRequest, recorder.Code)
		}},
	}

	for i := range testCases {
		tc := testCases[i]

		t.Run(tc.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			store := mockdb.NewMockStore(ctrl)
			tc.buildStubs(store)

			server := newTestServer(t, store)
			recorder := httptest.NewRecorder()

			url := fmt.Sprintf("/accounts/%d", tc.accountID)
			request, err := http.NewRequest(http.MethodGet, url, nil)
			require.NoError(t, err)

			server.router.ServeHTTP(recorder, request)
			tc.checkResponse(t, recorder)
		})
	}

}

func randomAccount(owner string) db.Account {
	return db.Account{
		ID:       util.RandomInt(1, 1000),
		Owner:    owner,
		Balance:  util.RandomMoney(),
		Currency: util.RandomCurrency(),
	}
}

func requireBodyMatchAccount(t *testing.T, body *bytes.Buffer, account db.Account) {
	data, err := io.ReadAll(body)
	require.NoError(t, err)

	var gotAccount db.Account
	err = json.Unmarshal(data, &gotAccount)
	require.NoError(t, err)
	require.Equal(t, account, gotAccount)
}
