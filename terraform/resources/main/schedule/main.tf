data "google_compute_default_service_account" "default" {
}

resource "google_cloud_scheduler_job" "schedules" {
  for_each = {
    for d in local.schedules : d.name => {
      resource_name = "${local.common_prefix}${d.name}"
      description   = d.description
      cron          = d.cron
    }
  }
  name        = each.value.resource_name
  description = each.value.description
  schedule    = each.value.cron

  http_target {
    http_method = "POST"
    uri = "https://${local.region}-run.googleapis.com/apis/run.googleapis.com/v1/namespaces/${local.project}/jobs/${each.value.resource_name}:run"
    oauth_token {
      service_account_email = data.google_compute_default_service_account.default.email
    }
  }

  depends_on = [
    google_project_iam_member.secret
  ]
}