package api

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/golang/mock/gomock"
	mockdb "github.com/pegasus7d/simplebank/db/mock"
	db "github.com/pegasus7d/simplebank/db/sqlc"
	"github.com/pegasus7d/simplebank/db/util"
	"github.com/stretchr/testify/require"
)

func TestGetAccountAPI(t *testing.T){
	account:=RandomAccount()

	testCases:=[]struct{
		name string 
		accountID int64
		buldSubs func(store *mockdb.MockStore)
		checkResponse func(t *testing.T,recorder *httptest.ResponseRecorder)
	}{
		{
			name: "OK",
			accountID:account.ID,
			buldSubs: func(store *mockdb.MockStore){

			},
			checkResponse: func(t *testing.T,recorder *httptest.ResponseRecorder){
				require.Equal(t,http.StatusOK,recorder.Code)
				requiredBodyMatchAccount(t,recorder.Body,account)
			},

		},
		

	}

	for i :=range testCases{

		tc:=testCases[i]
		t.Run(tc.name,func(t *testing.T) {
			ctrl:=gomock.NewController(t)
			defer ctrl.Finish()
			store:=mockdb.NewMockStore(ctrl)
			tc.buldSubs(store)

			store.EXPECT().
			GetAccount(gomock.Any(),gomock.Eq(account.ID)).
			Times(1).
			Return(account,nil)


			server:=NewServer(store)

			recorder:=httptest.NewRecorder()

			url:=fmt.Sprintf("/accounts/%d",account.ID)

			request,err:=http.NewRequest(http.MethodGet,url,nil)
			require.NoError(t,err)

			server.router.ServeHTTP(recorder,request)
			tc.checkResponse(t,recorder)
		})
		

	}
	




}

func RandomAccount()db.Account{
	return db.Account{
		ID: util.RandomInt(1,1000),
		Owner: util.RandomOwner(),
		Balance: util.RandomMoney(),
		Currency: util.RandomCurrency(),
	}
}


func requiredBodyMatchAccount(t *testing.T,body *bytes.Buffer,account db.Account){
	data,err:= ioutil.ReadAll(body)
	require.NoError(t,err)
	var gotAccount db.Account
	err=json.Unmarshal(data,&gotAccount)
	require.NoError(t,err)
	require.Equal(t,account,gotAccount)
}
