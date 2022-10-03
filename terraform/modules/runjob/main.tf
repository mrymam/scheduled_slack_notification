module "gcloud" {
  source  = "terraform-google-modules/gcloud/google"
  version = "~> 2.0"

  platform = "linux"
  additional_components = ["beta"]

  create_cmd_entrypoint  = "gcloud"
  create_cmd_body        = "beta run jobs create ${var.jobname} --image=gcr.io/${var.image_url} --project ${var.project} --region ${var.region}"
  destroy_cmd_entrypoint = "gcloud"
  destroy_cmd_body       = "beta run jobs delete ${var.jobname} --project ${var.project} --region ${var.region}"
}