variable "server_id" {
  description = "Existing Azure Server"
  type        = string
  default     = "/subscriptions/55f702f9-17ee-4d42-8da3-3f0bc97c4158/resourceGroups/Zab-Policies-Test/providers/Microsoft.Sql/servers/encrypt-terratest-2"
}

variable "db_encryption"{
  description = "DB Encryption status"
  type        = bool
  default     = true
}

variable "db_name"{
  description = "DB Name"
  type        = string
  default     = "database"
}