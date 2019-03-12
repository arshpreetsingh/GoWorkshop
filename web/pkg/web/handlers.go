package web

import (
	"net/http"
	"os"

	"google.golang.org/grpc"

	"github.com/gocraft/web"
	"github.com/gorilla/context"

	log "github.com/sirupsen/logrus"
	pbBounce "github.com/wizelineacademy/GoWorkshop/proto/bounce"
)

// Context struct
type Context struct {
	BounceService pbBounce.BounceClient
}

// ListenAndServe func
func ListenAndServe() {

	connBounce, errBounce := grpc.Dial(os.Getenv("SRV_BOUNCE_ADDR"), grpc.WithInsecure())
	if errBounce != nil {
		log.WithError(errBounce).Fatal("could not connect to bounce service")
	}

	ctx := new(Context)
	ctx.BounceService = pbBounce.NewBounceClient(connBounce)

	r := web.New(Context{}).
		Get("/", ctx.home).
		Post("/getbounce", ctx.getbounce)

	serveMux := http.NewServeMux()
	serveMux.Handle("/", r)

	log.Info("web service started")

	http.ListenAndServe(":8080", context.ClearHandler(serveMux))
}
