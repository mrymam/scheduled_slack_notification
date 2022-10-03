resource "google_container_registry" "registry" {
  project  = var.project
  location = var.region
}

module "gcloud" {
  source  = "terraform-google-modules/gcloud/google"
  version = "~> 2.0"

  platform = "linux"
  additional_components = []

  create_cmd_entrypoint  = "gcloud"
  create_cmd_body        = "builds submit ./sample --tag=gcr.io/${var.project}/${var.image_name} --region ${var.region}"
  destroy_cmd_entrypoint = "gcloud"
  destroy_cmd_body       = "container images delete --tag=gcr.io/${var.project}/${var.image_name}"
}