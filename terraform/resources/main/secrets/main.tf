resource "google_secret_manager_secret" "secrets" {
  for_each = local.secrets

  secret_id   = "${local.common_prefix}${each.value}"

  provisioner "local-exec" {
    command = "sleep 10"
  }

  replication {
    automatic = true
  }
}

resource "google_secret_manager_secret_version" "secret_versions" {
  for_each = google_secret_manager_secret.secrets

  secret = each.value.id
  secret_data = "change-me"

  depends_on = [
    google_secret_manager_secret.secrets
  ]
}
