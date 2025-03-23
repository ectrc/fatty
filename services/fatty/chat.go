package fatty

import (
	"fatty/helpers"
	"fmt"

	"github.com/brianvoe/gofakeit/v7"
)

type ChatSession struct {
	ID string
	User *FattyUser
}

func NewChatSession(client *helpers.ProxiedClient, user *FattyUser) (*ChatSession, error) {
	if client == nil || user == nil || user.AccessToken == nil || user.Location == nil {
		return nil, fmt.Errorf("invalid arguments")
	}

	session := &ChatSession{
		ID: gofakeit.UUID(),
		User: user,
	}

	response := client.Post("https://chatcontroller.foodtools.io/v1/chat/query?stream=false", helpers.JSON{
		"chatId": session.ID,
		"latitude": user.Location.Lat,
		"longitude": user.Location.Lon,
		"timeZoneId": user.Location.Timezone,
		"tenant": "uk",
		"postcode": fmt.Sprintf("%s 4FQ", user.Location.Zip),
		"data": helpers.JSON{
			"type": "None",
		},
		"prompt": "Hello",
	}, helpers.JSON{
		"application-version": user.Version,
		"authorization": fmt.Sprintf("Bearer %s", *user.AccessToken),
		"authorization2": "teamx-poc UngyHrXZGE3ociUOjJ1AKrqOy002Nv",
		"accept-tenant": "UK",
		"x-jet-functional-session-id": session.ID,
		"x-jet-personalised-session-id": session.ID,
		"x-jet-essential-session-id": session.ID,
		"x-je-auser": gofakeit.UUID(),
		"x-je-user-consent": "true",
		"x-jet-application": "OneAppIOS",
		"Cookies": "__cf_bm=RHCppJ66on29ub2xIhoCsUfIwjlxRHnqsWqxSlIC2ZE-1742756595-1.0.1.1-vbQ6EqZmwnEQ0N4nENaKyM50wu9u.GVtKBwkjKCvgOZCUg83EuppAOFwTl6Rkk07tVyQP2b8LG7j9294ZZB0R.Km961elhOeu96kwWLAJ.pAhSaJa5mJE9JTJf3uszff",
	})
	if response.Err != nil {
		return nil, response.Err
	}

	type ChatResponse struct {
		QuestionID string `json:"questionId"`
		Actions []struct{} `json:"actions"`
		Message string `json:"message"`
	}
	chatResponse := helpers.ToStruct[ChatResponse](response.Body)
	if chatResponse == nil {
		return nil, fmt.Errorf("failed to parse chat response: %s", response.Body)
	}
	if chatResponse.QuestionID == "" || len(chatResponse.Actions) == 0 {
		return nil, fmt.Errorf("invalid chat response: %s", response.Body)
	}

	fmt.Printf("%s\n", chatResponse.Message)
	
	return session, nil
}

func (c *ChatSession) HelpMeBail(client *helpers.ProxiedClient) error {
	if client == nil {
		return fmt.Errorf("invalid arguments")
	}

	response := client.Post("https://chatcontroller.foodtools.io/v1/chat/query?stream=false", helpers.JSON{
		"chatId": c.ID,
		"latitude": c.User.Location.Lat,
		"longitude": c.User.Location.Lon,
		"timeZoneId": c.User.Location.Timezone,
		"tenant": "uk",
		"postcode": fmt.Sprintf("%s 4FQ", c.User.Location.Zip),
		"data": helpers.JSON{
			"type": "None",
		},
		"prompt": "help me bail",
	}, helpers.JSON{
		"application-version": c.User.Version,
		"authorization": fmt.Sprintf("Bearer %s", *c.User.AccessToken),
		"authorization2": "teamx-poc UngyHrXZGE3ociUOjJ1AKrqOy002Nv",
		"accept-tenant": "UK",
		"x-jet-functional-session-id": c.ID,
		"x-jet-personalised-session-id": c.ID,
		"x-jet-essential-session-id": c.ID,
		"x-je-auser": gofakeit.UUID(),
		"x-je-user-consent": "true",
		"priority": "u=3, i",
		"application-id": "3",
		"User-Agent": fmt.Sprintf("Just Eat/%s (2548) iOS 18.0.1/iPhone", c.User.Version),
		"accept-charset": "utf-8",
		"x-je-applicationvariant": "live",
		"content-type": "application/json;v=2",
		"x-jet-application": "OneWeb",
	})
	if response.Err != nil {
		return response.Err
	}

	type ChatResponse struct {
		QuestionID string `json:"questionId"`
		Actions []struct{} `json:"actions"`
		Message string `json:"message"`
	}
	chatResponse := helpers.ToStruct[ChatResponse](response.Body)
	if chatResponse == nil {
		return fmt.Errorf("failed to parse chat response: %s", response.Body)
	}

	fmt.Printf("%s\n", chatResponse.Message)

	return nil
}