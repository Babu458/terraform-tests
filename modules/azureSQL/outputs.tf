output "db_encryption_status_test_one" {
  value = azurerm_mssql_database.test-one.transparent_data_encryption_enabled
}

# output "db_encryption_status_test_two" {
#   value = azurerm_mssql_database.test-two.transparent_data_encryption_enabled
# }

# output "db_encryption_status_test_three" {
#   value = azurerm_mssql_database.test-three.transparent_data_encryption_enabled
# }