package files

import (
	"context"
	"fmt"
	"log"
	"time"

	"cloud.google.com/go/storage"
)

// SetLifecycleManagement adds a lifecycle delete rule with the condition that the object is 2 days old.
func (r *repository) SetLifecycleManagement() error {
	ctx := context.Background()

	ctx, cancel := context.WithTimeout(ctx, time.Second*10)
	defer cancel()

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

	attrs, err := r.BucketHandle.Update(ctx, bucketAttrsToUpdate)
	if err != nil {
		return fmt.Errorf("Bucket.Update: %v", err)
	}
	log.Printf("Lifecycle management is enabled and the rules are:\n")
	for _, rule := range attrs.Lifecycle.Rules {
		log.Printf("Action: %v\n", rule.Action)
		log.Printf("Condition: %v\n", rule.Condition)
	}

	return nil
}

// DisableLifecycleManagement removes all existing lifecycle rules from the bucket.
func (r *repository) DisableLifecycleManagement() error {
	ctx := context.Background()

	ctx, cancel := context.WithTimeout(ctx, time.Second*10)
	defer cancel()

	bucketAttrsToUpdate := storage.BucketAttrsToUpdate{
		Lifecycle: &storage.Lifecycle{},
	}

	_, err := r.BucketHandle.Update(ctx, bucketAttrsToUpdate)
	if err != nil {
		return fmt.Errorf("Bucket.Update: %v", err)
	}
	log.Printf("Lifecycle management is disabled for bucket.\n")

	return nil
}
