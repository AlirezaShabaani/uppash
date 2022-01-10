package driving

import (
	"espad/back/uppash/config"
	"espad/back/uppash/internal/adapters/driven/minio"
	nats_handler "espad/back/uppash/internal/adapters/driven/nats"
	"espad/back/uppash/internal/adapters/driving/httpHandler"
	"espad/back/uppash/internal/core/service"
	"fmt"
	"net/http"
)

func StartEngine() {

	configs := config.LoadConfig()
	s3Client := minio.InitS3Client(configs.MinioEndPoint, configs.MinioAccessKey, configs.MinioSecretKey)
	natsCli := nats_handler.InitNats(configs.NatsUrl)
	tservice := service.New(s3Client,natsCli)
	handler := httpHandler.New(tservice)
	httpHandler.Routes(handler)

	//
	//
	//
	//store := s3store.New("nimble", s3Client)
	//
	//composer := tusd.NewStoreComposer()
	//store.UseIn(composer)
	//err := cli.SetupPreHooks(&tusd.Config{StoreComposer: composer})
	//if err != nil {
	//	panic(err)
	//}
	//handler, err := tusd.NewHandler(tusd.Config{
	//	BasePath:                "/nimble/",
	//	StoreComposer:           composer,
	//	NotifyCompleteUploads:   false,
	//	NotifyCreatedUploads:    false,
	//	NotifyTerminatedUploads: false,
	//	NotifyUploadProgress:    false,
	//})
	//if err != nil {
	//	panic(fmt.Errorf("Unable to create handler: %s", err))
	//}
	//
	//go func() {
	//	for {
	//		select {
	//		case event := <-handler.CompleteUploads:
	//			fmt.Printf("Upload %s finished\n", event.Upload.ID)
	//		case event := <-handler.CreatedUploads:
	//			fmt.Printf("Upload %s createdUpload\n", event.Upload.ID)
	//			broker := nats_handler.NewMessageBrokerHandler(natsCli)
	//			err := broker.Publisher(context.TODO(), "uploader", []byte(event.Upload.ID))
	//			if err != nil {
	//				log.Println("can't publish this message : " + event.Upload.ID)
	//			}
	//			fmt.Printf("Upload %s created\n", event.Upload.ID)
	//		case event := <-handler.UploadProgress:
	//			fmt.Printf("Upload %d progress\n", event.Upload.Offset)
	//		case event := <-handler.TerminatedUploads:
	//			fmt.Printf("Upload %s terminated\n", event.Upload.ID)
	//		}
	//	}
	//}()
	//
	//
	//http.Handle("/files/", http.StripPrefix("/files/", handler))

	err := http.ListenAndServe(":50000", nil)
	if err != nil {
		panic(fmt.Errorf("Unable to listen: %s", err))
	}
}
