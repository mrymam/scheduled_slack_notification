package query

import (
	"embed"
	"fmt"
	"os"
)

var (
	//go:embed files/*.sql
	files embed.FS
)

func LoadQuery(filename string) string {
	filepath := getFilepath(filename)
	buf, err := files.ReadFile(filepath)
	if err != nil {
		panic(err)
	}

	buf = []byte(os.ExpandEnv(string(buf)))
	return string(buf)
}

func getFilepath(filename string) string {
	return fmt.Sprintf("files/%s", filename)
}
