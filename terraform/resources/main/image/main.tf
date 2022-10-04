module "image" {
  source = "../../../modules/image_registory"

  project    = var.project
  region     = var.region
  image_name = "${var.common_prefix}notification"
  build_path = "./sample"
  gcr_location = var.gcr_location
}