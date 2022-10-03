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
  create_cmd_body        = "builds submit ${var.build_path} --tag=gcr.io/${var.project}/${var.image_name}"
  destroy_cmd_entrypoint = "gcloud"
  destroy_cmd_body       = "container images delete --tag=gcr.io/${var.project}/${var.image_name}"
}