package upsdk

import (
	"fmt"
	"net/http"
	"reflect"
	"testing"
)

func TestTagsService_List(t *testing.T) {
	setup()
	defer teardown()

	mux.HandleFunc("/tags", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "GET")
		x := ResponseMessage{Body: []Tag{{ID: string("123")}, {ID: string("1234")}}}
		fmt.Fprint(w, createJsonStringFromInterface(x))
	})

	tags, _, err := client.Tags.List()
	if err != nil {
		t.Errorf("Tags.List returned error: %v", err)
	}

	want := []Tag{{ID: string("123")}, {ID: string("1234")}}
	if !reflect.DeepEqual(tags, want) {
		t.Errorf("Tags.List returned %+v, want %+v", tags, want)
	}
}
