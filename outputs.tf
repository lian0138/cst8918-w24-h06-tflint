# Name of the Azure resource group
output "resource_group_name" {
  description = "The name of the resource group created for the web server."
  value       = azurerm_resource_group.rg.name
}

# Name of the Azure Linux virtual machine
output "vm_name" {
  description = "The name of the Linux virtual machine created."
  value       = azurerm_linux_virtual_machine.webserver.name
}

# Name of the network interface
output "nic_name" {
  description = "The name of the network interface attached to the virtual machine."
  value       = azurerm_network_interface.webserver.name
}

# Public IP address of the virtual machine
output "public_ip" {
  description = "The public IP address assigned to the virtual machine."
  value       = azurerm_public_ip.webserver.ip_address
}