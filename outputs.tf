output "action_group_id" {
  value = var.enableSMS ? azurerm_monitor_action_group.actionGroupSMS[0].id : var.enableEmail ? var.enableWebHook ? azurerm_monitor_action_group.actionGroupEmailAndWebHook[0].id : azurerm_monitor_action_group.actionGroupEmail[0].id : var.enableWebHook ? azurerm_monitor_action_group.actionGroupWebHook[0].id : ""
}