package setting

import (
	"embed"
	"log"
	"os"

	"github.com/joho/godotenv"
	"github.com/onyanko-pon/scheduled_slack_notification/app/pkg/config"
	"gopkg.in/yaml.v3"
)

var (
	//go:embed files/*.yaml
	files embed.FS
)

var setting Setting

func init() {
	if !config.GetEnv().IsTest() {
		err := godotenv.Load()
		if err != nil {
			log.Fatal("Error loading .env file")
		}
	}

	filepath := getFilepath()
	buf, err := files.ReadFile(filepath)
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
		return "files/setting.yaml"
	}
	return "files/test.yaml"
}
