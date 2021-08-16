package storage

import (
	"context"
	firebase "firebase.google.com/go/v4"
	"github.com/crashdump/netcp/pkg/entity"
	"github.com/google/uuid"
	"github.com/stretchr/testify/suite"
	"github.com/teris-io/shortid"
	"strings"
	"testing"
	"time"
)


var (
	testBlobs = []struct {
		blobMetadata entity.BlobMetadata
		blob         entity.Blob
	}{
		{
			blobMetadata: entity.BlobMetadata{
				ID:        uuid.MustParse("66ee4e6f-cae6-42b2-ad57-2bbb12b9b67c"),
				ShortID:   shortid.MustGenerate(),
				OwnerID:   uuid.MustParse("00000000-7357-1111-7357-000000000000"),
				Filename:  "foo.tar.gz",
				CreatedAt: time.Now(),
			},
			blob:         entity.Blob{
				ID:        uuid.MustParse("66ee4e6f-cae6-42b2-ad57-2bbb12b9b67c"),
				Content:   []byte(strings.Repeat("f", 128*4) + "\n"),
				},
		},
	}
)

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
		err := suite.blobRepository.Save(&b.blobMetadata, &b.blob)
		suite.NoError(err)

		blobByID, err := suite.blobRepository.GetByID(&b.blobMetadata)
		suite.NoError(err)
		suite.Equal(b.blob.Content, blobByID.Content)

		err = suite.blobRepository.Delete(&b.blobMetadata)
		suite.NoError(err)

		_, err = suite.blobRepository.GetByID(&b.blobMetadata)
		suite.Error(err)
	}
}