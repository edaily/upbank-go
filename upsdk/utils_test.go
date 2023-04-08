package upsdk

import (
	"fmt"
	"net/http"
	"reflect"
	"testing"
)

func TestUtilsService_List(t *testing.T) {
	setup()
	defer teardown()

	mux.HandleFunc("/util/ping", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "GET")
		x := PingResponse{Ping{ID: string("123")}}
		fmt.Fprint(w, createJsonStringFromInterface(x))
	})

	ping, _, err := client.Utils.Ping()
	if err != nil {
		t.Errorf("Utils.Ping returned error: %v", err)
	}

	want := &Ping{ID: string("123")}
	if !reflect.DeepEqual(ping, want) {
		t.Errorf("Utils.Ping returned %+v, want %+v", ping, want)
	}
}
