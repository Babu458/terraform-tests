provider "azurerm"{
  features {}
  subscription_id = "55f702f9-17ee-4d42-8da3-3f0bc97c4158"
}

terraform {
  backend "local" {
    path = "./terraform.tfstate"
  }
}

#Creating an Azure SQL database
resource "azurerm_mssql_database" "test-one" {
  name           = var.db_name
  server_id      = var.server_id
  collation      = "SQL_Latin1_General_CP1_CI_AS"
  license_type   = "LicenseIncluded"
  sku_name       = "DW100c"
  transparent_data_encryption_enabled = var.db_encryption
}