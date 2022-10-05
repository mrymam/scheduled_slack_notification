locals {
  tf = yamldecode(file("../../../../setting.yaml"))["terraform"]

  project       = local.tf["project"]
  region        = local.tf["region"]
  gcr_location  = local.tf["gcr_location"]
  common_prefix = local.tf["common_prefix"]
  bucket        = local.tf["bucket"]
}
