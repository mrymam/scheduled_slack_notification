variable "project" {
  type = string
}

variable "region" {
  type = string
}

variable "image_url" {
  type = string
}

variable "jobname" {
  type = string
}

variable "env_vars" {
  type = map(string)
}

variable "secrets" {
  type = map(string)
}