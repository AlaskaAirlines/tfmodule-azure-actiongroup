provider "azurerm" {
  # whilst the `version` attribute is optional, we recommend pinning to a given version of the Provider
  version = "=2.17.0"
  features {}
}

module "actionGroup" {
  source = "../../"

  resource-group-name = var.resource-group-name
  appName             = var.appName
  shortName           = var.shortName
  environment         = var.environment
  enableEmail         = var.enableEmail
  emailName           = var.emailName
  emailAddress        = var.emailAddress
}