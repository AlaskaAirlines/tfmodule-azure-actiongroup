provider "azurerm" {
  # whilst the `version` attribute is optional, we recommend pinning to a given version of the Provider
  version = "=2.0.0"
  features {}
}

data "azurerm_resource_group" "rg" {
  name = var.resource-group-name
}

resource "azurerm_monitor_action_group" "actionGroupEmail" {
  count = var.enableEmail && ! var.enableWebHook && ! var.enableWebHook ? 1 : 0

  name                = "${var.appName}-${var.environment}-actiongroup"
  resource_group_name = data.azurerm_resource_group.rg.name
  short_name          = var.shortName
  tags                = data.azurerm_resource_group.rg.tags

  email_receiver {
    name                    = var.emailName
    email_address           = var.emailAddress
    use_common_alert_schema = true
  }
}

resource "azurerm_monitor_action_group" "actionGroupSMS" {
  count = var.enableSMS && ! var.enableEmail && ! var.enableWebHook ? 1 : 0

  name                = "${var.appName}-${var.environment}-actiongroup"
  resource_group_name = data.azurerm_resource_group.rg.name
  short_name          = var.shortName
  tags                = data.azurerm_resource_group.rg.tags

  sms_receiver {
    name         = var.smsName
    country_code = var.smsCountryCode
    phone_number = var.smsPhoneNumber
  }
}

resource "azurerm_monitor_action_group" "actionGroupWebHook" {
  count = var.enableWebHook && ! var.enableEmail && ! var.enableSMS ? 1 : 0

  name                = "${var.appName}-${var.environment}-actiongroup"
  resource_group_name = data.azurerm_resource_group.rg.name
  short_name          = var.shortName
  tags                = data.azurerm_resource_group.rg.tags

  webhook_receiver {
    name                    = "callmyapiaswell"
    service_uri             = "http://example.com/alert"
    use_common_alert_schema = true
  }
}

resource "azurerm_monitor_action_group" "actionGroupEmailAndWebHook" {
  count = var.enableEmail && var.enableWebHook && ! var.enableSMS ? 1 : 0

  name                = "${var.appName}-${var.environment}-actiongroup"
  resource_group_name = data.azurerm_resource_group.rg.name
  short_name          = var.shortName
  tags                = data.azurerm_resource_group.rg.tags

  email_receiver {
    name                    = var.emailName
    email_address           = var.emailAddress
    use_common_alert_schema = true
  }

  webhook_receiver {
    name                    = "callmyapiaswell"
    service_uri             = "http://example.com/alert"
    use_common_alert_schema = true
  }
}