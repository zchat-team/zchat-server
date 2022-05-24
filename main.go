package main

import (
	"github.com/gin-gonic/gin"
	"github.com/zmicro-team/zchat-server/internal/router"
	"github.com/zmicro-team/zchat-server/pkg/runtime"
	"github.com/zmicro-team/zmicro"
	"github.com/zmicro-team/zmicro/core/log"
)

func main() {
	app := zmicro.New(
		zmicro.InitHttpServer(InitHttpServer),
		zmicro.Before(func() error {
			runtime.Setup()
			return nil
		}),
	)

	if err := app.Run(); err != nil {
		log.Fatal(err)
	}
}

func InitHttpServer(r *gin.Engine) error {
	router.Setup(r)
	return nil
}
