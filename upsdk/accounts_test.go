package upsdk

import (
	"fmt"
	"net/http"
	"reflect"
	"testing"
)

func TestAcountsService_List(t *testing.T) {
	setup()
	defer teardown()

	mux.HandleFunc("/accounts", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "GET")
		x := ResponseMessage{Body: []Account{{ID: string("123")}, {ID: string("1234")}}}
		fmt.Fprint(w, createJsonStringFromInterface(x))
	})

	accounts, _, err := client.Accounts.List()
	if err != nil {
		t.Errorf("Accounts.List returned error: %v", err)
	}

	want := []Account{{ID: string("123")}, {ID: string("1234")}}
	if !reflect.DeepEqual(accounts, want) {
		t.Errorf("Account.List returned %+v, want %+v", accounts, want)
	}
}
