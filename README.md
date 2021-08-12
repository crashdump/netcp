# netcp

Copy files and directories across systems without requiring a direct network line-of-sight. Netcp uses a cloud
storage endpoint to store the data while in flight.

# Client

## Usage 

Copy files from the source computer:

    computer-1> netcp upload SOURCE
    uploading: Dockerfile
    uploading: README.md
    uploading: go.mod
    complete: sucessfully uploaded, use code r7k-x23-9z2 to retrieve.

Paste files on the target computer:

    computer-2> netcp download DESTINATION --code r7k-x23-9z2
    downloading: Dockerfile
    downloading: README.md
    downloading: go.mod
    complete: sucessfully synchronised the files.

# Backend

netcp's server is built on top of Google Cloud and heavily relies on Firebase Authentication, Firestore and Storage. 

## Documentation

- [index](docs/index.md)
- [api](docs/api.md)
- [configuration](docs/configuration.md)
- [sdk](docs/sdks/index.md)
    - [go](docs/sdks/go.md)