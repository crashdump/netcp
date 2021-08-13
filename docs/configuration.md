# Configuration

The tool suite stores its configuration in the `.netcp` folder, located in the user's home directory.

## API

`~/.netcp/cli.ENV.yaml`

Example:

```yaml
server:
  hostname: 127.0.0.1
  port: 3000
  tls:
    enabled: false
  postgres:
    host: localhost
    user: postgres
    password: secret_here
    dbname: netcp-dev
    port: 5432
    sslmode: disable

auth:
  domain: netcp-dev.eu.auth0.com
  client_id: HlpkvBqBPwLLSTMTpaeI54Gh5H0R73NB
  audience: http://127.0.0.1:3000/srv/v1/
  api_key: secret_here

rollbar:
  enabled: false
  token: secret_here
```

## CLI

`~/.netcp/cli.ENV.yaml`

```yaml
api:
  url: http://127.0.0.1:3000/srv/v1

auth:
  refresh_token: secret_here
```