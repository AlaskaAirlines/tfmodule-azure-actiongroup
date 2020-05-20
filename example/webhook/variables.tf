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