package httpHandler

import "net/http"

func Routes(handler HttpHndlr) {
	http.Handle("/files/", http.StripPrefix("/files/", handler.services.Uploader()))
}
