package routes

import (
	"fmt"
	"net"

	"github.com/jinzhu/gorm"
	_grpc "github.com/muhammadisa/restful-api-boilerplate/api/foobar/delivery/grpc"
	_foobarRepo "github.com/muhammadisa/restful-api-boilerplate/api/foobar/repository"
	_foobarUsecase "github.com/muhammadisa/restful-api-boilerplate/api/foobar/usecase"

	"google.golang.org/grpc"
)

// GRPCConfigs struct
type GRPCConfigs struct {
	Port     string
	Protocol string
	DB       *gorm.DB
}

// IGRPCConfigs interface
type IGRPCConfigs interface {
	NewGRPC()
}

// NewGRPC grpc service initialization
func (gc GRPCConfigs) NewGRPC() {
	// Initialize grpc server
	listener, err := net.Listen(gc.Protocol, gc.Port)
	if err != nil {
		fmt.Println(fmt.Sprintf("Error while listening on %s", gc.Port))
	}
	fmt.Println(fmt.Sprintf("gRPC Server is Listening on %s", gc.Port))
	server := grpc.NewServer()

	// Init grpc services
	gc.initFoobarService(server)

	// Serve grpc
	err = server.Serve(listener)
	if err != nil {
		fmt.Println("Unexpected Error", err)
	}
}

//
func (gc GRPCConfigs) initFoobarService(server *grpc.Server) {
	foobarRepo := _foobarRepo.NewPostgresFoobarRepo(gc.DB)
	foobarUsecase := _foobarUsecase.NewFoobarUsecase(foobarRepo)
	_grpc.NewFoobarServerGrpc(server, foobarUsecase)
}
