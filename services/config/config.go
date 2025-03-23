package config

import (
	"fatty/helpers"
	"time"

	"github.com/joho/godotenv"
)

type FattyConfig struct {
	ENABLE_START_TIME bool
	START_TIME time.Time

	EXTRA_LOGGING bool
	EMAIL_DOMAIN string

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
		ENABLE_START_TIME: helpers.Enviroment("ENABLE_START_TIME", "false") == "true",
		START_TIME: helpers.ParseTime(helpers.Enviroment("START_TIME")),

		EXTRA_LOGGING: helpers.Enviroment("EXTRA_LOGGING", "false") == "true",
		EMAIL_DOMAIN: helpers.Enviroment("EMAIL_DOMAIN"),

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