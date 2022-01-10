package service

import (
	"context"
	nats_handler "espad/back/uppash/internal/adapters/driven/nats"
	"espad/back/uppash/internal/core/port"
	"fmt"
	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/nats-io/nats.go"
	"github.com/tus/tusd/cmd/tusd/cli"
	tusd "github.com/tus/tusd/pkg/handler"
	"github.com/tus/tusd/pkg/s3store"
	"log"
)

type service struct {
	s3Cli *s3.S3
	natsCli *nats.Conn
}

func New(s3Cli *s3.S3,natsCli *nats.Conn) port.TusSrv {
	return &service{s3Cli: s3Cli,natsCli : natsCli}
}

func (t *service)Uploader() *tusd.Handler {
	store := s3store.New("nimble", t.s3Cli)

	composer := tusd.NewStoreComposer()
	store.UseIn(composer)
	err := cli.SetupPreHooks(&tusd.Config{StoreComposer: composer})
	if err != nil {
		panic(err)
	}
	handler, err := tusd.NewHandler(tusd.Config{
		BasePath:                "/nimble/",
		StoreComposer:           composer,
		NotifyCompleteUploads:   true,
		NotifyCreatedUploads:    true,
		NotifyTerminatedUploads: true,
		NotifyUploadProgress:    true,
	})
	if err != nil {
		panic(fmt.Errorf("Unable to create handler: %s", err))
	}

	go func() {
		for {
			select {
			case event := <-handler.CompleteUploads:
				fmt.Printf("Upload %s finished\n", event.Upload.ID)
			case event := <-handler.CreatedUploads:
				fmt.Printf("Upload %s createdUpload\n", event.Upload.ID)
				broker := nats_handler.NewMessageBrokerHandler(t.natsCli)
				err := broker.Publisher(context.TODO(), "uploader", []byte(event.Upload.ID))
				if err != nil {
					log.Println("can't publish this message : " + event.Upload.ID)
				}
				fmt.Printf("Upload %s created\n", event.Upload.ID)
			case event := <-handler.UploadProgress:
				fmt.Printf("Upload %d progress\n", event.Upload.Offset)
			case event := <-handler.TerminatedUploads:
				fmt.Printf("Upload %s terminated\n", event.Upload.ID)
			}
		}
	}()

	return handler
}
