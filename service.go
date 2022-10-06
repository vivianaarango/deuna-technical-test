package main

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"
)

// Service struct with the arguments necessary to the api.
type Service struct {
	Router *mux.Router
}

// Initialize initializes the api with the routes.
func (s *Service) Initialize() {
	// create a new router.
	s.Router = mux.NewRouter()

	// set the defined routes.
	s.setRoutes()
}

// setRoutes init each route with determinate conf.
func (s *Service) setRoutes() {
	// create route service for get starships available according the passenger number.
	s.Router.HandleFunc("/api/v1/starships/available", getStarShipAvailabilityService).
		Queries("passengers", "{passengers:[0-9]+}").
		Methods(http.MethodGet)
}

// Run the app on it's router
func (s *Service) Run(port string) {
	log.Fatal(http.ListenAndServe(port, s.Router))
}

// getStarShipAvailabilityService service for get starships available according the passenger number.
func getStarShipAvailabilityService(w http.ResponseWriter, r *http.Request) {
	// obtain the passengers sent.
	passengers := mux.Vars(r)["passengers"]

	// obtain the starship available.
	starshipsNumber, err := getStarshipsAvailable()
	if err != nil {
		fmt.Fprintf(w, err.Error())
		return
	}

	// get the passenger number from query parameter.
	passengersNumber, err := strconv.Atoi(passengers)
	if err != nil {
		fmt.Fprintf(w, err.Error())
		return
	}

	// create struct with the response.
	starshipAvailability := &StarshipAvailabilityServiceResponse{
		PassengerNumber:     passengersNumber,
		ShipAvailableNumber: starshipsNumber,
	}

	// convert response into json.
	response, err := json.Marshal(starshipAvailability)
	if err != nil {
		fmt.Fprintf(w, err.Error())
		return
	}

	// set the default headers.
	w.Header().Set("Content-Type", "application/json")

	// return response for the user.
	_, err = w.Write(response)
	if err != nil {
		fmt.Fprintf(w, err.Error())
		return
	}
}

// getStarshipsAvailable obtain the number of starships available from star wars api.
func getStarshipsAvailable() (total int, err error) {
	// create client to consume star wars api for obtain the starships available
	client := &http.Client{}
	request, err := http.NewRequest(http.MethodGet, "https://swapi.dev/api/starships", nil)
	if err != nil {
		return total, err
	}

	// consume the specific service.
	response, err := client.Do(request)
	if err != nil {
		return total, err
	}

	// convert response in []byte.
	body, err := ioutil.ReadAll(response.Body)

	// parse the body into struct with the data necessary.
	var result StarshipsServiceResponse
	err = json.Unmarshal(body, &result)
	if err != nil {
		return total, err
	}

	// return the available starships.
	return result.Count, err
}
