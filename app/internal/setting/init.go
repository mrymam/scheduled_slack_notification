package setting

import (
	"os"

	"github.com/onyanko-pon/scheduled_slack_notification/app/pkg/config"
	"gopkg.in/yaml.v3"
)

var setting Setting

func init() {
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
	if config.GetEnv().IsProd() {
		return "../../config.yaml"
	}
	return "test.yaml"
}
