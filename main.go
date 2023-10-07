package main

import (
	"github.com/LCRERGO/drucssGO/pkg/domain/service"
	"github.com/LCRERGO/drucssGO/pkg/network/http"
	"github.com/gofiber/fiber/v2"
)

func main() {
	serv := service.NewService()
	fiberConn := fiber.New()
	h := http.NewHandler(serv)
	http.Configure(h, fiberConn)

	fiberConn.Listen(":3000")
}
