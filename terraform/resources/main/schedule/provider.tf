provider "google" {
  project = local.project
  region  = local.region
}

terraform {
  backend "gcs" {
    prefix = "terraform/main/scheduler"
  }
}
