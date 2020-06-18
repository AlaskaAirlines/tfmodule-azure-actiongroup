provider "azurerm" {
  # whilst the `version` attribute is optional, we recommend pinning to a given version of the Provider
  version = "=2.0.0"
  features {}
}

data "azurerm_resource_group" "rg" {
  name = var.resource-group-name
}

resource "azurerm_monitor_action_group" "actionGroup" {
  name                = "${var.appName}-${var.environment}-actiongroup"
  resource_group_name = data.azurerm_resource_group.rg.name
  short_name          = var.shortName
  tags                = data.azurerm_resource_group.rg.tags

  dynamic "email_receiver" {
    for_each = var.enableEmail ? ["email_receiver"] : []
    content {
      name                    = var.emailName
      email_address           = var.emailAddress
      use_common_alert_schema = true
    }
  }

  dynamic "sms_receiver" {
    for_each = var.enableSMS ? ["sms_receiver"] : []
    content {
      name         = var.smsName
      country_code = var.smsCountryCode
      phone_number = var.smsPhoneNumber
    }
  }

  dynamic "webhook_receiver" {
    for_each = var.enableWebHook ? ["webhook_receiver"] : []
    content {
      name                    = var.webhookName
      service_uri             = var.webhookServiceUri
      use_common_alert_schema = true
    }
  }

}