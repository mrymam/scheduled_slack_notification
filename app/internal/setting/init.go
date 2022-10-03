package setting

import (
	"log"
	"os"

	"github.com/joho/godotenv"
	"github.com/onyanko-pon/scheduled_slack_notification/app/pkg/config"
	"gopkg.in/yaml.v3"
)

var setting Setting

func init() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	filepath := getFilepath()
	buf, err := os.ReadFile(filepath)
	if err != nil {
		panic(err)
	}

	buf = []byte(os.ExpandEnv(string(buf)))

	if err := yaml.Unmarshal(buf, &setting); err != nil {
		panic(err)
	}
}

func Get() Setting {
	return setting
}

func getFilepath() string {
	if config.GetEnv().IsProd() || config.GetEnv().IsDev() {
		return "config.yaml"
	}
	return "test.yaml"
}
