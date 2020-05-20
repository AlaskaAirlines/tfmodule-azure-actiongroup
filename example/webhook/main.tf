module "actionGroup" {
  source = "../../"

  resource-group-name = var.resource-group-name
  appName             = var.appName
  shortName           = var.shortName
  environment         = var.environment
  enableWebHook       = var.enableWebHook
  webhookName         = var.webhookName
  webhookServiceUri   = var.webhookServiceUri
}