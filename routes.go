package main

import (
	"encoding/json"
	"fmt"
	"gopkg.in/mgo.v2"
	"html"
	"log"
	"net/http"
)

const ChannelCollectionName = "channels"

// -- Routes
type Routes struct {
	channelStore *channelStore
}

func NewRoutes(db *mgo.Database) *Routes {
	channelStore := &channelStore{db}
	return &Routes{channelStore}
}

func (routes *Routes) Index(w http.ResponseWriter, r *http.Request) error {
	fmt.Fprintf(w, "Hello, %q", html.EscapeString(r.URL.Path))
	return nil
}

func (routes *Routes) GetChannels(w http.ResponseWriter, r *http.Request) error {
	channels, err := routes.channelStore.GetAll()
	if err != nil {
		return err
	}

	return sendJSON(w, channels)
}

func (routes *Routes) CreateChannel(w http.ResponseWriter, r *http.Request) error {
	decoder := json.NewDecoder(r.Body)
	var channel *Channel
	err := decoder.Decode(&channel)
	if err != nil {
		// TODO: 400
		return err
	}

	channel, err = routes.channelStore.Create(channel)
	if err != nil {
		return err
	}

	return json.NewEncoder(w).Encode(channel)
}

func sendJSON(w http.ResponseWriter, data interface{}) error {
	w.Header().Add("Content-Type", "application/json")
	return json.NewEncoder(w).Encode(data)
}
