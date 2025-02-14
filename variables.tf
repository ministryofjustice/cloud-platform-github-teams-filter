variable "chart_version" {
  type        = string
  description = "the helm chart version to deploy"
}

variable "ecr_url" {
  type        = string
  description = "the url of the ecr repo"
}

variable "image_tag" {
  description = "the container image to deploy"
  type        = string
}

variable "hostname" {
  description = "the hostname to use for the service"
  type        = string
  default     = "github-teams-filter.apps.live.cloud-platform.service.justice.gov.uk"
}

variable "replica_count" {
  description = "the number of replicas to deploy"
  type        = number
  default     = 1
}

variable "filter_api_key" {
  description = "the api key to use for the filter"
  type        = string
  sensitive   = true
}

variable "logLevel" {
  description = "gin log level to use for the api"
  type        = string
  default     = "release"
}