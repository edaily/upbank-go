package upsdk

import (
	"fmt"
	"net/http"
	"reflect"
	"testing"
)

func TestCategoriesService_List(t *testing.T) {
	setup()
	defer teardown()

	mux.HandleFunc("/categories", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "GET")
		x := ResponseMessage{Body: []Category{{ID: string("123")}, {ID: string("1234")}}}
		fmt.Fprint(w, createJsonStringFromInterface(x))
	})

	categories, _, err := client.Categories.List()
	if err != nil {
		t.Errorf("Categories.List returned error: %v", err)
	}

	want := []Category{{ID: string("123")}, {ID: string("1234")}}
	if !reflect.DeepEqual(categories, want) {
		t.Errorf("Categories.List returned %+v, want %+v", categories, want)
	}
}
