package metadata

import (
	"context"
	"github.com/google/uuid"
	"log"

	"cloud.google.com/go/firestore"
	firebase "firebase.google.com/go/v4"
	"github.com/crashdump/netcp/pkg/entity"
)

type Repository interface {
	GetByID(ID uuid.UUID, owner uuid.UUID) (entity.BlobMetadata, error)
	GetByShortID(ID string, owner uuid.UUID) (entity.BlobMetadata, error)
	Save(blob *entity.BlobMetadata) error
	Delete(blob *entity.BlobMetadata) error
}

type repository struct {
	ctx             context.Context
	FirestoreClient *firestore.Client
}

//NewMetadataRepo is the single instance repo that is being created.
func NewMetadataRepo(fc *firebase.App) (Repository, error) {
	ctx := context.Background()

	s, err := fc.Firestore(ctx)
	if err != nil {
		log.Printf("error initialising firestore: %v", err)
		return nil, err
	}

	bir := &repository{
		ctx:             ctx,
		FirestoreClient: s,
	}

	//err = bir.SetLifecycleManagement()

	return bir, err
}

func (b repository) GetByID(id uuid.UUID, owner uuid.UUID) (entity.BlobMetadata, error) {
	var blob entity.BlobMetadata

	doc, err := b.FirestoreClient.
		Collection("users").
		Doc(owner.String()).
		Collection("blobs").
		Doc(id.String()).
		Get(b.ctx)
	if err != nil {
		log.Printf("GetByID(): %s", err.Error())
		return blob, err
	}

	if !doc.Exists() {
		log.Printf("GetByID() document not found: users/%s/blobs/%s, err: %s", blob.OwnerID, blob.ID, err)
		return blob, err
	}

	err = doc.DataTo(&blob)
	if err != nil {
		log.Printf("GetByID(): %s", err.Error())
		return blob, err
	}

	blob.ID = uuid.MustParse(doc.Ref.ID)
	blob.OwnerID = uuid.MustParse(doc.Ref.Parent.Parent.ID)

	log.Printf("GetByID() document found: users/%s/blobs/%s", blob.OwnerID, blob.ID)
	return blob, err
}

func (b repository) GetByShortID(id string, owner uuid.UUID) (entity.BlobMetadata, error) {
	var blob entity.BlobMetadata

	iter := b.FirestoreClient.
		Collection("users").
		Doc(owner.String()).
		Collection("blobs").
		Where("short_id", "==", id).
		Limit(1).
		Documents(b.ctx)

	docs, err := iter.GetAll()
	if err != nil {
		log.Printf("GetByShortID(): %s", err.Error())
		return blob, err
	}

	if !docs[0].Exists() {
		log.Printf("GetByShortID() document not found: users/%s/blobs/%s, err: %s", blob.OwnerID, blob.ID, err)
	}

	err = docs[0].DataTo(&blob)
	if err != nil {
		log.Printf("GetByShortID(): %s", err)
		return blob, err
	}

	blob.ID = uuid.MustParse(docs[0].Ref.ID)
	blob.OwnerID = uuid.MustParse(docs[0].Ref.Parent.Parent.ID)

	log.Printf("GetByShortID() document found: users/%s/blobs/%s", blob.OwnerID, blob.ID)
	return blob, nil
}

func (b repository) Save(blob *entity.BlobMetadata) error {
	ownerId := blob.OwnerID.String()
	docId := blob.ID.String()

	_, err := b.FirestoreClient.
		Collection("users").
		Doc(ownerId).
		Collection("blobs").
		Doc(docId).
		Set(b.ctx, blob)

	if err != nil {
		log.Printf("Save() document not saved: users/%s/blobs/%s, err: %s", ownerId, docId, err)
		return err
	}

	log.Printf("Save() document saved: users/%s/blobs/%s", ownerId, docId)
	return nil
}

func (b repository) Delete(blob *entity.BlobMetadata) error {
	ownerId := blob.OwnerID.String()
	docId := blob.ID.String()

	_, err := b.FirestoreClient.
		Collection("users").
		Doc(ownerId).
		Collection("blobs").
		Doc(docId).
		Delete(b.ctx)

	if err != nil {
		log.Printf("Delete(): %s", err.Error())
	}

	log.Printf("Delete() document deleted: users/%s/blobs/%s", ownerId, docId)
	return nil
}
