package upsdk

import (
	"fmt"
	"net/http"
	"reflect"
	"testing"
)

func TestTransactionsService_List(t *testing.T) {
	setup()
	defer teardown()

	mux.HandleFunc("/transactions", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "GET")
		x := ResponseMessage{Body: []Transaction{{ID: string("123")}, {ID: string("1234")}}}
		fmt.Fprint(w, createJsonStringFromInterface(x))
	})

	transactions, _, err := client.Transactions.List()
	if err != nil {
		t.Errorf("Transactions.List returned error: %v", err)
	}

	want := []Transaction{{ID: string("123")}, {ID: string("1234")}}
	if !reflect.DeepEqual(transactions, want) {
		t.Errorf("Transactions.List returned %+v, want %+v", transactions, want)
	}
}
