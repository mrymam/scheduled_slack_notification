locals {
  schedules = yamldecode(file("../../../schedules.yaml"))["schedules"]
}