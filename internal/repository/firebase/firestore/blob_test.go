package firestore

import (
	"context"
	firebase "firebase.google.com/go/v4"
	"github.com/crashdump/netcp/pkg/entity"
	"github.com/google/uuid"
	"github.com/stretchr/testify/suite"
	"github.com/teris-io/shortid"
	"testing"
	"time"
)

var testBlobs = entity.BlobMetadatas{
	{
		ID:        uuid.New(),
		ShortID:   shortid.MustGenerate(),
		OwnerID:   uuid.MustParse("00000000-7357-0000-7357-000000000000"),
		Filename:  "test.tar.gz",
		CreatedAt: time.Now(),
	},
}

type BlobRepositoryTestSuite struct {
	suite.Suite
	blobIndexRepository BlobIndexRepository
}

// We need this function to kick off the test suite,
// otherwise "go test" won't know about our tests
func TestRepositoryTestSuite(t *testing.T) {
	suite.Run(t, new(BlobRepositoryTestSuite))
}

func (suite *BlobRepositoryTestSuite) SetupTest() {
	var err error
	ctx := context.Background()

	fbcli, err := firebase.NewApp(ctx, &firebase.Config{
		ProjectID: "cloudcopy-it",
	})
	suite.NoError(err)

	suite.blobIndexRepository, err = NewBlobIndexRepo(fbcli)
	suite.NoError(err)
}

func (suite *BlobRepositoryTestSuite) TestBlobRepository() {
	for _, b := range testBlobs {
		err := suite.blobIndexRepository.Save(&b)
		suite.NoError(err)

		blobByID, err := suite.blobIndexRepository.GetByID(b.ID, b.OwnerID)
		suite.NoError(err)
		suite.Equal(b.ID, blobByID.ID)
		suite.Equal(b.ShortID, blobByID.ShortID)
		suite.Equal(b.Filename, blobByID.Filename)

		blobByShortID, err := suite.blobIndexRepository.GetByShortID(b.ShortID, b.OwnerID)
		suite.NoError(err)
		suite.Equal(b.ID, blobByShortID.ID)
		suite.Equal(b.ShortID, blobByShortID.ShortID)
		suite.Equal(b.Filename, blobByShortID.Filename)

		err = suite.blobIndexRepository.Delete(&b)
		suite.NoError(err)

		_, err = suite.blobIndexRepository.GetByID(b.ID, b.OwnerID)
		suite.Error(err)
	}
}