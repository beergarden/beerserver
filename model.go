package main

import (
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
	"time"
)

const ChannelCollectionName = "channels"
const DatapointCollectionName = "datapoints"

// -- Channel
type Channel struct {
	Id    bson.ObjectId `bson:"_id" json:"id"`
	Name  string        `bson:"name" json:"name"`
	Email string        `bson:"email" json:"email"`
}

type Channels []Channel

type ChannelStore interface {
	Get(id string) (*Channel, error)
	GetAll() (Channels, error)
	Create(channel *Channel) error
}

type channelStore struct {
	db *mgo.Database
}

func (store *channelStore) Get(id string) (*Channel, error) {
	channel := &Channel{}
	if !bson.IsObjectIdHex(id) {
		return nil, mgo.ErrNotFound
	}
	err := store.collection().FindId(bson.ObjectIdHex(id)).One(channel)
	if err != nil {
		return nil, err
	}

	return channel, nil
}

func (store *channelStore) GetAll() (Channels, error) {
	var channels Channels
	err := store.collection().Find(nil).All(&channels)
	if err != nil {
		return nil, err
	}
	if channels == nil {
		channels = make(Channels, 0)
	}

	return channels, nil
}

func (store *channelStore) Create(channel *Channel) error {
	// TODO: Validate input.
	channel.Id = bson.NewObjectId()
	err := store.collection().Insert(channel)
	if err != nil {
		return err
	}

	return nil
}

func (store *channelStore) collection() *mgo.Collection {
	return store.db.C(ChannelCollectionName)
}

// -- Datapoint
type Datapoint struct {
	Id        bson.ObjectId `bson:"_id" json:"id"`
	ChannelId bson.ObjectId `bson:"channelId" json:"channelId"`
	Value     float32       `bson:"value" json:"value"`
	At        time.Time     `bson:"at" json:"at"`
}

type Datapoints []Datapoint

type DatapointStore interface {
	GetAllOfChannel(channelId string) (Datapoints, error)
	Create(channelId string, datapoint *Datapoint) error
}

type datapointStore struct {
	db *mgo.Database
}

func (store *datapointStore) GetAllOfChannel(channelId string) (Datapoints, error) {
	if !bson.IsObjectIdHex(channelId) {
		return nil, mgo.ErrNotFound
	}
	var datapoints Datapoints
	query := bson.M{"channelId": bson.ObjectIdHex(channelId)}
	err := store.collection().Find(query).All(&datapoints)
	if err != nil {
		return nil, err
	}
	if datapoints == nil {
		datapoints = make(Datapoints, 0)
	}

	return datapoints, nil
}

func (store *datapointStore) Create(channelId string, datapoint *Datapoint) error {
	// TODO: Validate input.
	if !bson.IsObjectIdHex(channelId) {
		return mgo.ErrNotFound
	}
	datapoint.ChannelId = bson.ObjectIdHex(channelId)
	datapoint.Id = bson.NewObjectId()
	err := store.collection().Insert(datapoint)
	if err != nil {
		return err
	}

	return nil
}

func (store *datapointStore) collection() *mgo.Collection {
	return store.db.C(DatapointCollectionName)
}
