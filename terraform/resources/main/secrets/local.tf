locals {
  secrets = yamldecode(file("../../../schedules.yaml"))["secrets"]
}