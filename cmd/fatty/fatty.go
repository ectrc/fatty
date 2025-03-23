package fatty

import (
	"fatty/helpers"
	"fatty/services/fatty"
	"fmt"
)

type FattyCommand struct{}

func (f FattyCommand) Execute() error {
	return ProcessFatty()
}

func ProcessFatty() error {
	client := helpers.NewProxiedClient()
	client.SetProxy("http://ufb2602a657b905d1-zone-custom-region-gb:ufb2602a657b905d1@118.193.58.115:2334")

	user, err := fatty.NewFattyUser(client)
	if err != nil {
		return fmt.Errorf("failed to create user: %s", err)
	}

	err = user.Login(client)
	if err != nil {
		return fmt.Errorf("failed to login: %s", err)
	}

	err = user.Profile(client)
	if err != nil {
		return fmt.Errorf("failed to get consumer: %s", err)
	}

	err = user.EnableNewsletter(client)
	if err != nil {
		return fmt.Errorf("failed to set newsletter: %s", err)
	}

	fmt.Printf("%s:%s\n", user.Email, user.Password)

	chat, err := fatty.NewChatSession(client, user)
	if err != nil {
		return fmt.Errorf("failed to create chat session: %s", err)
	}

	err = chat.HelpMeBail(client)
	if err != nil {
		return fmt.Errorf("failed to help me bail: %s", err)
	}

	return nil
}