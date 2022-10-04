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
    }
  }

  project   = var.project
  region    = var.region
  image_url = data.terraform_remote_state.image.outputs.image_url
  jobname   = each.value.name
  schedule  = each.value.name
}