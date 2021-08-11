
## Development

### Directory structure

| Folder            | Description |
| ----------------- | ------------- |
| .github/workflows | CI - GitHub action manifests. |
| build             | Packaging and integration files. |
| cmd               | Main applications for this project. |
| cmd/app           | Backend server. |
| cmd/cli           | CLI client. |
| deployment        | Cloud, system and container orchestration deployment configurations and templates. |
| dist              | Ephemeral directory - contains the build result |
| docs              | Design, user documents and godoc generated documentation. |
| internal          | Private application and library code. |
| internal/controllers | Responds to the input and performs interactions on the data model objects. |
| internal/models   | Model files to reflect objects from the datasource. |
| internal/locale   | Messages and error codes for public API/Endpoints. |
| log               | Ephemeral directory - logs output. |
| pkg               | Library code that's ok to use by external applications. |
| pkg/sdk/v1        | Version 1 of our SDK. |
| ui                | TBD |

### Development Environment Setup

Make sure you have git LFS installed: https://git-lfs.github.com

#### Backend

Install all Go.mod dependencies

    go mod download

Install Swagger

    go get github.com/go-swagger/go-swagger/cmd/swagger

Install Air

    go get -u github.com/cosmtrek/air

Install Go linter

    brew install golangci-lint

#### Database Setup

The first thing you need to do is open up the "cmd/api/config/*.yml" file and edit it to use the correct
usernames, passwords, hosts, etc... that are appropriate for your environment.

You will also need to make sure that **you** start/install postgres.

Ok, so you've edited the "database.yml" file and started the database, now we can create the databases in
that file for you:

	buffalo pop create -a

#### Starting the Application

We use the utility `air` to watch our code and automatically rebuild the Go binary (and any assets). To do
that run the "air" command:

	air

If you point your browser to [http://127.0.0.1:3000](http://127.0.0.1:3000) you should see the home page.

**Congratulations!** the application up and running.

#### Get auth bearer token

Request a bearer token:

    export AUTH0_API_CLIENT_ID=XXX
    export AUTH0_API_CLIENT_SECRET=XXX
    curl --request POST \
           --url https://netcp-dev.eu.auth0.com/oauth/token \
           --header 'content-type: application/json' \
           --data '{"client_id":"${AUTH0_API_CLIENT_ID}","client_secret":"${AUTH0_API_CLIENT_SECRET}","audience":"http://127.0.0.1:3000/api/v1/","grant_type":"client_credentials"}'

Response:

    {
    "access_token": "XXX", 
    "token_type": "Bearer"
    }

Use the access_token above, for example, with curl:

    curl --header 'authorization: Bearer XXX' http://127.0.0.1:3000/api/v1/

#### Code Signing (macOS)

First you'll want to list all the signing identities on your machines

    security find-identity -v -p codesigning

Then edit the Makefile and replace the current command with your signature. This avoids MacOS to complain about the process listening to a port.