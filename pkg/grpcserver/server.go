package grpcserver

import (
	"log"
	"net"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

type Server interface {
	Serve() error
	Instance() *grpc.Server
}

type server struct {
	s             *grpc.Server
	listenAddress AddressURI
}

type AddressURI string

func NewServer(
	listenAddress AddressURI,
	interceptors ...grpc.UnaryServerInterceptor,
) Server {
	msgMaxSize := 20 * 1024 * 1024
	s := grpc.NewServer(
		grpc.ChainUnaryInterceptor(interceptors...),
		grpc.MaxSendMsgSize(msgMaxSize),
		grpc.MaxRecvMsgSize(msgMaxSize),
	)

	return server{
		s:             s,
		listenAddress: listenAddress,
	}
}

func (gs server) Instance() *grpc.Server {
	return gs.s
}

func (gs server) Serve() error {
	lis, err := net.Listen("tcp", string(gs.listenAddress))
	if err != nil {
		log.Printf("failed to create tcp listener: %v\n", err)
		return err
	}

	reflection.Register(gs.s)

	err = gs.s.Serve(lis)
	if err != nil {
		log.Printf("failed to start listen with grpc: %v\n", err)
		return err
	}
	return nil
}
