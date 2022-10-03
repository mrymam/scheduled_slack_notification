package config

import "os"

var gcp GCP

type GCP struct {
	ProjectID string
}

func loadGCP() {
	id := os.Getenv("GCP_PROJECT_ID")
	gcp = GCP{
		ProjectID: id,
	}
}

func GetGCP() GCP {
	return gcp
}
