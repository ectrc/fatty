package codes

import (
	"fatty/helpers"
	"fatty/services/config"
	"fatty/services/fatty"
	"fmt"
	"strings"
	"sync"
)

type CodeGeneratorCommand struct{}

func (g CodeGeneratorCommand) Execute() error {
	config := config.Config()

	input, err := helpers.File(config.CODE_GEN_INPUT_FILE_LOCATION)
	if err != nil {
		return fmt.Errorf("failed to open accounts file: %s", err)
	}
	defer input.Close()

	output, err := helpers.File(config.CODE_GEN_OUTPUT_FILE_LOCATION)
	if err != nil {
		return fmt.Errorf("failed to open accounts file: %s", err)
	}

	accounts, err := input.ReadAllLines()
	if err != nil {
		return fmt.Errorf("failed to read accounts file: %s", err)
	}

	waiter := sync.WaitGroup{}
	length := len(accounts) / config.CODE_GEN_THREAD_COUNT

	for i := 0; i < config.CODE_GEN_THREAD_COUNT; i++ {
		waiter.Add(1)
		subset := make([]string, length)
		copy(subset, accounts[i*length:(i+1)*length])

		go func(subset []string, watier *sync.WaitGroup) {
			for _, account := range subset {
				split := strings.Split(account, ":")
				if len(split) != 2 {
					fmt.Printf("invalid account: %s\n", account)
					continue
				}

				err := generateNewCode(output, split[0], split[1])
				if err != nil {
					fmt.Printf("failed to generate new code: %s\n", err)
				}
			}

			waiter.Done()
		}(subset, &waiter)
	}

	waiter.Wait()

	return nil
}

func generateNewCode(output *helpers.HelperFile, username, password string) error {
	config := config.Config()

	client := helpers.NewProxiedClient()
	if config.PROXY_ENABLED {
		client.SetProxy(config.PROXY_URL)
	}

	user, err := fatty.NewFattyUserFromUsernamePassword(client, username, password)
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

	chat, err := fatty.NewChatSession(client, user)
	if err != nil {
		return retryCreateChat(output, client, user, 0, 3)
	}

	if config.EXTRA_LOGGING {
		fmt.Printf("%s\n", helpers.StructToJSON(chat).Bytes())
	}

	message, err := chat.HelpMeBail(client)
	if err != nil {
		return fmt.Errorf("failed to help me bail: %s", err)
	}

	if strings.Contains(message, "Please try again") {
		return retryBail(output, client, user, chat, 0, 3)
	}

	if config.EXTRA_LOGGING {
		fmt.Printf("%s\n", helpers.StructToJSON(chat).Bytes())
	}

	output.Write([]byte(fmt.Sprintf("\n%s - %s:%s\n", message, user.Email, user.Password)))
	fmt.Printf("completed %s:%s\n", user.Email, user.Password)

	return nil
}

func retryCreateChat(output *helpers.HelperFile, client *helpers.ProxiedClient, user *fatty.FattyUser, current, max int) error {
	if current >= max {
		return fmt.Errorf("max retries done for email: %s", user.Email)
	}

	config := config.Config()

	chat, err := fatty.NewChatSession(client, user)
	if err != nil {
		return fmt.Errorf("failed to create chat session: %s", err)
	}

	if config.EXTRA_LOGGING {
		fmt.Printf("%s\n", helpers.StructToJSON(chat).Bytes())
	}

	message, err := chat.HelpMeBail(client)
	if err != nil {
		return fmt.Errorf("failed to help me bail: %s", err)
	}

	if strings.Contains(message, "Please try again") {
		return retryBail(output, client, user, chat, current+1, max)
	}

	if config.EXTRA_LOGGING {
		fmt.Printf("%s\n", helpers.StructToJSON(chat).Bytes())
	}

	output.Write([]byte(fmt.Sprintf("\n%s - %s:%s\n", message, user.Email, user.Password)))
	fmt.Printf("completed %s:%s\n", user.Email, user.Password)

	return nil
}

func retryBail(output *helpers.HelperFile, client *helpers.ProxiedClient, user *fatty.FattyUser, chat *fatty.ChatSession, current, max int) error {
	if current >= max {
		return fmt.Errorf("max retries done for email: %s", user.Email)
	}

	config := config.Config()

	message, err := chat.HelpMeBail(client)
	if err != nil {
		return fmt.Errorf("failed to help me bail: %s", err)
	}

	if strings.Contains(message, "Please try again") {
		return retryBail(output, client, user, chat, current+1, max)
	}

	if config.EXTRA_LOGGING {
		fmt.Printf("%s\n", helpers.StructToJSON(chat).Bytes())
	}

	output.Write([]byte(fmt.Sprintf("\n%s - %s:%s\n", message, user.Email, user.Password)))
	fmt.Printf("completed %s:%s\n", user.Email, user.Password)

	return nil
}