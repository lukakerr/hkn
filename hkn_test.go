package hkn

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"reflect"
	"testing"
)

var (
	mux    *http.ServeMux
	client *Client
	server *httptest.Server
)

func setup() {
	mux = http.NewServeMux()
	server = httptest.NewServer(mux)
	client = NewClient()

	client.BaseURL = server.URL
}

func teardown() {
	server.Close()
}

func TestGetItem(t *testing.T) {
	setup()
	defer teardown()

	jsonItem := `{
          "by" : "dhouston",
          "descendants" : 71,
          "id" : 8863,
          "kids" : [9224, 8952, 8917],
          "score" : 104,
          "time" : 1175714200,
          "title" : "My YC app: Dropbox - Throw away your USB drive",
          "type" : "story",
          "url" : "http://www.getdropbox.com/u/2/screencast.html"
        }`

	mux.HandleFunc("/item/8863.json", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, jsonItem)
	})

	expected := Item{}

	_ = json.Unmarshal([]byte(jsonItem), &expected)

	item, err := client.GetItem(8863)

	if err != nil {
		t.Errorf("Error for GetItem(8863) should have been nil. Was: %v", err)
	}

	if !reflect.DeepEqual(item, expected) {
		t.Errorf("GetItem(8863) returned %+v, was expecting %+v", item, expected)
	}
}

func TestGetUser(t *testing.T) {
	setup()
	defer teardown()

	jsonUser := `{
          "about" : "This is a test",
          "created" : 1173923446,
          "id" : "jl",
          "karma" : 4094,
          "submitted": [18498213, 16659709, 16659632, 16659556]
        }`

	mux.HandleFunc("/user/jl.json", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, jsonUser)
	})

	expected := User{}

	_ = json.Unmarshal([]byte(jsonUser), &expected)

	user, err := client.GetUser("jl")

	if err != nil {
		t.Errorf("Error for GetUser('jl') should have been nil. Was: %v", err)
	}

	if !reflect.DeepEqual(user, expected) {
		t.Errorf("GetUser('jl') returned %+v, was expecting %+v", user, expected)
	}
}

func TestGetMaxItemId(t *testing.T) {
	setup()
	defer teardown()

	maxItem := `123456`

	mux.HandleFunc("/maxitem.json", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, maxItem)
	})

	var expected int

	_ = json.Unmarshal([]byte(maxItem), &expected)

	id, err := client.GetMaxItemID()

	if err != nil {
		t.Errorf("Error for GetMaxItemId() should have been nil. Was: %v", err)
	}

	if !reflect.DeepEqual(id, expected) {
		t.Errorf("GetMaxItemId() returned %+v, was expecting %+v", id, expected)
	}
}

func TestGetUpdates(t *testing.T) {
	setup()
	defer teardown()

	jsonUpdates := `{
                "items" : [8423305, 8420805, 8423379, 8422504],
                "profiles" : ["thefox", "mdda", "plinkplonk", "GBond", "rqebmm", "neom"]
        }`

	mux.HandleFunc("/updates.json", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, jsonUpdates)
	})

	var expected Updates

	_ = json.Unmarshal([]byte(jsonUpdates), &expected)

	updates, err := client.GetUpdates()

	if err != nil {
		t.Errorf("Error for GetUpdates() should have been nil. Was: %v", err)
	}

	if !reflect.DeepEqual(updates, expected) {
		t.Errorf("GetUpdates() returned %+v, was expecting %+v", updates, expected)
	}
}
