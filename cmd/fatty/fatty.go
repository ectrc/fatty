package fatty

import (
	"fatty/helpers"
	"fatty/services/config"
	"fatty/services/fatty"
	"fmt"
)

type FattyCommand struct{}

func (f FattyCommand) Execute() error {
	return ProcessFatty()
}

func ProcessFatty() error {
	config := config.Config()

	codeOutput, err := helpers.File(config.CODE_GEN_OUTPUT_FILE_LOCATION)
	if err != nil {
		return fmt.Errorf("failed to open accounts file: %s", err)
	}

	accountOutput, err := helpers.File(config.ACC_GEN_FILE_LOCATION)
	if err != nil {
		return fmt.Errorf("failed to open accounts file: %s", err)
	}

	client := helpers.NewProxiedClient()
	if config.PROXY_ENABLED {
		client.SetProxy(config.PROXY_URL)
	}

	user, err := fatty.NewFattyUser(client)
	if err != nil {
		return fmt.Errorf("failed to create user: %s", err)
	}

	if config.EXTRA_LOGGING {
		fmt.Printf("%s\n", helpers.StructToJSON(user).Bytes())
	}

	err = user.Login(client)
	if err != nil {
		return fmt.Errorf("failed to login: %s", err)
	}

	if config.EXTRA_LOGGING {
		fmt.Printf("%s\n", *user.AccessToken)
	}

	err = user.Profile(client)
	if err != nil {
		return fmt.Errorf("failed to get consumer: %s", err)
	}

	err = user.EnableNewsletter(client)
	if err != nil {
		return fmt.Errorf("failed to set newsletter: %s", err)
	}

	chat, err := fatty.NewChatSession(client, user)
	if err != nil {
		return fmt.Errorf("failed to create chat session: %s", err)
	}

	chatMessage, err := chat.HelpMeBail(client)
	if err != nil {
		return fmt.Errorf("failed to help me bail: %s", err)
	}

	codeOutput.Write([]byte(fmt.Sprintf("\n%s - %s:%s\n", chatMessage, user.Email, user.Password)))
	fmt.Printf("%s:%s\n", user.Email, user.Password)
	accountOutput.Write([]byte(fmt.Sprintf("%s:%s\n", user.Email, user.Password)))

	return nil
}