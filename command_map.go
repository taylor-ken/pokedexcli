package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
)

type Config struct {
	Next     string `json:"next"`
	Previous string `json:"previous"`
}

type results struct {
	Name string `json:"name"`
	URL  string `json:"url"`
}

type response struct {
	Count    int       `json:"count"`
	Next     string    `json:"next"`
	Previous string    `json:"previous"`
	Results  []results `json:"results"`
}

func commandMap(cfg *Config) error {

	var res *http.Response
	var err error
	if len(cfg.Next) > 0 {
		res, err = http.Get(cfg.Next)
	} else {
		res, err = http.Get("https://pokeapi.co/api/v2/location-area/")
	}
	if err != nil {
		log.Fatal(err)
	}
	body, err := io.ReadAll(res.Body)
	res.Body.Close()
	if res.StatusCode > 299 {
		log.Fatalf("Response failed with status code: %d and\nbody: %s\n", res.StatusCode, body)
	}
	if err != nil {
		log.Fatal(err)
	}

	apiResponse := response{}
	err = json.Unmarshal(body, &apiResponse)
	if err != nil {
		fmt.Println(err)
		return err
	}
	for _, results := range apiResponse.Results {
		fmt.Printf("%s\n", results.Name)
	}
	cfg.Next = apiResponse.Next
	cfg.Previous = apiResponse.Previous

	return nil
}

func commandMapb(cfg *Config) error {

	var res *http.Response
	var err error
	if len(cfg.Previous) > 0 {
		res, err = http.Get(cfg.Previous)
	} else {
		fmt.Printf("you're on the first page\n")
		return nil
	}
	if err != nil {
		log.Fatal(err)
	}
	body, err := io.ReadAll(res.Body)
	res.Body.Close()
	if res.StatusCode > 299 {
		log.Fatalf("Response failed with status code: %d and\nbody: %s\n", res.StatusCode, body)
	}
	if err != nil {
		log.Fatal(err)
	}

	apiResponse := response{}
	err = json.Unmarshal(body, &apiResponse)
	if err != nil {
		fmt.Println(err)
		return err
	}
	for _, results := range apiResponse.Results {
		fmt.Printf("%s\n", results.Name)
	}
	cfg.Next = apiResponse.Next
	cfg.Previous = apiResponse.Previous

	return nil
}
