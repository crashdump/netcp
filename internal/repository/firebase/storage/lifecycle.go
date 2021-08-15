package storage

import (
	"context"
	"fmt"
	"log"
	"time"

	"cloud.google.com/go/storage"
)

// SetLifecycleManagement adds a lifecycle delete rule with the condition that the object is 2 days old.
func (r *blobRepository) SetLifecycleManagement() error {
	ctx := context.Background()

	ctx, cancel := context.WithTimeout(ctx, time.Second*10)
	defer cancel()

	bucket, err := r.StorageClient.Bucket(r.BucketName)
	if err != nil {
		return err
	}

	bucketAttrsToUpdate := storage.BucketAttrsToUpdate{
		Lifecycle: &storage.Lifecycle{
			Rules: []storage.LifecycleRule{
				{
					Action: storage.LifecycleAction{Type: "Delete"},
					Condition: storage.LifecycleCondition{
						AgeInDays: 2,
					},
				},
			},
		},
	}

	attrs, err := bucket.Update(ctx, bucketAttrsToUpdate)
	if err != nil {
		return fmt.Errorf("Bucket(%q).Update: %v", r.BucketName, err)
	}
	log.Printf("Lifecycle management is enabled for bucket %v\n and the rules are:\n", r.BucketName)
	for _, rule := range attrs.Lifecycle.Rules {
		log.Printf("Action: %v\n", rule.Action)
		log.Printf("Condition: %v\n", rule.Condition)
	}

	return nil
}

// DisableLifecycleManagement removes all existing lifecycle rules from the bucket.
func (r *blobRepository) DisableLifecycleManagement() error {
	ctx := context.Background()

	ctx, cancel := context.WithTimeout(ctx, time.Second*10)
	defer cancel()

	bucket, err := r.StorageClient.Bucket(r.BucketName)
	if err != nil {
		return err
	}

	bucketAttrsToUpdate := storage.BucketAttrsToUpdate{
		Lifecycle: &storage.Lifecycle{},
	}

	_, err = bucket.Update(ctx, bucketAttrsToUpdate)
	if err != nil {
		return fmt.Errorf("Bucket(%q).Update: %v", r.BucketName, err)
	}
	log.Printf("Lifecycle management is disabled for bucket %v.\n", r.BucketName)

	return nil
}
