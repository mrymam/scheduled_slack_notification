resource "google_secret_manager_secret" "secrets" {
  for_each = {
    for d in local.secrets : d => d
  }

  secret_id   = "${var.common_prefix}${each.value}"

  replication {
    automatic = true
  }
}

resource "google_secret_manager_secret_version" "secret_versions" {
  for_each = {
    for secret in google_secret_manager_secret.secrets : secret.id => {
      secret_id: secret.id
    }
  }

  secret = each.value.secret_id
  secret_data = "change-me"
}
