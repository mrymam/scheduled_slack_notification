locals {
  vars = yamldecode(file("../../../../setting.yaml"))
  tf = local.vars["terraform"]

  schedules     = local.vars["schedules"]
  project       = local.tf["project"]
  region        = local.tf["region"]
  common_prefix = local.tf["common_prefix"]
}