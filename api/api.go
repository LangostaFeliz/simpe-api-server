package api

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/gorilla/mux"
)

type allEvents []event

var events = allEvents{
	{
		ID:          "1",
		Title:       "This introduction to Golang language",
		Description: "Come join us for a chance to learn how golang works",
	},
	{
		ID:          "2",
		Title:       "This introduction to java language",
		Description: "Come join us for a chance to learn how java works",
	},
}

func CreateEvent(w http.ResponseWriter, r *http.Request) {
	var newEvent event
	reqBody, err := ioutil.ReadAll(r.Body)

	if err != nil {
		fmt.Fprintf(w, "Kindly enter data with the event title and description only in order to update")
	}

	json.Unmarshal(reqBody, &newEvent)
	events = append(events, newEvent)

	json.NewEncoder(w).Encode(newEvent)
}

func GetOneEvent(w http.ResponseWriter, r *http.Request) {
	eventID := mux.Vars(r)["id"]
	for _, singleEvent := range events {
		if singleEvent.ID == eventID {
			json.NewEncoder(w).Encode(singleEvent)
		}
	}
}

func GetAllEvent(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(events)
}
