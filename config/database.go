package config

import (
	"context"
	"log"
	"os"

	firebase "firebase.google.com/go"
	"firebase.google.com/go/auth"
	"google.golang.org/api/option"
)

type FirebaseClient struct {
	Auth *auth.Client
}

func NewFirebaseClient(config *Config) (*FirebaseClient, error) {
	ctx := context.Background()

	var opt option.ClientOption

	if firebaseConfigJSON := os.Getenv("FIREBASE_CONFIG_JSON"); firebaseConfigJSON != "" {
		opt = option.WithCredentialsJSON([]byte(firebaseConfigJSON))
	} else {
		opt = option.WithCredentialsFile(config.FirebaseConfigPath)
	}

	app, err := firebase.NewApp(ctx, nil, opt)
	if err != nil {
		log.Printf("Error initializing Firebase app: %v\n", err)
		return nil, err
	}

	authClient, err := app.Auth(ctx)
	if err != nil {
		log.Printf("Error initializing Firebase Auth: %v\n", err)
		return nil, err
	}

	log.Println("Firebase connection established successfully")
	return &FirebaseClient{
		Auth: authClient,
	}, nil
}
