variable "webapp_name"{
  description = "webapp name"
  type        = string
  default     = "App Server"
}

variable "webapp_location"{
  description = "location"
  type        = string
  default     = "East US"
}

variable "ftps_state"{
  description = "FTP state"
  type        = bool
  default     = FtpsOnly
}