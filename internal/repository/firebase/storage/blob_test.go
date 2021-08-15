package storage

import (
	"context"
	firebase "firebase.google.com/go/v4"
	"github.com/crashdump/netcp/pkg/entity"
	"github.com/google/uuid"
	"github.com/stretchr/testify/suite"
	"testing"
	"time"
)

var testBlobs = entity.Blobs{
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

func (suite *BlobRepositoryTestSuite) SetupTest() {
	var err error
	ctx := context.Background()

	projectId := "cloudcopy-it"
	bucketName := "cloudcopy-it.appspot.com"

	fbcli, err := firebase.NewApp(ctx, &firebase.Config{
		ProjectID: projectId,
	})
	suite.NoError(err)

	suite.blobRepository, err = NewBlobRepo(fbcli, bucketName)
	suite.NoError(err)
}

func (suite *BlobRepositoryTestSuite) TestBlobRepository() {
	for _, b := range testBlobs {
		err := suite.blobRepository.Save(&b)
		suite.NoError(err)

		blob, err := suite.blobRepository.GetByID(b.ID)
		suite.NoError(err)
		suite.Equal(b.Content, blob.Content)
	}
}
