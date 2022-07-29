provider "azurerm"{
  features {}
  subscription_id = "55f702f9-17ee-4d42-8da3-3f0bc97c4158"
}

resource "azurerm_mssql_server" "encrypt-terratest" {
  name                         = "encrypt-terratest-2"
  resource_group_name          = "Zab-Policies-Test"
  location                     = "East US"
  version                      = "12.0"
  administrator_login          = "4dm1n157r470r"
  administrator_login_password = "4-v3ry-53cr37-p455w0rd"
}

resource "azurerm_mssql_database" "test" {
  name           = "encrypt_test"
  server_id      = azurerm_mssql_server.encrypt-terratest.id
  collation      = "SQL_Latin1_General_CP1_CI_AS"
  license_type   = "LicenseIncluded"
  sku_name       = "DW100c"
  transparent_data_encryption_enabled = false
  tags = {
    foo = "bar"
  }
}