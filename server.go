package main

import (
	"encoding/json"
	"errors"
	"github.com/gorilla/mux"
	"github.com/ichtrojan/thoth"
	"github.com/joho/godotenv"
	"log"
	"net/http"
	"os"
)

var (
	logger, _ = thoth.Init("log")
)

type Error struct {
	Message string
}

type Bank struct {
	Name string `json:"name"`
	Slug string `json:"slug"`
	Code string `json:"code"`
	Logo string `json:"logo"`
}

func main() {
	if err := godotenv.Load(); err != nil {
		logger.Log(errors.New("no .env file found"))
		log.Fatal("No .env file found")
	}

	port, exist := os.LookupEnv("PORT")

	if !exist {
		logger.Log(errors.New("PORT not set in .env"))
		log.Fatal("PORT not set in .env")
	}

	route := mux.NewRouter()

	route.PathPrefix("/logo/").Handler(http.StripPrefix("/logo/", http.FileServer(http.Dir("./logos"))))

	route.NotFoundHandler = http.HandlerFunc(notFound)

	route.HandleFunc("/", func(writer http.ResponseWriter, request *http.Request) {

	})

	if err := http.ListenAndServe(":"+port, route); err != nil {
		logger.Log(err)
	}
}

func notFound(writer http.ResponseWriter, request *http.Request) {
	writer.WriteHeader(http.StatusNotFound)

	_ = json.NewEncoder(writer).Encode(Error{
		Message: "endpoint not found",
	})
}
