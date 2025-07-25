package firebase

import (
	"context"
	"fmt"
	"io"
	"os"

	"cloud.google.com/go/storage"
	firebase "firebase.google.com/go/v4"
	"firebase.google.com/go/v4/messaging"
	"github.com/Cenzios/pf-backend/pkg/logger"
	"google.golang.org/api/option"
)

var (
	firebaseApp     *firebase.App
	storageClient   *storage.Client
	messagingClient *messaging.Client
	bucketName      string
)

// Init initializes Firebase app, storage, and messaging clients
func Init() error {
	serviceAccount := os.Getenv("FIREBASE_SERVICE_ACCOUNT")
	bucketName = os.Getenv("FIREBASE_STORAGE_BUCKET")
	if serviceAccount == "" || bucketName == "" {
		return fmt.Errorf("FIREBASE_SERVICE_ACCOUNT or FIREBASE_STORAGE_BUCKET not set")
	}

	opt := option.WithCredentialsFile(serviceAccount)
	ctx := context.Background()

	app, err := firebase.NewApp(ctx, nil, opt)
	if err != nil {
		return fmt.Errorf("error initializing firebase app: %w", err)
	}
	firebaseApp = app

	storageClient, err = storage.NewClient(ctx, opt)
	if err != nil {
		return fmt.Errorf("error initializing storage client: %w", err)
	}

	messagingClient, err = app.Messaging(ctx)
	if err != nil {
		return fmt.Errorf("error initializing messaging client: %w", err)
	}

	logger.Info.Println("✅ Connected to Firebase")

	return nil
}

// UploadFile uploads a file to the Firebase storage bucket
func UploadFile(ctx context.Context, objectName string, file io.Reader) error {
	w := storageClient.Bucket(bucketName).Object(objectName).NewWriter(ctx)
	if _, err := io.Copy(w, file); err != nil {
		w.Close()
		return err
	}
	return w.Close()
}

// DownloadFile downloads a file from the Firebase storage bucket
func DownloadFile(ctx context.Context, objectName string) (io.ReadCloser, error) {
	r, err := storageClient.Bucket(bucketName).Object(objectName).NewReader(ctx)
	if err != nil {
		return nil, err
	}
	return r, nil
}

// SendPushNotification sends a push notification to a device
func SendPushNotification(ctx context.Context, token, title, body string, data map[string]string) (string, error) {
	msg := &messaging.Message{
		Token: token,
		Notification: &messaging.Notification{
			Title: title,
			Body:  body,
		},
		Data: data,
	}
	return messagingClient.Send(ctx, msg)
}
