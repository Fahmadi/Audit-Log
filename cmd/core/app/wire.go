// go:build wireinject
//go:build wireinject
// +build wireinject

package app

import (
	"auditLog/internal/inboundadpt/audit_log_grpc"

	"auditLog/db/storage/mongo_store"
	"auditLog/internal/core/port/inboundprt/audit_log_prt"
	repository_prt "auditLog/internal/core/port/outboundprt/repositoryprt"
	"auditLog/internal/core/service/audit_log_srv"
	"auditLog/internal/inboundadpt/audit_log_rabbit"
	"auditLog/internal/outboundadpt/audit/mongo"
	"auditLog/pkg/grpcserver"

	"github.com/google/wire"
)

type HTTPAddressURI string

var repoSet = wire.NewSet(
	wire.Bind(new(repository_prt.AuditLogRepository), new(mongo.AuditLogRepository)),
	mongo.New,
)

var serviceSet = wire.NewSet(
	audit_log_srv.New, wire.Bind(new(audit_log_prt.AuditLogService), new(audit_log_srv.Service)),
)

var grpcServiceSet = wire.NewSet(
	audit_log_grpc.New,
)

func InitializeApp(grpcAddress grpcserver.AddressURI, httpAddress HTTPAddressURI, dBURI mongo_store.DatabaseURI, brokerURI audit_log_rabbit.RabbitURI) (App, func(), error) {
	wire.Build(repoSet, serviceSet, grpcServiceSet, provideAuditBroker, provideAuditStorage, provideGrpcServer, provideGrpcClient, provideGatewayMUX, provideGatewayRegisteredMUX, provideGatewayServer, provideApp)
	return App{}, nil, nil
}
