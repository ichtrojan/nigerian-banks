package main

import (
	"encoding/json"
	"errors"
	"github.com/gorilla/mux"
	"github.com/ichtrojan/horus"
	"github.com/ichtrojan/thoth"
	"github.com/joho/godotenv"
	"github.com/rs/cors"
	"io/ioutil"
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

type BankJSON struct {
	Name string `json:"name"`
	Slug string `json:"slug"`
	Code string `json:"code"`
	USSD string `json:"ussd"`
}

type Bank struct {
	Name string `json:"name"`
	Slug string `json:"slug"`
	Code string `json:"code"`
	USSD string `json:"ussd"`
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

	host, exist := os.LookupEnv("HOST")

	if !exist {
		logger.Log(errors.New("HOST not set in .env"))
		log.Fatal("HOST not set in .env")
	}

	bankJson, err := ioutil.ReadFile("./banks.json")

	if err != nil {
		logger.Log(err)
	}

	horusDbUser, exist := os.LookupEnv("HORUS_DB_USER")

	if !exist {
		log.Fatal("HORUS_DB_USER not set in .env")
	}

	horusDbPass, exist := os.LookupEnv("HORUS_DB_PASS")

	if !exist {
		log.Fatal("HORUS_DB_PASS not set in .env")
	}

	horusDbHost, exist := os.LookupEnv("HORUS_DB_HOST")

	if !exist {
		log.Fatal("HORUS_DB_HOST not set in .env")
	}

	horusDbName, exist := os.LookupEnv("HORUS_DB_NAME")

	if !exist {
		log.Fatal("HORUS_DB_NAME not set in .env")
	}

	horusDbPort, exist := os.LookupEnv("HORUS_DB_PORT")

	if !exist {
		log.Fatal("HORUS_DB_PORT not set in .env")
	}

	listener, err := horus.Init("mysql", horus.Config{
		DbName:    horusDbName,
		DbHost:    horusDbHost,
		DbPssword: horusDbPass,
		DbPort:    horusDbPort,
		DbUser:    horusDbUser,
	})

	if err != nil {
		logger.Log(err)
	}

	var banks []BankJSON

	if err := json.Unmarshal(bankJson, &banks); err != nil {
		logger.Log(err)
	}

	route := mux.NewRouter()

	route.PathPrefix("/logo/").Handler(http.StripPrefix("/logo/", http.FileServer(http.Dir("./logos"))))

	route.NotFoundHandler = http.HandlerFunc(listener.Watch(notFound))

	route.HandleFunc("/", listener.Watch(func(writer http.ResponseWriter, request *http.Request) {
		writer.Header().Set("Content-Type", "application/json")

		var newBanks []Bank

		for _, bank := range banks {
			newBanks = append(newBanks, Bank{
				Name: bank.Name,
				Slug: bank.Slug,
				Code: bank.Code,
				USSD: bank.USSD,
				Logo: host + "/logo/" + getUrl(bank.Slug) + ".png",
			})
		}

		_ = json.NewEncoder(writer).Encode(newBanks)
	}))

	handler := cors.AllowAll().Handler(route)

	if err := http.ListenAndServe(":"+port, handler); err != nil {
		logger.Log(err)
	}
}

func notFound(writer http.ResponseWriter, request *http.Request) {
	writer.WriteHeader(http.StatusNotFound)

	writer.Header().Set("Content-Type", "application/json")

	_ = json.NewEncoder(writer).Encode(Error{
		Message: "endpoint not found",
	})
}

func getUrl(slug string) string {
	var files []string

	f, err := os.Open("./logos")

	if err != nil {
		logger.Log(err)
	}

	fileInfo, err := f.Readdir(0)

	_ = f.Close()

	if err != nil {
		logger.Log(err)
	}

	for _, file := range fileInfo {
		if file.Name() == ".DS_Store" {
			continue
		}

		files = append(files, file.Name())
	}

	_, found := find(files, slug+".png")

	if found {
		return slug
	}

	return "default-image"
}

func find(slice []string, val string) (int, bool) {
	for i, item := range slice {
		if item == val {
			return i, true
		}
	}
	return -1, false
}
