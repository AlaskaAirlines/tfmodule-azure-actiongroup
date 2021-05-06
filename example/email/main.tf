provider "azurerm" {
  # whilst the `version` attribute is optional, we recommend pinning to a given version of the Provider
  version = "=2.17.0"
  features {}
}

module "actionGroup" {
  source = "../../"

  resource_group_name = var.resource_group_name
  appName             = var.appName
  shortName           = var.shortName
  environment         = var.environment
  enableEmail         = var.enableEmail
  emailName           = var.emailName
  emailAddress        = var.emailAddress
}