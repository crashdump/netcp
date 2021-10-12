package main

import (
	"os"
	"testing"
)

func Test_isDirectory(t *testing.T) {
	tests := []struct {
		name      string
		path      string
		directory bool
		wantErr   bool
	}{
		{
			name:      "directory",
			path:      "/tmp/netcp-it-25t422e-directory",
			directory: true,
			wantErr:   false,
		},
		{
			name:      "file",
			path:      "/tmp/netcp-it-25t422e-file",
			directory: false,
			wantErr:   false,
		},
		{
			name:      "does-not-exist",
			path:      "/dev/nope",
			directory: false,
			wantErr:   true,
		},
	}

	// setup test directory
	os.Mkdir("/tmp/netcp-it-25t422e-directory", 0755)
	os.Create("/tmp/netcp-it-25t422e-file")

	// run tests
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotDirectory, err := isDirectory(tt.path)
			if (err != nil) != tt.wantErr {
				t.Errorf("isDirectory() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if gotDirectory != tt.directory {
				t.Errorf("isDirectory() gotDirectory = %v, directory %v", gotDirectory, tt.directory)
			}
		})
	}
}