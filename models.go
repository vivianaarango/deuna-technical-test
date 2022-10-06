package main

// StarshipsServiceResponse response for https://swapi.dev/api/starships service.
type StarshipsServiceResponse struct {
	Count int `json:"count"`
}

// StarshipAvailabilityServiceResponse response for /api/v1/starships/available service.
type StarshipAvailabilityServiceResponse struct {
	PassengerNumber     int `json:"number_of_passengers"`
	ShipAvailableNumber int `json:"number_of_ships_available"`
}
