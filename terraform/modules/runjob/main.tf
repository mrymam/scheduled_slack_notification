locals {
   env_vars_str = join(",", [for key, value in var.env_vars : "${key}=${value}"])
   secrets_str  = join(",", [for key, value in var.secrets : "${key}=${value}:latest"])
}

module "gcloud" {
  source  = "terraform-google-modules/gcloud/google"
  version = "~> 2.0"

  platform = "linux"
  additional_components = ["beta"]
  skip_download = true
  enabled       = true
  upgrade       = false

  create_cmd_entrypoint  = "gcloud"
  create_cmd_body        = "beta run jobs create ${var.jobname} --quiet --image=${var.image_url} --set-env-vars=${local.env_vars_str} --set-secrets=${local.secrets_str} --project ${var.project} --region ${var.region}"
  destroy_cmd_entrypoint = "gcloud"
  destroy_cmd_body       = "beta run jobs delete ${var.jobname} --quiet --project ${var.project} --region ${var.region}"
}