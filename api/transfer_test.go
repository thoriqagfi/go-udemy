package api

import (
	"bytes"
	"database/sql"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/require"
	mockdb "go-udemy.sqlc.dev/app/db/mock"
	db "go-udemy.sqlc.dev/app/db/sqlc"
	"go-udemy.sqlc.dev/app/util"
)

func TestGetTransferAPI(t *testing.T) {
	transfer := randomTransfer()

	testCases := []struct {
		name string
		transferID int64
		buildStubs func(store *mockdb.MockStore)
		checkResponse func(t *testing.T, recorder *httptest.ResponseRecorder)
	} {
		{
			name: "OK",
			transferID: transfer.ID,
			buildStubs: func(store *mockdb.MockStore) {
				store.EXPECT().
				GetTransfer(
					gomock.Any(),
					gomock.Eq(transfer.ID)).
					Times(1).Return(transfer, nil)
			},
			checkResponse: func(t *testing.T, recorder *httptest.ResponseRecorder) {
				require.Equal(t, http.StatusOK, recorder.Code)
				requireBodyMatchTransfer(t, recorder.Body, transfer)
			},
		},
		{
			name: "NotFound",
			transferID: transfer.ID,
			buildStubs: func(store *mockdb.MockStore) {
				store.EXPECT().
				GetTransfer(
					gomock.Any(),
					gomock.Eq(transfer.ID)).
					Times(1).
					Return(db.Transfer{}, sql.ErrNoRows)
			},
			checkResponse: func(t *testing.T, recorder *httptest.ResponseRecorder) {
				require.Equal(t, http.StatusNotFound, recorder.Code)
			},
		},
		{
			name: "InternalError",
			transferID: transfer.ID,
			buildStubs: func(store *mockdb.MockStore) {
				store.EXPECT().
				GetTransfer(
					gomock.Any(),
					gomock.Eq(transfer.ID)).
					Times(1).
					Return(db.Transfer{}, sql.ErrConnDone)
			},
			checkResponse: func(t *testing.T, recorder *httptest.ResponseRecorder) {
				require.Equal(t, http.StatusInternalServerError, recorder.Code)
			},
		},
		{
			name: "InvalidID",
			transferID: 0,
			buildStubs: func(store *mockdb.MockStore) {
				store.EXPECT().
				GetTransfer(
					gomock.Any(),
					gomock.Any()).
					Times(0)
			},
			checkResponse: func(t *testing.T, recorder *httptest.ResponseRecorder) {
				require.Equal(t, http.StatusBadRequest, recorder.Code)
			},
		},
	}

	for i := range testCases {
		tc := testCases[i]

		t.Run(tc.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			store := mockdb.NewMockStore(ctrl)
			tc.buildStubs(store)

			// Start test server and send request
			server := NewServer(store)
			recorder := httptest.NewRecorder()

			url := fmt.Sprintf("/transfers/%d", tc.transferID)
			request, err := http.NewRequest(http.MethodGet, url, nil)
			if err != nil {
				t.Error(err)
			}

			require.NoError(t, err)

			server.router.ServeHTTP(recorder, request)
		})
	}
}

func randomTransfer() db.Transfer {
	account1 := randomAccount()
	account2 := randomAccount()

	return db.Transfer {
		ID: int64(util.RandomInt(1, 1000)),
		FromAccountID: sql.NullInt64{Int64: account1.ID, Valid: true},
		ToAccountID: sql.NullInt64{Int64: account2.ID, Valid: true},
		Amount: int64(util.RandomMoney()),
	}
}

func requireBodyMatchTransfer(t *testing.T, body *bytes.Buffer, transfer db.Transfer) {
	data, err := ioutil.ReadAll(body)
	require.NoError(t, err)

	var gotTransfer db.Transfer
  err = json.Unmarshal(data, &gotTransfer)
	require.NoError(t, err)
	require.Equal(t, transfer, gotTransfer)
}