package api

import (
	"net/http"

	app "agner.io/boast"
	"agner.io/boast/log"
)

type ExportAPI struct {
	http.Handler
}

func NewTestAPI(statusPath string, strg app.Storage) *ExportAPI {
	handler, err := api("", statusPath, strg)
	if err != nil {
		log.Fatalln(err)
	}
	return &ExportAPI{
		Handler: handler,
	}
}
