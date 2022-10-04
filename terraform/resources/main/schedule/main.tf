data "google_compute_default_service_account" "default" {
}

resource "google_cloud_scheduler_job" "schedules" {
  for_each = {
    for d in local.schedules : d.name => {
      name        = d.name
      description = d.description
      cron        = d.cron
    }
  }
  name        = each.value.name
  description = each.value.description
  schedule    = each.value.cron

  http_target {
    http_method = "POST"
    uri = "https://${var.region}-run.googleapis.com/apis/run.googleapis.com/v1/namespaces/${var.project}/jobs/${each.value.name}:run"
    oauth_token {
      service_account_email = data.google_compute_default_service_account.default.email
    }
  }
}