provider "google" {
  project = local.project
  region  = local.region
}

terraform {
  backend "gcs" {
    bucket = local.bucket
    prefix = "terraform/main/services"
  }
}
