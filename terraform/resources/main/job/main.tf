data "terraform_remote_state" "image" {
  backend = "gcs"

  config = {
    bucket = var.bucket
    prefix = "terraform/main/image"
    # region = var.region
  }
}

module "jobs" {
  source = "../../../modules/runjob"
  for_each = {
    for d in local.schedules : d.name => {
      name = d.name
      resource_name = "${var.common_prefix}${d.name}"
      secrets = {
        for key, secret in d.secrets : key => "${var.common_prefix}${secret}"
      }
    }
  }

  project   = var.project
  region    = var.region
  image_url = data.terraform_remote_state.image.outputs.image_url
  jobname   = each.value.resource_name
  env_vars  = {
    "SCHEDULE": each.value.name,
    "GO_ENV": "prod"
    "GCP_PROJECT_ID": var.project,
  }
  secrets = each.value.secrets
}