locals {
  ns = "cloud-platform-github-teams-filter"
}

resource "random_password" "api-key" {
  length  = 16
  special = false
}

resource "kubernetes_namespace" "github-teams-filter" {
  metadata {
    name = local.ns

    labels = {
      "name"                               = local.ns
      "pod-security.kubernetes.io/enforce" = "restricted"
    }

    annotations = {
      "cloud-platform.justice.gov.uk/application"   = "Cloud Platform GitHub Teams Filter"
      "cloud-platform.justice.gov.uk/business-unit" = "cloud-platform"
      "cloud-platform.justice.gov.uk/owner"         = "Cloud Platform: platforms@digital.justice.gov.uk"
      "cloud-platform.justice.gov.uk/source-code"   = "https://github.com/ministryofjustice/cloud-platform-github-teams-filter"
    }
  }
}

resource "helm_release" "github-teams-filter" {
  name       = "github-teams-filter-api"
  namespace  = local.ns
  chart      = "cloud-platform-github-teams-filter"
  repository = "https://ministryofjustice.github.io/cloud-platform-helm-charts"
  version    = var.chart_version

  values = [templatefile("${path.module}/templates/values.yaml.tpl", {
    ecrUrl       = var.ecr_url
    imageTag     = var.image_tag
    hostName     = var.hostname
    replicaCount = var.replica_count
  })]

  set_sensitive {
    name  = "apiKey"
    value = random_password.api-key.result
  }
}