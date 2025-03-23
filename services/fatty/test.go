package fatty

import (
	"fatty/helpers"

	"fmt"
)

type LocationResponse struct {
	Status string `json:"status"`
	Country string `json:"country"`
	CountryCode string `json:"countryCode"`
	Region string `json:"region"`
	RegionName string `json:"regionName"`
	City string `json:"city"`
	Zip string `json:"zip"`
	Lat float64 `json:"lat"`
	Lon float64 `json:"lon"`
	Timezone string `json:"timezone"`
	Isp string `json:"isp"`
	Org string `json:"org"`
	As string `json:"as"`
	Query string `json:"query"`
}

func Location(client *helpers.ProxiedClient) (*LocationResponse, error) {
	result := client.Get("http://ip-api.com/json")
	if result.Err != nil {
		return nil, fmt.Errorf("result error: %s", result.Err)
	}

	locationResponse := helpers.ToStruct[LocationResponse](result.Body)
	if locationResponse == nil {
		return nil, fmt.Errorf("failed to parse location response")
	}

	if locationResponse.Status != "success" {
		return nil, fmt.Errorf("failed to get location, body: %s", result.Body)
	}

	fmt.Printf("%s %s %s\n", locationResponse.Country, locationResponse.As, locationResponse.Query)

	return locationResponse, nil
}