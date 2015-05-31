package main

import (
	"github.com/joho/godotenv"
	"gopkg.in/mgo.v2"
	"log"
	"net/http"
	"os"
	"regexp"
)

const DefaultPort = ":3000"
const EnvPort = "PORT"

func main() {
	// Load environmental variables.
	err := godotenv.Load()
	if err != nil {
		log.Println("Couldn't read .env file")
	}

	// Connect to database.
	connectionString := os.Getenv("MONGOLAB_URI")
	uriPattern := regexp.MustCompile("^(.+)\\/(\\w+)$")
	matches := uriPattern.FindStringSubmatch(connectionString)
	dbName := matches[2]
	log.Printf("Database URI: %v name: %v", connectionString, dbName)
	session, err := mgo.Dial(connectionString)
	if err != nil {
		panic(err)
	}
	defer session.Close()
	// ??
	session.SetMode(mgo.Monotonic, true)

	// Set up router.
	routes := NewRoutes(session.DB(dbName))
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
