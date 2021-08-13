package storage

import (
	"context"
	firebase "firebase.google.com/go/v4"
	"github.com/crashdump/netcp/pkg/entity"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
	"testing"
	"time"
)

var testBlobs = entity.Blobs {
	{
		ID:        uuid.New(),
		Name:      "test.tar.gz",
		CreatedAt: time.Now(),
		Content:   nil,
	},
}

type BlobRepositoryTestSuite struct {
	suite.Suite
	blobRepository BlobRepository
}

// We need this function to kick off the test suite,
// otherwise "go test" won't know about our tests
func TestRepositoryTestSuite(t *testing.T) {
	suite.Run(t, new(BlobRepositoryTestSuite))
}

func (suite *BlobRepositoryTestSuite) NewTest(t *testing.T) {
	var err error
	ctx := context.Background()

	fbcli, err := firebase.NewApp(ctx, &firebase.Config{
		ProjectID: "cloudcopy-it",
	})
	assert.NoError(t, err)

	suite.blobRepository, err = NewRepo(fbcli, "cloudcopy-it")
	assert.NoError(t, err)
}

func (suite *BlobRepositoryTestSuite) TestBlobRepository(t *testing.T) {
	for _, b := range testBlobs {
		err := suite.blobRepository.Save(&b)
		assert.NoError(t, err)

		blob, err := suite.blobRepository.GetByID(b.ID)
		assert.NoError(t, err)
		assert.Equal(t, b.Content, blob.Content)
	}
}