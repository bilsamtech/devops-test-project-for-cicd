package main

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gorilla/mux"
)

func Unit_TestGetItems(t *testing.T) {
	req, err := http.NewRequest("GET", "/api/items", nil)
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(GetItems)

	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusOK)
	}

	var got []Item
	err = json.Unmarshal(rr.Body.Bytes(), &got)
	if err != nil {
		t.Fatalf("Could not unmarshal response: %v", err)
	}

	want := []Item{
		{ID: "1", Name: "Item 1"},
		{ID: "2", Name: "Item 2"},
	}

	if len(got) != len(want) {
		t.Errorf("handler returned unexpected number of items: got %d want %d", len(got), len(want))
	}

	for i, item := range got {
		if item != want[i] {
			t.Errorf("handler returned unexpected item at index %d: got %v want %v", i, item, want[i])
		}
	}
}

func Unit_TestGetItem(t *testing.T) {
	req, err := http.NewRequest("GET", "/api/items/1", nil)
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	router := mux.NewRouter()
	router.HandleFunc("/api/items/{id}", GetItem)
	router.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusOK)
	}

	var got Item
	err = json.Unmarshal(rr.Body.Bytes(), &got)
	if err != nil {
		t.Fatalf("Could not unmarshal response: %v", err)
	}

	want := Item{ID: "1", Name: "Item 1"}

	if got != want {
		t.Errorf("handler returned unexpected item: got %v want %v", got, want)
	}
}

func Unit_TestGetItemNotFound(t *testing.T) {
	req, err := http.NewRequest("GET", "/api/items/999", nil)
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	router := mux.NewRouter()
	router.HandleFunc("/api/items/{id}", GetItem)
	router.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusNotFound {
		t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusNotFound)
	}
}
