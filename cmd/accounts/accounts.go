package accounts

import (
	"fatty/helpers"
	"fatty/services/config"
	"fatty/services/fatty"
	"fmt"
)

type AccountsGeneratorCommand struct{}

func (g AccountsGeneratorCommand) Execute() error {
	config := config.Config()

	file, err := helpers.File(config.ACC_GEN_FILE_LOCATION)
	if err != nil {
		return fmt.Errorf("failed to open accounts file: %s", err)
	}
	defer file.Close()

	for i := 0; i < config.ACC_GEN_THREAD_COUNT; i++ {
		go func(file *helpers.HelperFile) {
			for {
				err := generateNewAccount(file)
				if err != nil {
					fmt.Printf("failed to generate new account: %s\n", err)
				}
			}
		}(file)
	}

	<-make(chan bool)
	return nil
}

func generateNewAccount(file *helpers.HelperFile) error {
	config := config.Config()

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

	if config.ACC_GEN_ENABLE_NEWSLETTER {
		err = user.EnableNewsletter(client)
		if err != nil {
			return fmt.Errorf("failed to set newsletter: %s", err)
		}
	}

	file.Write([]byte(fmt.Sprintf("%s:%s\n",
		user.Email,
		user.Password,
	)))

	fmt.Printf("%s:%s\n", user.Email, user.Password)

	return nil
}
