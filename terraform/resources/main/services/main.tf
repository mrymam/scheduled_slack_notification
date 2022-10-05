locals {
  services = ["cloudbuild.googleapis.com", "secretmanager.googleapis.com", "secretmanager.googleapis.com", "cloudscheduler.googleapis.com", "run.googleapis.com"]
}

resource "google_project_service" "services" {
  for_each = toset(local.services)
  project = local.project
  service = each.value
}