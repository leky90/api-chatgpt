package firebase_client

import (
	"context"
	"os"
	"path"

	firebase "firebase.google.com/go"
	"firebase.google.com/go/auth"
	"google.golang.org/api/option"
)

var fbAuth *auth.Client

func InitFirebaseClient() *auth.Client {
	wd, err := os.Getwd()
	if err != nil {
		panic(err)
	}
	relPath := path.Join(wd, "credential.json")
	opt := option.WithCredentialsFile(relPath)
	ctx := context.Background()
	app, err := firebase.NewApp(ctx, nil, opt)
	if err != nil {
		panic(err)
	}

	// Lấy Firebase Authentication client
	client, err := app.Auth(ctx)
	if err != nil {
		panic(err)
	}

	fbAuth = client

	return client
}

func GetFirebaseClient() *auth.Client {
	return fbAuth
}
