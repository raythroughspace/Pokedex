package main
import (
	"fmt"
	"errors"
	"encoding/json"
	"net/http"
)

type parsedData struct {
	Count    int    `json:"count"`
	Next     *string `json:"next"`
	Previous *string    `json:"previous"`
	Results  []struct {
		Name string `json:"name"`
		URL  string `json:"url"`
	} `json:"results"`
}

type locationAreaPage struct {
	current string
	data parsedData
}	

func getParsedDataMap(url string, locCache *Cache) (parsedData, error) {
	elem, ok := locCache.Get(url)
	if ok {
		return elem, nil
	}
	res, err := http.Get(url)
	if err != nil {
		return parsedData{}, err
	}
	defer res.Body.Close()

	var data parsedData
	decoder := json.NewDecoder(res.Body)
	if err = decoder.Decode(&data); err != nil {
		return parsedData{}, err
	}
	return data, nil
}

func commandMap(pg *locationAreaPage, locCache *Cache, location string,
	pokename string, pokedex map[string]pokemon) error {
	var err error
	if pg.current == "" {
		// first page
		pg.current = "https://pokeapi.co/api/v2/location-area"
		data, err := getParsedDataMap(pg.current, locCache)
		if err != nil {
			return err
		}
		pg.data = data
	} else if pg.data.Next != nil {
		pg.current = *pg.data.Next
		fmt.Println(pg.current)
		pg.data, err = getParsedDataMap(pg.current, locCache)
		
		if err != nil {
			return err
		}
	} 
	for _, res := range pg.data.Results {
		fmt.Println(res.Name)
	}
	return nil
}

func commandMapb(pg *locationAreaPage, locCache *Cache, location string,
	pokename string, pokedex map[string]pokemon) error {
	var err error
	if pg.data.Previous == nil {
		return errors.New("you're on the first page\n")
	}
	pg.current = *pg.data.Previous
	fmt.Println(pg.current)
	pg.data, err = getParsedDataMap(pg.current, locCache)

	if err != nil {
		return err
	}
	for _, res := range pg.data.Results {
		fmt.Println(res.Name)
	}
	return nil
}

