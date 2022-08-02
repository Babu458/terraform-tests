provider "azurerm"{
  features {}
  subscription_id = "55f702f9-17ee-4d42-8da3-3f0bc97c4158"
}

terraform {
  backend "local" {
    path = "./terraform.tfstate"
  }
}

#Creating an Azure SQL database with encryption
resource "azurerm_mssql_database" "test-one" {
  name           = "happy-path"
  server_id      = "/subscriptions/55f702f9-17ee-4d42-8da3-3f0bc97c4158/resourceGroups/Zab-Policies-Test/providers/Microsoft.Sql/servers/encrypt-terratest-2"
  collation      = "SQL_Latin1_General_CP1_CI_AS"
  license_type   = "LicenseIncluded"
  sku_name       = "DW100c"
  transparent_data_encryption_enabled = true
}

# #Creating an Azure SQL database without encryption
# resource "azurerm_mssql_database" "test-two" {
#   name           = "sad-path"
#   server_id      = "/subscriptions/55f702f9-17ee-4d42-8da3-3f0bc97c4158/resourceGroups/Zab-Policies-Test/providers/Microsoft.Sql/servers/encrypt-terratest-2"
#   collation      = "SQL_Latin1_General_CP1_CI_AS"
#   license_type   = "LicenseIncluded"
#   sku_name       = "DW100c"
#   transparent_data_encryption_enabled = false
# }

#Toggling the encryption on an existing database to false
resource "azurerm_mssql_database" "test-three" {
  name           = azurerm_mssql_database.test-one.name
  id             ="/subscriptions/55f702f9-17ee-4d42-8da3-3f0bc97c4158/resourceGroups/Zab-Policies-Test/providers/Microsoft.Sql/servers/encrypt-terratest-2/databases/happy-path"
  transparent_data_encryption_enabled = false
}