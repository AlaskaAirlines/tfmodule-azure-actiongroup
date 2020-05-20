module "actionGroup" {
  source = "../../"

  resource-group-name = var.resource-group-name
  appName             = var.appName
  shortName           = var.shortName
  environment         = var.environment
  enableSMS           = var.enableSMS
  smsName             = var.smsName
  smsPhoneNumber      = var.smsPhoneNumber
}