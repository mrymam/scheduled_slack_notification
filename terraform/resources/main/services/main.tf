resource "google_project_service" "project" {
  project = var.project
  service = "secretmanager.googleapis.com"
}