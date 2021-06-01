package controllers

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
)

func CarsCreate(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "application/json;charset=UTF-8")
	w.WriteHeader(http.StatusOK)

	body, err := ioutil.ReadAll(r.Body)

	if err != nil {
		log.Fatal(err)
	}

	var car models.Car

	err = json.Unmarshal(body, &car)

	if err != nil {
		log.Fatal(err)
	}

	models.NewCar(&car)

	json.NewEncoder(w).Encode(car)
}
