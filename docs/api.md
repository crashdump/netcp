## API

## Definitions

You can find the swagger definition here: https://github.com/crashdump/netcp/docs/swagger.

## Error codes

In addition to descriptive error text, error messages contain machine-parseable codes. While the text for an error
message may change, the codes will stay the same.

The following table describes the codes which may appear when working with the Netcp API. If an error response
is not listed in the table, fall back to examining the HTTP status codes above in order to determine the best way
to address the issue. Please also use the above tables for troubleshooting tips for each corresponding HTTP status code.

Code | Message | Description
-----|------|-------------
1000 | Could not authenticate you. | Corresponds with HTTP 401. There was an issue with the authentication data for the request.
1001 | Invalid or expired token. | Corresponds with HTTP 403. The access token used in the request is incorrect or has expired.
1001 | User has been suspended. | Corresponds with HTTP 403 The user account has been suspended and information cannot be retrieved.
1020 | Rate limit exceeded. | Corresponds with HTTP 429. The request limit for this resource has been reached for the current rate limit window.
4000 | Sorry, that endpoint does not exist. | Corresponds with HTTP 404. The specified resource was not found.
4001 | User not found. | Corresponds with HTTP 404. The user is not found.
4002 | Query parameters are missing. | Corresponds with HTTP 404.