data "terraform_remote_state" "image" {
  backend = "gcs"

  config = {
    bucket = local.bucket
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
    }
  }

  project   = local.project
  region    = local.region
  image_url = data.terraform_remote_state.image.outputs.image_url
  jobname   = each.value.resource_name
  env_vars  = {
    "SCHEDULE": each.value.name,
    "GO_ENV": "prod"
    "GCP_PROJECT_ID": local.project,
  }
  secrets = {
    for key, secret in local.secrets : key => "${local.common_prefix}${secret}"
  }
}