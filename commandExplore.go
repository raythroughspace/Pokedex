package main
import (
	"fmt"
	"encoding/json"
	"net/http"
)

type locationAreaPageDetailed struct {
	EncounterMethodRates []struct {
		EncounterMethod struct {
			Name string `json:"name"`
			URL  string `json:"url"`
		} `json:"encounter_method"`
		VersionDetails []struct {
			Rate    int `json:"rate"`
			Version struct {
				Name string `json:"name"`
				URL  string `json:"url"`
			} `json:"version"`
		} `json:"version_details"`
	} `json:"encounter_method_rates"`
	GameIndex int `json:"game_index"`
	ID        int `json:"id"`
	Location  struct {
		Name string `json:"name"`
		URL  string `json:"url"`
	} `json:"location"`
	Name  string `json:"name"`
	Names []struct {
		Language struct {
			Name string `json:"name"`
			URL  string `json:"url"`
		} `json:"language"`
		Name string `json:"name"`
	} `json:"names"`
	PokemonEncounters []struct {
		Pokemon struct {
			Name string `json:"name"`
			URL  string `json:"url"`
		} `json:"pokemon"`
		VersionDetails []struct {
			EncounterDetails []struct {
				Chance          int   `json:"chance"`
				ConditionValues []any `json:"condition_values"`
				MaxLevel        int   `json:"max_level"`
				Method          struct {
					Name string `json:"name"`
					URL  string `json:"url"`
				} `json:"method"`
				MinLevel int `json:"min_level"`
			} `json:"encounter_details"`
			MaxChance int `json:"max_chance"`
			Version   struct {
				Name string `json:"name"`
				URL  string `json:"url"`
			} `json:"version"`
		} `json:"version_details"`
	} `json:"pokemon_encounters"`
}

func commandExplore(pg *locationAreaPage, locCache *Cache, location string,
	pokename string, pokedex map[string]pokemon) error {
	explorePg, err := getParsedDataExplore(fmt.Sprintf("https://pokeapi.co/api/v2/location-area/%v", location))
	if err != nil {
		return err
	}
	fmt.Println("Exploring ...", explorePg.Location.Name)
	fmt.Println("Found Pokemon:")
	for _, poke := range explorePg.PokemonEncounters {
		fmt.Println("- ", poke.Pokemon.Name)
	}
	return nil
}

func getParsedDataExplore(url string) (locationAreaPageDetailed, error) {
	res, err := http.Get(url)
	if err != nil {
		return locationAreaPageDetailed{}, err
	}
	defer res.Body.Close()

	var data locationAreaPageDetailed
	decoder := json.NewDecoder(res.Body)
	if err = decoder.Decode(&data); err != nil {
		return locationAreaPageDetailed{}, err
	}
	return data, nil
}