package app

import (
	"auditLog/internal/inboundadpt/audit_log_rabbit"
	"auditLog/pkg/grpcserver"

	"github.com/gofiber/fiber/v2"
)

type App struct {
	grpcServer    grpcserver.Server
	httpServer    *fiber.App
	listenHttpURI HTTPAddressURI
	rabbitBroker  *audit_log_rabbit.RabbitMQ
}

func provideApp(listenHttpURI HTTPAddressURI, grpcServer grpcserver.Server, httpServer *fiber.App, rabbitBroker *audit_log_rabbit.RabbitMQ) App {
	return App{
		grpcServer:    grpcServer,
		httpServer:    httpServer,
		listenHttpURI: listenHttpURI,
		rabbitBroker:  rabbitBroker,
	}
}

func (a App) ServeGRPC() error {
	return a.grpcServer.Serve()
}

func (a App) ServeHTTP() error {
	return a.httpServer.Listen(string(a.listenHttpURI))
}
