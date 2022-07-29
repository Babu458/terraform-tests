provider "azurerm"{
  features {}
  subscription_id = "55f702f9-17ee-4d42-8da3-3f0bc97c4158"

}

resource "azurerm_mssql_database" "test" {
  name           = "encrypt_test_2"
  server_id      = "/subscriptions/55f702f9-17ee-4d42-8da3-3f0bc97c4158/resourceGroups/Zab-Policies-Test/providers/Microsoft.Sql/servers/encrypt-terratest-2"
  collation      = "SQL_Latin1_General_CP1_CI_AS"
  license_type   = "LicenseIncluded"
  sku_name       = "DW100c"
  transparent_data_encryption_enabled = false
  tags = {
    foo = "bar"
  }
}