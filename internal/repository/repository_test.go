package repository

import (
	"context"
	firebase "firebase.google.com/go/v4"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestOpen(t *testing.T) {
	var err error
	ctx := context.Background()

	fbcli, err := firebase.NewApp(ctx, &firebase.Config{ProjectID: "cloudcopy-it"})
	assert.NoError(t, err)

	err = Open(fbcli)
	assert.NoError(t, err)
}