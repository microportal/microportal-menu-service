package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"github.com/subosito/gotenv"
	"log"
	"microportal-menu-service/auth"
	"microportal-menu-service/controller"
	"net/http"
	"os"
)

var (
	port           string
	pathPrefix     string
	clientName     string
	clientPath     string
	clientImageUrl string
	clientIndexUrl string
	clientStoreUrl string
	clientService  string
)

func init() {
	_ = gotenv.Load()
	port = os.Getenv("PORT")
	pathPrefix = os.Getenv("PATH_PREFIX")

	clientName = os.Getenv("PORTAL_CLIENT_NAME")
	clientPath = os.Getenv("PORTAL_CLIENT_PATH")
	clientImageUrl = os.Getenv("PORTAL_CLIENT_IMAGEURL")
	clientIndexUrl = os.Getenv("PORTAL_CLIENT_INDEXURL")
	clientStoreUrl = os.Getenv("PORTAL_CLIENT_STOREURL")
	clientService = os.Getenv("PORTAL_CLIENT_SERVICE")

	body, err := json.Marshal(map[string]string{
		"name":     clientName,
		"path":     clientPath,
		"imageUrl": clientImageUrl,
		"indexUrl": clientIndexUrl,
		"storeUrl": clientStoreUrl,
		"service":  clientService,
	})
	if err != nil {
		log.Fatal(err)
	}

	response, err := http.Post("http://localhost:9000/core-service/applications", "application/json", bytes.NewBuffer(body))
	if err != nil {
		log.Fatal(err)
	}
	if response.StatusCode != 200 {
		log.Fatal(response.Status)
	}

	defer response.Body.Close()
}

func main() {
	if port == "" {
		port = "8080"
	}
	if pathPrefix == "" {
		pathPrefix = "/menu-service"
	}
	addr := fmt.Sprint(":", port)

	router := mux.NewRouter().PathPrefix(pathPrefix).Subrouter()
	router.Use(auth.JwtAuthentication)

	mc := controller.MenuController{}

	router.HandleFunc("/modules/{id}/menu", mc.FindMenuByModuleID).Methods(http.MethodGet)

	fmt.Println("Server listening on port: ", port)
	log.Fatal(http.ListenAndServe(addr, router))
}
