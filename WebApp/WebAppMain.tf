provider "azurerm" {
  features {}
  subscription_id = "bbd10d6a-717c-4f01-9d3f-6d6f810957ef"
}

terraform {
  backend "local" {
    path = "./terraform.tfstate"
  }
}
provider "azurerm" {
  features {}
}

resource "azurerm_resource_group" "test" {
  name     = var.webapp_name
  location = var.webapp_location
}

resource "azurerm_service_plan" "test" {
  name                = var.webapp_name
  resource_group_name = azurerm_resource_group.test.name
  location            = var.webapp_location
  sku_name            = "F1"
  os_type             = "Windows"
}

resource "azurerm_windows_web_app" "test" {
  name                = var.webapp_name
  resource_group_name = azurerm_resource_group.test.name
  location            = var.webapp_location
  service_plan_id     = azurerm_service_plan.test.id

  site_config {
    ftps_state                = "FtpsOnly"
  }
}