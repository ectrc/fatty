package fatty

import (
	"fatty/helpers"
	"fmt"
)

func GetVersion(client *helpers.ProxiedClient) (string, error) {
	version := client.Get("https://serviceinformationassets-production.je-apis.com/ios/service-information-en_GB.json")
	if version.Err != nil {
		return "", fmt.Errorf("failed to get version: %s", version.Err)
	}

	type VersionResponse struct {
		Version string `json:"minimumRecommendedVersion"`
	}

	versionResponse := helpers.ToStruct[VersionResponse](version.Body)
	if versionResponse == nil {
		return "", fmt.Errorf("failed to parse version response")
	}

	if versionResponse.Version == "" {
		return "", fmt.Errorf("failed to get version, body: %s", version.Body)
	}

	return versionResponse.Version, nil
}