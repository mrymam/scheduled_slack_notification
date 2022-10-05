module "image" {
  source = "../../../modules/image_registory"

  project    = local.project
  region     = local.region
  image_name = "${local.common_prefix}notification"
  build_path = "./sample"
  gcr_location = local.gcr_location
}