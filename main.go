package main

import (
	"net/http"
	"os"

	"gopkg.in/redis.v4"

	"golang.org/x/net/context"

	"github.com/go-kit/kit/log"

	httptransport "github.com/go-kit/kit/transport/http"
)

func main() {
	ctx := context.Background()
	logger := log.NewLogfmtLogger(os.Stderr)

	var svc ProphetService
	svc = &prophetService{
		redis.NewClient(&redis.Options{
			Addr:		"redis:6379",
			Password:	"",
			DB:			0,
		}),
	}

	svc.Prepare()

	consultHandler := httptransport.NewServer(
		ctx,
		makeConsultEndpoint(svc),
		decodeConsultRequest,
		encodeResponse,
	)

	http.Handle("/prophet/consult", consultHandler)
	logger.Log("err", http.ListenAndServe(":8080", nil))
}
