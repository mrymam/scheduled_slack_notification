resource "google_project_service" "project" {
  project = local.project
  service = "secretmanager.googleapis.com"
}