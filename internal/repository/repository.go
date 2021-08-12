package repository

import (
	"cloud.google.com/go/firestore"
	"context"
	firebase "firebase.google.com/go/v4"
	"firebase.google.com/go/v4/auth"
	"firebase.google.com/go/v4/storage"
	"google.golang.org/appengine/log"
)

type Fclient struct {
	auth *auth.Client
	firestore *firestore.Client
	storage *storage.Client
}

var fclient Fclient

// Open establishes connection to firebase and saves its handler into fclient *firebase.App
func Open(fa *firebase.App) error {
	var err error
	ctx := context.Background()

	fclient.auth, err = fa.Auth(ctx)
	if err != nil {
		log.Criticalf(ctx, "error initialising auth storage: %v", err)
		return err
	}

	fclient.firestore, err = fa.Firestore(ctx)
	if err != nil {
		log.Criticalf(ctx, "error initialising firestore storage: %v", err)
		return err
	}

	fclient.storage, err = fa.Storage(ctx)
	if err != nil {
		log.Criticalf(ctx, "error initialising firebase storage: %v", err)
		return err
	}

	return nil
}

func Firebase() Fclient {
	return fclient
}