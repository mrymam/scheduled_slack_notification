locals {
  image_url = "gcr.io/${var.project}/${var.image_name}"
}

resource "google_container_registry" "registry" {
  project  = var.project
  location = "ASIA"
}

module "gcloud" {
  source  = "terraform-google-modules/gcloud/google"
  version = "~> 2.0"

  platform = "linux"
  additional_components = []

  create_cmd_entrypoint  = "gcloud"
  create_cmd_body        = "builds submit ${var.build_path} --tag=${local.image_url}"
  destroy_cmd_entrypoint = "gcloud"
  destroy_cmd_body       = "container images delete ${local.image_url} --quiet"
}