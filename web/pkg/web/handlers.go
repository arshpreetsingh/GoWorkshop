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
	BounceService pbBounce.BounceClinet
}

// ListenAndServe func
func ListenAndServe() {

	connBounce, errBounce := grpc.Dial(os.Getenv("SRV_BOUNCE_ADDR"), grpc.WithInsecure())
	if errUsers != nil {
		log.WithError(errUsers).Fatal("could not connect to users service")
	}

	ctx := new(Context)
	ctx.BounceService = pbUsers.NewUsersClient(connBounce)

	r := web.New(Context{}).
		Get("/", ctx.home).
		Post("/", ctx.createUser).
		Get("/user/:id", ctx.user).
		Post("/user/:id", ctx.user)

	serveMux := http.NewServeMux()
	serveMux.Handle("/", r)

	log.Info("web service started")

	http.ListenAndServe(":8080", context.ClearHandler(serveMux))
}
