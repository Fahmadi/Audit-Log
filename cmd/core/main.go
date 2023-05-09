package main

import (
	"log"
	"os"
	"os/signal"
	"syscall"

	"auditLog/cmd/config"
	"auditLog/cmd/core/app"
	"auditLog/db/storage/mongo_store"
	"auditLog/internal/inboundadpt/audit_log_rabbit"
	"auditLog/pkg/grpcserver"

	_ "github.com/denisenkom/go-mssqldb"
)

func main() {
	if err := run(); err != nil {
		log.Fatal(err)
		os.Exit(1)
	}
}

func run() error {
	cfg := config.LoadConfig()

	// ToDo: load path of migration and other from config's file
	// if cfg.AutoMigrate {
	// 	if err := migrator.Up(cfg.DatabaseURI, "file://./db/migration", "migrations"); err != nil {
	// 		return err
	// 	}
	// }

	a, cleanup, err := app.InitializeApp(grpcserver.AddressURI(cfg.GrpcURI), app.HTTPAddressURI(cfg.HttpURI), mongo_store.DatabaseURI(cfg.DatabaseURI), audit_log_rabbit.RabbitURI(cfg.BrokerURI))
	defer cleanup()
	if err != nil {
		return err
	}

	log.Println("Starting servers")
	go func() {
		log.Printf("Serving gRPC on %v\n", cfg.GrpcURI)
		if err := a.ServeGRPC(); err != nil {
			log.Fatal(err)
		}
	}()

	go func() {
		log.Printf("Serving HTTP on %v\n", cfg.HttpURI)
		if err := a.ServeHTTP(); err != nil {
			log.Fatal(err)
		}
	}()

	quitChannel := make(chan os.Signal, 1)
	signal.Notify(quitChannel, syscall.SIGINT, syscall.SIGTERM)
	<-quitChannel

	return nil
}
