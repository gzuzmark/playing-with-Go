package main

import (
	"./app"
	"log"
	"net/http"
)

var privKey []byte

func init() {
	key, err := ioutil.ReadFile("id_rsa")
	if err != nil {
		log.Fatal(err)
	}
	privKey = key
}

func main() {
	var configPath = "config.json"

	a, err := app.NewApp(configPath)
	if err != nil {
		panic(err)
	}

	defer func() {
		a.Connection.Session.Close()
	}()

	log.Printf("Running on port %s", a.Config.Port)
	if err := http.ListenAndServe(a.Config.Port, nil); err != nil {
		log.Println(err.Error())
	}
}
