variable "resource-group-name" {
  type        = string
  description = "Name of the resource group that exists in Azure"
}

variable "appName" {
  type        = string
  description = "The base name of the application used in the naming convention."
}

variable "environment" {
  type        = string
  description = "Name of the environment ex (Dev, Test, QA, Prod)"
}

variable "shortName" {
  type        = string
  description = "Required shorthand name for SMS texts."
}

variable "enableEmail" {
  type        = bool
  description = "Enable email alert capabilities"
  default     = false
}

variable "emailName" {
  type        = string
  description = "Friendly Name for email address"
  default     = ""
}

variable "emailAddress" {
  type        = string
  description = "email address to send alerts to"
  default     = ""
}

variable "enableSMS" {
  type        = bool
  description = "Enable Texting Alerts"
  default     = false
}

variable "smsName" {
  type        = string
  description = "Friendly Name for phone number"
  default     = ""
}

variable "smsCountryCode" {
  type        = number
  description = "Country Code for phone number"
  default     = 1
}

variable "smsPhoneNumber" {
  type        = number
  description = "Phone number for text alerts"
  default     = 0
}

variable "enableWebHook" {
  type        = bool
  description = "Enable Web Hook Alerts"
  default     = false
}

variable "webhookName" {
  type        = string
  description = "Friendly Name for web hook"
  default     = ""
}

variable "webhookServiceUri" {
  type        = string
  description = "The full URI for the webhook"
  default     = ""
}