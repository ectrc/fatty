package config

import (
	"fatty/helpers"

	"github.com/joho/godotenv"
)

type FattyConfig struct {
	EXTRA_LOGGING bool

	ACC_GEN_FILE_LOCATION string
	ACC_GEN_THREAD_COUNT int
	ACC_GEN_ENABLE_NEWSLETTER bool

	CODE_GEN_INPUT_FILE_LOCATION string
	CODE_GEN_OUTPUT_FILE_LOCATION string
	CODE_GEN_THREAD_COUNT int

	PROXY_ENABLED bool
	PROXY_URL string
}

func NewFattyConfig() *FattyConfig {
	godotenv.Load(".env")

	return &FattyConfig{
		EXTRA_LOGGING: helpers.Enviroment("EXTRA_LOGGING", "false") == "true",

		ACC_GEN_FILE_LOCATION: helpers.Enviroment("ACC_GEN_FILE_LOCATION"),
		ACC_GEN_THREAD_COUNT: helpers.Parse[int](helpers.Enviroment("ACC_GEN_THREAD_COUNT", "0")),
		ACC_GEN_ENABLE_NEWSLETTER: helpers.Enviroment("ACC_GEN_ENABLE_NEWSLETTER", "false") == "true",

		CODE_GEN_INPUT_FILE_LOCATION: helpers.Enviroment("CODE_GEN_INPUT_FILE_LOCATION"),
		CODE_GEN_OUTPUT_FILE_LOCATION: helpers.Enviroment("CODE_GEN_OUTPUT_FILE_LOCATION"),
		CODE_GEN_THREAD_COUNT: helpers.Parse[int](helpers.Enviroment("CODE_GEN_THREAD_COUNT", "0")),

		PROXY_ENABLED: helpers.Enviroment("PROXY_ENABLED", "false") == "true",
		PROXY_URL: helpers.Enviroment("PROXY_URL"),
	}
}

var (
	config = NewFattyConfig()
)

func Config() *FattyConfig {
	return config
}