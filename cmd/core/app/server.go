package app

import (
	"auditLog/db/storage/mongo_store"
	repository_prt "auditLog/internal/core/port/outboundprt/repositoryprt"
	"auditLog/internal/inboundadpt/audit_log_grpc"
	"auditLog/internal/inboundadpt/audit_log_rabbit"

	"auditLog/pkg/grpcserver"
	"context"

	"github.com/gofiber/adaptor/v2"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type GRPCGatewayMux *runtime.ServeMux

func provideGrpcServer(listenAddress grpcserver.AddressURI, auditRepo repository_prt.AuditLogRepository) grpcserver.Server {

	interceptors := []grpc.UnaryServerInterceptor{}

	return grpcserver.NewServer(listenAddress, interceptors...)
}

func provideGrpcClient(serverAddr grpcserver.AddressURI) (*grpc.ClientConn, error) {
	// Create a client connection to the gRPC server we just started
	// This is where the gRPC-Gateway proxies the requests
	conn, err := grpc.DialContext(
		context.Background(),
		string(serverAddr),
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	)
	if err != nil {
		return nil, err
	}
	return conn, nil
}

func provideGatewayMUX() *runtime.ServeMux {
	return runtime.NewServeMux()
}

func provideGatewayRegisteredMUX(
	gwmux *runtime.ServeMux,
	conn *grpc.ClientConn,
	auditLogSrv audit_log_grpc.Service,

) (GRPCGatewayMux, error) {
	// ctx := context.Background()
	// if err := auditLogSrv.InitGateway(ctx, gwmux, conn); err != nil {
	// 	return nil, err
	// }

	return GRPCGatewayMux(gwmux), nil
}

func provideGatewayServer(gwmux GRPCGatewayMux) *fiber.App {
	app := fiber.New()
	app.Use(cors.New(cors.Config{
		AllowOrigins: "*",
		AllowHeaders: "*",
	}))

	app.Get("/swagger/apidocs.swagger.json", func(c *fiber.Ctx) error {
		return c.SendFile("swagger/apidocs.swagger.json")
	})
	app.Static("/swagger", "swagger/swagger-ui", fiber.Static{
		Index: "index.html",
	})
	app.Use(adaptor.HTTPHandler((*runtime.ServeMux)(gwmux)))

	return app
}

func provideAuditStorage(dbURI mongo_store.DatabaseURI) (mongo_store.Options, error) {
	return mongo_store.New(dbURI)
}
func provideAuditBroker(
	rURI audit_log_rabbit.RabbitURI,
	auditLogGrpcSrv audit_log_grpc.Service,
) (*audit_log_rabbit.RabbitMQ, error) {
	return audit_log_rabbit.New(rURI, auditLogGrpcSrv)
}
