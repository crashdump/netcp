# netcp

Copy files and directories across systems without requiring a direct network line-of-sight. Netcp uses a cloud
storage endpoint to store the data while in flight.

# Client

## Usage 

Copy files from the source computer:

    computer-1> netcp up FILENAME
    uploading: foo.tar.gz

    computer-1> netcp ls
    1334  11 Aug 19:17  foo.tar.gz
     212  11 Aug 22:14  bar.tar.gz

Paste files on the target computer:

    computer-2> netcp down FILENAME DESTINATION
    downloading: foo.tar.gz

Note: files are automatically deleted after 48h.

## Install

Pre-built clients for Linux, macOS and Windows are available [here], or you can choose to build them from source. I am
working to get them available on popular repositories such as Homebrew, Debian, ArchLinux and CentOS.

# Backend

netcp's server is built on top of Google Cloud and heavily relies on Firebase Authentication, Firestore and
Storage. However, I've tried to keep clean interfaces and porting it to other backend shouldn't be too complicated.

## Documentation

- [index](docs/index.md)
- [api](docs/api.md)
- [configuration](docs/configuration.md)
- [sdk](docs/sdks/index.md)
    - [go](docs/sdks/go.md)