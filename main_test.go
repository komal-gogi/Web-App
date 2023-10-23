package main

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gorilla/mux"
)

func TestCreateEvent(t *testing.T) {
	router := mux.NewRouter()
	router.HandleFunc("/event", createEvent).Methods("POST")
	ts := httptest.NewServer(router)
	defer ts.Close()

	newEvent := event{
		ID:          "3",
		Title:       "Test Event",
		Description: "Testing createEvent",
	}

	jsonEvent, err := json.Marshal(newEvent)
	if err != nil {
		t.Fatal(err)
	}

	resp, err := http.Post(ts.URL+"/event", "application/json", bytes.NewBuffer(jsonEvent))
	if err != nil {
		t.Fatal(err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusCreated {
		t.Errorf("Expected status 201, got %d", resp.StatusCode)
	}

	var createdEvent event
	err = json.NewDecoder(resp.Body).Decode(&createdEvent)
	if err != nil {
		t.Fatal(err)
	}

	if createdEvent.ID != newEvent.ID || createdEvent.Title != newEvent.Title || createdEvent.Description != newEvent.Description {
		t.Errorf("Created event does not match the expected event")
	}
}

func TestGetAllEvents(t *testing.T) {
	router := mux.NewRouter()
	router.HandleFunc("/events", getAllEvents).Methods("GET")
	ts := httptest.NewServer(router)
	defer ts.Close()

	resp, err := http.Get(ts.URL + "/events")
	if err != nil {
		t.Fatal(err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		t.Errorf("Expected status 200, got %d", resp.StatusCode)
	}

	var eventList allEvents
	err = json.NewDecoder(resp.Body).Decode(&eventList)
	if err != nil {
		t.Fatal(err)
	}

	if len(eventList) != len(events) {
		t.Errorf("Expected %d events, got %d", len(events), len(eventList))
	}
}

func TestGetOneEvent(t *testing.T) {
	router := mux.NewRouter()
	router.HandleFunc("/event/{id}", getOneEvent).Methods("GET")
	ts := httptest.NewServer(router)
	defer ts.Close()

	targetEvent := events[0]

	resp, err := http.Get(ts.URL + "/event/" + targetEvent.ID)
	if err != nil {
		t.Fatal(err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		t.Errorf("Expected status 200, got %d", resp.StatusCode)
	}

	var retrievedEvent event
	err = json.NewDecoder(resp.Body).Decode(&retrievedEvent)
	if err != nil {
		t.Fatal(err)
	}

	if retrievedEvent.ID != targetEvent.ID {
		t.Errorf("Expected event with ID %s, got ID %s", targetEvent.ID, retrievedEvent.ID)
	}
}

func TestUpdateEvent(t *testing.T) {
	router := mux.NewRouter()
	router.HandleFunc("/event/{id}", updateEvent).Methods("PATCH")
	ts := httptest.NewServer(router)
	defer ts.Close()

	targetEvent := events[0]
	updatedEvent := event{
		ID:          targetEvent.ID,
		Title:       "Updated Event",
		Description: "Testing updateEvent",
	}

	jsonEvent, err := json.Marshal(updatedEvent)
	if err != nil {
		t.Fatal(err)
	}

	req, err := http.NewRequest("PATCH", ts.URL+"/event/"+targetEvent.ID, bytes.NewBuffer(jsonEvent))
	if err != nil {
		t.Fatal(err)
	}

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		t.Fatal(err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		t.Errorf("Expected status 200, got %d", resp.StatusCode)
	}

	var updated event
	err = json.NewDecoder(resp.Body).Decode(&updated)
	if err != nil {
		t.Fatal(err)
	}

	if updated.ID != updatedEvent.ID || updated.Title != updatedEvent.Title || updated.Description != updatedEvent.Description {
		t.Errorf("Updated event does not match the expected event")
	}
}

func TestDeleteEvent(t *testing.T) {
	router := mux.NewRouter()
	router.HandleFunc("/event/{id}", deleteEvent).Methods("DELETE")
	ts := httptest.NewServer(router)
	defer ts.Close()

	targetEvent := events[0]

	req, err := http.NewRequest("DELETE", ts.URL+"/event/"+targetEvent.ID, nil)
	if err != nil {
		t.Fatal(err)
	}

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		t.Fatal(err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		t.Errorf("Expected status 200, got %d", resp.StatusCode)
	}

	// Check that the event has been deleted
	for _, e := range events {
		if e.ID == targetEvent.ID {
			t.Errorf("Event with ID %s was not deleted", targetEvent.ID)
		}
	}
}

func TestMain(m *testing.M) {
	// Initialize any setup code here if needed.
	// Run the tests.
	code := m.Run()
	// Perform any teardown or cleanup here if necessary.
	os.Exit(code)
}
