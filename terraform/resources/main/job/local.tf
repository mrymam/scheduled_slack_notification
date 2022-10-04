locals {
  schedules = yamldecode(file("../../../../setting/setting.yaml"))["schedules"]
}