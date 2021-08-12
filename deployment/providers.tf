terraform {
  required_providers {
    auth0 = {
      source  = "alexkappa/auth0"
      version = "0.21.0"
    }
  }
}

provider "auth0" {
  domain = "netcp-dev.eu.auth0.com"
  client_id = "<client-id>"
  debug = false
}