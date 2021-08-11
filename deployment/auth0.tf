resource "auth0_client" "auth0" {
  name = "Netcp"
  description = ""
  app_type = "spa"
  custom_login_page_on = true
  is_first_party = true
  is_token_endpoint_ip_header_trusted = true
  grant_types = [ "authorization_code", "http://auth0.com/oauth/grant-type/password-realm", "implicit", "password", "refresh_token" ]
  callbacks = [ "http://127.0.0.1:3000/ui/auth/callback" ]
  allowed_origins = [ "http://127.0.0.1:3000" ]
  allowed_logout_urls = [ "http://127.0.0.1:3000" ]
  web_origins = [ "http://127.0.0.1:3000" ]
  jwt_configuration {
    lifetime_in_seconds = 300
    secret_encoded = true
    alg = "RS256"
  }
}

resource "auth0_client" "auth0" {
  name = "Netcp API"
  description = ""
  app_type = "non_interactive"
  jwt_configuration {
    lifetime_in_seconds = 300
    secret_encoded = true
    alg = "RS256"
  }
}