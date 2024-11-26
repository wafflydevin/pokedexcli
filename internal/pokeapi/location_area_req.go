package pokeapi

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

func (c *Client) ListLocationAreas(pageURL *string) (LocationAreasResp, error) {
	fullURL := baseURL + "/location-area/"
	if pageURL != nil {
		fullURL = *pageURL
	}

	data, cached := c.cache.Get(fullURL)
	if cached {
		fmt.Println("cache hit")
		locationAreasResp := LocationAreasResp{}
		err := json.Unmarshal(data, &locationAreasResp)
		if err != nil {
			return LocationAreasResp{}, err
		}

		return locationAreasResp, nil
	}
	fmt.Println("cache miss")

	req, err := http.NewRequest("GET", fullURL, nil)
	if err != nil {
		return LocationAreasResp{}, err
	}

	res, err := c.httpClient.Do(req)
	if err != nil {
		return LocationAreasResp{}, nil
	}
	defer res.Body.Close()

	if res.StatusCode > 399 {
		return LocationAreasResp{}, fmt.Errorf("bad status code: %v", res.StatusCode)
	}

	data, err = io.ReadAll(res.Body)
	if err != nil {
		return LocationAreasResp{}, err
	}

	locationAreasResp := LocationAreasResp{}
	err = json.Unmarshal(data, &locationAreasResp)
	if err != nil {
		return LocationAreasResp{}, err
	}

	c.cache.Add(fullURL, data)

	return locationAreasResp, nil
}

func (c *Client) GetLocationAreas(locationAreaName *string) (LocationArea, error) {
	fullURL := baseURL + "/location-area/" + *locationAreaName

	data, cached := c.cache.Get(fullURL)
	if cached {
		fmt.Println("cache hit")
		locationAreas := LocationArea{}
		err := json.Unmarshal(data, &locationAreas)
		if err != nil {
			return LocationArea{}, err
		}

		return locationAreas, nil
	}
	fmt.Println("cache miss")

	req, err := http.NewRequest("GET", fullURL, nil)
	if err != nil {
		return LocationArea{}, err
	}

	res, err := c.httpClient.Do(req)
	if err != nil {
		return LocationArea{}, nil
	}
	defer res.Body.Close()

	if res.StatusCode > 399 {
		return LocationArea{}, fmt.Errorf("bad status code: %v", res.StatusCode)
	}

	data, err = io.ReadAll(res.Body)
	if err != nil {
		return LocationArea{}, err
	}

	locationAreas := LocationArea{}
	err = json.Unmarshal(data, &locationAreas)
	if err != nil {
		return LocationArea{}, err
	}

	c.cache.Add(fullURL, data)

	return locationAreas, nil
}
