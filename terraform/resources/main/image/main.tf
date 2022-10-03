module "image" {
  source = "../../../modules/image_registory"

  project    = var.project
  region     = var.region
  image_name = "notification"
}