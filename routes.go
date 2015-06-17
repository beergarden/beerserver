package main

import (
	"encoding/json"
	"gopkg.in/mgo.v2"
	"log"
	"net/http"
)

// -- Routes
// TODO: Store these in context.
type Routes struct {
	channelStore   ChannelStore
	datapointStore DatapointStore
	mailer         Mailer
}

func NewRoutes(db *mgo.Database, mailer Mailer) *Routes {
	channelStore := &channelStore{db}
	datapointStore := &datapointStore{db}
	return &Routes{channelStore, datapointStore, mailer}
}

func (routes *Routes) GetChannel(w http.ResponseWriter, r *http.Request, params RouteParams) error {
	id := params["id"]
	channel, err := routes.channelStore.Get(id)
	if err == mgo.ErrNotFound {
		log.Printf("Channel not found for id: %v", id)
		http.NotFound(w, r)
		return nil
	}
	if err != nil {
		log.Printf("Failed to get channel with id: %v", id)
		log.Println(err)
		return err
	}

	return sendJSON(w, channel)
}

func (routes *Routes) GetChannels(w http.ResponseWriter, r *http.Request, params RouteParams) error {
	channels, err := routes.channelStore.GetAll()
	if err != nil {
		log.Printf("Failed to get channels")
		log.Println(err)
		return err
	}

	return sendJSON(w, channels)
}

func (routes *Routes) CreateChannel(w http.ResponseWriter, r *http.Request, params RouteParams) error {
	decoder := json.NewDecoder(r.Body)
	var channel *Channel
	err := decoder.Decode(&channel)
	if err != nil {
		// TODO: 400
		log.Printf("Failed to decode JSON body: %v", r.Body)
		log.Println(err)
		return err
	}

	err = routes.channelStore.Create(channel)
	if err != nil {
		log.Printf("Failed to create channel: %v", channel)
		log.Println(err)
		return err
	}

	w.WriteHeader(201)
	return sendJSON(w, channel)
}

func (routes *Routes) GetDatapoints(w http.ResponseWriter, r *http.Request, params RouteParams) error {
	if !routes.checkChannelId(w, r, params) {
		return nil
	}
	channelId := params["channelId"]
	datapoints, err := routes.datapointStore.GetAllOfChannel(channelId)
	if err != nil {
		log.Printf("Failed to get datapoints for channel: %v", channelId)
		log.Println(err)
		return err
	}

	return sendJSON(w, datapoints)
}

func (routes *Routes) CreateDatapoint(w http.ResponseWriter, r *http.Request, params RouteParams) error {
	if !routes.checkChannelId(w, r, params) {
		return nil
	}
	channelId := params["channelId"]

	decoder := json.NewDecoder(r.Body)
	var datapoint *Datapoint
	err := decoder.Decode(&datapoint)
	if err != nil {
		// TODO: 400
		log.Printf("Failed to decode JSON body: %v", r.Body)
		log.Println(err)
		return err
	}

	err = routes.datapointStore.Create(channelId, datapoint)
	if err != nil {
		log.Printf("Failed to create datapoint: %v", datapoint)
		log.Println(err)
		return err
	}

	if datapoint.Value > 26.0 {
		go routes.sendNotification(channelId, datapoint)
	}

	w.WriteHeader(201)
	return sendJSON(w, datapoint)
}

// TODO: Implement as a middleware?
func (routes *Routes) checkChannelId(w http.ResponseWriter, r *http.Request, params RouteParams) bool {
	channelId := params["channelId"]
	_, err := routes.channelStore.Get(channelId)
	if err == mgo.ErrNotFound {
		log.Printf("Channel not found for id: %v", channelId)
		http.NotFound(w, r)
		return false
	}
	return true
}

func sendJSON(w http.ResponseWriter, data interface{}) error {
	w.Header().Add("Content-Type", "application/json")
	return json.NewEncoder(w).Encode(data)
}

func (routes *Routes) sendNotification(channelId string, datapoint *Datapoint) error {
	channel, err := routes.channelStore.Get(channelId)
	if err != nil {
		log.Printf("Failed to get channel for id: %v", channelId)
		log.Println(err)
		return err
	}

	return routes.mailer.SendNotification(channel, datapoint)
}
