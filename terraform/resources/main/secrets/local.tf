locals {
  secrets = yamldecode(file("../../../../setting/setting.yaml"))["secrets"]
}