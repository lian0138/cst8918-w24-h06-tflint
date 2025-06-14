# Prefix for resource names to ensure uniqueness
variable "label_prefix" {
  description = "Prefix used for naming resources to ensure uniqueness."
  type        = string
}

# Azure region for resource deployment
variable "region" {
  description = "The Azure region where resources will be deployed."
  type        = string
}

# Admin username for the virtual machine
variable "admin_username" {
  description = "The admin username for the Linux virtual machine."
  type        = string
}