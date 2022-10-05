resource "google_secret_manager_secret" "secrets" {
  for_each = {
    for key, secret in local.secrets : key => secret
  }

  secret_id   = "${local.common_prefix}${each.value}"

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
