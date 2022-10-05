locals {
  vars = yamldecode(file("../../../../setting.yaml"))
  tf = local.vars["terraform"]

  project       = local.tf["project"]
  region        = local.tf["region"]
  bucket        = local.tf["bucket"]
}