package main

import (
	"replapp/route"

	"github.com/joho/godotenv"
	log "github.com/sirupsen/logrus"
)

/*
https://gateway.marvel.com:443/v1/public/characters?name=thor&orderBy=name&limit=10&apikey=2fb4abe381548aa744b2f197ffeb9889
*/
func main() {
	log.SetFormatter(&log.JSONFormatter{})
	err := godotenv.Load(".env")
	if err != nil {
		log.Printf("Some error occured. Err: %s", err)
		return
	}
	route.InitRoute()
}
