
output "sql_database_name" {
  value = azurerm_mssql_database.test.name
}

# output "sql_server_id" {
#   value = azurerm_sql_server.sqlserver.id
# }

# output "sql_server_name" {
#   value = azurerm_sql_server.sqlserver.name
# }

# output "sql_server_full_domain_name" {
#   value = azurerm_sql_server.sqlserver.fully_qualified_domain_name
# }

# output "sql_server_admin_login" {
#   value = azurerm_sql_server.sqlserver.administrator_login
# }

# output "sql_server_admin_login_pass" {
#   value     = azurerm_sql_server.sqlserver.administrator_login_password
#   sensitive = true
# }

output "db_encryption_status_test_one" {
  value = azurerm_mssql_database.test-one.transparent_data_encryption_enabled
}

output "db_encryption_status_test_two" {
  value = azurerm_mssql_database.test-two.transparent_data_encryption_enabled
}

output "db_encryption_status_test_three" {
  value = azurerm_mssql_database.test-three.transparent_data_encryption_enabled
}