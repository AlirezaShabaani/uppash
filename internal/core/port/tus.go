package port

import (
	tusd "github.com/tus/tusd/pkg/handler"
)

type TusSrv interface {
	Uploader() *tusd.Handler
}
