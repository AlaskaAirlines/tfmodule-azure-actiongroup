module "actionGroup" {
  source = "../../"

  resource-group-name = var.resource-group-name
  appName             = var.appName
  shortName           = var.shortName
  environment         = var.environment
  enableEmail         = var.enableEmail
  emailName           = var.emailName
  emailAddress        = var.emailAddress
  enableWebHook       = var.enableWebHook
  webhookName         = var.webhookName
  webhookServiceUri   = var.webhookServiceUri
}