data "terraform_remote_state" "image" {
  backend = "gcs"

  config = {
    bucket = var.bucket
    prefix = "terraform/main/image"
    # region = var.region
  }
}

module "image" {
  source = "../../../modules/runjob"

  project   = var.project
  region    = var.region
  image_url = data.terraform_remote_state.image.outputs.image_url
  jobname   = "notification"
}