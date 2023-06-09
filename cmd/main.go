package main

import (
	"net/http"

	"github.com/andreashanson/go-template/internal/infra/mongo"
	handler "github.com/andreashanson/go-template/internal/transport/http"
	"github.com/andreashanson/go-template/internal/user"
)

func main() {
	mongoConn := mongo.NewConnection()
	mongoRepo := mongo.NewUserRepo(mongoConn)
	userSvc := user.New(mongoRepo)

	uh := handler.NewUserHandler(userSvc)
	h := handler.NewHandler(uh).Routes()

	if err := http.ListenAndServe(":5500", h); err != nil {
		panic(err)
	}
}
