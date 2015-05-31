package main

import (
	"github.com/joho/godotenv"
	"gopkg.in/mgo.v2"
	"log"
	"net/http"
	"os"
	// "gopkg.in/mgo.v2/bson"
)

const DefaultPort = ":3000"
const EnvPort = "PORT"
const DatabaseName = "beerserver"

func main() {
	// Load environmental variables.
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	// Connect to database.
	session, err := mgo.Dial(os.Getenv("MONGOHQ_URL"))
	if err != nil {
		panic(err)
	}
	defer session.Close()
	// ??
	session.SetMode(mgo.Monotonic, true)

	// Set up router.
	routes := NewRoutes(session.DB(DatabaseName))
	router := NewRouter()
	router.GET("/", routes.Index)
	router.GET("/channels", routes.GetChannels)
	router.POST("/channels", routes.CreateChannel)
	router.GET("/channels/{id}", routes.GetChannel)
	router.GET("/channels/{channelId}/datapoints", routes.GetDatapoints)
	router.POST("/channels/{channelId}/datapoints", routes.CreateDatapoint)

	// Start server.
	port := DefaultPort
	if os.Getenv(EnvPort) != "" {
		port = ":" + os.Getenv(EnvPort)
	}
	log.Fatal(http.ListenAndServe(port, router))

	log.Println("Listening on %v", port)
}
