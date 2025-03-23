package fatty

import (
	"encoding/base64"
	"fatty/helpers"
	"fatty/services/config"
	"fmt"
	"net/url"
	"regexp"

	"github.com/brianvoe/gofakeit/v7"
)

type FattyUser struct {
	Person *gofakeit.PersonInfo
	Location *LocationResponse

	Email string
	Password string
	Device string
	Version string

	OtacToken *string
	AccessToken *string
}

func NewFattyUser(client *helpers.ProxiedClient) (*FattyUser, error) {
	config := config.Config()

	user := &FattyUser{
		Person: gofakeit.Person(),
		Device: gofakeit.UUID(),
	}
	user.Email = fmt.Sprintf("%s%s%d@%s", user.Person.FirstName, user.Person.LastName, gofakeit.Int8(), config.EMAIL_DOMAIN)
	user.Password = gofakeit.Password(true, true, true, true, false, 12)

	version, err := GetVersion(client)
	if err != nil {
		return nil, fmt.Errorf("failed to get version: %s", err)
	}
	user.Version = version

	location, err := Location(client)
	if err != nil {
		return nil, fmt.Errorf("failed to get location: %s", err)
	}
	user.Location = location

	response := client.Post("https://uk.api.just-eat.io/consumers/uk", helpers.JSON{
		"emailAddress": user.Email,
		"fullName": fmt.Sprintf("%s %s", user.Person.FirstName, user.Person.LastName),
		"password": user.Password,
		"registrationSource": "native",
	}, helpers.JSON{
		"content-type": "application/json;v=2",
		"x-jet-application": "OneWeb",
	})
	if response.Err != nil {
		return nil, response.Err
	}

	type OtacResponse struct {
		Type string `json:"type"`
		Token string `json:"token"`
	}
	otacResponse := helpers.ToStruct[OtacResponse](response.Body)
	if otacResponse == nil {
		return nil, fmt.Errorf("failed to parse response")
	}
	if otacResponse.Type == "" || otacResponse.Token == "" {
		return nil, fmt.Errorf("failed to get token, body: %s", response.Body)
	}

	user.OtacToken = &otacResponse.Token
	return user, nil
}

func NewFattyUserFromUsernamePassword(client *helpers.ProxiedClient, username, password string) (*FattyUser, error) {
	user := &FattyUser{
		Email: username,
		Password: password,
		Device: gofakeit.UUID(),
	}

	version, err := GetVersion(client)
	if err != nil {
		return nil, fmt.Errorf("failed to get version: %s", err)
	}
	user.Version = version

	location, err := Location(client)
	if err != nil {
		return nil, fmt.Errorf("failed to get location: %s", err)
	}
	user.Location = location

	return user, nil
}

func (u FattyUser) DeviceData() string {
	data := base64.StdEncoding.EncodeToString(helpers.JSON{
		"DeviceName": "iPhone",
		"DeviceId": u.Device,
		"DeviceType": "iPhone11,8",
	}.Bytes())

	result := regexp.MustCompile(`\+`).ReplaceAllString(data, "-")
	result = regexp.MustCompile(`\/`).ReplaceAllString(result, "_")
	result = regexp.MustCompile(`=`).ReplaceAllString(result, "")

	return result
} 

func (u *FattyUser) Login(client *helpers.ProxiedClient) error {
	response := client.PostForm("https://auth.just-eat.co.uk/connect/token", url.Values{
		"grant_type": []string{"password"},
		"client_id": []string{"consumer_ios_je"},
		"scope": []string{"openid mobile_scope offline_access"},
		"tenant": []string{"uk"},
		"acr_values": []string{fmt.Sprintf("language:en-GB tenant:UK device:%s deviceId:rvnios-%s", u.DeviceData(), u.Device)},
		"username": []string{u.Email},
		"password": []string{u.Password},
	}, helpers.JSON{
		"x-jet-application": "OneWeb",
	})
	if response.Err != nil {
		return response.Err
	}

	type LoginResponse struct {
		AccessToken string `json:"access_token"`
	}
	loginResponse := helpers.ToStruct[LoginResponse](response.Body)
	if loginResponse == nil {
		return fmt.Errorf("failed to parse login response")
	}
	if loginResponse.AccessToken == "" {
		return fmt.Errorf("failed to get access token, body: %s", response.Body)
	}

	u.AccessToken = &loginResponse.AccessToken
	return nil
}

func (u *FattyUser) Profile(client *helpers.ProxiedClient) error {
	response := client.Get("https://uk.api.just-eat.io/consumer", helpers.JSON{
		"authorization": fmt.Sprintf("Bearer %s", *u.AccessToken),
		"x-jet-application": "OneWeb",
	})
	if response.Err != nil {
		return response.Err
	}

	data := helpers.ToStruct[helpers.JSON](response.Body)
	if data == nil {
		return fmt.Errorf("failed to parse response")
	}

	if config.Config().EXTRA_LOGGING {
		fmt.Printf("%s\n", data.Bytes())
	}
	
	return nil
}

func (u *FattyUser) EnableNewsletter(client *helpers.ProxiedClient) error {
	response := client.Put("https://uk.api.just-eat.io/consumers/uk/me/communication-preferences/marketing", helpers.JSON{
		"subscribedChannels": []string{"email"},
	}, helpers.JSON{
		"authorization": fmt.Sprintf("Bearer %s", *u.AccessToken),
		"x-jet-application": "OneWeb",
	})
	if response.Err != nil {
		return response.Err
	}
		
	return nil
}