package main

import (
	"gopkg.in/mgo.v2"
)

type Channel struct {
	Name string
}

type Channels []Channel

type ChannelStore interface {
	GetAll() (Channels, error)
	Create(name string) (*Channel, error)
}

type channelStore struct {
	db *mgo.Database
}

func (store *channelStore) GetAll() (Channels, error) {
	var channels Channels
	err := store.db.C(ChannelCollectionName).Find(nil).All(&channels)
	if err != nil {
		return nil, err
	}
	// Return an empty array when no data.
	if channels == nil {
		channels = make([]Channel, 0)
	}

	return channels, nil
}

func (store *channelStore) Create(channel *Channel) (*Channel, error) {
	c := store.db.C(ChannelCollectionName)

	err := c.Insert(channel)
	if err != nil {
		return nil, err
	}

	return channel, nil
}
